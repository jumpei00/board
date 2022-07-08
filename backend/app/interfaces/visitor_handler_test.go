package interfaces_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/interfaces"
	appError "github.com/jumpei00/board/backend/app/library/error"
	mock_application "github.com/jumpei00/board/backend/app/mock/application"
	"github.com/pkg/errors"
)

func TestVisitorHandler_get(t *testing.T) {
	// setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	r := gin.Default()

	visitorApplication := mock_application.NewMockVisitorApplication(mockCtrl)

	type mockField struct {
		visitorApplication *mock_application.MockVisitorApplication
	}

	field := mockField{
		visitorApplication: visitorApplication,
	}

	visitorHandler := interfaces.NewVisitorsHandler(visitorApplication)
	visitorHandler.SetupRouter(r.Group("/api/visitor"))

	// execute
	var (
		yesterday = 0
		today     = 0
		sum       = 0
		visitor   = domain.Visitor{
			YesterdayVisitor: &yesterday,
			TodayVisitor:     &today,
			VisitorSum:       &sum,
		}
	)
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		statusCode int
	}{
		{
			testCase: "訪問者データが取得できれば200を返す",
			prepare: func(mf *mockField) {
				mf.visitorApplication.EXPECT().GetVisitorsStat().Return(&visitor, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			testCase: "訪問者データが存在しなければ404となる",
			prepare: func(mf *mockField) {
				mf.visitorApplication.EXPECT().GetVisitorsStat().Return(nil, appError.ErrNotFound)
			},
			statusCode: http.StatusNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			response := executeHttpTest(r, http.MethodGet, "/api/visitor", nil)

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestVisitorHandler_visited(t *testing.T) {
	// setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	r := gin.Default()

	visitorApplication := mock_application.NewMockVisitorApplication(mockCtrl)

	type mockField struct {
		visitorApplication *mock_application.MockVisitorApplication
	}

	field := mockField{
		visitorApplication: visitorApplication,
	}

	visitorHandler := interfaces.NewVisitorsHandler(visitorApplication)
	visitorHandler.SetupRouter(r.Group("/api/visitor"))

	// execute
	var (
		yesterday = 0
		today     = 0
		sum       = 0
		visitor   = domain.Visitor{
			YesterdayVisitor: &yesterday,
			TodayVisitor:     &today,
			VisitorSum:       &sum,
		}
	)
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		statusCode int
	}{
		{
			testCase: "訪問者のカウントアップに成功したら200となる",
			prepare: func(mf *mockField) {
				mf.visitorApplication.EXPECT().CountupVisitors().Return(&visitor, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			testCase: "訪問者のカウントアップに失敗したら500となる",
			prepare: func(mf *mockField) {
				mf.visitorApplication.EXPECT().CountupVisitors().Return(nil, errors.New("Internal Server Error"))
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			response := executeHttpTest(r, http.MethodPut, "/api/visitor", nil)

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}

func TestVisitorHandler_reset(t *testing.T) {
	// setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	r := gin.Default()

	visitorApplication := mock_application.NewMockVisitorApplication(mockCtrl)

	type mockField struct {
		visitorApplication *mock_application.MockVisitorApplication
	}

	field := mockField{
		visitorApplication: visitorApplication,
	}

	visitorHandler := interfaces.NewVisitorsHandler(visitorApplication)
	visitorHandler.SetupRouter(r.Group("/api/visitor"))

	// execute
	cases := []struct {
		testCase   string
		prepare    func(*mockField)
		statusCode int
	}{
		{
			testCase: "訪問者のリセットに成功したら204となる",
			prepare: func(mf *mockField) {
				mf.visitorApplication.EXPECT().ResetVisitors().Return(nil, nil)
			},
			statusCode: http.StatusNoContent,
		},
		{
			testCase: "訪問者のリセットに失敗したら500となる",
			prepare: func(mf *mockField) {
				mf.visitorApplication.EXPECT().ResetVisitors().Return(nil, errors.New("Internal Server Error"))
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			response := executeHttpTest(r, http.MethodPut, "/api/visitor/reset", nil)

			if response.Code != c.statusCode {
				t.Errorf("different status code.\nwant: %d\ngot: %d", c.statusCode, response.Code)
			}
		})
	}
}
