package auth

import (
	"strconv"

	"omics/internal/domain/configuration"

	"golang.org/x/crypto/bcrypt"
)

type PasswordCrypter interface {
	Hash(plainPassword string) (string, error)
	Compare(hashedPassword, plainPassword string) bool
}

type bcrypterCrypter struct {
	configRepo configuration.Repository
}

func DefaultCrypter(configRepo configuration.Repository) PasswordCrypter {
	return &bcrypterCrypter{
		configRepo: configRepo,
	}
}

func (c *bcrypterCrypter) Hash(plainPassword string) (string, error) {
	cost, err := strconv.Atoi(c.configRepo.Get("password_crypter.cost"))
	if err != nil {
		return "", ErrNull
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), cost)
	if err != nil {
		return "", ErrNull
	}
	return string(hash), nil
}

func (c *bcrypterCrypter) Compare(hashedPassword, plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return false
	}
	return true
}
