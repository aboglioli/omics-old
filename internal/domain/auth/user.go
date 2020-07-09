package auth

import (
	"omics/pkg/models"
	"time"
)

type Permission struct {
	Permission string
	Module     string
}

type Role struct {
	Code        string
	Permissions []Permission
}

// User is an aggregate root
type User struct {
	ID        models.ID
	Username  string
	Email     string
	Password  string
	Validated bool
	LastLogin time.Time
	Role      Role
	Provider  string
}

func NewUser(username, email string) *User {
	return &User{
		Username:  username,
		Email:     email,
		Validated: false,
	}
}

func (u *User) ChangePassword(crypter PasswordCrypter, oldPwd, newPwd string) error {
	if u.Password != "" && !u.ValidatePassword(crypter, oldPwd) {
		return ErrNull
	}

	hashedPassword, err := crypter.Hash(newPwd)
	if err != nil {
		return ErrNull
	}

	u.Password = hashedPassword

	return nil
}

func (u *User) ValidatePassword(crypter PasswordCrypter, passwd string) bool {
	return crypter.Compare(u.Password, passwd)
}

func (u *User) Validate() {
	u.Validated = true
}

func (u *User) Login() {
	u.LastLogin = time.Now()
}
