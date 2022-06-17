package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/jumpei00/board/backend/app/params"
)

type Comment struct {
	Key         string    `gorm:"column:comment_key"`
	ThreadKey   string    `gorm:"column:thread_key"`
	Contributor string    `gorm:"column:contributor"`
	Comment     string    `gorm:"comment"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}

func NewComment(param *params.CreateCommentDomainLayerParam) *Comment {
	comment := &Comment{
		ThreadKey:   param.ThreadKey,
		Contributor: param.Contributor,
		Comment:     param.Comment,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	comment.setCommentKey()
	return comment
}

func (c *Comment) UpdateComment(param *params.EditCommentDomainLayerParam) *Comment {
	c.Comment = param.Comment
	c.UpdatedAt = time.Now()
	return c
}

func (c *Comment) GetKey() string {
	return c.Key
}

func (c *Comment) GetThreadKey() string {
	return c.ThreadKey
}

func (c *Comment) GetContributor() string {
	return c.Contributor
}

func (c *Comment) GetComment() string {
	return c.Comment
}

func (c *Comment) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *Comment) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}

func (c *Comment) setCommentKey() {
	c.Key = uuid.New().String()
}

func (c *Comment) IsNotSameContritubor(contributor string) bool {
	return c.Contributor != contributor
}

func (t *Comment) FormatCreatedDate() string {
	return t.CreatedAt.Format("2006/01/02 15:04")
}

func (c *Comment) FormatUpdateDate() string {
	return c.UpdatedAt.Format("2006/01/02 15:04")
}
