package users

import (
	"context"

	"omics/pkg/errors"
)

type Service interface {
	Register(ctx context.Context, req *RegisterRequest) *UserDTO
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Register(ctx context.Context, req *RegisterRequest) (*UserDTO, error) {
	existingUser, err := s.repository.Find(ctx, FindArgs{
		Username: &req.Username,
	})
	if existingUser != nil || err == nil {
		return nil, errors.NewApplication("existing.username").SetPath("register")
	}

	existingUser, err = s.repository.Find(ctx, FindArgs{
		Email: &req.Email,
	})
	if existingUser != nil || err == nil {
		return nil, errors.NewApplication("existing.email").SetPath("register")
	}

	return nil, errors.New(errors.INTERNAL, "not implemented yet")
}
