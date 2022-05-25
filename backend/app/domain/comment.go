package domain

import "time"

type Comment struct {
	threadKey string
	commentKey string
	contributer string
	comment string
	updateDate time.Time
}

func NewComment() *Comment {
	return &Comment{}
}

func (c *Comment) ThreadKey() string {
	return c.threadKey
}

func (c *Comment) CommentKey() string {
	return c.commentKey
}

func (c *Comment) Contributer() string {
	return c.contributer
}

func (c *Comment) Comment() string {
	return c.comment
}

func (c *Comment) Updatedate() time.Time {
	return c.updateDate
}