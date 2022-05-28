package domain

import "time"

type Comment struct {
	threadKey string
	commentKey string
	contributor string
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

func (c *Comment) Contributor() string {
	return c.contributor
}

func (c *Comment) Comment() string {
	return c.comment
}

func (c *Comment) Updatedate() time.Time {
	return c.updateDate
}