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

func TestCommentApp_GetAllByThreadKey(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		correctThreadKey = "correct-key"
		wrongThreadKey   = "wrong-key"
		comments         = []domain.Comment{}
	)
	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	// mock
	var threadKey string
	commentRepository.EXPECT().GetAllByKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*[]domain.Comment, error) {
			if threadKey == wrongThreadKey {
				return nil, appError.ErrNotFound
			}
			return &comments, nil
		},
	)

	commentApplication := application.NewCommentApplication(threadRepository, commentRepository)

	//
	// execute
	//
	cases := []struct {
		name             string
		input            string
		expectedComments *[]domain.Comment
		expectedError    error
	}{
		{
			name:             "間違ったスレッドキーの場合は失敗する",
			input:            wrongThreadKey,
			expectedComments: nil,
			expectedError:    appError.ErrNotFound,
		},
		{
			name:             "正しいスレッドキーの場合は成功する",
			input:            correctThreadKey,
			expectedComments: &comments,
			expectedError:    nil,
		},
	}

	for _, c := range cases {
		comments, err := commentApplication.GetAllByThreadKey(c.input)
		if comments != c.expectedComments {
			t.Errorf(
				"comment application, get all by thread key, different comment, name: %s, want: %v, got: %v",
				c.name, c.expectedComments, comments,
			)
		}
		if !isSameError(err, c.expectedError) {
			t.Errorf(
				"comment application, get all by thread key, different error, name: %s, want: %s, got: %s",
				c.name, c.expectedError, err,
			)
		}
	}
}

func TestCommentApp_CreateComment(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
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
	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	var threadKey string
	threadRepository.EXPECT().GetByKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*domain.Thread, error) {
			if threadKey == wrongThreadKey {
				return nil, appError.ErrNotFound
			}
			return &thread, nil
		},
	)
	threadRepository.EXPECT().Update(gomock.AssignableToTypeOf(&domain.Thread{})).AnyTimes().DoAndReturn(
		func(thread *domain.Thread) (*domain.Thread, error) {
			return thread, nil
		},
	)
	commentRepository.EXPECT().Insert(gomock.AssignableToTypeOf(&domain.Comment{})).AnyTimes().DoAndReturn(
		func(comment *domain.Comment) (*domain.Comment, error) {
			// init
			comments = []domain.Comment{}

			comments = append(comments, *comment)
			return comment, nil
		},
	)
	commentRepository.EXPECT().GetAllByKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*[]domain.Comment, error) {
			if threadKey == wrongThreadKey {
				return nil, appError.ErrNotFound
			}
			return &comments, nil
		},
	)

	commentApplication := application.NewCommentApplication(threadRepository, commentRepository)

	//
	// execute
	//
	cases := []struct {
		name             string
		input            params.CreateCommentAppLayerParam
		expectedComments *[]domain.Comment
		expectedError    error
	}{
		{
			name:             "間違ったスレッドキーの場合は失敗する",
			input:            params.CreateCommentAppLayerParam{ThreadKey: wrongThreadKey, Comment: comment, Contributor: contibutor},
			expectedComments: nil,
			expectedError:    appError.ErrNotFound,
		},
		{
			name:             "正しいスレッドキーの場合は成功する",
			input:            params.CreateCommentAppLayerParam{ThreadKey: correctThreadKey, Comment: comment, Contributor: contibutor},
			expectedComments: &[]domain.Comment{{ThreadKey: correctThreadKey, Comment: comment, Contributor: contibutor}},
			expectedError:    nil,
		},
	}

	opt := cmpopts.IgnoreFields(domain.Comment{}, "Key", "CreatedAt", "UpdatedAt")
	for _, c := range cases {
		comments, err := commentApplication.CreateComment(&c.input)
		if diff := cmp.Diff(comments, c.expectedComments, opt); diff != "" {
			t.Errorf(
				"comment application, create comment, different comments, name: %s, diff: %s, want: %v, got: %v",
				c.name, diff, c.expectedComments, comments,
			)
		}
		if !isSameError(err, c.expectedError) {
			t.Errorf(
				"comment application, create comment, different error, name: %s, want: %s, got: %s",
				c.name, c.expectedError, err,
			)
		}
	}
}

