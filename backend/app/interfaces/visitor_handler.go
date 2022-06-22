package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/domain"
)

type VisitorsHandler struct {
	visitorApplication application.VisitorApplication
}

func NewVisitorsHandler(va application.VisitorApplication) *VisitorsHandler {
	return &VisitorsHandler{
		visitorApplication: va,
	}
}

func (v *VisitorsHandler) SetupRouter(r *gin.RouterGroup) {
	r.GET("/", v.get)
	r.PUT("/", v.visited)
	r.PUT("/reset", v.reset)
}

// Visitor godoc
// @Summary 訪問者統計の取得
// @Description サイトへの訪問者情報を取得する
// @Tags visitor
// @Accept json
// @Produce json
// @Success 200 {object} responseVisitor
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/visitor [get]
// Visitor godoc
func (v *VisitorsHandler) get(c *gin.Context) {
	visitors, err := v.visitorApplication.GetVisitorsStat()
	if err != nil {
		handleError(c, err)
		return
	}

	res := NewResponseVisitor(visitors)
	c.JSON(http.StatusOK, res)
}

// Visitor godoc
// @Summary 訪問者のカウントアップ
// @Description サイトへの訪問回数をカウントアップさせる
// @Tags visitor
// @Accept json
// @Produce json
// @Success 200 {object} responseVisitor
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/visitor [put]
// Visitor godoc
func (v *VisitorsHandler) visited(c *gin.Context) {
	visitors, err := v.visitorApplication.CountupVisitors()
	if err != nil {
		handleError(c, err)
		return
	}

	res := NewResponseVisitor(visitors)
	c.JSON(http.StatusOK, res)
}

// Visitor godoc
// @Summary 訪問者のリセット
// @Description 昨日の訪問者を今日の訪問者で上書きしリセットさせる
// @Tags visitor
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/visitor/reset [put]
// Visitor godoc
func (v *VisitorsHandler) reset(c *gin.Context) {
	_, err := v.visitorApplication.ResetVisitors()
	if err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type responseVisitor struct {
	Yesterday int `json:"yesterday"`
	Today     int `json:"today"`
	Sum       int `json:"sum"`
}

func NewResponseVisitor(visitor *domain.Visitor) *responseVisitor {
	return &responseVisitor{
		Yesterday: visitor.GetYesterdayVisitor(),
		Today: visitor.GetTodayVisitor(),
		Sum: visitor.GetVisitorSum(),
	}
}
