package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/jumpei00/board/backend/app/params"
)

type UserHandler struct {
	userApp application.UserApplication
}

func NewUserHandler(ua application.UserApplication) *UserHandler {
	return &UserHandler{
		userApp: ua,
	}
}

func (u *UserHandler) SetupRouter(r *gin.RouterGroup) {
	r.GET("/me", u.me)
	r.POST("/signup", u.signup)
	r.POST("/signin", u.signin)
	r.DELETE("/signout", u.signout)
}

func (u *UserHandler) me(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (u *UserHandler) signup(c *gin.Context) {
	var req requestSignUp
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("requesting json bind error on signup", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.UserSignUpApplicationLayerParam{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := u.userApp.CreateUser(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	res := NewResponseSignUp(user)
	c.JSON(http.StatusOK, res)
}

func (u *UserHandler) signin(c *gin.Context) {
	var req requestSignIn
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("requesting json bind error on signin", "error", err, "binded_request", req)
		handleError(c, err)
		return
	}

	param := params.UserSignInApplicationLayerParam{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := u.userApp.ValidateUser(&param)
	if err != nil {
		handleError(c, err)
		return
	}

	res := NewResponseSignIn(user)
	c.JSON(http.StatusOK, res)
}

func (u *UserHandler) signout(c *gin.Context) {
	// TODO: sessionの開発が完了後、こちらも修正する
	c.Status(http.StatusOK)
}

type requestSignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type requestSignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type responseSignUp struct {
	Username string `json:"username"`
}

func NewResponseSignUp(user *domain.User) *responseSignUp {
	return &responseSignUp{
		Username: user.GetUsername(),
	}
}

type responseSignIn struct {
	Username string `json:"username"`
}

func NewResponseSignIn(user *domain.User) *responseSignIn {
	return &responseSignIn{
		Username: user.GetUsername(),
	}
}
