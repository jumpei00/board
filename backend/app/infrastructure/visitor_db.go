package infrastructure

import "github.com/jumpei00/board/backend/app/domain"

type VisitorDB struct {}

func NewVisitorDB() *VisitorDB {
	return &VisitorDB{}
}

func (v *VisitorDB) Get() *domain.Visitors {
	return &domain.Visitors{}
}

func (v *VisitorDB) CountUp() {}