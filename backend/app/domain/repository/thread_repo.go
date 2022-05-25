package repository

import "github.com/jumpei00/board/backend/app/domain"

type ThreadRepository interface {
	GetAll() []*domain.Thread
	Create(t *domain.Thread)
	Edit(t *domain.Thread)
	Delete(threadKey string)
}