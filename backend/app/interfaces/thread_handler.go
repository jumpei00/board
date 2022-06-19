package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/domain"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/jumpei00/board/backend/app/params"
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
	r.GET("/", t.getAll)
	r.GET("/:thread_key", t.get)
	r.POST("/", t.create)
	r.PUT("/:thread_key", t.edit)
	r.DELETE("/:thread_key", t.delete)
}

// Thread godoc
// @Summary スレッドを全て取得
// @Description スレッドを全て取得します
// @Tags thread
// @Accept json
// @Produce json
// @Success 200 {object} responseThreads
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/thread [get]
// Thread godoc
func (t *ThreadHandler) getAll(c *gin.Context) {
	threads, err := t.threadApplication.GetAllThread()

	if err != nil {
		handleError(c, err)
		return
	}

	var res responseThreads
	for _, thread := range *threads {
		res.Threads = append(res.Threads, NewResponseThread(&thread))
	}

	c.JSON(http.StatusOK, res)
}

// Thread godoc
// @Summary 指定のスレッドを取得
// @Description スレッドキーに当てはまるスレッドを取得
// @Tags thread
// @Accept json
// @Produce json
// @Param thread_key path string true "スレッドキー"
// @Success 200 {object} responseThread
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/thread/{thread_key} [get]
// Thread godoc
func (t *ThreadHandler) get(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		logger.Warning("thread get, but not thread key")
		handleError(c, appError.NewErrBadRequest(appError.Message().NotThreadKey, "not thread key"))
		return
	}

	thread, err := t.threadApplication.GetByThreadKey(threadKey)

	if err != nil {
		handleError(c, err)
		return
	}

	responseThread := NewResponseThread(thread)
	c.JSON(http.StatusOK, responseThread)
}

// Thread godoc
// @Summary スレッドを新規作成
// @Description 新しいスレッドを作成します
// @Tags thread
// @Accept json
// @Produce json
// @Param body body requestThreadCreate true "スレッド作成情報"
// @Success 200 {object} responseThread
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/thread [post]
// Thread godoc
func (t *ThreadHandler) create(c *gin.Context) {
	var req requestThreadCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("thread create, requesting json bind error", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.CreateThreadAppLayerParam{
		Title:       req.Title,
		Contributor: req.Contributor,
	}

	thread, err := t.threadApplication.CreateThread(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	res := NewResponseThread(thread)
	c.JSON(http.StatusOK, res)
}

// Thread godoc
// @Summary 指定のスレッドを更新
// @Description 指定されたスレッドを編集し更新する
// @Tags thread
// @Accept json
// @Produce json
// @Param thread_key path string true "スレッドキー"
// @Param body body requestThreadEdit true "スレッド編集情報"
// @Success 200 {object} responseThread
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/thread/{thread_key} [put]
// Thread godoc
func (t *ThreadHandler) edit(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		logger.Warning("thread edit, but not thread key")
		handleError(c, appError.NewErrBadRequest(appError.Message().NotThreadKey, "not thread key"))
		return
	}

	var req requestThreadEdit
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("thread edit, requesting json bind error", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.EditThreadAppLayerParam{
		ThreadKey:   threadKey,
		Title:       req.Title,
		Contributor: req.Contributor,
	}

	thread, err := t.threadApplication.EditThread(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	res := NewResponseThread(thread)
	c.JSON(http.StatusOK, res)
}

// Thread godoc
// @Summary 指定のスレッドを削除
// @Description 指定されたスレッドを削除し、それに紐づいているコメントも同時に削除する
// @Tags thread
// @Accept json
// @Produce json
// @Param thread_key path string true "スレッドキー"
// @Param body body requestThreadDelete true "スレッド削除情報"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/thread/{thread_key} [delete]
// Thread godoc
func (t *ThreadHandler) delete(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		logger.Warning("thread delete, but not thread key")
		handleError(c, appError.NewErrBadRequest(appError.Message().NotThreadKey, "not thread key"))
		return
	}

	var req requestThreadDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("thread delete, requesting json bind error", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.DeleteThreadAppLayerParam{
		ThreadKey:   threadKey,
		Contributor: req.Contributor,
	}

	if err := t.threadApplication.DeleteThread(&param); err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type requestThreadCreate struct {
	Title       string `json:"title"`
	Contributor string `json:"contributor"`
}

type requestThreadEdit struct {
	Title       string `json:"title"`
	Contributor string `json:"contributor"`
}

type requestThreadDelete struct {
	Contributor string `json:"contributor"`
}

type responseThreads struct {
	Threads []*responseThread `json:"threads"`
}

type responseThread struct {
	ThreadKey   string `json:"thread_key"`
	Title       string `json:"title"`
	Contributor string `json:"contributor"`
	UpdateDate  string `json:"update_date"`
	Views       int    `json:"views"`
	CommentSum  int    `json:"comment_sum"`
}

func NewResponseThread(thread *domain.Thread) *responseThread {
	return &responseThread{
		ThreadKey:   thread.GetKey(),
		Title:       thread.GetTitle(),
		Contributor: thread.GetContributor(),
		UpdateDate:  thread.FormatUpdatedDate(),
		Views:       thread.GetViews(),
		CommentSum:  thread.GetCommentSum(),
	}
}
