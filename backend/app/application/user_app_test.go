package application_test

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jumpei00/board/backend/app/application"
	"github.com/jumpei00/board/backend/app/domain"
	appError "github.com/jumpei00/board/backend/app/library/error"
	mock_repository "github.com/jumpei00/board/backend/app/mock/repository"
	"github.com/jumpei00/board/backend/app/params"
	"golang.org/x/crypto/bcrypt"
)

func TestUserApp_GetUserByID(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		correntID = "correct-id"
		wrongID   = "wrong-id"
		user      = domain.User{ID: correntID}
	)
	userReposotory := mock_repository.NewMockUserRepository(mockCtrl)

	// mock
	var id string
	userReposotory.EXPECT().GetByID(gomock.AssignableToTypeOf(id)).AnyTimes().DoAndReturn(
		func(id string) (*domain.User, error) {
			if id == wrongID {
				return nil, appError.ErrNotFound
			}
			return &user, nil
		},
	)

	userApplication := application.NewUserApplication(userReposotory)

	//
	// execute
	//
	cases := []struct {
		name          string
		input         string
		expectedUser  *domain.User
		expectedError error
	}{
		{
			name:          "異なるIDの場合は失敗する",
			input:         wrongID,
			expectedUser:  nil,
			expectedError: appError.ErrNotFound,
		},
		{
			name:          "正しいIDの場合は成功する",
			input:         correntID,
			expectedUser:  &user,
			expectedError: nil,
		},
	}
	for _, c := range cases {
		user, err := userApplication.GetUserByID(c.input)
		if user != c.expectedUser {
			t.Errorf(
				"user application, get user by id, different user, name: %s, want: %s, got: %s",
				c.name, c.expectedUser, user,
			)
		}
		if !isSameError(err, c.expectedError) {
			t.Errorf(
				"user application, get user by id, different error, name: %s, want: %s, got: %s",
				c.name, c.expectedError, err,
			)
		}
	}
}

func TestUserApp_GetUserByUsername(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		correntUsername = "correct-username"
		wrongUsername   = "wrong-username"
		user            = domain.User{Username: correntUsername}
	)
	userReposotory := mock_repository.NewMockUserRepository(mockCtrl)

	// mock
	var username string
	userReposotory.EXPECT().GetByUsername(gomock.AssignableToTypeOf(username)).AnyTimes().DoAndReturn(
		func(username string) (*domain.User, error) {
			if username == wrongUsername {
				return nil, appError.ErrNotFound
			}
			return &user, nil
		},
	)

	userApplication := application.NewUserApplication(userReposotory)

	//
	// execute
	//
	cases := []struct {
		name          string
		input         string
		expectedUser  *domain.User
		expectedError error
	}{
		{
			name:          "異なるユーザー名の場合は失敗する",
			input:         wrongUsername,
			expectedUser:  nil,
			expectedError: appError.ErrNotFound,
		},
		{
			name:          "正しいユーザー名の場合は成功する",
			input:         correntUsername,
			expectedUser:  &user,
			expectedError: nil,
		},
	}
	for _, c := range cases {
		user, err := userApplication.GetUserByUsername(c.input)
		if user != c.expectedUser {
			t.Errorf(
				"user application, get user by id, different user, name: %s, want: %s, got: %s",
				c.name, c.expectedUser, user,
			)
		}
		if !isSameError(err, c.expectedError) {
			t.Errorf(
				"user application, get user by id, different error, name: %s, want: %s, got: %s",
				c.name, c.expectedError, err,
			)
		}
	}
}

