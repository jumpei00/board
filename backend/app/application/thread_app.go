package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
)

type ThreadApplication struct {
	threadRepo repository.ThreadRepository
}

func NewThreadApplication(tr repository.ThreadRepository) *ThreadApplication {
	return &ThreadApplication{
		threadRepo: tr,
	}
}

func (ta *ThreadApplication) GetAllThread() []*domain.Thread {
	threads := ta.threadRepo.GetAll()
	return threads
}