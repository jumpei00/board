package infrastructure

import "github.com/jumpei00/board/backend/app/domain"

type ThreadDB struct{}

func NewThreadDB() *ThreadDB {
	return &ThreadDB{}
}

func (t *ThreadDB) GetAll() []*domain.Thread {
	return []*domain.Thread{}
}

func (t *ThreadDB) GetByKey(threadKey string) *domain.Thread {
	return &domain.Thread{}
}

func (t *ThreadDB) Insert(thread *domain.Thread) {}

func (t *ThreadDB) Update(thread *domain.Thread) {}

func (t *ThreadDB) Delete(threadKey string) {}