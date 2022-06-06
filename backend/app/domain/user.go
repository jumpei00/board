package domain

import (
	"bytes"

	"github.com/jumpei00/board/backend/app/params"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	username string
	password []byte
}

func NewUser(param *params.UserSignUpDomainLayerParam) (*User, error) {
	newUser := &User{username: param.Username}

	if err := newUser.setHashingPassword(param.Password); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *User) Validate(param *params.UserSignInDomainLayerParam) error {
	return u.validatePassword(param.Password)
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() []byte {
	return u.password
}

func (u *User) setHashingPassword(password string) error {
	buf := bytes.NewBufferString(password)
	pass, err := bcrypt.GenerateFromPassword(buf.Bytes(), 10)
	
	if err != nil {
		return err
	}

	u.password = pass
	return nil
}

func (u *User) validatePassword(password string) error {
	buf := bytes.NewBufferString(password)
	return bcrypt.CompareHashAndPassword(u.password, buf.Bytes())
}