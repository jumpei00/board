package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/domain"
	mock_repository "github.com/jumpei00/board/backend/app/mock/repository"
)

func TestVisiorApp_GetVisitorsStat(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	visitorRepository := mock_repository.NewMockVisitorRepository(mockCtrl)

	type mockField struct {
		visitorRepository *mock_repository.MockVisitorRepository
	}

	field := mockField{
		visitorRepository: visitorRepository,
	}

	visitorApplication := application.NewVisitorApplication(visitorRepository)

	//
	// exucute
	//
	var visitor = domain.Visitor{}
	cases := []struct {
		testCase        string
		prepare         func(*mockField)
		expectedVisitor *domain.Visitor
		expectedError   error
	}{
		{
			testCase: "訪問者データが取得できれば成功する",
			prepare: func(mf *mockField) {
				mf.visitorRepository.EXPECT().Get().Return(&visitor, nil)
			},
			expectedVisitor: &visitor,
			expectedError:   nil,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			v, err := visitorApplication.GetVisitorsStat()

			if v != c.expectedVisitor {
				t.Errorf("different visitor.\nwant: %v\ngot: %v", c.expectedVisitor, v)
			}
			if err != c.expectedError {
				t.Errorf("different error.\nwant: %s.\ngot: %s", c.expectedError, err)
			}
		})
	}
}

func TestVisiorApp_CountupVisitors(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	visitorRepository := mock_repository.NewMockVisitorRepository(mockCtrl)

	type mockField struct {
		visitorRepository *mock_repository.MockVisitorRepository
	}

	field := mockField{
		visitorRepository: visitorRepository,
	}

	visitorApplication := application.NewVisitorApplication(visitorRepository)

	//
	// exucute
	//
	var (
		yesterdayVisitor     = 0
		todayVisitor         = 0
		sumVisitor           = 0
		counupedTodayVisitor = todayVisitor + 1
		coutupedSumVisitor   = sumVisitor + 1
		originalVisitor      = &domain.Visitor{ID: 1, YesterdayVisitor: &yesterdayVisitor, TodayVisitor: &todayVisitor, VisitorSum: &sumVisitor}
		expectedVisitor      = &domain.Visitor{ID: 1, YesterdayVisitor: &yesterdayVisitor, TodayVisitor: &counupedTodayVisitor, VisitorSum: &coutupedSumVisitor}
	)
	cases := []struct {
		testCase        string
		prepare         func(*mockField)
		expectedVisitor *domain.Visitor
		expectedError   error
	}{
		{
			testCase: "カウントアップが行えたら成功する",
			prepare: func(mf *mockField) {
				mf.visitorRepository.EXPECT().Get().Return(originalVisitor, nil)
				mf.visitorRepository.EXPECT().Update(gomock.Any()).DoAndReturn(
					func(visitor *domain.Visitor) (*domain.Visitor, error) {
						return visitor, nil
					},
				)
			},
			expectedVisitor: expectedVisitor,
			expectedError:   nil,
		},
	}

	opt := cmpopts.IgnoreFields(domain.Visitor{}, "CreatedAt", "UpdatedAt")

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			v, err := visitorApplication.CountupVisitors()

			if diff := cmp.Diff(v, expectedVisitor, opt); diff != "" {
				t.Errorf("different visitor.\ndiff: %s\nwant: %v\ngot: %v", diff, expectedVisitor, v)
			}
			if err != c.expectedError {
				t.Errorf("different error.\nwant: %s\ngot: %s", c.expectedError, err)
			}
		})
	}
}

func TestVisitorApp_ResetVisitors(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	visitorRepository := mock_repository.NewMockVisitorRepository(mockCtrl)

	type mockField struct {
		visitorRepository *mock_repository.MockVisitorRepository
	}

	field := mockField{
		visitorRepository: visitorRepository,
	}

	visitorApplication := application.NewVisitorApplication(visitorRepository)

	//
	// exucute
	//
	var (
		yesterdayVisitor    = 50
		todayVisitor        = 100
		sumVisitor          = 1000
		setYesterdayVisitor = todayVisitor
		resetTodayVisitor   = 0
		originalVisitor     = &domain.Visitor{ID: 1, YesterdayVisitor: &yesterdayVisitor, TodayVisitor: &todayVisitor, VisitorSum: &sumVisitor}
		expectedVisitor     = &domain.Visitor{ID: 1, YesterdayVisitor: &setYesterdayVisitor, TodayVisitor: &resetTodayVisitor, VisitorSum: &sumVisitor}
	)
	cases := []struct {
		testCase        string
		prepare         func(*mockField)
		expectedVisitor *domain.Visitor
		expectedError   error
	}{
		{
			testCase: "訪問者のリセット処理が完了すると成功する",
			prepare: func(mf *mockField) {
				mf.visitorRepository.EXPECT().Get().Return(originalVisitor, nil)
				mf.visitorRepository.EXPECT().Update(gomock.Any()).DoAndReturn(
					func(visitor *domain.Visitor) (*domain.Visitor, error) {
						return visitor, nil
					},
				)
			},
		},
	}

	opt := cmpopts.IgnoreFields(domain.Visitor{}, "CreatedAt", "UpdatedAt")

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			v, err := visitorApplication.ResetVisitors()

			if diff := cmp.Diff(v, expectedVisitor, opt); diff != "" {
				t.Errorf("different visitor.\ndiff: %s\nwant: %v\ngot: %v", diff, expectedVisitor, v)
			}
			if err != c.expectedError {
				t.Errorf("different error.\nwant: %s\ngot: %s", c.expectedError, err)
			}
		})
	}
}
