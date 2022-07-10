package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/pkg/errors"
)

func handleError(c *gin.Context, err error) {
	if err == nil {
		c.Status(http.StatusOK)
		return
	}

	cause := errors.Cause(err)

	if appError.IsBadRequest(cause) {
		c.JSON(http.StatusBadRequest, cause)
		logger.Info("bad request", "cause", cause)
		return
	}

	switch cause {
	case appError.ErrNotFound:
		logger.Info("not found", "cause", cause)
		c.Status(http.StatusNotFound)
	case appError.ErrUnauthorized:
		logger.Info("unauthorized", "cause", cause)
		c.Status(http.StatusUnauthorized)
	default:
		logger.Error("internal server error", "cause", cause)
		c.Status(http.StatusInternalServerError)
	}
}
