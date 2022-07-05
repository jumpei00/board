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
	"github.com/jumpei00/board/backend/app/params"
	"github.com/pkg/errors"
)

func TestCommentHandler_getAll(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	r := gin.Default()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)
	commentApplication := mock_application.NewMockCommentApplication(mockCtrl)

	var (
		correctThreadkey = "correct-thread-key"
		wrongThreadKey1  = "wrong-thread-key1"
		wrongThreadKey2  = "wrong-thread-key2"
		wrongThreadKey3  = "wrong-thread-key3"
		initViews        = 0
		commentSum       = 0
	)

	var threadKey string

	threadApplication.EXPECT().GetByThreadKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*domain.Thread, error) {
			if threadKey == wrongThreadKey1 {
				return nil, appError.ErrNotFound
			}
			return &domain.Thread{Views: &initViews, CommentSum: &commentSum}, nil
		},
	)

	commentApplication.EXPECT().GetAllByThreadKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*[]domain.Comment, error) {
			switch threadKey {
			case wrongThreadKey2:
				return nil, appError.ErrNotFound
			case wrongThreadKey3:
				return nil, errors.New("Internal Server Error")
			default:
				return &[]domain.Comment{}, nil
			}
		},
	)

	commentHandler := interfaces.NewCommentHandler(sessionManager, threadApplication, commentApplication)
	commentHandler.SetupRouter(r.Group("/api/comment"))

	//
	// execute
	//
	cases := []struct {
		name       string
		threadKey  string
		statusCode int
	}{
		{
			name:       "スレッドもしくはコメントが存在するようなスレッドキーの場合は200となる",
			threadKey:  correctThreadkey,
			statusCode: http.StatusOK,
		},
		{
			name:       "スレッドが存在しないようなスレッドキーの場合は404となる",
			threadKey:  wrongThreadKey1,
			statusCode: http.StatusNotFound,
		},
		{
			name:       "コメントが存在しないようなスレッドキーの場合は200となる",
			threadKey:  wrongThreadKey2,
			statusCode: http.StatusOK,
		},
		{
			name:       "コメントを取得時にNotFound以外のエラーを発生させるとそれに応じたエラーとなる",
			threadKey:  wrongThreadKey3,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			path := "/api/comment/" + c.threadKey
			response := executeHttpTest(r, "GET", path, nil)

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestCommenHandler_create(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	r := gin.Default()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)
	commentApplication := mock_application.NewMockCommentApplication(mockCtrl)

	var (
		correctThreadKey = "correct-thread-key"
		wrongThreadKey1  = "wrong-thread-key1"
		wrongThreadKey2  = "wrong-thread-key2"
		notSession       bool
		errSession       bool
		initViews        = 0
		commentSum       = 0
	)

	var threadKey string

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

	threadApplication.EXPECT().GetByThreadKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*domain.Thread, error) {
			if threadKey == wrongThreadKey2 {
				return nil, appError.ErrNotFound
			}
			return &domain.Thread{Views: &initViews, CommentSum: &commentSum}, nil
		},
	)

	commentApplication.EXPECT().CreateComment(gomock.AssignableToTypeOf(&params.CreateCommentAppLayerParam{})).AnyTimes().DoAndReturn(
		func(params *params.CreateCommentAppLayerParam) (*[]domain.Comment, error) {
			if params.ThreadKey == wrongThreadKey1 {
				return nil, errors.New("Internal Server Error")
			}
			return &[]domain.Comment{}, nil
		},
	)

	commentHandler := interfaces.NewCommentHandler(sessionManager, threadApplication, commentApplication)
	commentHandler.SetupRouter(r.Group("/api/comment"))

	//
	// execute
	//
	cases := []struct {
		name       string
		notSession bool
		errSession bool
		threadKey  string
		request    request.RequestCommentCreate
		statusCode int
	}{
		{
			name:       "コメントが欠けている場合は500となる",
			notSession: false,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentCreate{Contributor: "user"},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "投稿者が欠けている場合は500となる",
			notSession: false,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentCreate{Comment: "comment"},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "セッションが無い場合は401となる",
			notSession: true,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentCreate{Comment: "comment", Contributor: "user"},
			statusCode: http.StatusUnauthorized,
		},
		{
			name:       "セッションがエラーの場合は500になる",
			notSession: false,
			errSession: true,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentCreate{Comment: "comment", Contributor: "user"},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "コメント生成もスレッド取得も問題ない場合200となる",
			notSession: false,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentCreate{Comment: "comment", Contributor: "user"},
			statusCode: http.StatusOK,
		},
		{
			name:       "コメント作成に失敗した場合は500となる",
			notSession: false,
			errSession: false,
			threadKey:  wrongThreadKey1,
			request:    request.RequestCommentCreate{Comment: "comment", Contributor: "user"},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "コメント作成後のスレッド取得に失敗した場合は404となる",
			notSession: false,
			errSession: false,
			threadKey:  wrongThreadKey2,
			request:    request.RequestCommentCreate{Comment: "comment", Contributor: "user"},
			statusCode: http.StatusNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			notSession = c.notSession
			errSession = c.errSession
			path := "/api/comment/" + c.threadKey
			j, _ := json.Marshal(&c.request)

			response := executeHttpTest(r, "POST", path, bytes.NewBuffer(j))

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestCommentService_edit(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	r := gin.Default()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)
	commentApplication := mock_application.NewMockCommentApplication(mockCtrl)

	var (
		correctThreadKey = "correct-thread-key"
		wrongThreadKey1  = "wrong-thread-key1"
		wrongThreadKey2  = "wrong-thread-key2"
		notSession       bool
		errSession       bool
		initViews        = 0
		commentSum       = 0
	)

	var threadKey string

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

	threadApplication.EXPECT().GetByThreadKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*domain.Thread, error) {
			if threadKey == wrongThreadKey2 {
				return nil, appError.ErrNotFound
			}
			return &domain.Thread{Views: &initViews, CommentSum: &commentSum}, nil
		},
	)

	commentApplication.EXPECT().EditComment(gomock.AssignableToTypeOf(&params.EditCommentAppLayerParam{})).AnyTimes().DoAndReturn(
		func(params *params.EditCommentAppLayerParam) (*[]domain.Comment, error) {
			if params.ThreadKey == wrongThreadKey1 {
				return nil, errors.New("Internal Server Error")
			}
			return &[]domain.Comment{}, nil
		},
	)

	commentHandler := interfaces.NewCommentHandler(sessionManager, threadApplication, commentApplication)
	commentHandler.SetupRouter(r.Group("/api/comment"))

	//
	// execute
	//
	cases := []struct {
		name       string
		notSession bool
		errSession bool
		threadKey  string
		request    request.RequestCommentEdit
		StatusCode int
	}{
		{
			name:       "コメントキーが欠けている場合は500となる",
			notSession: false,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentEdit{Comment: "comment", Contributor: "user"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			name:       "コメントが欠けている場合は500となる",
			notSession: false,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentEdit{CommentKey: "key", Contributor: "user"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			name:       "投稿者が欠けている場合は500となる",
			notSession: false,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentEdit{CommentKey: "key", Comment: "comment"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			name:       "セッションが無い場合は401となる",
			notSession: true,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentEdit{CommentKey: "key", Comment: "comment", Contributor: "user"},
			StatusCode: http.StatusUnauthorized,
		},
		{
			name:       "セッションがエラーの場合は500となる",
			notSession: false,
			errSession: true,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentEdit{CommentKey: "key", Comment: "comment", Contributor: "user"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			name:       "コメントの編集に成功した場合は200となる",
			notSession: false,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentEdit{CommentKey: "key", Comment: "comment", Contributor: "user"},
			StatusCode: http.StatusOK,
		},
		{
			name:       "コメントの編集に失敗した場合は500となる",
			notSession: false,
			errSession: false,
			threadKey:  wrongThreadKey1,
			request:    request.RequestCommentEdit{CommentKey: "key", Comment: "comment", Contributor: "user"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			name:       "コメント編集後のスレッド取得に失敗した場合は404となる",
			notSession: false,
			errSession: false,
			threadKey:  wrongThreadKey2,
			request:    request.RequestCommentEdit{CommentKey: "key", Comment: "comment", Contributor: "user"},
			StatusCode: http.StatusNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			notSession = c.notSession
			errSession = c.errSession
			path := "/api/comment/" + c.threadKey
			j, _ := json.Marshal(&c.request)

			response := executeHttpTest(r, "PUT", path, bytes.NewBuffer(j))

			if response.Code != c.StatusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.StatusCode, response.Code)
			}
		})
	}
}

func TestCommentService_delete(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	r := gin.Default()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)
	commentApplication := mock_application.NewMockCommentApplication(mockCtrl)

	var (
		correctThreadKey = "correct-thread-key"
		wrongThreadKey1  = "wrong-thread-key1"
		wrongThreadKey2  = "wrong-thread-key2"
		notSession       bool
		errSession       bool
		initViews        = 0
		commentSum       = 0
	)

	var threadKey string

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

	threadApplication.EXPECT().GetByThreadKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*domain.Thread, error) {
			if threadKey == wrongThreadKey2 {
				return nil, appError.ErrNotFound
			}
			return &domain.Thread{Views: &initViews, CommentSum: &commentSum}, nil
		},
	)

	commentApplication.EXPECT().DeleteComment(gomock.AssignableToTypeOf(&params.DeleteCommentAppLayerParam{})).AnyTimes().DoAndReturn(
		func(params *params.DeleteCommentAppLayerParam) (*[]domain.Comment, error) {
			if params.ThreadKey == wrongThreadKey1 {
				return nil, errors.New("Internal Server Error")
			}
			return &[]domain.Comment{}, nil
		},
	)

	commentHandler := interfaces.NewCommentHandler(sessionManager, threadApplication, commentApplication)
	commentHandler.SetupRouter(r.Group("/api/comment"))

	//
	// execute
	//
	cases := []struct {
		name       string
		notSession bool
		errSession bool
		threadKey  string
		request    request.RequestCommentDelete
		StatusCode int
	}{
		{
			name:       "コメントキーが欠けている場合は500となる",
			notSession: false,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentDelete{Contributor: "user"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			name:       "投稿者が欠けている場合は500となる",
			notSession: false,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentDelete{CommentKey: "key"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			name:       "セッションが無い場合は401となる",
			notSession: true,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentDelete{CommentKey: "key", Contributor: "user"},
			StatusCode: http.StatusUnauthorized,
		},
		{
			name:       "セッションがエラーの場合は500となる",
			notSession: false,
			errSession: true,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentDelete{CommentKey: "key", Contributor: "user"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			name:       "コメントの削除に成功した場合は200となる",
			notSession: false,
			errSession: false,
			threadKey:  correctThreadKey,
			request:    request.RequestCommentDelete{CommentKey: "key", Contributor: "user"},
			StatusCode: http.StatusOK,
		},
		{
			name:       "コメントの削除に失敗した場合は500となる",
			notSession: false,
			errSession: false,
			threadKey:  wrongThreadKey1,
			request:    request.RequestCommentDelete{CommentKey: "key", Contributor: "user"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			name:       "コメント編集後のスレッド取得に失敗した場合は404となる",
			notSession: false,
			errSession: false,
			threadKey:  wrongThreadKey2,
			request:    request.RequestCommentDelete{CommentKey: "key", Contributor: "user"},
			StatusCode: http.StatusNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			notSession = c.notSession
			errSession = c.errSession
			path := "/api/comment/" + c.threadKey
			j, _ := json.Marshal(&c.request)

			response := executeHttpTest(r, "DELETE", path, bytes.NewBuffer(j))

			if response.Code != c.StatusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.StatusCode, response.Code)
			}
		})
	}
}
