package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/params"
)

type CommentHandler struct {
	threadApplication  application.ThreadApplication
	commentApplication application.CommentApplication
}

func NewCommentHandler(ta application.ThreadApplication, ca application.CommentApplication) *CommentHandler {
	return &CommentHandler{
		threadApplication:  ta,
		commentApplication: ca,
	}
}

func (co *CommentHandler) SetupRouter(r *gin.RouterGroup) {
	r.GET("/:thread_key", co.getAll)
	r.POST("/:thread_key", co.create)
	r.PUT("/:thread_key", co.edit)
	r.DELETE("/:thread_key", co.delete)
}

func (co *CommentHandler) getAll(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		handleError(c)
		return
	}

	comments, err := co.commentApplication.GetAllByThreadKey(threadKey)
	if err != nil {
		handleError(c)
		return
	}

	thread, err := co.threadApplication.GetByThreadKey(threadKey)
	if err != nil {
		handleError(c)
		return
	}

	var res responseThreadAndComments
	res.Thread = NewResponseThread(thread)
	for _, comment := range *comments {
		res.Comments = append(res.Comments, NewResponseComment(&comment))
	}

	c.JSON(http.StatusOK, res)
}

func (co *CommentHandler) create(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		handleError(c)
		return
	}

	var req requestCommentCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c)
		return
	}

	param := params.CreateCommentAppLayerParam{
		ThreadKey:   threadKey,
		Comment:     req.Comment,
		Contributor: req.Contributor,
	}

	comments, err := co.commentApplication.CreateComment(&param)
	if err != nil {
		handleError(c)
		return
	}

	thread, err := co.threadApplication.GetByThreadKey(threadKey)
	if err != nil {
		handleError(c)
		return
	}

	var res responseThreadAndComments
	res.Thread = NewResponseThread(thread)
	for _, comment := range *comments {
		res.Comments = append(res.Comments, NewResponseComment(&comment))
	}

	c.JSON(http.StatusOK, res)
}

func (co *CommentHandler) edit(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		handleError(c)
		return
	}

	var req requestCommentEdit
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c)
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
		handleError(c)
		return
	}

	thread, err := co.threadApplication.GetByThreadKey(threadKey)
	if err != nil {
		handleError(c)
		return
	}

	var res responseThreadAndComments
	res.Thread = NewResponseThread(thread)
	for _, comment := range *comments {
		res.Comments = append(res.Comments, NewResponseComment(&comment))
	}

	c.JSON(http.StatusOK, res)
}

func (co *CommentHandler) delete(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		handleError(c)
		return
	}

	var req requestCommentDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c)
		return
	}

	param := params.DeleteCommentAppLayerParam{
		ThreadKey:   threadKey,
		CommentKey:  req.CommentKey,
		Contributor: req.Contributor,
	}

	comments, err := co.commentApplication.DeleteComment(&param)
	if err != nil {
		handleError(c)
		return
	}

	thread, err := co.threadApplication.GetByThreadKey(threadKey)
	if err != nil {
		handleError(c)
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
