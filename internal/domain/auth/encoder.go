package auth

import (
	"omics/internal/domain/configuration"

	"github.com/dgrijalva/jwt-go"
)

type Encoder interface {
	Encode(tokenID string) (string, error)
	Decode(token string) (string, error)
}

type jwtEncoder struct {
	configRepo configuration.Repository
}

func DefaultEncoder(configRepo configuration.Repository) Encoder {
	return &jwtEncoder{
		configRepo: configRepo,
	}
}

func (e *jwtEncoder) Encode(tokenID string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": tokenID,
	})

	jwtSecret := e.configRepo.Get("jwtsecret")

	tokenStr, err := jwtToken.SignedString(jwtSecret)
	if err != nil {
		return "", ErrNull
	}

	return tokenStr, nil
}

func (e *jwtEncoder) Decode(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrNull
		}

		jwtSecret := e.configRepo.Get("jwtsecret")

		return jwtSecret, nil
	})
	if err != nil {
		return "", ErrNull
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", ErrNull
	}

	id, ok := claims["id"].(string)
	if !ok {
		return "", ErrNull
	}

	return id, nil
}
