package repository

import "github.com/jumpei00/board/backend/app/domain"

type UserRepository interface {
	GetByID(ID string) (*domain.User, error)
	GetByUsername(key string) (*domain.User, error)
	Insert(user *domain.User) (*domain.User, error)
}
