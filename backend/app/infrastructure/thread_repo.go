package infrastructure

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/library/logger"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type threadRepository struct{
	db *gorm.DB
}

func NewThreadRepository(dbSession *gorm.DB) *threadRepository {
	return &threadRepository{
		db: dbSession,
	}
}

func (t *threadRepository) GetAll() (*[]domain.Thread, error) {
	var threads []domain.Thread

	// TODO: スレッド数が増えてきたらLIMITをかけてページネーションなどにする
	result := t.db.Order("updated_at desc").Find(&threads)

	if result.Error != nil {
		logger.Error("thread get all failed -> err: %s", result.Error)
		return nil, errors.WithStack(result.Error)
	}

	if result.RowsAffected == 0 {
		logger.Info("thread get all, but not thread")
		return nil, appError.NewErrNotFound("thread get all, but not thread")
	}

	return &threads, nil
}

func (t *threadRepository) GetByKey(threadKey string) (*domain.Thread, error) {
	var thread domain.Thread

	if err := t.db.Where("key = ?", threadKey).First(&thread).Error; err != nil {
		if errors.Cause(err) == gorm.ErrRecordNotFound {
			logger.Info("no thread by target key", "key", threadKey)
			return nil, appError.NewErrNotFound("target key thread get failed -> err: %s, key: %s", err, threadKey)
		}
		errors.WithStack(err)
		return nil, err
	}

	return &thread, nil
}

func (t *threadRepository) Insert(thread *domain.Thread) (*domain.Thread, error) {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(thread).Error; err != nil {
			logger.Error("thread insert failed -> err: %s, target thread: %s", err, thread)
			return err
		}
		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return thread, nil
}

func (t *threadRepository) Update(thread *domain.Thread) (*domain.Thread, error) {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(thread).Error; err != nil {
			logger.Error("thread update failed -> err: %s, target thread: %s", err, thread)
			return err
		}
		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return thread, nil
}

func (t *threadRepository) Delete(thread *domain.Thread) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(thread).Error; err != nil {
			logger.Error("thread delete failed -> err: %s, target thread: %s", err, thread)
			return err
		}
		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (t *threadRepository) DeleteThreadAndComments(thread *domain.Thread, comments *[]domain.Comment) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(thread).Error; err != nil {
			logger.Error("thread delete failed -> err: %s, target thread: %s", err, thread)
			return err
		}

		if err := tx.Delete(comments).Error; err != nil {
			logger.Error("targeted all comments delete failed -> err: %s, target thread_key: %s", err, thread.GetKey())
			return err
		}

		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}