package infrastructure

import (
	"github.com/jumpei00/board/backend/app/domain"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type VisitorDB struct {
	db *gorm.DB
}

func NewVisitorDB(dbPool *gorm.DB) *VisitorDB {
	return &VisitorDB{
		db: dbPool,
	}
}

func (v *VisitorDB) Get() (*domain.Visitors, error) {
	var visitors domain.Visitors

	if err := v.db.First(&visitors).Error; err != nil {
		if errors.Cause(err) == gorm.ErrRecordNotFound {
			logger.Error("not found visitor stat", "error", err)
			return nil, appError.NewErrNotFound("not found visitor stat -> error: %s", err)
		}
		return nil, errors.WithStack(err)
	}

	return &visitors, nil
}

func (v *VisitorDB) Update(visitors *domain.Visitors) (*domain.Visitors, error) {
	err := v.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(visitors).Error; err != nil {
			logger.Error("visitor stat update failed", "error", err, "visitor_stat", visitors)
			return err
		}
		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return visitors, nil
}