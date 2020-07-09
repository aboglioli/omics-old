package auth

import (
	"time"

	"omics/pkg/models"

	"github.com/google/uuid"
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

type Token struct {
	ID   string
	User *User
}

func NewToken() *Token {
	return &Token{
		ID: uuid.New().String(),
	}
}

func (t *Token) Encode(enc Encoder) (string, error) {
	token, err := enc.Encode(t.ID)
	if err != nil {
		return "", ErrNull
	}

	return token, nil
}

func DecodeToken(enc Encoder, tokenStr string) (*Token, error) {
	tokenID, err := enc.Decode(tokenStr)
	if err != nil {
		return nil, ErrNull
	}

	return &Token{
		ID: tokenID,
	}, nil
}
