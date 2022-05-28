package infrastructure

import "github.com/jumpei00/board/backend/app/domain"

type UserDB struct{}

func NewUserDB() *UserDB {
	return &UserDB{}
}

func (u *UserDB) GetByKey(key string) *domain.User {
	return &domain.User{}
}

func (u *UserDB) Insert(user *domain.User) {}
