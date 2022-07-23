package response

import "github.com/jumpei00/board/backend/app/domain"

type ResponseThreads struct {
	Threads []*ResponseThread `json:"threads"`
}

type ResponseThread struct {
	ThreadKey   string `json:"thread_key"`
	Title       string `json:"title"`
	Contributor string `json:"contributor"`
	Views       int    `json:"views"`
	CommentSum  int    `json:"comment_sum"`
	CreateDate  string `json:"create_date"`
	UpdateDate  string `json:"update_date"`
}

func NewResponseThread(thread *domain.Thread) *ResponseThread {
	return &ResponseThread{
		ThreadKey:   thread.GetKey(),
		Title:       thread.GetTitle(),
		Contributor: thread.GetContributor(),
		Views:       thread.GetViews(),
		CommentSum:  thread.GetCommentSum(),
		CreateDate:  thread.FormatCreatedDate(),
		UpdateDate:  thread.FormatUpdatedDate(),
	}
}
