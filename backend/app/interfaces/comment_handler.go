package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/interfaces/middleware"
	"github.com/jumpei00/board/backend/app/interfaces/request"
	"github.com/jumpei00/board/backend/app/interfaces/response"
	"github.com/jumpei00/board/backend/app/interfaces/session"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/jumpei00/board/backend/app/params"
	"github.com/pkg/errors"
)

type CommentHandler struct {
	sessionManager     session.Manager
	threadApplication  application.ThreadApplication
	commentApplication application.CommentApplication
}

func NewCommentHandler(sm session.Manager, ta application.ThreadApplication, ca application.CommentApplication) *CommentHandler {
	return &CommentHandler{
		sessionManager:     sm,
		threadApplication:  ta,
		commentApplication: ca,
	}
}

func (co *CommentHandler) SetupRouter(r *gin.RouterGroup) {
	operatePermissionMiddleware := middleware.NewOperatePermissionMiddleware(co.sessionManager)

	r.GET("/:threadKey", co.getAll)
	r.POST("/:threadKey", operatePermissionMiddleware, co.create)
	r.PUT("/:threadKey", operatePermissionMiddleware, co.edit)
	r.DELETE("/:threadKey", operatePermissionMiddleware, co.delete)
}

// Comment godoc
// @Summary コメントを全て取得
// @Description スレッドに紐づくコメントを全て取得
// @Tags comment
// @Accept json
// @Produce json
// @Param threadKey path string true "スレッドキー"
// @Success 200 {object} responseThreadAndComments
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/comment/{threadKey} [get]
// Comment godoc
func (co *CommentHandler) getAll(c *gin.Context) {
	threadKey := c.Param("threadKey")

	thread, err := co.threadApplication.GetByThreadKey(threadKey)
	if err != nil {
		handleError(c, err)
		return
	}

	comments, err := co.commentApplication.GetAllByThreadKey(threadKey)

	// もしエラーがNotFound以外のエラーだった場合はそこでエラーを返す
	if err != nil && errors.Cause(err) != appError.ErrNotFound {
		handleError(c, err)
		return
	}

	var res response.ResponseThreadAndComments
	res.Thread = response.NewResponseThread(thread)

	if comments != nil {
		for _, comment := range *comments {
			res.Comments = append(res.Comments, response.NewResponseComment(&comment))
		}
	}

	c.JSON(http.StatusOK, res)
}

// Comment godoc
// @Summary スレッドへの新規コメントを作成
// @Description スレッドに対するコメントを作成する
// @Tags comment
// @Accept json
// @Produce json
// @Param threadKey path string true "スレッドキー"
// @Param body body request.requestCommentCreate true "コメント作成情報"
// @Success 200 {object} response.responseThreadAndComments
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/comment/{threadKey} [post]
// Comment godoc
func (co *CommentHandler) create(c *gin.Context) {
	threadKey := c.Param("threadKey")

	var req request.RequestCommentCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("comment create, requesting json bind error", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.CreateCommentAppLayerParam{
		ThreadKey:   threadKey,
		Comment:     req.Comment,
		Contributor: req.Contributor,
	}

	comments, err := co.commentApplication.CreateComment(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	thread, err := co.threadApplication.GetByThreadKey(threadKey)
	if err != nil {
		handleError(c, err)
		return
	}

	var res response.ResponseThreadAndComments
	res.Thread = response.NewResponseThread(thread)
	for _, comment := range *comments {
		res.Comments = append(res.Comments, response.NewResponseComment(&comment))
	}

	c.JSON(http.StatusOK, res)
}

// Comment godoc
// @Summary 指定されたコメントを更新
// @Description 指定されたコメントを編集し更新する
// @Tags comment
// @Accept json
// @Produce json
// @Param threadKey path string true "スレッドキー"
// @Param body body request.requestCommentEdit true "コメント編集情報"
// @Success 200 {object} response.responseThreadAndComments
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/comment/{threadKey} [put]
// Comment godoc
func (co *CommentHandler) edit(c *gin.Context) {
	threadKey := c.Param("threadKey")

	var req request.RequestCommentEdit
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("comment edit, requesting json bind error", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.EditCommentAppLayerParam{
		ThreadKey:   threadKey,
		CommentKey:  req.CommentKey,
		Comment:     req.Comment,
		Contributor: req.Contributor,
	}

	comments, err := co.commentApplication.EditComment(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	thread, err := co.threadApplication.GetByThreadKey(threadKey)
	if err != nil {
		handleError(c, err)
		return
	}

	var res response.ResponseThreadAndComments
	res.Thread = response.NewResponseThread(thread)
	for _, comment := range *comments {
		res.Comments = append(res.Comments, response.NewResponseComment(&comment))
	}

	c.JSON(http.StatusOK, res)
}

// Comment godoc
// @Summary 指定されたコメントを削除
// @Description 指定されたコメントを削除する
// @Tags comment
// @Accept json
// @Produce json
// @Param threadKey path string true "スレッドキー"
// @Param body body requestCommentDelete true "コメント削除情報"
// @Success 200 {object} response.responseThreadAndComments
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/comment/{threadKey} [delete]
// Comment godoc
func (co *CommentHandler) delete(c *gin.Context) {
	threadKey := c.Param("threadKey")

	var req request.RequestCommentDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("comment delete, requesting json bind error", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.DeleteCommentAppLayerParam{
		ThreadKey:   threadKey,
		CommentKey:  req.CommentKey,
		Contributor: req.Contributor,
	}

	comments, err := co.commentApplication.DeleteComment(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	thread, err := co.threadApplication.GetByThreadKey(threadKey)
	if err != nil {
		handleError(c, err)
		return
	}

	var res response.ResponseThreadAndComments
	res.Thread = response.NewResponseThread(thread)
	for _, comment := range *comments {
		res.Comments = append(res.Comments, response.NewResponseComment(&comment))
	}

	c.JSON(http.StatusOK, res)
}
