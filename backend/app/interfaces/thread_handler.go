package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/domain"
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

func (t *ThreadHandler) getAll(c *gin.Context) {
	threads, err := t.threadApplication.GetAllThread()

	if err != nil {
		handleError(c)
		return
	}

	var res responseThreads
	for _, thread := range threads {
		res.Threads = append(res.Threads, NewResponseThread(thread))
	}

	c.JSON(http.StatusOK, res)
}

func (t *ThreadHandler) get(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		handleError(c)
		return
	}

	thread, err := t.threadApplication.GetByThreadKey(threadKey)

	if err != nil {
		handleError(c)
		return
	}

	responseThread := NewResponseThread(thread)
	c.JSON(http.StatusOK, responseThread)
}

func (t *ThreadHandler) create(c *gin.Context) {
	var req requestThreadCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c)
		return
	}

	param := params.CreateThreadAppLayerParam{
		Title: req.Title,
		Contributor: req.Contributor,
	}

	thread, err := t.threadApplication.CreateThread(&param)
	if err != nil {
		handleError(c)
		return
	}

	res := NewResponseThread(thread)
	c.JSON(http.StatusOK, res)
}

func (t *ThreadHandler) edit(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		handleError(c)
		return
	}

	var req requestThreadEdit
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c)
		return
	}

	param := params.EditThreadAppLayerParam{
		ThreadKey: threadKey,
		Title: req.Title,
		Contributor: req.Contributor,
	}

	thread, err := t.threadApplication.EditThread(&param)
	if err != nil {
		handleError(c)
		return
	}

	res := NewResponseThread(thread)
	c.JSON(http.StatusOK, res)
}

func (t *ThreadHandler) delete(c *gin.Context) {
	threadKey := c.Param("thread_key")
	if threadKey == "" {
		handleError(c)
		return
	}

	var req requestThreadDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c)
		return
	}

	param := params.DeleteThreadAppLayerParam{
		ThreadKey: threadKey,
		Contributor: req.Contributor,
	}

	if err := t.threadApplication.DeleteThread(&param); err != nil {
		handleError(c)
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
	PostDate    string `json:"post_date"`
	Views       int    `json:"views"`
	SumComment  int    `json:"sum_comment"`
}

func NewResponseThread(thread *domain.Thread) *responseThread {
	return &responseThread{
		ThreadKey:   thread.ThreadKey(),
		Title:       thread.Title(),
		Contributor: thread.Contributor(),
		PostDate:    thread.FormatPostDate(),
		Views:       thread.Views(),
		SumComment:  thread.SumComment(),
	}
}
