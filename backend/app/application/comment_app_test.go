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
)

func TestCommentApp_GetAllByThreadKey(t *testing.T) {
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

	commentApplication := application.NewCommentApplication(threadRepository, commentRepository)

	//
	// execute
	//
	var (
		correctThreadKey = "correct-key"
		wrongThreadKey   = "wrong-key"
		comments         = []domain.Comment{}
	)
	cases := []struct {
		testCase         string
		input            string
		prepare          func(*mockField)
		expectedComments *[]domain.Comment
		expectedError    error
	}{
		{
			testCase: "間違ったスレッドキーの場合は失敗する",
			input:    wrongThreadKey,
			prepare: func(mf *mockField) {
				mf.commentRepository.EXPECT().GetAllByKey(wrongThreadKey).Return(nil, appError.ErrNotFound)
			},
			expectedComments: nil,
			expectedError:    appError.ErrNotFound,
		},
		{
			testCase: "正しいスレッドキーの場合は成功する",
			input:    correctThreadKey,
			prepare: func(mf *mockField) {
				mf.commentRepository.EXPECT().GetAllByKey(correctThreadKey).Return(&comments, nil)
			},
			expectedComments: &comments,
			expectedError:    nil,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			comments, err := commentApplication.GetAllByThreadKey(c.input)
			if comments != c.expectedComments {
				t.Errorf("different comment.\nwant: %v\ngot: %v", c.expectedComments, comments)
			}
			if !isSameError(err, c.expectedError) {
				t.Errorf("different error.\nwant: %s\ngot: %s", c.expectedError, err)
			}
		})
	}
}

func TestCommentApp_CreateComment(t *testing.T) {
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

	commentApplication := application.NewCommentApplication(threadRepository, commentRepository)

	//
	// execute
	//
	var (
		correctThreadKey  = "correct-thread-key"
		wrongThreadKey    = "wrong-thread-key"
		comment           = "comment-contents"
		contibutor        = "comment-user"
		initialCommentSum = 0
		thread            = domain.Thread{Key: correctThreadKey, CommentSum: &initialCommentSum}
		comments          = []domain.Comment{}
	)
	cases := []struct {
		testCase         string
		input            params.CreateCommentAppLayerParam
		prepare          func(*mockField)
		expectedComments *[]domain.Comment
		expectedError    error
	}{
		{
			testCase: "間違ったスレッドキーの場合は失敗する",
			input:    params.CreateCommentAppLayerParam{ThreadKey: wrongThreadKey, Comment: comment, Contributor: contibutor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(wrongThreadKey).Return(nil, appError.ErrNotFound)
			},
			expectedComments: nil,
			expectedError:    appError.ErrNotFound,
		},
		{
			testCase: "正しいスレッドキーの場合は成功する",
			input:    params.CreateCommentAppLayerParam{ThreadKey: correctThreadKey, Comment: comment, Contributor: contibutor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correctThreadKey).Return(&thread, nil)
				mf.threadRepository.EXPECT().Update(gomock.Any()).Return(nil, nil)
				mf.commentRepository.EXPECT().Insert(gomock.Any()).DoAndReturn(
					func(comment *domain.Comment) (*domain.Comment, error) {
						comments = append(comments, *comment)
						return comment, nil
					},
				)
				mf.commentRepository.EXPECT().GetAllByKey(correctThreadKey).Return(&comments, nil)
			},
			expectedComments: &[]domain.Comment{{ThreadKey: correctThreadKey, Comment: comment, Contributor: contibutor}},
			expectedError:    nil,
		},
	}

	opt := cmpopts.IgnoreFields(domain.Comment{}, "Key", "CreatedAt", "UpdatedAt")

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			comments, err := commentApplication.CreateComment(&c.input)

			if diff := cmp.Diff(comments, c.expectedComments, opt); diff != "" {
				t.Errorf("different comments.\ndiff: %s\nwant: %v\ngot: %v", diff, c.expectedComments, comments)
			}
			if !isSameError(err, c.expectedError) {
				t.Errorf("different error.\nwant: %s\ngot: %s", c.expectedError, err)
			}
		})
	}
}

