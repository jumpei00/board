package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/interfaces/session"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/pkg/errors"
)

func NewOperatePermissionMiddleware(sessionManager session.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := sessionManager.Get(c)

		if errors.Cause(err) == appError.ErrNotFound {
			c.Abort()
			c.Status(http.StatusUnauthorized)
			logger.Error("cannot operate, please login", "session", session)
			return
		}

		if err != nil {
			c.Abort()
			c.Status(http.StatusInternalServerError)
			logger.Error("session get failed", "error", err, "session", session)
			return
		}
	}
}