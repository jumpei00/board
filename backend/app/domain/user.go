package domain

type User struct {
	username string
	password string
}

func NewUser(username, password string) *User {
	return &User{
		username: username,
		password: password,
	}
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}