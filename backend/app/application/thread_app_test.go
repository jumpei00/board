package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/domain"
	appError "github.com/jumpei00/board/backend/app/library/error"
	mock_repository "github.com/jumpei00/board/backend/app/mock/repository"
	"github.com/jumpei00/board/backend/app/params"
)

func TestThreadApp_GetAllThread(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	testThreads := []domain.Thread{
		{Key: "1", Title: "Test1", Contributor: "test-name-1"},
		{Key: "2", Title: "Test2", Contributor: "test-name-2"},
	}

	// mock
	threadRepository.EXPECT().GetAll().Return(&testThreads, nil)

	threadApplication := application.NewThreadApplication(threadRepository, commentRepository)

	//
	// execute
	//
	threads, err := threadApplication.GetAllThread()
	if err != nil {
		t.Errorf("thread application, get all error, expected: nil, actual: %s", err)
	}
	for i, thread := range *threads {
		if diff := cmp.Diff(thread, testThreads[i]); diff != "" {
			t.Errorf("thread application, get all different contents, index: %v, diff: %s, want: %v, got: %v",
				i, diff, testThreads[i], thread,
			)
		}
	}
}

func TestThreadApp_GetByThreadKey(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		correctKey = "correct-thread-key"
		wrongKey   = "wrong-thread-key"
		thread     = domain.Thread{Key: correctKey, Title: "test-title", Contributor: "test-user"}
	)

	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	// mock
	var threadKey string
	threadRepository.EXPECT().GetByKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*domain.Thread, error) {
			if threadKey == correctKey {
				return &thread, nil
			}
			return nil, appError.ErrNotFound
		},
	)

	threadApplication := application.NewThreadApplication(threadRepository, commentRepository)

	//
	// execute
	//
	cases := []struct {
		name           string
		input          string
		expectedThread *domain.Thread
		expectedError  error
	}{
		{
			name:           "正しいスレッドキーのテスト",
			input:          correctKey,
			expectedThread: &thread,
			expectedError:  nil,
		},
		{
			name:           "間違ったスレッドキーのテスト",
			input:          wrongKey,
			expectedThread: nil,
			expectedError:  appError.ErrNotFound,
		},
	}

	for _, c := range cases {
		thread, err := threadApplication.GetByThreadKey(c.input)
		if diff := cmp.Diff(c.expectedThread, thread); diff != "" {
			t.Errorf("thread application, get by thread-key wrong thread got, name: %s, diff: %s, want: %v, got: %v",
				c.name, diff, c.expectedThread, thread,
			)
		}
		if !isSameError(c.expectedError, err) {
			t.Errorf("thread application, get bey thread-key wrong got error, name: %s, want: %s, got: %s",
				c.name, c.expectedError, err,
			)
		}
	}
}

func TestThreadApp_CreateThread(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		initialVews       = 0
		initialCommentSum = 0
		title             = "test-title"
		contibutor        = "test-user"
	)
	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	// mock
	threadRepository.EXPECT().Insert(gomock.AssignableToTypeOf(&domain.Thread{})).AnyTimes().DoAndReturn(
		func(thread *domain.Thread) (*domain.Thread, error) {
			return thread, nil
		},
	)

	threadApplication := application.NewThreadApplication(threadRepository, commentRepository)

	//
	// execute
	//
	cases := []struct {
		name           string
		input          *params.CreateThreadAppLayerParam
		expectedThread *domain.Thread
	}{
		{
			name:           "適切なパラメータのテスト",
			input:          &params.CreateThreadAppLayerParam{Title: title, Contributor: contibutor},
			expectedThread: &domain.Thread{Title: title, Contributor: contibutor, Views: &initialVews, CommentSum: &initialCommentSum},
		},
	}

	opt := cmpopts.IgnoreFields(domain.Thread{}, "Key", "Views", "CommentSum", "CreatedAt", "UpdatedAt")
	for _, c := range cases {
		thread, err := threadApplication.CreateThread(c.input)
		if diff := cmp.Diff(thread, c.expectedThread, opt); diff != "" {
			t.Errorf(
				"thread application, create wrong return thread, name: %s, diff: %s, want: %v, got %v",
				c.name, diff, c.expectedThread, thread,
			)
		}
		if *c.expectedThread.Views != *thread.Views {
			t.Errorf("thread application, create views difference, name: %s, want: %v, got: %v",
				c.name, *c.expectedThread.Views, *thread.Views,
			)
		}
		if *c.expectedThread.CommentSum != *thread.CommentSum {
			t.Errorf(
				"thread application, create comment-sum difference, name: %s, want: %v, got: %v",
				c.name, *c.expectedThread.CommentSum, *thread.CommentSum,
			)
		}
		if err != nil {
			t.Errorf("thread application, create error return, name: %s, error: %s",
				c.name, err,
			)
		}
	}
}

