package users

import (
	"context"
	"omics/pkg/models"
)

type FindArgs struct {
	Username *string
	Email    *string
}

type Repository interface {
	Find(ctx context.Context, args FindArgs) (*models.User, error)
}
