package infrastructure

import "github.com/jumpei00/board/backend/app/domain"

type VisitorDB struct {}

func NewVisitorDB() *VisitorDB {
	return &VisitorDB{}
}

func (v *VisitorDB) Get() (*domain.Visitors, error) {
	return &domain.Visitors{}, nil
}

func (v *VisitorDB) Update(visitors *domain.Visitors) (*domain.Visitors, error) {
	return nil, nil
}