func TestCommentApp_EditComment(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		correctThreadKey   = "correct-thread-key"
		correctCommentKey  = "correct-comment-key"
		correctContributor = "correct-contributor"
		wrongThreadKey     = "wrong-thread-key"
		wrongCommentKey    = "wrong-comment-key"
		wrongContributor   = "wrong-contributor"
		initialComment     = "initial-comment"
		editedComment      = "edited-comment"
		thread             = domain.Thread{Key: correctThreadKey}
		comment            = domain.Comment{Key: correctCommentKey, ThreadKey: correctThreadKey, Contributor: correctContributor}
		comments           = []domain.Comment{}
	)
	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	var threadKey string
	var commentKey string
	threadRepository.EXPECT().GetByKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*domain.Thread, error) {
			if threadKey == wrongThreadKey {
				return nil, appError.ErrNotFound
			}
			return &thread, nil
		},
	)
	threadRepository.EXPECT().Update(gomock.AssignableToTypeOf(&domain.Thread{})).AnyTimes().DoAndReturn(
		func(thread *domain.Thread) (*domain.Thread, error) {
			return thread, nil
		},
	)
	commentRepository.EXPECT().GetByKey(gomock.AssignableToTypeOf(commentKey)).AnyTimes().DoAndReturn(
		func(commentKey string) (*domain.Comment, error) {
			if commentKey == wrongCommentKey {
				return nil, appError.ErrNotFound
			}

			// init
			comment.Comment = initialComment

			return &comment, nil
		},
	)
	commentRepository.EXPECT().Insert(gomock.AssignableToTypeOf(&domain.Comment{})).AnyTimes().DoAndReturn(
		func(comment *domain.Comment) (*domain.Comment, error) {
			// init
			comments = []domain.Comment{}

			comments = append(comments, *comment)
			return comment, nil
		},
	)
	commentRepository.EXPECT().GetAllByKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*[]domain.Comment, error) {
			if threadKey == wrongThreadKey {
				return nil, appError.ErrNotFound
			}
			return &comments, nil
		},
	)

	commentApplication := application.NewCommentApplication(threadRepository, commentRepository)

	//
	// exucute
	//
	cases := []struct {
		name             string
		input            params.EditCommentAppLayerParam
		expectedComments *[]domain.Comment
		expectedError    error
	}{
		{
			name:             "間違ったスレッドキーの場合失敗する",
			input:            params.EditCommentAppLayerParam{ThreadKey: wrongThreadKey},
			expectedComments: nil,
			expectedError:    appError.ErrNotFound,
		},
		{
			name:             "間違ったコメントキーの場合失敗する",
			input:            params.EditCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: wrongCommentKey},
			expectedComments: nil,
			expectedError:    appError.ErrNotFound,
		},
		{
			name:             "違うユーザーは編集できない",
			input:            params.EditCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: correctCommentKey, Contributor: wrongContributor},
			expectedComments: nil,
			expectedError:    &appError.BadRequest{},
		},
		{
			name:             "正常なパラメータであれば成功する",
			input:            params.EditCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: correctCommentKey, Contributor: correctContributor, Comment: editedComment},
			expectedComments: &[]domain.Comment{{Key: correctCommentKey, ThreadKey: correctThreadKey, Contributor: correctContributor, Comment: editedComment}},
			expectedError:    nil,
		},
	}

	opt := cmpopts.IgnoreFields(domain.Comment{}, "CreatedAt", "UpdatedAt")
	for _, c := range cases {
		comments, err := commentApplication.EditComment(&c.input)
		if diff := cmp.Diff(comments, c.expectedComments, opt); diff != "" {
			t.Errorf(
				"comment application, edit different comments, name: %s, diff: %s, want: %v, got: %v",
				c.name, diff, c.expectedComments, comments,
			)
		}
		if !isSameError(err, c.expectedError) {
			t.Errorf(
				"comment application, edit different error, name: %s, want: %s, got: %s",
				c.name, c.expectedError, err,
			)
		}
	}
}

func TestCommentApp_DeleteComment(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
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
	threadRepository := mock_repository.NewMockThreadRepository(mockCtrl)
	commentRepository := mock_repository.NewMockCommentRepository(mockCtrl)

	var threadKey string
	var commentKey string
	threadRepository.EXPECT().GetByKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*domain.Thread, error) {
			if threadKey == wrongThreadKey {
				return nil, appError.ErrNotFound
			}
			return &thread, nil
		},
	)
	threadRepository.EXPECT().Update(gomock.AssignableToTypeOf(&domain.Thread{})).AnyTimes().DoAndReturn(
		func(thread *domain.Thread) (*domain.Thread, error) {
			return thread, nil
		},
	)
	commentRepository.EXPECT().GetByKey(gomock.AssignableToTypeOf(commentKey)).AnyTimes().DoAndReturn(
		func(commentKey string) (*domain.Comment, error) {
			if commentKey == wrongCommentKey {
				return nil, appError.ErrNotFound
			}
			return &comment, nil
		},
	)
	commentRepository.EXPECT().Delete(gomock.AssignableToTypeOf(&domain.Comment{})).AnyTimes().DoAndReturn(
		func(comment *domain.Comment) error {
			return nil
		},
	)
	commentRepository.EXPECT().GetAllByKey(gomock.AssignableToTypeOf(threadKey)).AnyTimes().DoAndReturn(
		func(threadKey string) (*[]domain.Comment, error) {
			if threadKey == wrongThreadKey {
				return nil, appError.ErrNotFound
			}
			return &comments, nil
		},
	)

	commentApplication := application.NewCommentApplication(threadRepository, commentRepository)

	//
	// exucute
	//
	cases := []struct {
		name             string
		input            params.DeleteCommentAppLayerParam
		expectedComments *[]domain.Comment
		expectedError    error
	}{
		{
			name: "間違ったスレッドキーの場合失敗する",
			input: params.DeleteCommentAppLayerParam{ThreadKey: wrongThreadKey},
			expectedComments: nil,
			expectedError: appError.ErrNotFound,
		},
		{
			name: "間違ったコメントキーの場合失敗する",
			input: params.DeleteCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: wrongCommentKey},
			expectedComments: nil,
			expectedError: appError.ErrNotFound,
		},
		{
			name: "違うユーザーは削除できない",
			input: params.DeleteCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: correctCommentKey, Contributor: wrongContributor},
			expectedComments: nil,
			expectedError: &appError.BadRequest{},
		},
		{
			name: "正常なパラメータであれば成功する",
			input: params.DeleteCommentAppLayerParam{ThreadKey: correctThreadKey, CommentKey: correctCommentKey, Contributor: correctContributor},
			expectedComments: &comments,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		comments, err := commentApplication.DeleteComment(&c.input)
		if comments != c.expectedComments {
			t.Errorf(
				"comment application, delete different comments, name: %s, want: %v, got: %v",
				c.name, c.expectedComments, comments,
			)
		}
		if !isSameError(err, c.expectedError) {
			t.Errorf(
				"comment application, delete different error, name: %s, want: %s, got: %s",
				c.name, c.expectedError, err,
			)
		}
	}
}
