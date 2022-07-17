package request

type RequestThreadCreate struct {
	Title       string `json:"title" binding:"required"`
	Contributor string `json:"contributor" binding:"required"`
}

type RequestThreadEdit struct {
	Title       string `json:"title" binding:"required"`
	Contributor string `json:"contributor" binding:"required"`
}
