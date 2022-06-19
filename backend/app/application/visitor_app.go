package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
)

type VisitorApplication interface {
	GetVisitorsStat() (*domain.Visitors, error)
	CountupVisitors() (*domain.Visitors, error)
	ResetVisitors() (*domain.Visitors, error)
}

type visitorApplication struct {
	visitorRepo repository.VisitorRepository
}

func NewVisitorApplication(vr repository.VisitorRepository) *visitorApplication {
	return &visitorApplication{
		visitorRepo: vr,
	}
}

func (v *visitorApplication) GetVisitorsStat() (*domain.Visitors, error) {
	visitors, err := v.visitorRepo.Get()
	if err != nil {
		return nil, err
	}

	return visitors, nil
}

func (v *visitorApplication) CountupVisitors() (*domain.Visitors, error) {
	visitors, err := v.visitorRepo.Get()
	if err != nil {
		return nil, err
	}

	visitors.CoutupTodayVisitors()
	visitors.CountupSumVisitor()

	updatedVisitors, err := v.visitorRepo.Update(visitors)
	if err != nil {
		return nil, err
	}

	return updatedVisitors, nil
}

func (v *visitorApplication) ResetVisitors() (*domain.Visitors, error) {
	visitors, err := v.visitorRepo.Get()
	if err != nil {
		return nil, err
	}

	visitors.SetYesterdayVisitors(visitors.GetTodayVisitors())
	visitors.ResetTodayVisitors(0)

	updatedVisitors, err := v.visitorRepo.Update(visitors)
	if err != nil {
		return nil, err
	}

	return updatedVisitors, nil
}