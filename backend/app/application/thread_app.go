package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
)

type ThreadApplication interface {
	GetAllThread() []*domain.Thread
	GetByThreadKey(threadKey string) *domain.Thread
	CreateThread()
	EditThread()
	DeleteThread(threadKey string)
}

type threadApplication struct {
	threadRepo repository.ThreadRepository
}

func NewThreadApplication(tr repository.ThreadRepository) *threadApplication {
	return &threadApplication{
		threadRepo: tr,
	}
}

func (t *threadApplication) GetAllThread() []*domain.Thread {
	threads := t.threadRepo.GetAll()
	return threads
}

func (t *threadApplication) GetByThreadKey(threadKey string) *domain.Thread {
	return t.GetByThreadKey(threadKey)
}

func (t *threadApplication) CreateThread() {}

func (t *threadApplication) EditThread() {}

func (t *threadApplication) DeleteThread(threadKey string) {}
