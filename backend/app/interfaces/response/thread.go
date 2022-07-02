package response

import "github.com/jumpei00/board/backend/app/domain"

type ResponseThreads struct {
	Threads []*ResponseThread `json:"threads"`
}

type ResponseThread struct {
	ThreadKey   string `json:"thread_key"`
	Title       string `json:"title"`
	Contributor string `json:"contributor"`
	UpdateDate  string `json:"update_date"`
	Views       int    `json:"views"`
	CommentSum  int    `json:"comment_sum"`
}

func NewResponseThread(thread *domain.Thread) *ResponseThread {
	return &ResponseThread{
		ThreadKey:   thread.GetKey(),
		Title:       thread.GetTitle(),
		Contributor: thread.GetContributor(),
		UpdateDate:  thread.FormatUpdatedDate(),
		Views:       thread.GetViews(),
		CommentSum:  thread.GetCommentSum(),
	}
}
