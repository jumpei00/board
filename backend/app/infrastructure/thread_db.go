package infrastructure

import "github.com/jumpei00/board/backend/app/domain"

type ThreadDB struct{}

func NewThreadDB() *ThreadDB {
	return &ThreadDB{}
}

func (t *ThreadDB) GetAll() ([]*domain.Thread, error) {
	return []*domain.Thread{}, nil
}

func (t *ThreadDB) GetByKey(threadKey string) (*domain.Thread, error) {
	return &domain.Thread{}, nil
}

func (t *ThreadDB) Insert(thread *domain.Thread) (*domain.Thread, error) {
	return &domain.Thread{}, nil
}

func (t *ThreadDB) Update(thread *domain.Thread) (*domain.Thread, error) {
	return &domain.Thread{}, nil
}

func (t *ThreadDB) Delete(threadKey string) error {
	return nil
}