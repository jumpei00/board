package request

type RequestCommentCreate struct {
	Comment     string `json:"comment" binding:"required"`
	Contributor string `json:"contributor" binding:"required"`
}

type RequestCommentEdit struct {
	CommentKey  string `json:"comment_key" binding:"required"`
	Comment     string `json:"comment" binding:"required"`
	Contributor string `json:"contributor" binding:"required"`
}

type RequestCommentDelete struct {
	CommentKey  string `json:"comment_key" binding:"required"`
	Contributor string `json:"contributor" binding:"required"`
}