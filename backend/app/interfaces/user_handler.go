package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/application/params"
	"github.com/jumpei00/board/backend/app/interfaces/request"
	"github.com/jumpei00/board/backend/app/interfaces/session"
	"github.com/jumpei00/board/backend/app/library/logger"
)

type UserHandler struct {
	sessionManager  session.Manager
	userApplication application.UserApplication
}

func NewUserHandler(sm session.Manager, ua application.UserApplication) *UserHandler {
	return &UserHandler{
		sessionManager:  sm,
		userApplication: ua,
	}
}

func (u *UserHandler) SetupRouter(r *gin.RouterGroup) {
	r.GET("/helthcheck", u.helthcheck)
	r.GET("/me", u.me)
	r.POST("/signup", u.signup)
	r.POST("/signin", u.signin)
	r.DELETE("/signout", u.signout)
}

func (u *UserHandler) helthcheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

// User godoc
// @Summary ユーザー情報の取得
// @Description セッション情報からユーザーを取得する
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} domain.User
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/user/me [get]
// User godoc
func (u *UserHandler) me(c *gin.Context) {
	session, err := u.sessionManager.Get(c)
	if err != nil {
		handleError(c, err)
		return
	}

	user, err := u.userApplication.GetUserByID(session.UserID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// User godoc
// @Summary ユーザーの新規作成
// @Description 新規ユーザーを作成する
// @Tags user
// @Accept json
// @Produce json
// @Param body body request.RequestSignUp true "新規ユーザー作成情報"
// @Success 200 {object} domain.User
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/user/signup [post]
// User godoc
func (u *UserHandler) signup(c *gin.Context) {
	var req request.RequestSignUp
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("requesting json bind error on signup", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.UserSignUpApplicationLayerParam{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := u.userApplication.CreateUser(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	if _, err := u.sessionManager.Create(c, user); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// User godoc
// @Summary ログイン
// @Description ユーザーがログインできるか検証する
// @Tags user
// @Accept json
// @Produce json
// @Param body body request.RequestSignIn true "ログイン情報"
// @Success 200 {object} domain.User
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/user/signin [post]
// User godoc
func (u *UserHandler) signin(c *gin.Context) {
	var req request.RequestSignIn
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("requesting json bind error on signin", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.UserSignInApplicationLayerParam{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := u.userApplication.ValidateUser(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	if _, err := u.sessionManager.Create(c, user); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// User godoc
// @Summary ログアウト
// @Description ユーザーをログアウトさせる
// @Tags user
// @Accept json
// @Produce json
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/user/signout [delete]
// User godoc
func (u *UserHandler) signout(c *gin.Context) {
	if err := u.sessionManager.Delete(c); err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