func TestCommentApp_EditComment(t *testing.T) {
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

	commentApplication := application.NewCommentApplication(threadRepository, commentRepository)

	//
	// exucute
	//
	var (
		correctThreadKey   = "correct-thread-key"
		correctCommentKey  = "correct-comment-key"
		correctContributor = "correct-contributor"
		wrongThreadKey     = "wrong-thread-key"
		wrongCommentKey    = "wrong-comment-key"
		wrongContributor   = "wrong-contributor"
		editedComment      = "edited-comment"
		thread             = domain.Thread{Key: correctThreadKey}
		comment            = domain.Comment{Key: correctCommentKey, ThreadKey: correctThreadKey, Contributor: correctContributor}
		comments           = []domain.Comment{}
	)
	cases := []struct {
		testCase         string
		input            params.EditCommentAppLayerParam
		prepare          func(*mockField)
		expectedComments *[]domain.Comment
		expectedError    error
	}{
		{
			testCase: "間違ったスレッドキーの場合失敗する",
			input:    params.EditCommentAppLayerParam{ThreadKey: wrongThreadKey},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(wrongThreadKey).Return(nil, appError.ErrNotFound)
			},
			expectedComments: nil,
			expectedError:    appError.ErrNotFound,
		},
		{
			testCase: "間違ったコメントキーの場合失敗する",
			input:    params.EditCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: wrongCommentKey},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correctThreadKey).Return(&thread, nil)
				mf.commentRepository.EXPECT().GetByKey(wrongCommentKey).Return(nil, appError.ErrNotFound)
			},
			expectedComments: nil,
			expectedError:    appError.ErrNotFound,
		},
		{
			testCase: "違うユーザーは編集できない",
			input:    params.EditCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: correctCommentKey, Contributor: wrongContributor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correctThreadKey).Return(&thread, nil)
				mf.commentRepository.EXPECT().GetByKey(correctCommentKey).Return(&comment, nil)
			},
			expectedComments: nil,
			expectedError:    &appError.BadRequest{},
		},
		{
			testCase: "同じユーザーの場合は成功する",
			input:    params.EditCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: correctCommentKey, Contributor: correctContributor, Comment: editedComment},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correctThreadKey).Return(&thread, nil)
				mf.commentRepository.EXPECT().GetByKey(correctCommentKey).Return(&comment, nil)
				mf.commentRepository.EXPECT().Insert(gomock.Any()).DoAndReturn(
					func(comment *domain.Comment) (*domain.Comment, error) {
						comments = append(comments, *comment)
						return comment, nil
					},
				)
				mf.threadRepository.EXPECT().Update(gomock.Any()).Return(nil, nil)
				mf.commentRepository.EXPECT().GetAllByKey(correctThreadKey).Return(&comments, nil)
			},
			expectedComments: &[]domain.Comment{{Key: correctCommentKey, ThreadKey: correctThreadKey, Contributor: correctContributor, Comment: editedComment}},
			expectedError:    nil,
		},
	}

	//
	// execute
	//
	opt := cmpopts.IgnoreFields(domain.Comment{}, "CreatedAt", "UpdatedAt")

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			comments, err := commentApplication.EditComment(&c.input)

			if diff := cmp.Diff(comments, c.expectedComments, opt); diff != "" {
				t.Errorf("different comments.\ndiff: %s\nwant: %v, got: %v", diff, c.expectedComments, comments)
			}
			if !isSameError(err, c.expectedError) {
				t.Errorf("different error.\nwant: %s\ngot: %s", c.expectedError, err)
			}
		})
	}
}

func TestCommentApp_DeleteComment(t *testing.T) {
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

	commentApplication := application.NewCommentApplication(threadRepository, commentRepository)

	//
	// exucute
	//
	var (
		correctThreadKey   = "correct-thread-key"
		correctCommentKey  = "correct-comment-key"
		correctContributor = "correct-contributor"
		wrongThreadKey     = "wrong-thread-key"
		wrongCommentKey    = "wrong-comment-key"
		wrongContributor   = "wrong-contributor"
		initialCommentSum  = 100
		thread             = domain.Thread{Key: correctThreadKey, CommentSum: &initialCommentSum}
		comment            = domain.Comment{Key: correctCommentKey, ThreadKey: correctThreadKey, Contributor: correctContributor}
		comments           = []domain.Comment{}
	)
	cases := []struct {
		testCase         string
		input            params.DeleteCommentAppLayerParam
		prepare          func(*mockField)
		expectedComments *[]domain.Comment
		expectedError    error
	}{
		{
			testCase: "間違ったスレッドキーの場合失敗する",
			input:    params.DeleteCommentAppLayerParam{ThreadKey: wrongThreadKey},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(wrongThreadKey).Return(nil, appError.ErrNotFound)
			},
			expectedComments: nil,
			expectedError:    appError.ErrNotFound,
		},
		{
			testCase: "間違ったコメントキーの場合失敗する",
			input:    params.DeleteCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: wrongCommentKey},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correctThreadKey).Return(&thread, nil)
				mf.commentRepository.EXPECT().GetByKey(wrongCommentKey).Return(nil, appError.ErrNotFound)
			},
			expectedComments: nil,
			expectedError:    appError.ErrNotFound,
		},
		{
			testCase: "違うユーザーは削除できない",
			input:    params.DeleteCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: correctCommentKey, Contributor: wrongContributor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correctThreadKey).Return(&thread, nil)
				mf.commentRepository.EXPECT().GetByKey(correctCommentKey).Return(&comment, nil)
			},
			expectedComments: nil,
			expectedError:    &appError.BadRequest{},
		},
		{
			testCase: "正常なパラメータであれば成功する",
			input:    params.DeleteCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: correctCommentKey, Contributor: correctContributor},
			prepare: func(mf *mockField) {
				mf.threadRepository.EXPECT().GetByKey(correctThreadKey).Return(&thread, nil)
				mf.commentRepository.EXPECT().GetByKey(correctCommentKey).Return(&comment, nil)
				mf.commentRepository.EXPECT().Delete(gomock.Any()).Return(nil)
				mf.threadRepository.EXPECT().Update(gomock.Any()).Return(nil, nil)
				mf.commentRepository.EXPECT().GetAllByKey(correctThreadKey).Return(&comments, nil)
			},
			expectedComments: &comments,
			expectedError:    nil,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			comments, err := commentApplication.DeleteComment(&c.input)

			if comments != c.expectedComments {
				t.Errorf("different comments.\nwant: %v\ngot: %v", c.expectedComments, comments)
			}
			if !isSameError(err, c.expectedError) {
				t.Errorf("different error.\nwant: %s\ngot: %s", c.expectedError, err)
			}
		})
	}
}
