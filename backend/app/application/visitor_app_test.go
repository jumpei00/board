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
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var visitor domain.Visitor

	// mock
	visitorRepository := mock_repository.NewMockVisitorRepository(mockCtrl)
	visitorRepository.EXPECT().Get().Return(&visitor, nil)

	visitorApplication := application.NewVisitorApplication(visitorRepository)

	//
	// exucute
	//
	v, err := visitorApplication.GetVisitorsStat()
	if v != &visitor {
		t.Errorf("visitor application, get visitors stat, different visitor, want: %v, got: %v", visitor, v)
	}
	if err != nil {
		t.Errorf("visitor application, get visitors stat, different error, want: nil, got: %s", err)
	}
}

func TestVisiorApp_CountupVisitors(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		yesterdayVisitor     = 0
		todayVisitor         = 0
		sumVisitor           = 0
		counupedTodayVisitor = todayVisitor + 1
		coutupedSumVisitor   = sumVisitor + 1
		visitor              = domain.Visitor{ID: 1, YesterdayVisitor: &yesterdayVisitor, TodayVisitor: &todayVisitor, VisitorSum: &sumVisitor}
	)

	// mock
	visitorRepository := mock_repository.NewMockVisitorRepository(mockCtrl)
	visitorRepository.EXPECT().Get().Return(&visitor, nil)
	visitorRepository.EXPECT().Update(gomock.AssignableToTypeOf(&domain.Visitor{})).DoAndReturn(
		func(v *domain.Visitor) (*domain.Visitor, error) {
			return v, nil
		},
	)

	visitorApplication := application.NewVisitorApplication(visitorRepository)

	//
	// exucute
	//
	expectedVisitor := &domain.Visitor{ID: 1, YesterdayVisitor: &yesterdayVisitor, TodayVisitor: &counupedTodayVisitor, VisitorSum: &coutupedSumVisitor}

	opt := cmpopts.IgnoreFields(domain.Visitor{}, "CreatedAt", "UpdatedAt")
	v, err := visitorApplication.CountupVisitors()
	if diff := cmp.Diff(v, expectedVisitor, opt); diff != "" {
		t.Errorf(
			"visitor application, coutup visitor, different visitor, diff: %s, want: %v, got: %v",
			diff, expectedVisitor, v,
		)
	}
	if err != nil {
		t.Errorf("visitor application, coutup visitor, different error, want: nil, got: %s", err)
	}
}

func TestVisitorApp_ResetVisitors(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var (
		yesterdayVisitor    = 50
		todayVisitor        = 100
		sumVisitor          = 1000
		setYesterdayVisitor = todayVisitor
		resetTodayVisitor   = 0
		visitor             = domain.Visitor{ID: 1, YesterdayVisitor: &yesterdayVisitor, TodayVisitor: &todayVisitor, VisitorSum: &sumVisitor}
	)

	// mock
	visitorRepository := mock_repository.NewMockVisitorRepository(mockCtrl)
	visitorRepository.EXPECT().Get().Return(&visitor, nil)
	visitorRepository.EXPECT().Update(gomock.AssignableToTypeOf(&domain.Visitor{})).DoAndReturn(
		func(v *domain.Visitor) (*domain.Visitor, error) {
			return v, nil
		},
	)

	visitorApplication := application.NewVisitorApplication(visitorRepository)

	//
	// exucute
	//
	expectedVisitor := &domain.Visitor{ID: 1, YesterdayVisitor: &setYesterdayVisitor, TodayVisitor: &resetTodayVisitor, VisitorSum: &sumVisitor}

	opt := cmpopts.IgnoreFields(domain.Visitor{}, "CreatedAt", "UpdatedAt")
	v, err := visitorApplication.ResetVisitors()
	if diff := cmp.Diff(v, expectedVisitor, opt); diff != "" {
		t.Errorf(
			"visitor application, reset visitors, different visitor, diff: %s, want: %v, got: %v",
			diff, expectedVisitor, v,
		)
	}
	if err != nil {
		t.Errorf("visitor application, reset visitors, different error, want: nil, got: %s", err)
	}
}
