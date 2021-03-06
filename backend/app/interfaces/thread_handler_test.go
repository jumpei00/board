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

func TestThreadHandler_getAll(t *testing.T) {
	//
	// setup
	//
	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)

	type mockField struct {
		sessionManager    *mock_session.MockManager
		threadApplication *mock_application.MockThreadApplication
	}

	field := mockField{
		sessionManager:    sessionManager,
		threadApplication: threadApplication,
	}

	threadHandler := interfaces.NewThreadHandler(sessionManager, threadApplication)
	threadHandler.SetupRouter(r.Group("/api/threads"))

	//
	// exucute
	//
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		statusCode int
	}{
		{
			testCase: "空のスレッドの時のテスト",
			prepare: func(mf *mockField) {
				mf.threadApplication.EXPECT().GetAllThread().Return(nil, appError.ErrNotFound)
			},
			statusCode: http.StatusOK,
		},
		{
			testCase: "空でない時のテスト",
			prepare: func(mf *mockField) {
				mf.threadApplication.EXPECT().GetAllThread().Return(&[]domain.Thread{}, nil)
			},
			statusCode: http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			response := executeHttpTest(r, http.MethodGet, "/api/threads", nil)

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestThreadHandler_get(t *testing.T) {
	// mock
	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)

	type mockField struct {
		sessionManager    *mock_session.MockManager
		threadApplication *mock_application.MockThreadApplication
	}

	field := mockField{
		sessionManager:    sessionManager,
		threadApplication: threadApplication,
	}

	threadHandler := interfaces.NewThreadHandler(sessionManager, threadApplication)
	threadHandler.SetupRouter(r.Group("/api/threads"))

	//
	// execute
	//
	var (
		wrongThreadKey   = "wrong-key"
		correctThreadKey = "correct-key"
		initView         = 0
		commentSum       = 0
		thread           = &domain.Thread{Views: &initView, CommentSum: &commentSum}
	)
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		threadKey  string
		statucCode int
	}{
		{
			testCase: "スレッドキーに対するスレッドが存在しない場合はstatus codeが404になる",
			prepare: func(mf *mockField) {
				mf.threadApplication.EXPECT().GetByThreadKey(wrongThreadKey).Return(nil, appError.ErrNotFound)
			},
			threadKey:  wrongThreadKey,
			statucCode: http.StatusNotFound,
		},
		{
			testCase: "スレッドキーに対するスレッドが存在する場合はstatus codeが200となる",
			prepare: func(mf *mockField) {
				mf.threadApplication.EXPECT().GetByThreadKey(correctThreadKey).Return(thread, nil)
			},
			threadKey:  correctThreadKey,
			statucCode: http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			path := "/api/threads/" + c.threadKey
			response := executeHttpTest(r, http.MethodGet, path, nil)

			if response.Code != c.statucCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statucCode, response.Code)
			}
		})
	}
}

