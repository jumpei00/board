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

func TestCommentHandler_getAll(t *testing.T) {
	//
	// setup
	//
	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)
	commentApplication := mock_application.NewMockCommentApplication(mockCtrl)

	type mockField struct {
		sessionManager     *mock_session.MockManager
		threadApplication  *mock_application.MockThreadApplication
		commentApplication *mock_application.MockCommentApplication
	}

	field := mockField{
		sessionManager:     sessionManager,
		threadApplication:  threadApplication,
		commentApplication: commentApplication,
	}

	commentHandler := interfaces.NewCommentHandler(sessionManager, threadApplication, commentApplication)
	commentHandler.SetupRouter(r.Group("/api/threads"))

	//
	// execute
	//
	var (
		correctThreadkey = "correct-thread-key"
		wrongThreadKey1  = "wrong-thread-key1"
		wrongThreadKey2  = "wrong-thread-key2"
		wrongThreadKey3  = "wrong-thread-key3"
		initViews        = 0
		commentSum       = 0
		thread           = &domain.Thread{Views: &initViews, CommentSum: &commentSum}
	)
	cases := []struct {
		testCase   string
		threadKey  string
		prepare    func(*mockField)
		statusCode int
	}{
		{
			testCase:  "スレッドもしくはコメントが存在するようなスレッドキーの場合は200となる",
			threadKey: correctThreadkey,
			prepare: func(mf *mockField) {
				mf.threadApplication.EXPECT().GetByThreadKey(correctThreadkey).Return(thread, nil)
				mf.commentApplication.EXPECT().GetAllByThreadKey(correctThreadkey).Return(&[]domain.Comment{}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			testCase:  "スレッドが存在しないようなスレッドキーの場合は404となる",
			threadKey: wrongThreadKey1,
			prepare: func(mf *mockField) {
				mf.threadApplication.EXPECT().GetByThreadKey(wrongThreadKey1).Return(nil, appError.ErrNotFound)
			},
			statusCode: http.StatusNotFound,
		},
		{
			testCase:  "コメントが存在しないようなスレッドキーの場合は200となる",
			threadKey: wrongThreadKey2,
			prepare: func(mf *mockField) {
				mf.threadApplication.EXPECT().GetByThreadKey(wrongThreadKey2).Return(thread, nil)
				mf.commentApplication.EXPECT().GetAllByThreadKey(wrongThreadKey2).Return(nil, appError.ErrNotFound)
			},
			statusCode: http.StatusOK,
		},
		{
			testCase:  "コメントを取得時にNotFound以外のエラーを発生させるとそれに応じたエラーとなる",
			threadKey: wrongThreadKey3,
			prepare: func(mf *mockField) {
				mf.threadApplication.EXPECT().GetByThreadKey(wrongThreadKey3).Return(thread, nil)
				mf.commentApplication.EXPECT().GetAllByThreadKey(wrongThreadKey3).Return(nil, errors.New("Internal Server Error"))
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			path := "/api/threads/" + c.threadKey + "/comments"
			response := executeHttpTest(r, http.MethodGet, path, nil)

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestCommenHandler_create(t *testing.T) {
	//
	// setup
	//
	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)
	commentApplication := mock_application.NewMockCommentApplication(mockCtrl)

	type mockField struct {
		sessionManager     *mock_session.MockManager
		threadApplication  *mock_application.MockThreadApplication
		commentApplication *mock_application.MockCommentApplication
	}

	field := mockField{
		sessionManager:     sessionManager,
		threadApplication:  threadApplication,
		commentApplication: commentApplication,
	}

	commentHandler := interfaces.NewCommentHandler(sessionManager, threadApplication, commentApplication)
	commentHandler.SetupRouter(r.Group("/api/threads"))

	//
	// execute
	//
	var (
		correctThreadKey = "correct-thread-key"
		wrongThreadKey   = "wrong-thread-key"
		initViews        = 0
		commentSum       = 0
		thread           = &domain.Thread{Views: &initViews, CommentSum: &commentSum}
	)
	cases := []struct {
		testCase   string
		threadKey  string
		prepare    func(*mockField)
		body       request.RequestCommentCreate
		statusCode int
	}{
		{
			testCase: "コメントが欠けている場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, nil)
			},
			threadKey:  correctThreadKey,
			body:       request.RequestCommentCreate{Contributor: "user"},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase: "投稿者が欠けている場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, nil)
			},
			threadKey:  correctThreadKey,
			body:       request.RequestCommentCreate{Comment: "comment"},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase: "セッションが無い場合は401となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, appError.ErrNotFound)
			},
			threadKey:  correctThreadKey,
			body:       request.RequestCommentCreate{Comment: "comment", Contributor: "user"},
			statusCode: http.StatusUnauthorized,
		},
		{
			testCase: "セッションがエラーの場合は500になる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			threadKey:  correctThreadKey,
			body:       request.RequestCommentCreate{Comment: "comment", Contributor: "user"},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase: "コメント生成もスレッド取得も問題ない場合200となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, nil)
				mf.commentApplication.EXPECT().CreateComment(gomock.Any()).Return(&[]domain.Comment{}, nil)
				mf.threadApplication.EXPECT().GetByThreadKey(correctThreadKey).Return(thread, nil)
			},
			threadKey:  correctThreadKey,
			body:       request.RequestCommentCreate{Comment: "comment", Contributor: "user"},
			statusCode: http.StatusOK,
		},
		{
			testCase: "コメント作成に失敗した場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, nil)
				mf.commentApplication.EXPECT().CreateComment(gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			threadKey:  wrongThreadKey,
			body:       request.RequestCommentCreate{Comment: "comment", Contributor: "user"},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase: "コメント作成後のスレッド取得に失敗した場合は404となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, nil)
				mf.commentApplication.EXPECT().CreateComment(gomock.Any()).Return(&[]domain.Comment{}, nil)
				mf.threadApplication.EXPECT().GetByThreadKey(wrongThreadKey).Return(nil, appError.ErrNotFound)
			},
			threadKey:  wrongThreadKey,
			body:       request.RequestCommentCreate{Comment: "comment", Contributor: "user"},
			statusCode: http.StatusNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			path := "/api/threads/" + c.threadKey + "/comments"
			j, _ := json.Marshal(&c.body)

			response := executeHttpTest(r, http.MethodPost, path, bytes.NewBuffer(j))

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestCommentService_edit(t *testing.T) {
	//
	// setup
	//
	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)
	commentApplication := mock_application.NewMockCommentApplication(mockCtrl)

	type mockField struct {
		sessionManager     *mock_session.MockManager
		threadApplication  *mock_application.MockThreadApplication
		commentApplication *mock_application.MockCommentApplication
	}

	field := mockField{
		sessionManager:     sessionManager,
		threadApplication:  threadApplication,
		commentApplication: commentApplication,
	}

	commentHandler := interfaces.NewCommentHandler(sessionManager, threadApplication, commentApplication)
	commentHandler.SetupRouter(r.Group("/api/threads"))

	//
	// execute
	//
	var (
		threadKey  = "thread-key"
		initViews  = 0
		commentSum = 0
		thread     = &domain.Thread{Views: &initViews, CommentSum: &commentSum}
	)
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		threadKey  string
		body       request.RequestCommentEdit
		StatusCode int
	}{
		{
			testCase: "コメントが欠けている場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, nil)
			},
			threadKey:  threadKey,
			body:       request.RequestCommentEdit{Contributor: "user"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			testCase: "投稿者が欠けている場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, nil)
			},
			threadKey:  threadKey,
			body:       request.RequestCommentEdit{Comment: "comment"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			testCase: "セッションが無い場合は401となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, appError.ErrNotFound)
			},
			threadKey:  threadKey,
			body:       request.RequestCommentEdit{Comment: "comment", Contributor: "user"},
			StatusCode: http.StatusUnauthorized,
		},
		{
			testCase: "セッションがエラーの場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			threadKey:  threadKey,
			body:       request.RequestCommentEdit{Comment: "comment", Contributor: "user"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			testCase: "コメントの編集に成功した場合は200となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, nil)
				mf.commentApplication.EXPECT().EditComment(gomock.Any()).Return(&[]domain.Comment{}, nil)
				mf.threadApplication.EXPECT().GetByThreadKey(threadKey).Return(thread, nil)
			},
			threadKey:  threadKey,
			body:       request.RequestCommentEdit{Comment: "comment", Contributor: "user"},
			StatusCode: http.StatusOK,
		},
		{
			testCase: "コメントの編集に失敗した場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, nil)
				mf.commentApplication.EXPECT().EditComment(gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			threadKey:  threadKey,
			body:       request.RequestCommentEdit{Comment: "comment", Contributor: "user"},
			StatusCode: http.StatusInternalServerError,
		},
		{
			testCase: "コメント編集後のスレッド取得に失敗した場合は404となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, nil)
				mf.commentApplication.EXPECT().EditComment(gomock.Any()).Return(&[]domain.Comment{}, nil)
				mf.threadApplication.EXPECT().GetByThreadKey(threadKey).Return(nil, appError.ErrNotFound)
			},
			threadKey:  threadKey,
			body:       request.RequestCommentEdit{Comment: "comment", Contributor: "user"},
			StatusCode: http.StatusNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			path := "/api/threads/" + c.threadKey + "/comments/comment-key"
			j, _ := json.Marshal(&c.body)

			response := executeHttpTest(r, http.MethodPut, path, bytes.NewBuffer(j))

			if response.Code != c.StatusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.StatusCode, response.Code)
			}
		})
	}
}

