package domain

import (
	"time"

	"github.com/jumpei00/board/backend/app/params"
)

type Thread struct {
	threadKey   string
	title       string
	contributor string
	postDate    time.Time
	views       int
	sumComment  int
}

func NewThread(param *params.ThreadCreateDomainLayerParam) *Thread {
	thread := &Thread{
		title:       param.Title,
		contributor: param.Contributor,
		postDate:    time.Now(),
		views:       0,
		sumComment:  0,
	}
	thread.setRandomThreadKey()

	return thread
}

func (t *Thread) UpdateThread(param *params.ThreadEditDomainLayerParam) *Thread {
	t.threadKey = param.ThreadKey
	t.title = param.Title
	t.contributor = param.Contributor
	t.postDate = param.PostDate
	t.views = param.Views
	t.sumComment = param.SumComment

	return t
}

func (t *Thread) ThreadKey() string {
	return t.threadKey
}

func (t *Thread) setRandomThreadKey() {
	t.threadKey = "hello"
}

func (t *Thread) Title() string {
	return t.title
}

func (t *Thread) Contributor() string {
	return t.contributor
}

func (t *Thread) IsNotSameContritubor(person string) bool {
	return t.contributor != person
}

func (t *Thread) PostDate() time.Time {
	return t.postDate
}

func (t *Thread) FormatPostDate() string {
	return t.postDate.Format("2006/01/02 15:04")
}

func (t *Thread) Views() int {
	return t.views
}

func (t *Thread) SumComment() int {
	return t.sumComment
}