func TestThreaHandler_create(t *testing.T) {
	//
	// setup
	//
	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)

	type mockField struct {
		sessionManager    *mock_session.MockManager
		threadApplication *mock_application.MockThreadApplication
	}

	field := mockField{
		sessionManager:    sessionManager,
		threadApplication: threadApplication,
	}

	threadHandler := interfaces.NewThreadHandler(sessionManager, threadApplication)
	threadHandler.SetupRouter(r.Group("/api/threads"))

	//
	// execute
	//
	var (
		initView   = 0
		commentSum = 0
		thread     = domain.Thread{Views: &initView, CommentSum: &commentSum}
		session    = session.Session{UserID: "userID"}
	)
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		body       request.RequestThreadCreate
		statusCode int
	}{
		{
			testCase: "タイトルが欠如してリクエストされた場合は500になる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(&session, nil)
			},
			body:       request.RequestThreadCreate{},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase: "スレッドの作成に失敗した場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).MaxTimes(2).Return(&session, nil)
				mf.threadApplication.EXPECT().CreateThread(gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			body:       request.RequestThreadCreate{Title: "test-title"},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase: "sessionがない場合は401となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, appError.ErrNotFound)
			},
			body:       request.RequestThreadCreate{Title: "test-title"},
			statusCode: http.StatusUnauthorized,
		},
		{
			testCase: "session層でエラーが発生した場合は500となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			body:       request.RequestThreadCreate{Title: "test-title"},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase: "スレッドの作成に成功した場合は200となる",
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).MaxTimes(2).Return(&session, nil)
				mf.threadApplication.EXPECT().CreateThread(gomock.Any()).Return(&thread, nil)
			},
			body:       request.RequestThreadCreate{Title: "test-title"},
			statusCode: http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			j, _ := json.Marshal(&c.body)

			response := executeHttpTest(r, http.MethodPost, "/api/threads", bytes.NewBuffer(j))

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestThreadHandler_edit(t *testing.T) {
	//
	// setup
	//

	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)

	type mockField struct {
		sessionManager    *mock_session.MockManager
		threadApplication *mock_application.MockThreadApplication
	}

	field := mockField{
		sessionManager:    sessionManager,
		threadApplication: threadApplication,
	}

	threadHandler := interfaces.NewThreadHandler(sessionManager, threadApplication)
	threadHandler.SetupRouter(r.Group("/api/threads"))

	//
	// execute
	//
	var (
		correctThreadKey = "correct-thread-key"
		wrongThreadKey   = "wrong-thread-key"
		initView         = 0
		commentSum       = 0
		thread           = domain.Thread{Views: &initView, CommentSum: &commentSum}
		session          = session.Session{UserID: "userID"}
	)
	cases := []struct {
		testCase   string
		threadKey  string
		prepare    func(*mockField)
		body       request.RequestThreadEdit
		statusCode int
	}{
		{
			testCase:  "タイトルが欠如してリクエストされた場合は500になる",
			threadKey: correctThreadKey,
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).MaxTimes(2).Return(&session, nil)
			},
			body:       request.RequestThreadEdit{},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase:  "スレッドキーに対するスレッドが存在しない場合は404となる",
			threadKey: wrongThreadKey,
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).MaxTimes(2).Return(&session, nil)
				mf.threadApplication.EXPECT().EditThread(gomock.Any()).Return(nil, appError.ErrNotFound)
			},
			body:       request.RequestThreadEdit{Title: "test-title"},
			statusCode: http.StatusNotFound,
		},
		{
			testCase:  "sessionがない場合は401となる",
			threadKey: correctThreadKey,
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, appError.ErrNotFound)
			},
			body:       request.RequestThreadEdit{Title: "test-title"},
			statusCode: http.StatusUnauthorized,
		},
		{
			testCase:  "session層でエラーが発生した場合は500となる",
			threadKey: correctThreadKey,
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			body:       request.RequestThreadEdit{Title: "test-title"},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase:  "スレッドの編集に成功したら200となる",
			threadKey: correctThreadKey,
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).MaxTimes(2).Return(&session, nil)
				mf.threadApplication.EXPECT().EditThread(gomock.Any()).Return(&thread, nil)
			},
			body:       request.RequestThreadEdit{Title: "test-title"},
			statusCode: http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			path := "/api/threads/" + c.threadKey
			j, _ := json.Marshal(&c.body)

			response := executeHttpTest(r, http.MethodPut, path, bytes.NewBuffer(j))

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestThreadHandler_delete(t *testing.T) {
	//
	// setup
	//
	r := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sessionManager := mock_session.NewMockManager(mockCtrl)
	threadApplication := mock_application.NewMockThreadApplication(mockCtrl)

	type mockField struct {
		sessionManager    *mock_session.MockManager
		threadApplication *mock_application.MockThreadApplication
	}

	field := mockField{
		sessionManager:    sessionManager,
		threadApplication: threadApplication,
	}

	threadHandler := interfaces.NewThreadHandler(sessionManager, threadApplication)
	threadHandler.SetupRouter(r.Group("/api/threads"))

	//
	// execute
	//
	var (
		correctThreadKey = "correct-thread-key"
		wrongThreadKey   = "wrong-thread-key"
		session          = session.Session{UserID: "userID"}
	)
	cases := []struct {
		testCase   string
		threadKey  string
		prepare    func(*mockField)
		statusCode int
	}{
		{
			testCase:  "スレッドキーに対するスレッドが存在しない場合は404となる",
			threadKey: wrongThreadKey,
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).MaxTimes(2).Return(&session, nil)
				mf.threadApplication.EXPECT().DeleteThread(gomock.Any()).Return(appError.ErrNotFound)
			},
			statusCode: http.StatusNotFound,
		},
		{
			testCase:  "sessionがない場合は401となる",
			threadKey: correctThreadKey,
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, appError.ErrNotFound)
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			testCase:  "session層でエラーが発生した場合は500となる",
			threadKey: correctThreadKey,
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			testCase:  "スレッドの削除に成功した場合は204となる",
			threadKey: correctThreadKey,
			prepare: func(mf *mockField) {
				mf.sessionManager.EXPECT().Get(gomock.Any()).MaxTimes(2).Return(&session, nil)
				mf.threadApplication.EXPECT().DeleteThread(gomock.Any()).Return(nil)
			},
			statusCode: http.StatusNoContent,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			path := "/api/threads/" + c.threadKey

			response := executeHttpTest(r, http.MethodDelete, path, nil)

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}
