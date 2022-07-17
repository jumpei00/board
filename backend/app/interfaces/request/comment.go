package request

type RequestCommentCreate struct {
	Comment     string `json:"comment" binding:"required"`
	Contributor string `json:"contributor" binding:"required"`
}

type RequestCommentEdit struct {
	Comment     string `json:"comment" binding:"required"`
	Contributor string `json:"contributor" binding:"required"`
}
