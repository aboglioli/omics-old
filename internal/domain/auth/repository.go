package auth

import (
	"context"
	"omics/pkg/models"
)

type UserRepository interface {
	FindByID(ctx context.Context, id models.ID) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) error
}

type AuthRepository interface {
	FindByTokenID(ctx context.Context, tokenID string) (*User, error)
	Save(ctx context.Context, tokenID string, user *User) error
	Delete(ctx context.Context, tokenID string) error
}
