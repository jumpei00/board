package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
)

type VisitorApplication interface {
	GetVisitorsStat() *domain.Visitors
	CountUpVisitors()
}

type visitorApplication struct {
	visitorRepo repository.VisitorRepository
}

func NewVisitorApplication(vr repository.VisitorRepository) *visitorApplication {
	return &visitorApplication{
		visitorRepo: vr,
	}
}

func (v *visitorApplication) GetVisitorsStat() *domain.Visitors {
	return v.visitorRepo.Get()
}

func (v *visitorApplication) CountUpVisitors() {}