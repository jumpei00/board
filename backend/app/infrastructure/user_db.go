package infrastructure

import "github.com/jumpei00/board/backend/app/domain"

type UserDB struct{}

func NewUserDB() *UserDB {
	return &UserDB{}
}

func (u *UserDB) GetByKey(key string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (u *UserDB) Insert(user *domain.User) (*domain.User, error) {
	return &domain.User{}, nil
}
