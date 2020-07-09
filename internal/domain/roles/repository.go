package roles

import "context"

type Repository interface {
	FindByCode(ctx context.Context, code string) (*Role, error)
}
