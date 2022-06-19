package repository

import "github.com/jumpei00/board/backend/app/domain"

type VisitorRepository interface {
	Get() (*domain.Visitor, error)
	Update(visitors *domain.Visitor) (*domain.Visitor, error)
}