package auth

import "github.com/google/uuid"

type Token struct {
	ID string
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
