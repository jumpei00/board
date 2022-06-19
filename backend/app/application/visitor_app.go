package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
)

type VisitorApplication interface {
	GetVisitorsStat() (*domain.Visitor, error)
	CountupVisitors() (*domain.Visitor, error)
	ResetVisitors() (*domain.Visitor, error)
}

type visitorApplication struct {
	visitorRepo repository.VisitorRepository
}

func NewVisitorApplication(vr repository.VisitorRepository) *visitorApplication {
	return &visitorApplication{
		visitorRepo: vr,
	}
}

func (v *visitorApplication) GetVisitorsStat() (*domain.Visitor, error) {
	visitor, err := v.visitorRepo.Get()
	if err != nil {
		return nil, err
	}

	return visitor, nil
}

func (v *visitorApplication) CountupVisitors() (*domain.Visitor, error) {
	visitor, err := v.visitorRepo.Get()
	if err != nil {
		return nil, err
	}

	visitor.CoutupTodayVisitor()
	visitor.CountupSumVisitor()

	updatedVisitors, err := v.visitorRepo.Update(visitor)
	if err != nil {
		return nil, err
	}

	return updatedVisitors, nil
}

func (v *visitorApplication) ResetVisitors() (*domain.Visitor, error) {
	visitor, err := v.visitorRepo.Get()
	if err != nil {
		return nil, err
	}

	visitor.SetYesterdayVisitor(visitor.GetTodayVisitor())
	visitor.ResetTodayVisitor(0)

	updatedVisitors, err := v.visitorRepo.Update(visitor)
	if err != nil {
		return nil, err
	}

	return updatedVisitors, nil
}