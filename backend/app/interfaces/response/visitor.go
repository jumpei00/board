package response

import "github.com/jumpei00/board/backend/app/domain"

type ResponseVisitor struct {
	Yesterday int `json:"yesterday"`
	Today     int `json:"today"`
	Sum       int `json:"sum"`
}

func NewResponseVisitor(visitor *domain.Visitor) *ResponseVisitor {
	return &ResponseVisitor{
		Yesterday: visitor.GetYesterdayVisitor(),
		Today:     visitor.GetTodayVisitor(),
		Sum:       visitor.GetVisitorSum(),
	}
}
