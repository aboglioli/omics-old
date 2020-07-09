package users

import "omics/pkg/models"

// User is an aggregate root
type User struct {
	models.Base
	Username     string
	Email        string
	Name         string
	Lastname     string
	Gender       string
	Publications int
	Suscription  bool
}

func (u *User) IsAuthor() bool {
	return u.Publications > 0
}

func (u *User) IsSuscribed() bool {
	return u.Suscription
}
