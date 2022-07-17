package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/application/params"
	"github.com/jumpei00/board/backend/app/interfaces/middleware"
	"github.com/jumpei00/board/backend/app/interfaces/request"
	"github.com/jumpei00/board/backend/app/interfaces/response"
	"github.com/jumpei00/board/backend/app/interfaces/session"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/pkg/errors"
)

type ThreadHandler struct {
	sessionManager    session.Manager
	threadApplication application.ThreadApplication
}

func NewThreadHandler(sm session.Manager, ta application.ThreadApplication) *ThreadHandler {
	return &ThreadHandler{
		sessionManager:    sm,
		threadApplication: ta,
	}
}

func (t *ThreadHandler) SetupRouter(r *gin.RouterGroup) {
	operatePermissionMiddleware := middleware.NewOperatePermissionMiddleware(t.sessionManager)

	r.GET("", t.getAll)
	r.GET("/:threadKey", t.get)
	r.POST("", operatePermissionMiddleware, t.create)
	r.PUT("/:threadKey", operatePermissionMiddleware, t.edit)
	r.DELETE("/:threadKey", operatePermissionMiddleware, t.delete)
}

// Thread godoc
// @Summary スレッドを全て取得
// @Description スレッドを全て取得します
// @Tags thread
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseThreads
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/threads [get]
// Thread godoc
func (t *ThreadHandler) getAll(c *gin.Context) {
	threads, err := t.threadApplication.GetAllThread()

	if err != nil && errors.Cause(err) != appError.ErrNotFound {
		handleError(c, err)
		return
	}

	var res response.ResponseThreads
	if threads != nil {
		for _, thread := range *threads {
			res.Threads = append(res.Threads, response.NewResponseThread(&thread))
		}
	}

	c.JSON(http.StatusOK, res)
}

// Thread godoc
// @Summary 指定のスレッドを取得
// @Description スレッドキーに当てはまるスレッドを取得
// @Tags thread
// @Accept json
// @Produce json
// @Param threadKey path string true "スレッドキー"
// @Success 200 {object} response.ResponseThread
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/threads/{threadKey} [get]
// Thread godoc
func (t *ThreadHandler) get(c *gin.Context) {
	threadKey := c.Param("threadKey")

	thread, err := t.threadApplication.GetByThreadKey(threadKey)

	if err != nil {
		handleError(c, err)
		return
	}

	responseThread := response.NewResponseThread(thread)
	c.JSON(http.StatusOK, responseThread)
}

// Thread godoc
// @Summary スレッドを新規作成
// @Description 新しいスレッドを作成します
// @Tags thread
// @Accept json
// @Produce json
// @Param body body request.RequestThreadCreate true "スレッド作成情報"
// @Success 200 {object} response.ResponseThread
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/threads [post]
// Thread godoc
func (t *ThreadHandler) create(c *gin.Context) {
	var req request.RequestThreadCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("thread create, requesting json bind error", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	user, _ := t.sessionManager.Get(c)

	param := params.CreateThreadAppLayerParam{
		Title:  req.Title,
		UserID: user.UserID,
	}

	thread, err := t.threadApplication.CreateThread(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	res := response.NewResponseThread(thread)
	c.JSON(http.StatusOK, res)
}

// Thread godoc
// @Summary 指定のスレッドを更新
// @Description 指定されたスレッドを編集し更新する
// @Tags thread
// @Accept json
// @Produce json
// @Param threadKey path string true "スレッドキー"
// @Param body body request.RequestThreadEdit true "スレッド編集情報"
// @Success 200 {object} response.ResponseThread
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/threads/{threadKey} [put]
// Thread godoc
func (t *ThreadHandler) edit(c *gin.Context) {
	threadKey := c.Param("threadKey")
	user, _ := t.sessionManager.Get(c)

	var req request.RequestThreadEdit
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("thread edit, requesting json bind error", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.EditThreadAppLayerParam{
		ThreadKey: threadKey,
		Title:     req.Title,
		UserID:    user.UserID,
	}

	thread, err := t.threadApplication.EditThread(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	res := response.NewResponseThread(thread)
	c.JSON(http.StatusOK, res)
}

// Thread godoc
// @Summary 指定のスレッドを削除
// @Description 指定されたスレッドを削除し、それに紐づいているコメントも同時に削除する
// @Tags thread
// @Accept json
// @Produce json
// @Param threadKey path string true "スレッドキー"
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/threads/{threadKey} [delete]
// Thread godoc
func (t *ThreadHandler) delete(c *gin.Context) {
	threadKey := c.Param("threadKey")
	user, _ := t.sessionManager.Get(c)

	param := params.DeleteThreadAppLayerParam{
		ThreadKey: threadKey,
		UserID:    user.UserID,
	}

	if err := t.threadApplication.DeleteThread(&param); err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
