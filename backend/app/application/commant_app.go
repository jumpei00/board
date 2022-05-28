package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
)

type CommentApplication interface {
	GetAllByThreadKey(threadKey string) []*domain.Comment
	CreateComment()
	EditComment()
	DeleteComment(threadKey, commentKey string)
}

type commentApplication struct {
	commentRepo repository.CommentRepository
}

func NewCommentApplication(cr repository.CommentRepository) *commentApplication {
	return &commentApplication{
		commentRepo: cr,
	}
}

func (c *commentApplication) GetAllByThreadKey(threadKey string) []*domain.Comment {
	return c.commentRepo.GetAllByKey(threadKey)
}

func (c *commentApplication) CreateComment() {}

func (c *commentApplication) EditComment() {}

func (c *commentApplication) DeleteComment(threadKey, commentKey string) {} 