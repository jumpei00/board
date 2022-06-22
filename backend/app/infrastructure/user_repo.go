package infrastructure

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/library/logger"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserDB struct{
	db *gorm.DB
}

func NewUserDB(dbPool *gorm.DB) *UserDB {
	return &UserDB{
		db: dbPool,
	}
}

func (u *UserDB) GetByID(ID string) (*domain.User, error) {
	var user domain.User

	if err := u.db.Where(&domain.User{ID: ID}).First(&user).Error; err != nil {
		if errors.Cause(err) == gorm.ErrRecordNotFound {
			logger.Info("search user by ID, but not found", "ID", ID)
			return nil, appError.NewErrNotFound("no serarch user by ID -> ID: %s", ID)
		}
		logger.Error("search user by ID error", "error" , err, "ID", ID)
		return nil, errors.WithStack(err)
	}

	return &user, nil
}

func (u *UserDB) GetByUsername(username string) (*domain.User, error) {
	var user domain.User

	if err := u.db.Where(&domain.User{Username: username}).First(&user).Error; err != nil {
		if errors.Cause(err) == gorm.ErrRecordNotFound {
			logger.Info("search user by targeted username, but not found", "username", username)
			return nil, appError.NewErrNotFound("no search user by targeted username -> username: %s", username)
		}
		logger.Error("search user by targeted username error", "error", err, "username", username)
		return nil, errors.WithStack(err)
	}

	return &user, nil
}

func (u *UserDB) Insert(user *domain.User) (*domain.User, error) {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			logger.Error("new user create failed", "error", err, "user", user)
			return err
		}
		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}
