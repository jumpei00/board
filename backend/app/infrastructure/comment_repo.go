package infrastructure

import (
	"github.com/jumpei00/board/backend/app/domain"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CommentDB struct {
	db *gorm.DB
}

func NewCommentDB(dbPool *gorm.DB) *CommentDB {
	return &CommentDB{
		db: dbPool,
	}
}

func (c *CommentDB) GetAllByKey(threadKey string) (*[]domain.Comment, error) {
	var comments []domain.Comment

	// コメント数が増えるとページネーションにする必要がある
	result := c.db.Order("updated_at desc").Where(&domain.Comment{ThreadKey: threadKey}).Find(&comments)

	if result.Error != nil {
		logger.Error("comments get all failed at targed thread_key", "error", result.Error, "thread_key", threadKey)
		return nil, errors.WithStack(result.Error)
	}

	if result.RowsAffected == 0 {
		logger.Info("comment get all at taeged thread_key, but no comments", "thread_key", threadKey)
		return nil, appError.NewErrNotFound("comment get all at targed thread_key, but no comments -> thread_key: %s", threadKey)
	}

	return &comments, nil
}

func (c *CommentDB) GetByKey(commentKey string) (*domain.Comment, error) {
	var comment domain.Comment

	if err := c.db.Where(&domain.Comment{Key: commentKey}).First(&comment).Error; err != nil {
		if errors.Cause(err) == gorm.ErrRecordNotFound {
			logger.Info("comment get at targed comment_key, but no comment", "comment_key", commentKey)
			return nil, appError.NewErrNotFound("comment get at targed comment_key, but no comment -> comment_key: %s", commentKey)
		}
		logger.Error("comment get at targed comment_key failed", "error", err, "comment_key", commentKey)
		return nil, errors.WithStack(err)
	}

	return &comment, nil
}

func (c *CommentDB) Insert(comment *domain.Comment) (*domain.Comment, error) {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(comment).Error; err != nil {
			logger.Error("comment insert failed", "error", err)
			return err
		}
		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return comment, nil
}

func (c *CommentDB) Update(comment *domain.Comment) (*domain.Comment, error) {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(comment).Error; err != nil {
			logger.Error("comment update failed", "error", err)
			return err
		}
		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return comment, nil
}

func (c *CommentDB) Delete(comment *domain.Comment) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(comment).Error; err != nil {
			logger.Error("comment delete failed", "error", err)
			return err
		}
		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
