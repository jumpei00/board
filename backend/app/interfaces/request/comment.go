package request

type RequestCommentCreate struct {
	Comment     string `json:"comment" binding:"required"`
}

type RequestCommentEdit struct {
	Comment     string `json:"comment" binding:"required"`
}
