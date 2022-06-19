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

func (v *VisitorDB) Get() (*domain.Visitor, error) {
	var visitor domain.Visitor

	if err := v.db.First(&visitor).Error; err != nil {
		if errors.Cause(err) == gorm.ErrRecordNotFound {
			logger.Error("not found visitor stat", "error", err)
			return nil, appError.NewErrNotFound("not found visitor stat -> error: %s", err)
		}
		return nil, errors.WithStack(err)
	}

	return &visitor, nil
}

func (v *VisitorDB) Update(visitor *domain.Visitor) (*domain.Visitor, error) {
	err := v.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(visitor).Error; err != nil {
			logger.Error("visitor stat update failed", "error", err, "visitor_stat", visitor)
			return err
		}
		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return visitor, nil
}