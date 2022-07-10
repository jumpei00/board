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
	"github.com/jumpei00/board/backend/app/application/params"
	"golang.org/x/crypto/bcrypt"
)

func TestUserApp_GetUserByID(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userReposotory := mock_repository.NewMockUserRepository(mockCtrl)

	type mockField struct {
		userRepository *mock_repository.MockUserRepository
	}

	field := mockField{
		userRepository: userReposotory,
	}

	userApplication := application.NewUserApplication(userReposotory)

	//
	// execute
	//
	var (
		correntID = "correct-id"
		wrongID   = "wrong-id"
		user      = domain.User{ID: correntID}
	)
	cases := []struct {
		testCase      string
		input         string
		prepare       func(*mockField)
		expectedUser  *domain.User
		expectedError error
	}{
		{
			testCase: "IDに対するユーザー見つからない場合は失敗する",
			input:    wrongID,
			prepare: func(mf *mockField) {
				mf.userRepository.EXPECT().GetByID(wrongID).Return(nil, appError.ErrNotFound)
			},
			expectedUser:  nil,
			expectedError: appError.ErrNotFound,
		},
		{
			testCase: "IDに対するユーザーが見つかった場合は成功する",
			input:    correntID,
			prepare: func(mf *mockField) {
				mf.userRepository.EXPECT().GetByID(correntID).Return(&user, nil)
			},
			expectedUser:  &user,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			user, err := userApplication.GetUserByID(c.input)

			if user != c.expectedUser {
				t.Errorf("different user.\nwant: %s\ngot: %s", c.expectedUser, user)
			}
			if !isSameError(err, c.expectedError) {
				t.Errorf("different error.\nwant: %s\ngot: %s", c.expectedError, err)
			}
		})
	}
}

func TestUserApp_GetUserByUsername(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userReposotory := mock_repository.NewMockUserRepository(mockCtrl)

	type mockField struct {
		userRepository *mock_repository.MockUserRepository
	}

	field := mockField{
		userRepository: userReposotory,
	}

	userApplication := application.NewUserApplication(userReposotory)

	//
	// execute
	//
	var (
		correntUsername = "correct-username"
		wrongUsername   = "wrong-username"
		user            = domain.User{Username: correntUsername}
	)
	cases := []struct {
		testCase      string
		input         string
		prepare       func(*mockField)
		expectedUser  *domain.User
		expectedError error
	}{
		{
			testCase: "ユーザー名に対するユーザーが存在しない場合は失敗する",
			input:    wrongUsername,
			prepare: func(mf *mockField) {
				mf.userRepository.EXPECT().GetByUsername(wrongUsername).Return(nil, appError.ErrNotFound)
			},
			expectedUser:  nil,
			expectedError: appError.ErrNotFound,
		},
		{
			testCase: "ユーザー名に対するユーザーが存在する場合は成功する",
			input:    correntUsername,
			prepare: func(mf *mockField) {
				mf.userRepository.EXPECT().GetByUsername(correntUsername).Return(&user, nil)
			},
			expectedUser:  &user,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			user, err := userApplication.GetUserByUsername(c.input)

			if user != c.expectedUser {
				t.Errorf("different user.\nwant: %s.\ngot: %s", c.expectedUser, user)
			}
			if !isSameError(err, c.expectedError) {
				t.Errorf("different error.\nwant: %s\ngot: %s", c.expectedError, err)
			}
		})
	}
}

