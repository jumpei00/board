package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/application/params"
	"github.com/jumpei00/board/backend/app/domain"
	appError "github.com/jumpei00/board/backend/app/library/error"
	mock_repository "github.com/jumpei00/board/backend/app/mock/repository"
	"github.com/pkg/errors"
)

func TestThreadApp_GetAllThread(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	type mockField struct {
		threadRepository  *mock_repository.MockThreadRepository
		commentRepository *mock_repository.MockCommentRepository
	}

	field := mockField{
		threadRepository:  threadRepository,
		commentRepository: commentRepository,
	}

	threadApplication := application.NewThreadApplication(threadRepository, commentRepository)

	//
	// execute
	//
	var (
		testThreads = []domain.Thread{
			{Key: "1", Title: "Test1", Contributor: "test-name-1"},
			{Key: "2", Title: "Test2", Contributor: "test-name-2"},
		}
	)
	cases := []struct {
		testCase        string
		prepare         func(*mockField)
		expectedThreads *[]domain.Thread
		expectedError   error
	}{
		{
			testCase: "スレッドが存在していて取得できれば成功する",
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetAll().Return(&testThreads, nil)
			},
			expectedThreads: &testThreads,
			expectedError:   nil,
		},
		{
			testCase: "スレッドが存在していなければ失敗する",
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetAll().Return(nil, appError.ErrNotFound)
			},
			expectedThreads: nil,
			expectedError:   appError.ErrNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			threads, err := threadApplication.GetAllThread()

			if err != c.expectedError {
				t.Errorf("different error.\nwant: nil\ngot: %s", err)
			}

			if diff := cmp.Diff(threads, c.expectedThreads); diff != "" {
				t.Errorf("different threads.\ndiff: %s\nwant: %v\ngot: %v", diff, c.expectedThreads, threads)
			}
		})
	}
}

func TestThreadApp_GetByThreadKey(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	type mockField struct {
		threadRepository  *mock_repository.MockThreadRepository
		commentRepository *mock_repository.MockCommentRepository
	}

	field := mockField{
		threadRepository:  threadRepository,
		commentRepository: commentRepository,
	}

	threadApplication := application.NewThreadApplication(threadRepository, commentRepository)

	//
	// execute
	//
	var (
		correctKey = "correct-thread-key"
		wrongKey   = "wrong-thread-key"
		thread     = domain.Thread{Key: correctKey, Title: "test-title", Contributor: "test-user"}
	)
	cases := []struct {
		testCase       string
		input          string
		prepare        func(*mockField)
		expectedThread *domain.Thread
		expectedError  error
	}{
		{
			testCase: "キーに対するスレッドが存在する場合は成功するテスト",
			input:    correctKey,
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correctKey).Return(&thread, nil)
			},
			expectedThread: &thread,
			expectedError:  nil,
		},
		{
			testCase: "キーに対するスレッドが存在しない場合は失敗するテスト",
			input:    wrongKey,
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(wrongKey).Return(nil, appError.ErrNotFound)
			},
			expectedThread: nil,
			expectedError:  appError.ErrNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			thread, err := threadApplication.GetByThreadKey(c.input)

			if diff := cmp.Diff(c.expectedThread, thread); diff != "" {
				t.Errorf("different thread.\ndiff: %s\nwant: %v\ngot: %v", diff, c.expectedThread, thread)
			}
			if !isSameError(c.expectedError, err) {
				t.Errorf("different error.\nwant: %s\ngot: %s", c.expectedError, err)
			}
		})
	}
}

func TestThreadApp_CreateThread(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	type mockField struct {
		threadRepository  *mock_repository.MockThreadRepository
		commentRepository *mock_repository.MockCommentRepository
	}

	field := mockField{
		threadRepository:  threadRepository,
		commentRepository: commentRepository,
	}

	threadApplication := application.NewThreadApplication(threadRepository, commentRepository)

	//
	// execute
	//
	var (
		initialVews       = 0
		initialCommentSum = 0
		title             = "test-title"
		contibutor        = "test-user"
		Error             = errors.New("Internal Server Error")
	)
	cases := []struct {
		testCase       string
		input          params.CreateThreadAppLayerParam
		prepare        func(*mockField)
		expectedThread *domain.Thread
		expectedError  error
	}{
		{
			testCase: "スレッドが作成された場合は成功する",
			input:    params.CreateThreadAppLayerParam{Title: title, Contributor: contibutor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().Insert(gomock.Any()).DoAndReturn(
					func(thread *domain.Thread) (*domain.Thread, error) {
						return thread, nil
					},
				)
			},
			expectedThread: &domain.Thread{Title: title, Contributor: contibutor, Views: &initialVews, CommentSum: &initialCommentSum},
			expectedError:  nil,
		},
		{
			testCase: "スレッド作成時にエラーが発生した場合は失敗する",
			input:    params.CreateThreadAppLayerParam{Title: title, Contributor: contibutor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().Insert(gomock.Any()).Return(nil, Error)
			},
			expectedThread: nil,
			expectedError:  Error,
		},
	}

	opt := cmpopts.IgnoreFields(domain.Thread{}, "Key", "Views", "CommentSum", "CreatedAt", "UpdatedAt")

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			thread, err := threadApplication.CreateThread(&c.input)

			if diff := cmp.Diff(thread, c.expectedThread, opt); diff != "" {
				t.Errorf("different thread.\ndiff: %s\nwant: %v\ngot %v", diff, c.expectedThread, thread)
			}
			if !isSameError(c.expectedError, err) {
				t.Errorf("different error.\nwant: %s\ngot: %s", c.expectedError, err)
			}
		})
	}
}

