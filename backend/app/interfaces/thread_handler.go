package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
)

type ThreadHandler struct {
	threadApplication application.ThreadApplication
}

func NewThreadHandler(ta application.ThreadApplication) *ThreadHandler {
	return &ThreadHandler{
		threadApplication: ta,
	}
}

func (t *ThreadHandler) SetupRouter(r *gin.RouterGroup) {
	r.GET("/", t.GetAll)
}

func (t *ThreadHandler) GetAll(c *gin.Context) {
	threads := t.threadApplication.GetAllThread()

	c.JSON(http.StatusOK, threads)
}