package auth

import "context"

type Service interface {
	GetToken(ctx context.Context) (*Token, error)
	GetUser(ctx context.Context) (*User, error)
}

type service struct {
	tokenEncoder Encoder
	authRepo     AuthRepository
}

func (s *service) GetToken(ctx context.Context) (*Token, error) {
	tokenStr, ok := ctx.Value("authToken").(string)
	if !ok {
		return nil, ErrNull
	}

	token, err := DecodeToken(s.tokenEncoder, tokenStr)
	if err != nil {
		return nil, ErrNull
	}

	return token, nil
}

func (s *service) GetUser(ctx context.Context) (*User, error) {
	token, err := s.GetToken(ctx)
	if err != nil {
		return nil, ErrNull
	}

	user, err := s.authRepo.FindByTokenID(ctx, token.ID)
	if err != nil {
		return nil, ErrNull
	}

	return user, nil
}
