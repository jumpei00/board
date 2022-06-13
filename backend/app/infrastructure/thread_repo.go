package infrastructure

import (
	"github.com/jumpei00/board/backend/app/domain"
	"gorm.io/gorm"
)

type threadRepository struct{
	db *gorm.DB
}

func NewThreadRepository(dbSession *gorm.DB) *threadRepository {
	return &threadRepository{
		db: dbSession,
	}
}

func (t *threadRepository) GetAll() ([]*domain.Thread, error) {
	return []*domain.Thread{}, nil
}

func (t *threadRepository) GetByKey(threadKey string) (*domain.Thread, error) {
	return &domain.Thread{}, nil
}

func (t *threadRepository) Insert(thread *domain.Thread) (*domain.Thread, error) {
	return &domain.Thread{}, nil
}

func (t *threadRepository) Update(thread *domain.Thread) (*domain.Thread, error) {
	return &domain.Thread{}, nil
}

func (t *threadRepository) Delete(thread *domain.Thread) error {
	return nil
}