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

func NewThreadRepository(dbPool *gorm.DB) *threadRepository {
	return &threadRepository{
		db: dbPool,
	}
}

func (t *threadRepository) GetAll() (*[]domain.Thread, error) {
	var threads []domain.Thread

	// TODO: スレッド数が増えてきたらLIMITをかけてページネーションなどにする
	result := t.db.Order("updated_at desc").Find(&threads)

	if result.Error != nil {
		logger.Error("thread get all failed", "error", result.Error)
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

	if err := t.db.Where(&domain.Thread{Key: threadKey}).First(&thread).Error; err != nil {
		if errors.Cause(err) == gorm.ErrRecordNotFound {
			logger.Info("thread get at targed thread_key, but no thread", "thread_key", threadKey)
			return nil, appError.NewErrNotFound("thread get at targed thread_key, but no thread -> thread_key: %s", threadKey)
		}
		logger.Error("thread get at targed thread_key failed", "error", err, "thread_key", threadKey)
		return nil, errors.WithStack(err)
	}

	return &thread, nil
}

func (t *threadRepository) Insert(thread *domain.Thread) (*domain.Thread, error) {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(thread).Error; err != nil {
			logger.Error("thread insert failed", "error", err, "thread", thread)
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
			logger.Error("thread update failed", "error", err, "thread", thread)
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
			logger.Error("thread delete failed", "error", err, "thread", thread)
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
		if err := tx.Delete(comments).Error; err != nil {
			logger.Error("targeted all comments delete failed", "error", err, "thread_key", thread.GetKey())
			return err
		}

		if err := tx.Delete(thread).Error; err != nil {
			logger.Error("thread delete failed", "error", err, "thread", thread)
			return err
		}

		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}