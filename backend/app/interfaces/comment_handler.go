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

	r.GET("/:threadKey/comments", co.getAll)
	r.POST("/:threadKey/comments", operatePermissionMiddleware, co.create)
	r.PUT("/:threadKey/comments/:commentKey", operatePermissionMiddleware, co.edit)
	r.DELETE("/:threadKey/comments/:commentKey", operatePermissionMiddleware, co.delete)
}

// Comment godoc
// @Summary コメントを全て取得
// @Description スレッドに紐づくコメントを全て取得
// @Tags comment
// @Accept json
// @Produce json
// @Param threadKey path string true "スレッドキー"
// @Success 200 {object} response.ResponseThreadAndComments
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/threads/{threadKey}/comments [get]
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
// @Param body body request.RequestCommentCreate true "コメント作成情報"
// @Success 200 {object} response.ResponseComment
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/threads/{threadKey}/comments [post]
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

	comment, err := co.commentApplication.CreateComment(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	res := response.NewResponseComment(comment)

	c.JSON(http.StatusOK, res)
}

// Comment godoc
// @Summary 指定されたコメントを更新
// @Description 指定されたコメントを編集し更新する
// @Tags comment
// @Accept json
// @Produce json
// @Param threadKey path string true "スレッドキー"
// @Param commentKey path string true "コメントキー"
// @Param body body request.RequestCommentEdit true "コメント編集情報"
// @Success 200 {object} response.ResponseComment
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/threads/{threadKey}/comments/{commentKey} [put]
// Comment godoc
func (co *CommentHandler) edit(c *gin.Context) {
	threadKey := c.Param("threadKey")
	commentKey := c.Param("commentKey")

	var req request.RequestCommentEdit
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("comment edit, requesting json bind error", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.EditCommentAppLayerParam{
		ThreadKey:   threadKey,
		CommentKey:  commentKey,
		Comment:     req.Comment,
		Contributor: req.Contributor,
	}

	comment, err := co.commentApplication.EditComment(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	res := response.NewResponseComment(comment)

	c.JSON(http.StatusOK, res)
}

// Comment godoc
// @Summary 指定されたコメントを削除
// @Description 指定されたコメントを削除する
// @Tags comment
// @Accept json
// @Produce json
// @Param threadKey path string true "スレッドキー"
// @Param commentKey path string true "コメントキー"
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/threads/{threadKey}/comments/{commentKey} [delete]
// Comment godoc
func (co *CommentHandler) delete(c *gin.Context) {
	threadKey := c.Param("threadKey")
	commentKey := c.Param("commentKey")
	user, _ := co.sessionManager.Get(c)

	param := params.DeleteCommentAppLayerParam{
		ThreadKey:  threadKey,
		CommentKey: commentKey,
		UserID:     user.UserID,
	}

	if err := co.commentApplication.DeleteComment(&param); err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
