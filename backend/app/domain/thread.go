package domain

import (
	"time"

	"github.com/google/uuid"

	"github.com/jumpei00/board/backend/app/params"
)

type Thread struct {
	Key         string    `gorm:"primaryKey;column:thread_key"`
	Title       string    `gorm:"column:title"`
	Contributor string    `gorm:"column:contributor"`
	Views       *int      `gorm:"column:views;default:0"`
	CommentSum  *int      `gorm:"column:comment_sum;default:0"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func NewThread(param *params.CreateThreadDomainLayerParam) *Thread {
	initViews := 0
	initCommentSum := 0

	thread := &Thread{
		Title:       param.Title,
		Contributor: param.Contributor,
		Views:       &initViews,
		CommentSum:  &initCommentSum,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	thread.setKey()

	return thread
}

func (t *Thread) UpdateThread(param *params.EditThreadDomainLayerParam) *Thread {
	t.Title = param.Title
	t.UpdatedAt = time.Now()

	return t
}

func (t *Thread) GetKey() string {
	return t.Key
}

func (t *Thread) GetTitle() string {
	return t.Title
}

func (t *Thread) GetContributor() string {
	return t.Contributor
}

func (t *Thread) GetCreatedAt() time.Time {
	return t.CreatedAt
}

func (t *Thread) GetUpdatedAt() time.Time {
	return t.UpdatedAt
}

func (t *Thread) GetViews() int {
	return *t.Views
}

func (t *Thread) GetCommentSum() int {
	return *t.CommentSum
}

func (t *Thread) setKey() {
	t.Key = uuid.New().String()
}

func (t *Thread) IsNotSameContributor(contributor string) bool {
	return t.Contributor != contributor
}

func (t *Thread) FormatCreatedDate() string {
	return t.CreatedAt.Format("2006/01/02 15:04")
}

func (t *Thread) FormatUpdatedDate() string {
	return t.UpdatedAt.Format("2006/01/02 15:04")
}

func (t *Thread) UpdateLatestUpdatedDate() {
	t.UpdatedAt = time.Now()
}

func (t *Thread) CountupViews() {
	*t.Views += 1
}

func (t *Thread) CountupCommentSum() {
	*t.CommentSum += 1
}

func (t *Thread) CountdownCommentSum() {
	*t.CommentSum -= 1
}
