package infrastructure

import "github.com/jumpei00/board/backend/app/domain"

type ThreadDB struct{}

func NewThreadDB() *ThreadDB {
	return &ThreadDB{}
}

func (tb *ThreadDB) GetAll() []*domain.Thread {
	return []*domain.Thread{}
}

func (tb *ThreadDB) Create(t *domain.Thread) {}

func (tb *ThreadDB) Edit(t *domain.Thread) {}

func (tb *ThreadDB) Delete(threadKey string) {}