func TestCommentService_delete(t *testing.T) {
	//
	// setup
	//
	r := gin.Default()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)
	commentApplication := mock_application.NewMockCommentApplication(mockCtrl)

	type mockField struct {
		sessionManager     *mock_session.MockManager
		threadApplication  *mock_application.MockThreadApplication
		commentApplication *mock_application.MockCommentApplication
	}

	field := mockField{
		sessionManager:     sessionManager,
		threadApplication:  threadApplication,
		commentApplication: commentApplication,
	}

	commentHandler := interfaces.NewCommentHandler(sessionManager, threadApplication, commentApplication)
	commentHandler.SetupRouter(r.Group("/api/threads"))

	//
	// execute
	//
	var (
		threadKey  = "thread-key"
		session = session.Session{UserID: "userID"}
	)
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		StatusCode int
	}{
		{
			testCase: "セッションが無い場合は401となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, appError.ErrNotFound)
			},
			StatusCode: http.StatusUnauthorized,
		},
		{
			testCase: "セッションがエラーの場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			StatusCode: http.StatusInternalServerError,
		},
		{
			testCase: "コメントの削除に成功した場合は204となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).MaxTimes(2).Return(&session, nil)
				mf.commentApplication.EXPECT().DeleteComment(gomock.Any()).Return(nil)
			},
			StatusCode: http.StatusNoContent,
		},
		{
			testCase: "コメントの削除に失敗した場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).MaxTimes(2).Return(&session, nil)
				mf.commentApplication.EXPECT().DeleteComment(gomock.Any()).Return(errors.New("Internal Server Error"))
			},
			StatusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			path := "/api/threads/" + threadKey + "/comments/comment-key"

			response := executeHttpTest(r, http.MethodDelete, path, nil)

			if response.Code != c.StatusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.StatusCode, response.Code)
			}
		})
	}
}