func TestThreadApp_EditThread(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	type mockField struct {
		threadRepository  *mock_repository.MockThreadRepository
		commentRepository *mock_repository.MockCommentRepository
	}

	field := mockField{
		threadRepository:  threadRepository,
		commentRepository: commentRepository,
	}

	threadApplication := application.NewThreadApplication(threadRepository, commentRepository)

	//
	// exucute
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
	cases := []struct {
		testCase       string
		input          params.EditThreadAppLayerParam
		prepare        func(*mockField)
		expectedThread *domain.Thread
		expectedError  error
	}{
		{
			testCase: "スレッドが存在しない場合は失敗する",
			input:    params.EditThreadAppLayerParam{ThreadKey: wrongKey, Title: editedTitle, Contributor: originalContributor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(wrongKey).Return(nil, appError.ErrNotFound)
			},
			expectedThread: nil,
			expectedError:  appError.ErrNotFound,
		},
		{
			testCase: "違う投稿者が編集しようとすると失敗する",
			input:    params.EditThreadAppLayerParam{ThreadKey: correctKey, Title: editedTitle, Contributor: differentContributor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correctKey).Return(&originalThread, nil)
			},
			expectedThread: nil,
			expectedError:  &appError.BadRequest{},
		},
		{
			testCase: "編集が正しく終了した場合は成功する",
			input:    params.EditThreadAppLayerParam{ThreadKey: correctKey, Title: editedTitle, Contributor: originalContributor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correctKey).Return(&originalThread, nil)
				mf.threadRepository.EXPECT().Update(gomock.Any()).DoAndReturn(
					func(thread *domain.Thread) (*domain.Thread, error) {
						return thread, nil
					},
				)
			},
			expectedThread: &domain.Thread{Key: correctKey, Title: editedTitle, Contributor: originalContributor, Views: &initialViews, CommentSum: &initialCommentSum},
			expectedError:  nil,
		},
	}

	opt := cmpopts.IgnoreFields(domain.Thread{}, "CreatedAt", "UpdatedAt")

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			thread, err := threadApplication.EditThread(&c.input)

			if diff := cmp.Diff(c.expectedThread, thread, opt); diff != "" {
				t.Errorf("different thread.\ndiff: %s\nwant: %v\ngot %v", diff, c.expectedThread, thread)
			}
			if !isSameError(c.expectedError, err) {
				t.Errorf("different error.\nwant: %v\ngot: %v", c.expectedError, err)
			}
		})
	}
}

func TestThreadApp_DeleteThread(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	type mockField struct {
		threadRepository  *mock_repository.MockThreadRepository
		commentRepository *mock_repository.MockCommentRepository
	}

	field := mockField{
		threadRepository:  threadRepository,
		commentRepository: commentRepository,
	}

	threadApplication := application.NewThreadApplication(threadRepository, commentRepository)

	//
	// execute
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
	cases := []struct {
		testCase      string
		input         params.DeleteThreadAppLayerParam
		prepare       func(*mockField)
		expectedError error
	}{
		{
			testCase: "スレッドが存在しない場合は失敗する",
			input:    params.DeleteThreadAppLayerParam{ThreadKey: wrongThreadKey, Contributor: originalContributor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(wrongThreadKey).Return(nil, appError.ErrNotFound)
			},
			expectedError: appError.ErrNotFound,
		},
		{
			testCase: "異なる投稿者の場合は失敗する",
			input:    params.DeleteThreadAppLayerParam{ThreadKey: correntThreadKey, Contributor: differentContributor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correntThreadKey).Return(&thread, nil)
			},
			expectedError: &appError.BadRequest{},
		},
		{
			testCase: "スレッドに対応するコメントが無い場合は成功する",
			input:    params.DeleteThreadAppLayerParam{ThreadKey: correctNoCommentThreadKey, Contributor: originalContributor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correctNoCommentThreadKey).Return(&thread, nil)
				mf.commentRepository.EXPECT().GetAllByKey(correctNoCommentThreadKey).Return(nil, appError.ErrNotFound)
				mf.threadRepository.EXPECT().Delete(&thread).Return(nil)
			},
			expectedError: nil,
		},
		{
			testCase: "スレッドに対応するコメントがあって、削除できた場合は成功する",
			input:    params.DeleteThreadAppLayerParam{ThreadKey: correntThreadKey, Contributor: originalContributor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correntThreadKey).Return(&thread, nil)
				mf.commentRepository.EXPECT().GetAllByKey(correntThreadKey).Return(&comments, nil)
				mf.threadRepository.EXPECT().DeleteThreadAndComments(&thread, &comments).Return(nil)
			},
			expectedError: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			err := threadApplication.DeleteThread(&c.input)

			if !isSameError(c.expectedError, err) {
				t.Errorf("different error.\nwant: %s\ngot: %s", c.expectedError, err)
			}
		})
	}
}
