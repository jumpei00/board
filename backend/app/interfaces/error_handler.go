package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context) {
	c.Status(http.StatusBadRequest)
}