func TestUserApp_CreateUser(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		newUsername     = "new-username"
		existedUsername = "existed-username"
		password        = "password"
	)
	userReposotory := mock_repository.NewMockUserRepository(mockCtrl)

	// mock
	var username string
	userReposotory.EXPECT().GetByUsername(gomock.AssignableToTypeOf(username)).AnyTimes().DoAndReturn(
		func(username string) (*domain.User, error) {
			if username == existedUsername {
				return &domain.User{Username: existedUsername}, nil
			}
			return nil, appError.ErrNotFound
		},
	)
	userReposotory.EXPECT().Insert(gomock.AssignableToTypeOf(&domain.User{})).AnyTimes().DoAndReturn(
		func(user *domain.User) (*domain.User, error) {
			return user, nil
		},
	)

	userApplication := application.NewUserApplication(userReposotory)

	//
	// execute
	//
	cases := []struct {
		name          string
		input         params.UserSignUpApplicationLayerParam
		expectedUser  *domain.User
		expectedError error
	}{
		{
			name:          "既に登録済みのユーザーの場合は失敗する",
			input:         params.UserSignUpApplicationLayerParam{Username: existedUsername, Password: password},
			expectedUser:  nil,
			expectedError: &appError.BadRequest{},
		},
		{
			name:          "新規ユーザーの場合は成功する",
			input:         params.UserSignUpApplicationLayerParam{Username: newUsername, Password: password},
			expectedUser:  &domain.User{Username: newUsername},
			expectedError: nil,
		},
	}

	opt := cmpopts.IgnoreFields(domain.User{}, "ID", "Password", "CreatedAt", "UpdatedAt")
	for _, c := range cases {
		user, err := userApplication.CreateUser(&c.input)
		if diff := cmp.Diff(user, c.expectedUser, opt); diff != "" {
			t.Errorf(
				"user application, create user, different user, name: %s, diff: %s, want: %v, got: %v",
				c.name, diff, c.expectedUser, user,
			)
		}
		if !isSameError(err, c.expectedError) {
			t.Errorf(
				"user application, create user, different error, name: %s, want: %v, got: %v",
				c.name, c.expectedError, err,
			)
		}
	}
}

func TestUserApp_ValidateUser(t *testing.T) {
	// mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//
	// setup
	//
	var (
		id              = "1"
		correctUsername = "correct-username"
		correctPassword = "correct-password"
		wrongUsername   = "wrong-username"
		wrongPassword   = "wrong-password"
		user            = domain.User{ID: id, Username: correctUsername, Password: generatePassword(correctPassword)}
	)
	userReposotory := mock_repository.NewMockUserRepository(mockCtrl)

	// mock
	var username string
	userReposotory.EXPECT().GetByUsername(gomock.AssignableToTypeOf(username)).AnyTimes().DoAndReturn(
		func(username string) (*domain.User, error) {
			if username == wrongUsername {
				return nil, appError.ErrNotFound
			}
			return &user, nil
		},
	)

	userApplication := application.NewUserApplication(userReposotory)

	//
	// execute
	//
	cases := []struct {
		name          string
		input         params.UserSignInApplicationLayerParam
		expectedUser  *domain.User
		expectedError error
	}{
		{
			name:          "違うユーザー名の場合は失敗する",
			input:         params.UserSignInApplicationLayerParam{Username: wrongUsername, Password: correctPassword},
			expectedUser:  nil,
			expectedError: &appError.BadRequest{},
		},
		{
			name:          "違うパスワードの場合は失敗する",
			input:         params.UserSignInApplicationLayerParam{Username: correctUsername, Password: wrongPassword},
			expectedUser:  nil,
			expectedError: &appError.BadRequest{},
		},
		{
			name:          "正しいデータの場合は成功する",
			input:         params.UserSignInApplicationLayerParam{Username: correctUsername, Password: correctPassword},
			expectedUser:  &user,
			expectedError: nil,
		},
	}

	opt := cmpopts.IgnoreFields(domain.User{}, "CreatedAt", "UpdatedAt")
	for _, c := range cases {
		user, err := userApplication.ValidateUser(&c.input)
		if diff := cmp.Diff(user, c.expectedUser, opt); diff != "" {
			t.Errorf(
				"user application, validate user, different user, name: %s, diff: %s, want: %v, got: %v",
				c.name, diff, c.expectedUser, user,
			)
		}
		if !isSameError(err, c.expectedError) {
			t.Errorf(
				"user application, validate user, different error, name: %s, want: %v, got: %v",
				c.name, c.expectedError, err,
			)
		}
	}
}

func generatePassword(pass string) []byte {
	buf := bytes.NewBufferString(pass)
	password, _ := bcrypt.GenerateFromPassword(buf.Bytes(), 10)
	return password
}
