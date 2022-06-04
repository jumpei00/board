package infrastructure

import "github.com/jumpei00/board/backend/app/domain"

type CommentDB struct{}

func NewCommentDB() *CommentDB {
	return &CommentDB{}
}

func (c *CommentDB) GetAllByKey(threadKey string) ([]*domain.Comment, error) {
	return []*domain.Comment{}, nil
}

func (c *CommentDB) GetByKey(commentKey string) (*domain.Comment, error) {
	return &domain.Comment{}, nil
}

func (c *CommentDB) Insert(comment *domain.Comment) (*domain.Comment, error) {
	return &domain.Comment{}, nil
}

func (c *CommentDB) Update(comment *domain.Comment) (*domain.Comment, error) {
	return &domain.Comment{}, nil
}

func (c *CommentDB) Delete(comment *domain.Comment) error {
	return nil
}
