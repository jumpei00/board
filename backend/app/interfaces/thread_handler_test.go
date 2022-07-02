package interfaces_test

import (
	"bytes"
	"encoding/json"
	"errors"
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
	"github.com/jumpei00/board/backend/app/params"
)

func TestThreadHandler_getAll(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	r := gin.Default()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)

	var nothing bool
	threadApplication.EXPECT().GetAllThread().AnyTimes().DoAndReturn(
		func() (*[]domain.Thread, error) {
			if nothing {
				return nil, appError.ErrNotFound
			}
			return &[]domain.Thread{}, nil
		},
	)

	threadHandler := interfaces.NewThreadHandler(sessionManager, threadApplication)
	threadHandler.SetupRouter(r.Group("/api/thread"))

	//
	// exucute
	//
	cases := []struct {
		name       string
		input      bool
		statusCode int
	}{
		{
			name:       "空のスレッドの時のテスト",
			input:      true,
			statusCode: http.StatusNotFound,
		},
		{
			name:       "空でない時のテスト",
			input:      false,
			statusCode: http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			nothing = c.input
			response := executeHttpTest(r, "GET", "/api/thread", nil)

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestThreadHandler_get(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// mock
	var (
		wrongThreadKey   = "wrong-key"
		correctThreadKey = "correct-key"
		initView         = 0
		commentSum       = 0
	)
	r := gin.Default()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)

	var threadKey string
	threadApplication.EXPECT().GetByThreadKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*domain.Thread, error) {
			if threadKey == wrongThreadKey {
				return nil, appError.ErrNotFound
			}
			return &domain.Thread{Views: &initView, CommentSum: &commentSum}, nil
		},
	)

	threadHandler := interfaces.NewThreadHandler(sessionManager, threadApplication)
	threadHandler.SetupRouter(r.Group("/api/thread"))

	//
	// execute
	//
	cases := []struct {
		name       string
		threadKey  string
		statucCode int
	}{
		{
			name:       "スレッドキーが空の場合はstatus codeが301になる",
			threadKey:  "",
			statucCode: http.StatusMovedPermanently,
		},
		{
			name:       "スレッドキーに対するスレッドが存在しない場合はstatus codeが404になる",
			threadKey:  wrongThreadKey,
			statucCode: http.StatusNotFound,
		},
		{
			name:       "スレッドキーに対するスレッドが存在する場合はstatus codeが200となる",
			threadKey:  correctThreadKey,
			statucCode: http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			path := "/api/thread/" + c.threadKey
			response := executeHttpTest(r, "GET", path, nil)

			if response.Code != c.statucCode {
				t.Errorf("different status code.\npath: %s\nwant: %d\ngot: %d", path, c.statucCode, response.Code)
			}
		})
	}
}

