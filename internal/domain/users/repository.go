package users

import "context"

type Repository interface {
	FindByID(ctx context.Context, userID string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) error
}