func TestThreadApp_EditThread(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		initialViews         = 0
		initialCommentSum    = 0
		correctKey           = "correct-key"
		wrongKey             = "wrong-key"
		originalTitle        = "original-title"
		editedTitle          = "edited-title"
		originalContributor  = "original-contributor"
		differentContributor = "different-contributor"
		originalThread       = domain.Thread{Key: correctKey, Title: originalTitle, Contributor: originalContributor, Views: &initialViews, CommentSum: &initialCommentSum}
	)
	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	// mock
	var threadKey string
	threadRepository.EXPECT().GetByKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*domain.Thread, error) {
			if threadKey == correctKey {
				return &originalThread, nil
			}
			return nil, appError.ErrNotFound
		},
	)
	threadRepository.EXPECT().Update(gomock.AssignableToTypeOf(&domain.Thread{})).AnyTimes().DoAndReturn(
		func(thread *domain.Thread) (*domain.Thread, error) {
			return thread, nil
		},
	)

	threadApplication := application.NewThreadApplication(threadRepository, commentRepository)

	//
	// exucute
	//
	cases := []struct {
		name           string
		input          *params.EditThreadAppLayerParam
		expectedThread *domain.Thread
		expectedError  error
	}{
		{
			name:           "間違ったスレッドキーの場合は失敗する",
			input:          &params.EditThreadAppLayerParam{ThreadKey: wrongKey, Title: editedTitle, Contributor: originalContributor},
			expectedThread: nil,
			expectedError:  appError.ErrNotFound,
		},
		{
			name:           "違う投稿者が編集しようとすると失敗する",
			input:          &params.EditThreadAppLayerParam{ThreadKey: correctKey, Title: editedTitle, Contributor: differentContributor},
			expectedThread: nil,
			expectedError:  &appError.BadRequest{},
		},
		{
			name:           "正常な投稿は成功する",
			input:          &params.EditThreadAppLayerParam{ThreadKey: correctKey, Title: editedTitle, Contributor: originalContributor},
			expectedThread: &domain.Thread{Key: correctKey, Title: editedTitle, Contributor: originalContributor, Views: &initialViews, CommentSum: &initialCommentSum},
			expectedError:  nil,
		},
	}

	opt := cmpopts.IgnoreFields(domain.Thread{}, "CreatedAt", "UpdatedAt")
	for _, c := range cases {
		thread, err := threadApplication.EditThread(c.input)
		if diff := cmp.Diff(c.expectedThread, thread, opt); diff != "" {
			t.Errorf(
				"thread application, edit wrong thread, name: %s, diff: %s, want: %v, got %v",
				c.name, diff, c.expectedThread, thread,
			)
		}
		if !isSameError(c.expectedError, err) {
			t.Errorf(
				"thread application, edit wrong error, name: %s, want: %v, got: %v",
				c.name, c.expectedError, err,
			)
		}
	}
}

func TestThreadApp_DeleteThread(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		correntThreadKey          = "correct-key"
		correctNoCommentThreadKey = "correct-no-comment-key"
		wrongThreadKey            = "wrong-key"
		originalContributor       = "original-contributor"
		differentContributor      = "different-contributor"
		thread                    = domain.Thread{Key: correntThreadKey, Contributor: originalContributor}
		comments                  = []domain.Comment{}
	)
	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	// mock
	var threadKey string
	threadRepository.EXPECT().GetByKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*domain.Thread, error) {
			if threadKey == wrongThreadKey {
				return nil, appError.ErrNotFound
			}
			return &thread, nil
		},
	)
	threadRepository.EXPECT().Delete(gomock.AssignableToTypeOf(&domain.Thread{})).MinTimes(1).Return(nil)
	threadRepository.EXPECT().DeleteThreadAndComments(gomock.AssignableToTypeOf(&domain.Thread{}), gomock.AssignableToTypeOf(&[]domain.Comment{})).Return(nil)
	commentRepository.EXPECT().GetAllByKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*[]domain.Comment, error) {
			if threadKey == correctNoCommentThreadKey {
				return nil, appError.ErrNotFound
			}
			return &comments, nil
		},
	)

	threadApplication := application.NewThreadApplication(threadRepository, commentRepository)

	//
	// execute
	//
	cases := []struct {
		name          string
		input         *params.DeleteThreadAppLayerParam
		expectedError error
	}{
		{
			name:          "間違ったスレッドキーの場合は失敗する",
			input:         &params.DeleteThreadAppLayerParam{ThreadKey: wrongThreadKey, Contributor: originalContributor},
			expectedError: appError.ErrNotFound,
		},
		{
			name: "異なる投稿者の場合は失敗する",
			input: &params.DeleteThreadAppLayerParam{ThreadKey: correntThreadKey, Contributor: differentContributor},
			expectedError: &appError.BadRequest{},
		},
		{
			name: "スレッドに対応するコメントが無い場合は成功する",
			input: &params.DeleteThreadAppLayerParam{ThreadKey: correctNoCommentThreadKey, Contributor: originalContributor},
			expectedError: nil,
		},
		{
			name: "スレッドに対応するコメントがある場合は成功する",
			input: &params.DeleteThreadAppLayerParam{ThreadKey: correntThreadKey, Contributor: originalContributor},
			expectedError: nil,
		},
	}

	for _, c := range cases {
		err := threadApplication.DeleteThread(c.input)
		if !isSameError(c.expectedError, err) {
			t.Errorf("thread application, wrong error return, name: %s, want: %s, got: %s", c.name, c.expectedError, err)
		}
	}
}
