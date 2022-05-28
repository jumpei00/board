package repository

import "github.com/jumpei00/board/backend/app/domain"

type CommentRepository interface {
	GetAllByKey(threadKey string) []*domain.Comment
	Insert(comment *domain.Comment)
	Update(comment *domain.Comment)
	Delete(threadKey, commentKey string)
}