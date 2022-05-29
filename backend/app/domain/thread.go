package domain

import (
	"time"

	"github.com/jumpei00/board/backend/app/params"
)

type Thread struct {
	threadKey   string
	title       string
	contributor string
	updateDate    time.Time
	views       int
	sumComment  int
}

func NewThread(param *params.CreateThreadDomainLayerParam) *Thread {
	thread := &Thread{
		title:       param.Title,
		contributor: param.Contributor,
		updateDate:    time.Now(),
		views:       0,
		sumComment:  0,
	}
	thread.setRandomThreadKey()

	return thread
}

func (t *Thread) UpdateThread(param *params.EditThreadDomainLayerParam) *Thread {
	t.threadKey = param.ThreadKey
	t.title = param.Title
	t.contributor = param.Contributor
	t.updateDate = param.UpdateDate
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

func (t *Thread) UpdateDate() time.Time {
	return t.updateDate
}

func (t *Thread) FormatUpdateDate() string {
	return t.updateDate.Format("2006/01/02 15:04")
}

func (t *Thread) UpdateLatestUpdateDate() {
	t.updateDate = time.Now()
}

func (t *Thread) Views() int {
	return t.views
}

func (t *Thread) CountupViews() {
	t.views += 1
}

func (t *Thread) SumComment() int {
	return t.sumComment
}

func (t *Thread) CountupSumComment() {
	t.sumComment += 1
}
