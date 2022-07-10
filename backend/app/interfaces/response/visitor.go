package response

import "github.com/jumpei00/board/backend/app/domain"

type responseVisitor struct {
	Yesterday int `json:"yesterday"`
	Today     int `json:"today"`
	Sum       int `json:"sum"`
}

func NewResponseVisitor(visitor *domain.Visitor) *responseVisitor {
	return &responseVisitor{
		Yesterday: visitor.GetYesterdayVisitor(),
		Today:     visitor.GetTodayVisitor(),
		Sum:       visitor.GetVisitorSum(),
	}
}
