package request

type RequestThreadCreate struct {
	Title       string `json:"title" binding:"required"`
}

type RequestThreadEdit struct {
	Title       string `json:"title" binding:"required"`
}