func TestUserApp_CreateUser(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userReposotory := mock_repository.NewMockUserRepository(mockCtrl)

	type mockField struct {
		userRepository *mock_repository.MockUserRepository
	}

	field := mockField{
		userRepository: userReposotory,
	}

	userApplication := application.NewUserApplication(userReposotory)

	//
	// execute
	//
	var (
		newUsername     = "new-username"
		existedUsername = "existed-username"
		password        = "password"
	)
	cases := []struct {
		testCase      string
		input         params.UserSignUpApplicationLayerParam
		prepare       func(*mockField)
		expectedUser  *domain.User
		expectedError error
	}{
		{
			testCase: "既に登録済みのユーザーの場合は失敗する",
			input:    params.UserSignUpApplicationLayerParam{Username: existedUsername, Password: password},
			prepare: func(mf *mockField) {
				mf.userRepository.EXPECT().GetByUsername(existedUsername).Return(nil, nil)
			},
			expectedUser:  nil,
			expectedError: &appError.BadRequest{},
		},
		{
			testCase: "新規ユーザーの場合は成功する",
			input:    params.UserSignUpApplicationLayerParam{Username: newUsername, Password: password},
			prepare: func(mf *mockField) {
				mf.userRepository.EXPECT().GetByUsername(newUsername).Return(nil, appError.ErrNotFound)
				mf.userRepository.EXPECT().Insert(gomock.Any()).DoAndReturn(
					func(user *domain.User) (*domain.User, error) {
						return user, nil
					},
				)
			},
			expectedUser:  &domain.User{Username: newUsername},
			expectedError: nil,
		},
	}

	opt := cmpopts.IgnoreFields(domain.User{}, "ID", "Password", "CreatedAt", "UpdatedAt")

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			user, err := userApplication.CreateUser(&c.input)

			if diff := cmp.Diff(user, c.expectedUser, opt); diff != "" {
				t.Errorf("different user.\ndiff: %s, want: %v, got: %v", diff, c.expectedUser, user)
			}
			if !isSameError(err, c.expectedError) {
				t.Errorf("different error.\nwant: %v\ngot: %v", c.expectedError, err)
			}
		})
	}
}

func TestUserApp_ValidateUser(t *testing.T) {
	//
	// setup
	//
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userReposotory := mock_repository.NewMockUserRepository(mockCtrl)

	type mockField struct {
		userRepository *mock_repository.MockUserRepository
	}

	field := mockField{
		userRepository: userReposotory,
	}

	userApplication := application.NewUserApplication(userReposotory)

	//
	// execute
	//
	var (
		id              = "1"
		correctUsername = "correct-username"
		correctPassword = "correct-password"
		wrongUsername   = "wrong-username"
		wrongPassword   = "wrong-password"
		user            = domain.User{ID: id, Username: correctUsername, Password: generatePassword(correctPassword)}
	)
	cases := []struct {
		testCase      string
		input         params.UserSignInApplicationLayerParam
		prepare       func(*mockField)
		expectedUser  *domain.User
		expectedError error
	}{
		{
			testCase: "違うユーザー名の場合は失敗する",
			input:    params.UserSignInApplicationLayerParam{Username: wrongUsername, Password: correctPassword},
			prepare: func(mf *mockField) {
				mf.userRepository.EXPECT().GetByUsername(wrongUsername).Return(nil, appError.ErrNotFound)
			},
			expectedUser:  nil,
			expectedError: &appError.BadRequest{},
		},
		{
			testCase: "違うパスワードの場合は失敗する",
			input:    params.UserSignInApplicationLayerParam{Username: correctUsername, Password: wrongPassword},
			prepare: func(mf *mockField) {
				mf.userRepository.EXPECT().GetByUsername(correctUsername).Return(&user, nil)
			},
			expectedUser:  nil,
			expectedError: &appError.BadRequest{},
		},
		{
			testCase: "ログイン可能な場合は成功する",
			input:    params.UserSignInApplicationLayerParam{Username: correctUsername, Password: correctPassword},
			prepare: func(mf *mockField) {
				mf.userRepository.EXPECT().GetByUsername(correctUsername).Return(&user, nil)
			},
			expectedUser:  &user,
			expectedError: nil,
		},
	}

	opt := cmpopts.IgnoreFields(domain.User{}, "CreatedAt", "UpdatedAt")

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			c.prepare(&field)
			user, err := userApplication.ValidateUser(&c.input)

			if diff := cmp.Diff(user, c.expectedUser, opt); diff != "" {
				t.Errorf("different user.\ndiff: %s\nwant: %v\ngot: %v", diff, c.expectedUser, user)
			}
			if !isSameError(err, c.expectedError) {
				t.Errorf("different error.\nwant: %v\ngot: %v", c.expectedError, err)
			}
		})
	}
}

func generatePassword(pass string) []byte {
	buf := bytes.NewBufferString(pass)
	password, _ := bcrypt.GenerateFromPassword(buf.Bytes(), 10)
	return password
}
