package infrastructure

import "github.com/jumpei00/board/backend/app/domain"

type CommentDB struct{}

func NewCommentDB() *CommentDB {
	return &CommentDB{}
}

func (c *CommentDB) GetAllByKey(threadKey string) []*domain.Comment {
	return []*domain.Comment{}
}

func (c *CommentDB) Insert(comment *domain.Comment) {}

func (c *CommentDB) Update(comment *domain.Comment) {}

func (c *CommentDB) Delete(threadKey, commentKey string) {}
