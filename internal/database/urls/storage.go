package urls

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, u *URLs) error
	FindLong(ctx context.Context,  u URLs) (URLs, error)
	FindShort(ctx context.Context,  u URLs) (URLs, error)
}