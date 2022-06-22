package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/interfaces/middleware"
	"github.com/jumpei00/board/backend/app/interfaces/session"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/jumpei00/board/backend/app/params"
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

	r.GET("/:thread_key", co.getAll)
	r.POST("/:thread_key", operatePermissionMiddleware, co.create)
	r.PUT("/:thread_key", operatePermissionMiddleware, co.edit)
	r.DELETE("/:thread_key", operatePermissionMiddleware, co.delete)
}

// Comment godoc
// @Summary コメントを全て取得
// @Description スレッドに紐づくコメントを全て取得
// @Tags comment
// @Accept json
// @Produce json
// @Param thread_key path string true "スレッドキー"
// @Success 200 {object} responseThreadAndComments
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/comment/{thread_key} [get]
// Comment godoc
func (co *CommentHandler) getAll(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		logger.Warning("comment get all, but not thread key")
		handleError(c, appError.NewErrBadRequest(appError.Message().NotThreadKey, "not thread key"))
		return
	}

	comments, err := co.commentApplication.GetAllByThreadKey(threadKey)
	if err != nil {
		handleError(c, err)
		return
	}

	thread, err := co.threadApplication.GetByThreadKey(threadKey)
	if err != nil {
		handleError(c, err)
		return
	}

	var res responseThreadAndComments
	res.Thread = NewResponseThread(thread)
	for _, comment := range *comments {
		res.Comments = append(res.Comments, NewResponseComment(&comment))
	}

	c.JSON(http.StatusOK, res)
}

// Comment godoc
// @Summary スレッドへの新規コメントを作成
// @Description スレッドに対するコメントを作成する
// @Tags comment
// @Accept json
// @Produce json
// @Param thread_key path string true "スレッドキー"
// @Param body body requestCommentCreate true "コメント作成情報"
// @Success 200 {object} responseThreadAndComments
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/comment/{thread_key} [post]
// Comment godoc
func (co *CommentHandler) create(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		logger.Warning("comment create, but not thread key")
		handleError(c, appError.NewErrBadRequest(appError.Message().NotThreadKey, "not thread key"))
		return
	}

	var req requestCommentCreate
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

	var res responseThreadAndComments
	res.Thread = NewResponseThread(thread)
	for _, comment := range *comments {
		res.Comments = append(res.Comments, NewResponseComment(&comment))
	}

	c.JSON(http.StatusOK, res)
}

// Comment godoc
// @Summary 指定されたコメントを更新
// @Description 指定されたコメントを編集し更新する
// @Tags comment
// @Accept json
// @Produce json
// @Param thread_key path string true "スレッドキー"
// @Param body body requestCommentEdit true "コメント編集情報"
// @Success 200 {object} responseThreadAndComments
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/comment/{thread_key} [put]
// Comment godoc
func (co *CommentHandler) edit(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		logger.Warning("comment edit, but not thread key")
		handleError(c, appError.NewErrBadRequest(appError.Message().NotThreadKey, "not thread key"))
		return
	}

	var req requestCommentEdit
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("comment edit, requesting json bind error", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	if req.CommentKey == "" {
		logger.Warning("comment edit, but not comment key")
		handleError(c, appError.NewErrBadRequest(appError.Message().NotCommentKey, "not comment key"))
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

	var res responseThreadAndComments
	res.Thread = NewResponseThread(thread)
	for _, comment := range *comments {
		res.Comments = append(res.Comments, NewResponseComment(&comment))
	}

	c.JSON(http.StatusOK, res)
}

// Comment godoc
// @Summary 指定されたコメントを削除
// @Description 指定されたコメントを削除する
// @Tags comment
// @Accept json
// @Produce json
// @Param thread_key path string true "スレッドキー"
// @Param body body requestCommentDelete true "コメント削除情報"
// @Success 200 {object} responseThreadAndComments
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/comment/{thread_key} [delete]
// Comment godoc
func (co *CommentHandler) delete(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		logger.Warning("comment delete, but not thread key")
		handleError(c, appError.NewErrBadRequest(appError.Message().NotThreadKey, "not thread key"))
		return
	}

	var req requestCommentDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("comment delete, requesting json bind error", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	if req.CommentKey == "" {
		logger.Warning("comment delete, but not comment key")
		handleError(c, appError.NewErrBadRequest(appError.Message().NotCommentKey, "not comment key"))
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

	var res responseThreadAndComments
	res.Thread = NewResponseThread(thread)
	for _, comment := range *comments {
		res.Comments = append(res.Comments, NewResponseComment(&comment))
	}

	c.JSON(http.StatusOK, res)
}

type requestCommentCreate struct {
	Comment     string `json:"comment"`
	Contributor string `json:"contributor"`
}

type requestCommentEdit struct {
	CommentKey  string `json:"comment_key"`
	Comment     string `json:"comment"`
	Contributor string `json:"contributor"`
}

type requestCommentDelete struct {
	CommentKey  string `json:"comment_key"`
	Contributor string `json:"contributor"`
}

type responseThreadAndComments struct {
	Thread   *responseThread    `json:"thread"`
	Comments []*responseComment `json:"comments"`
}

type responseComment struct {
	CommentKey  string `joson:"comment_key"`
	Contributor string `json:"contributor"`
	Comment     string `json:"comment"`
	UpdateDate  string `json:"update_date"`
}

func NewResponseComment(comment *domain.Comment) *responseComment {
	return &responseComment{
		CommentKey:  comment.GetKey(),
		Contributor: comment.GetContributor(),
		Comment:     comment.GetComment(),
		UpdateDate:  comment.FormatUpdateDate(),
	}
}
