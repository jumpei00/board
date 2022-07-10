package domain

import (
	"bytes"
	"time"

	"github.com/google/uuid"
	"github.com/jumpei00/board/backend/app/domain/params"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"       gorm:"primaryKey;column:id"`
	Username  string    `json:"username" gorm:"column:username"`
	Password  []byte    `json:"-"        gorm:"column:password"`
	CreatedAt time.Time `json:"-"        gorm:"column:created_at"`
	UpdatedAt time.Time `json:"-"        gorm:"column:updated_at"`
}

func NewUser(param *params.UserSignUpDomainLayerParam) (*User, error) {
	newUser := &User{
		Username:  param.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newUser.setId()

	if err := newUser.setHashingPassword(param.Password); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *User) Validate(param *params.UserSignInDomainLayerParam) error {
	return u.validatePassword(param.Password)
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetPassword() []byte {
	return u.Password
}

func (u *User) FormatCreatedDate() string {
	return u.CreatedAt.Format("2006/01/02 15:04")
}

func (u *User) FormatUpdateDate() string {
	return u.UpdatedAt.Format("2006/01/02 15:04")
}

func (u *User) setId() {
	u.ID = uuid.New().String()
}

func (u *User) setHashingPassword(password string) error {
	buf := bytes.NewBufferString(password)
	pass, err := bcrypt.GenerateFromPassword(buf.Bytes(), 10)

	if err != nil {
		logger.Error("new user password generating failed", "error", err)
		return errors.WithStack(err)
	}

	u.Password = pass
	return nil
}

func (u *User) validatePassword(password string) error {
	buf := bytes.NewBufferString(password)
	return bcrypt.CompareHashAndPassword(u.Password, buf.Bytes())
}
