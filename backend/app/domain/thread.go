package domain

import "time"

type Thread struct {
	threadKey         string
	title       string
	contributer string
	postDate    time.Time
	views       int
	sumComment  int
}

func NewThread() *Thread {
	return &Thread{}
}

func (t *Thread) ThreadKey() string {
	return t.threadKey
}

func (t *Thread) Title() string {
	return t.title
}

func (t *Thread) Contributer() string {
	return t.contributer
}

func (t *Thread) PostDate() time.Time {
	return t.postDate
}

func (t *Thread) Views() int {
	return t.views
}

func (t *Thread) SumComment() int {
	return t.sumComment
}
