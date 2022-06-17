package repository

import "github.com/jumpei00/board/backend/app/domain"

type CommentRepository interface {
	GetAllByKey(threadKey string) (*[]domain.Comment, error)
	GetByKey(commentKey string) (*domain.Comment, error)
	Insert(comment *domain.Comment) (*domain.Comment, error)
	Update(comment *domain.Comment) (*domain.Comment, error)
	Delete(comment *domain.Comment) error
}