package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/domain"
)

type VisitorsHandler struct {
	visitorApp application.VisitorApplication
}

func NewVisitorsHandler(va application.VisitorApplication) *VisitorsHandler {
	return &VisitorsHandler{
		visitorApp: va,
	}
}

func (v *VisitorsHandler) SetupRouter(r *gin.RouterGroup) {
	r.GET("/", v.get)
	r.PUT("/", v.visited)
	r.PUT("/reset", v.reset)
}

func (v *VisitorsHandler) get(c *gin.Context) {
	visitors, err := v.visitorApp.GetVisitorsStat()
	if err != nil {
		handleError(c, err)
		return
	}

	res := NewResponseVisitors(visitors)
	c.JSON(http.StatusOK, res)
}

func (v *VisitorsHandler) visited(c *gin.Context) {
	visitors, err := v.visitorApp.CountupVisitors()
	if err != nil {
		handleError(c, err)
		return
	}

	res := NewResponseVisitors(visitors)
	c.JSON(http.StatusOK, res)
}

func (v *VisitorsHandler) reset(c *gin.Context) {
	_, err := v.visitorApp.ResetVisitors()
	if err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type responseVisitors struct {
	Yesterday int `json:"yesterday"`
	Today     int `json:"today"`
	Sum       int `json:"sum"`
}

func NewResponseVisitors(visitor *domain.Visitor) *responseVisitors {
	return &responseVisitors{
		Yesterday: visitor.GetYesterdayVisitor(),
		Today: visitor.GetTodayVisitor(),
		Sum: visitor.GetVisitorSum(),
	}
}
