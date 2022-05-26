package repository

import "github.com/jumpei00/board/backend/app/domain"

type ThreadRepository interface {
	GetAll() []*domain.Thread
	GetByKey(threadKey string) *domain.Thread
	Insert(thread *domain.Thread)
	Update(thread *domain.Thread)
	Delete(threadKey string)
}