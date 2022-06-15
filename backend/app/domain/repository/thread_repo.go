package repository

import "github.com/jumpei00/board/backend/app/domain"

type ThreadRepository interface {
	GetAll() (*[]domain.Thread, error)
	GetByKey(threadKey string) (*domain.Thread, error)
	Insert(thread *domain.Thread) (*domain.Thread, error)
	Update(thread *domain.Thread) (*domain.Thread, error)
	Delete(thread *domain.Thread) error
}