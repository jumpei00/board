package domain

import (
	"time"

	"github.com/jumpei00/board/backend/app/params"
)

type Comment struct {
	threadKey   string
	commentKey  string
	contributor string
	comment     string
	updateDate  time.Time
}

func NewComment(param *params.CreateCommentDomainLayerParam) *Comment {
	comment := &Comment{
		threadKey: param.ThreadKey,
		contributor: param.Contributor,
		comment: param.Comment,
		updateDate: time.Now(),
	}
	comment.setCommentKey()
	return comment
}

func (c *Comment) UpdateComment(param *params.EditCommentDomainLayerParam) *Comment {
	c.comment = param.Comment
	c.updateDate = time.Now()
	return c
}

func (c *Comment) ThreadKey() string {
	return c.threadKey
}

func (c *Comment) CommentKey() string {
	return c.commentKey
}

func (c *Comment) Contributor() string {
	return c.contributor
}

func (c *Comment) Comment() string {
	return c.comment
}

func (c *Comment) UpdateDate() time.Time {
	return c.updateDate
}

func (c *Comment) setCommentKey() {
	c.commentKey = "key"
}

func (c *Comment) IsNotSameContritubor(contributor string) bool {
	return c.contributor != contributor
}

func (c *Comment) FormatUpdateDate() string {
	return c.updateDate.Format("2006/01/02 15:04")
}
