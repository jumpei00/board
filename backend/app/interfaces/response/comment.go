package response

import "github.com/jumpei00/board/backend/app/domain"

type ResponseThreadAndComments struct {
	Thread   *ResponseThread    `json:"thread"`
	Comments []*ResponseComment `json:"comments"`
}

type ResponseComment struct {
	CommentKey  string `joson:"comment_key"`
	Contributor string `json:"contributor"`
	Comment     string `json:"comment"`
	UpdateDate  string `json:"update_date"`
}

func NewResponseComment(comment *domain.Comment) *ResponseComment {
	return &ResponseComment{
		CommentKey:  comment.GetKey(),
		Contributor: comment.GetContributor(),
		Comment:     comment.GetComment(),
		UpdateDate:  comment.FormatUpdateDate(),
	}
}
