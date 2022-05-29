package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
	"github.com/jumpei00/board/backend/app/params"
)

type CommentApplication interface {
	GetAllByThreadKey(threadKey string) ([]*domain.Comment, error)
	CreateComment(param *params.CreateCommentAppLayerParam) (*domain.Comment, error)
	EditComment(param *params.EditCommentAppLayerParam) (*domain.Comment, error)
	DeleteComment(param *params.DeleteCommentAppLayerParam) error
}

type commentApplication struct {
	commentRepo repository.CommentRepository
}

func NewCommentApplication(cr repository.CommentRepository) *commentApplication {
	return &commentApplication{
		commentRepo: cr,
	}
}

func (c *commentApplication) GetAllByThreadKey(threadKey string) ([]*domain.Comment, error) {
	comments, err := c.commentRepo.GetAllByKey(threadKey)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *commentApplication) CreateComment(param *params.CreateCommentAppLayerParam) (*domain.Comment, error) {

}

func (c *commentApplication) EditComment(param *params.EditCommentAppLayerParam) (*domain.Comment, error){

}

func (c *commentApplication) DeleteComment(params *params.DeleteCommentAppLayerParam) error {
	
}