func TestThreaHandler_create(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		notFound   bool
		notSession bool
		errSession bool
	)

	r := gin.Default()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)

	threadApplication.EXPECT().CreateThread(gomock.AssignableToTypeOf(&params.CreateThreadAppLayerParam{})).AnyTimes().DoAndReturn(
		func(params *params.CreateThreadAppLayerParam) (*domain.Thread, error) {
			if notFound {
				return nil, appError.ErrNotFound
			}
			return &domain.Thread{}, nil
		},
	)

	sessionManager.EXPECT().Get(gomock.AssignableToTypeOf(&gin.Context{})).AnyTimes().DoAndReturn(
		func(c *gin.Context) (*session.Session, error) {
			if notSession {
				return nil, appError.ErrNotFound
			}
			if errSession {
				return nil, errors.New("internal server error")
			}
			return &session.Session{}, nil
		},
	)

	threadHandler := interfaces.NewThreadHandler(sessionManager, threadApplication)
	threadHandler.SetupRouter(r.Group("/api/thread"))

	//
	// execute
	//
	cases := []struct {
		name       string
		notFound   bool
		notSession bool
		errSession bool
		request    request.RequestThreadCreate
		statusCode int
	}{
		{
			name:       "タイトルが欠如してリクエストされた場合は500になる",
			notFound:   false,
			notSession: false,
			errSession: false,
			request:    request.RequestThreadCreate{Contributor: "test-user"},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "投稿者が欠如してリクエストされた場合は500になる",
			notFound:   false,
			notSession: false,
			errSession: false,
			request:    request.RequestThreadCreate{Title: "test-title"},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "application層でNotFoundが発生した場合は404となる",
			notFound:   true,
			notSession: false,
			errSession: false,
			request:    request.RequestThreadCreate{Title: "test-title", Contributor: "test-user"},
			statusCode: http.StatusNotFound,
		},
		{
			name:       "sessionがない場合は401となる",
			notFound:   false,
			notSession: true,
			errSession: false,
			request:    request.RequestThreadCreate{Title: "test-title", Contributor: "test-user"},
			statusCode: http.StatusUnauthorized,
		},
		{
			name:       "session層でエラーが発生した場合は500となる",
			notFound:   false,
			notSession: false,
			errSession: true,
			request:    request.RequestThreadCreate{Title: "test-title", Contributor: "test-user"},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			notFound = c.notFound
			notSession = c.notSession
			errSession = c.errSession
			j, _ := json.Marshal(&c.request)

			response := executeHttpTest(r, "POST", "/api/thread", bytes.NewBuffer(j))

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestThreadHandler_edit(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		correctThreadKey = "correct-thread-key"
		wrongThreadKey   = "wrong-thread-key"
		notSession       bool
		errSession       bool
	)

	r := gin.Default()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)

	threadApplication.EXPECT().EditThread(gomock.AssignableToTypeOf(&params.EditThreadAppLayerParam{})).AnyTimes().DoAndReturn(
		func(params *params.EditThreadAppLayerParam) (*domain.Thread, error) {
			if params.ThreadKey == wrongThreadKey {
				return nil, appError.ErrNotFound
			}
			return &domain.Thread{}, nil
		},
	)

	sessionManager.EXPECT().Get(gomock.AssignableToTypeOf(&gin.Context{})).AnyTimes().DoAndReturn(
		func(c *gin.Context) (*session.Session, error) {
			if notSession {
				return nil, appError.ErrNotFound
			}
			if errSession {
				return nil, errors.New("internal server error")
			}
			return &session.Session{}, nil
		},
	)

	threadHandler := interfaces.NewThreadHandler(sessionManager, threadApplication)
	threadHandler.SetupRouter(r.Group("/api/thread"))

	//
	// execute
	//
	cases := []struct {
		name       string
		threadKey  string
		notSession bool
		errSession bool
		request    request.RequestThreadEdit
		statusCode int
	}{
		{
			name:       "タイトルが欠如してリクエストされた場合は500になる",
			threadKey:  correctThreadKey,
			notSession: false,
			errSession: false,
			request:    request.RequestThreadEdit{Contributor: "test-user"},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "投稿者が欠如してリクエストされた場合は500になる",
			threadKey:  correctThreadKey,
			notSession: false,
			errSession: false,
			request:    request.RequestThreadEdit{Title: "test-title"},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "スレッドキーに対するスレッドが存在しない場合は404となる",
			threadKey:  wrongThreadKey,
			notSession: false,
			errSession: false,
			request:    request.RequestThreadEdit{Title: "test-title", Contributor: "test-user"},
			statusCode: http.StatusNotFound,
		},
		{
			name:       "sessionがない場合は401となる",
			threadKey:  correctThreadKey,
			notSession: true,
			errSession: false,
			request:    request.RequestThreadEdit{Title: "test-title", Contributor: "test-user"},
			statusCode: http.StatusUnauthorized,
		},
		{
			name:       "session層でエラーが発生した場合は500となる",
			threadKey:  correctThreadKey,
			notSession: false,
			errSession: true,
			request:    request.RequestThreadEdit{Title: "test-title", Contributor: "test-user"},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			path := "/api/thread/" + c.threadKey
			notSession = c.notSession
			errSession = c.errSession
			j, _ := json.Marshal(&c.request)

			response := executeHttpTest(r, "PUT", path, bytes.NewBuffer(j))

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestThreadHandler_delete(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		correctThreadKey = "correct-thread-key"
		wrongThreadKey   = "wrong-thread-key"
		notSession       bool
		errSession       bool
	)

	r := gin.Default()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)

	threadApplication.EXPECT().DeleteThread(gomock.AssignableToTypeOf(&params.DeleteThreadAppLayerParam{})).AnyTimes().DoAndReturn(
		func(params *params.DeleteThreadAppLayerParam) error {
			if params.ThreadKey == wrongThreadKey {
				return appError.ErrNotFound
			}
			return nil
		},
	)

	sessionManager.EXPECT().Get(gomock.AssignableToTypeOf(&gin.Context{})).AnyTimes().DoAndReturn(
		func(c *gin.Context) (*session.Session, error) {
			if notSession {
				return nil, appError.ErrNotFound
			}
			if errSession {
				return nil, errors.New("internal server error")
			}
			return &session.Session{}, nil
		},
	)

	threadHandler := interfaces.NewThreadHandler(sessionManager, threadApplication)
	threadHandler.SetupRouter(r.Group("/api/thread"))

	//
	// execute
	//
	cases := []struct {
		name       string
		threadKey  string
		notSession bool
		errSession bool
		request    request.RequestThreadDelete
		statusCode int
	}{
		{
			name:       "投稿者が欠如してリクエストされた場合は500になる",
			threadKey:  correctThreadKey,
			notSession: false,
			errSession: false,
			request:    request.RequestThreadDelete{},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "スレッドキーに対するスレッドが存在しない場合は404となる",
			threadKey:  wrongThreadKey,
			notSession: false,
			errSession: false,
			request:    request.RequestThreadDelete{Contributor: "test-user"},
			statusCode: http.StatusNotFound,
		},
		{
			name:       "sessionがない場合は401となる",
			threadKey:  correctThreadKey,
			notSession: true,
			errSession: false,
			request:    request.RequestThreadDelete{Contributor: "test-user"},
			statusCode: http.StatusUnauthorized,
		},
		{
			name:       "session層でエラーが発生した場合は500となる",
			threadKey:  correctThreadKey,
			notSession: false,
			errSession: true,
			request:    request.RequestThreadDelete{Contributor: "test-user"},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			path := "/api/thread/" + c.threadKey
			notSession = c.notSession
			errSession = c.errSession
			j, _ := json.Marshal(&c.request)

			response := executeHttpTest(r, "DELETE", path, bytes.NewBuffer(j))

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}
