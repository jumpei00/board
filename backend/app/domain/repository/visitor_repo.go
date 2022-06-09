package repository

import "github.com/jumpei00/board/backend/app/domain"

type VisitorRepository interface {
	Get() (*domain.Visitors, error)
	Update(visitors *domain.Visitors) (*domain.Visitors, error)
}