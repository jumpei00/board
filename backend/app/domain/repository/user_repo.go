package repository

import "github.com/jumpei00/board/backend/app/domain"

type UserRepository interface {
	GetByKey(key string) *domain.User
	Insert(user *domain.User)
}