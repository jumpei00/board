package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
	"github.com/jumpei00/board/backend/app/params"
)

type UserApplication interface {
	GetUserByKey(key string) (*domain.User, error)
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

func (u *userApplication) GetUserByKey(key string) (*domain.User, error) {
	user, err := u.userRepo.GetByKey(key)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userApplication) CreateUser(param *params.UserSignUpApplicationLayerParam) (*domain.User, error) {
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
	user, err := u.userRepo.GetByKey(param.Username)
	if err != nil {
		return nil, err
	}

	domainParam := params.UserSignInDomainLayerParam{
		Username: param.Username,
		Password: param.Password,
	}

	if err := user.Validate(&domainParam); err != nil {
		return nil, err
	}

	return user, nil
}