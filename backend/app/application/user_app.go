package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
)

type UserApplication interface {
	GetUserByKey(key string) *domain.User
	CreateUser()
}

type userApplication struct {
	userRepo repository.UserRepository
}

func (u *userApplication) GetUserByKey(key string) *domain.User {
	return u.userRepo.GetByKey(key)
}

func (u *userApplication) CreateUser() {}