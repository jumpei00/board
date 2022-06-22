package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/jumpei00/board/backend/app/params"
	"github.com/pkg/errors"
)

type UserApplication interface {
	GetUserByID(id string) (*domain.User, error)
	GetUserByUsername(key string) (*domain.User, error)
	CreateUser(param *params.UserSignUpApplicationLayerParam) (*domain.User, error)
	ValidateUser(param *params.UserSignInApplicationLayerParam) (*domain.User, error)
}

type userApplication struct {
	userRepo repository.UserRepository
}

func NewUserApplication(ur repository.UserRepository) *userApplication {
	return &userApplication{
		userRepo: ur,
	}
}

func (u *userApplication) GetUserByID(id string) (*domain.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *userApplication) GetUserByUsername(username string) (*domain.User, error) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userApplication) CreateUser(param *params.UserSignUpApplicationLayerParam) (*domain.User, error) {
	// 登録しようとしているユーザー名は既に登録済みではないか調べる必要がある
	_, err := u.userRepo.GetByUsername(param.Username)

	// nilの場合は該当ユーザーが存在しているのでエラーを返す必要がある
	if err == nil {
		logger.Info("requesting username is already registered", "username", param.Username)
		return nil, appError.NewErrBadRequest(appError.Message().AlreadyUsernameExist, "username is already exist -> username: %s", param.Username)
	}

	// ユーザーが存在しない以外のエラーだった場合はエラーをそのまま返す
	if err != nil && errors.Cause(err) != appError.ErrNotFound {
		return nil, err
	}

	domainParam := params.UserSignUpDomainLayerParam{
		Username: param.Username,
		Password: param.Password,
	}

	newUser, err := domain.NewUser(&domainParam)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.Insert(newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userApplication) ValidateUser(param *params.UserSignInApplicationLayerParam) (*domain.User, error) {
	// ユーザー名かパスワードが違う場合はエラーメッセージを返す必要がある
	user, err := u.userRepo.GetByUsername(param.Username)
	if err != nil {
		if errors.Cause(err) == appError.ErrNotFound {
			logger.Info("requesting usrename is not found", "username", param.Username)
			return nil, appError.NewErrBadRequest(
				appError.Message().SignInBadRequest,
				"requesting username is not found -> username: %s", param.Username,
			)
		}
		return nil, err
	}

	domainParam := params.UserSignInDomainLayerParam{
		Username: param.Username,
		Password: param.Password,
	}

	if err := user.Validate(&domainParam); err != nil {
		logger.Info("requesting password is not matched", "error", err)
		return nil, appError.NewErrBadRequest(
			appError.Message().SignInBadRequest,
			"requesting password is not matched -> err: %s", err,
		)
	}

	return user, nil
}