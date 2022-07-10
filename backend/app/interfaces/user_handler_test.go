package interfaces_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/interfaces"
	"github.com/jumpei00/board/backend/app/interfaces/request"
	"github.com/jumpei00/board/backend/app/interfaces/session"
	appError "github.com/jumpei00/board/backend/app/library/error"
	mock_application "github.com/jumpei00/board/backend/app/mock/application"
	mock_session "github.com/jumpei00/board/backend/app/mock/session"
	"github.com/pkg/errors"
)

func TestUserHandler_me(t *testing.T) {
	//
	// setup
	//
	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	userApplication := mock_application.NewMockUserApplication(mockCtrl)

	type mockField struct {
		sessionManager  *mock_session.MockManager
		userApplication *mock_application.MockUserApplication
	}

	field := mockField{
		sessionManager:  sessionManager,
		userApplication: userApplication,
	}

	userHandler := interfaces.NewUserHandler(sessionManager, userApplication)
	userHandler.SetupRouter(r.Group("/api/user"))

	//
	// execute
	//
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		statusCode int
	}{
		{
			testCase: "セッションが確認できない場合は404となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, appError.ErrNotFound)
			},
			statusCode: http.StatusNotFound,
		},
		{
			testCase: "セッションのユーザーIDに対するユーザーが存在しない時は404となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(&session.Session{}, nil)
				mf.userApplication.EXPECT().GetUserByID(gomock.Any()).Return(nil, appError.ErrNotFound)
			},
			statusCode: http.StatusNotFound,
		},
		{
			testCase: "セッションのユーザーIDに対するユーザーが存在する時は200となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(&session.Session{}, nil)
				mf.userApplication.EXPECT().GetUserByID(gomock.Any()).Return(&domain.User{}, nil)
			},
			statusCode: http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)

			response := executeHttpTest(r, http.MethodGet, "/api/user/me", nil)

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestUserHandler_signup(t *testing.T) {
	//
	// setup
	//
	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	userApplication := mock_application.NewMockUserApplication(mockCtrl)

	type mockField struct {
		sessionManager  *mock_session.MockManager
		userApplication *mock_application.MockUserApplication
	}

	field := mockField{
		sessionManager:  sessionManager,
		userApplication: userApplication,
	}

	userHandler := interfaces.NewUserHandler(sessionManager, userApplication)
	userHandler.SetupRouter(r.Group("/api/user"))

	//
	// execute
	//
	var (
		username = "username"
		password = "password"
	)
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		body       request.RequestSignUp
		statusCode int
	}{
		{
			testCase: "ユーザー登録が成功したら200となる",
			prepare: func(mf *mockField) {
				mf.userApplication.EXPECT().CreateUser(gomock.Any()).Return(&domain.User{}, nil)
				mf.sessionManager.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			body:       request.RequestSignUp{Username: username, Password: password},
			statusCode: http.StatusOK,
		},
		{
			testCase:   "ユーザー名が空文字の場合は500となる",
			prepare:    func(mf *mockField) {},
			body:       request.RequestSignUp{Username: "", Password: password},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase:   "パスワードが空文字の場合は500となる",
			prepare:    func(mf *mockField) {},
			body:       request.RequestSignUp{Username: username, Password: ""},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase: "ユーザーの登録時にエラーとなった場合は500となる",
			prepare: func(mf *mockField) {
				mf.userApplication.EXPECT().CreateUser(gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			body:       request.RequestSignUp{Username: username, Password: password},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase: "セッション登録時にエラーとなった場合は500となる",
			prepare: func(mf *mockField) {
				mf.userApplication.EXPECT().CreateUser(gomock.Any()).Return(&domain.User{}, nil)
				mf.sessionManager.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			body:       request.RequestSignUp{Username: username, Password: password},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)

			j, _ := json.Marshal(&c.body)
			response := executeHttpTest(r, http.MethodPost, "/api/user/signup", bytes.NewBuffer(j))

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestUserHandler_signin(t *testing.T) {
	//
	// setup
	//
	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	userApplication := mock_application.NewMockUserApplication(mockCtrl)

	type mockField struct {
		sessionManager  *mock_session.MockManager
		userApplication *mock_application.MockUserApplication
	}

	field := mockField{
		sessionManager:  sessionManager,
		userApplication: userApplication,
	}

	userHandler := interfaces.NewUserHandler(sessionManager, userApplication)
	userHandler.SetupRouter(r.Group("/api/user"))

	//
	// execute
	//
	var (
		username = "username"
		password = "password"
	)
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		body       request.RequestSignIn
		statusCode int
	}{
		{
			testCase: "ログインが成功したら200となる",
			prepare: func(mf *mockField) {
				mf.userApplication.EXPECT().ValidateUser(gomock.Any()).Return(&domain.User{}, nil)
				mf.sessionManager.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			body:       request.RequestSignIn{Username: username, Password: password},
			statusCode: http.StatusOK,
		},
		{
			testCase:   "ユーザー名が空文字の場合は500となる",
			prepare:    func(mf *mockField) {},
			body:       request.RequestSignIn{Username: "", Password: password},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase:   "パスワードが空文字の場合は500となる",
			prepare:    func(mf *mockField) {},
			body:       request.RequestSignIn{Username: username, Password: ""},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase: "ログインに失敗した場合は400となる",
			prepare: func(mf *mockField) {
				mf.userApplication.EXPECT().ValidateUser(gomock.Any()).Return(nil, &appError.BadRequest{})
			},
			body:       request.RequestSignIn{Username: username, Password: password},
			statusCode: http.StatusBadRequest,
		},
		{
			testCase: "セッション登録時にエラーとなった場合は500となる",
			prepare: func(mf *mockField) {
				mf.userApplication.EXPECT().ValidateUser(gomock.Any()).Return(&domain.User{}, nil)
				mf.sessionManager.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			body:       request.RequestSignIn{Username: username, Password: password},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)

			j, _ := json.Marshal(&c.body)
			response := executeHttpTest(r, http.MethodPost, "/api/user/signin", bytes.NewBuffer(j))

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestUserHandler_signout(t *testing.T) {
	//
	// setup
	//
	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	userApplication := mock_application.NewMockUserApplication(mockCtrl)

	type mockField struct {
		sessionManager  *mock_session.MockManager
		userApplication *mock_application.MockUserApplication
	}

	field := mockField{
		sessionManager:  sessionManager,
		userApplication: userApplication,
	}

	userHandler := interfaces.NewUserHandler(sessionManager, userApplication)
	userHandler.SetupRouter(r.Group("/api/user"))

	//
	// execute
	//
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		statusCode int
	}{
		{
			testCase: "セッションの削除に失敗した場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Delete(gomock.Any()).Return(errors.New("Internal Server Error"))
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase: "セッションの削除に成功した場合は204となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Delete(gomock.Any()).Return(nil)
			},
			statusCode: http.StatusNoContent,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)

			response := executeHttpTest(r, http.MethodDelete, "/api/user/signout", nil)

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}
