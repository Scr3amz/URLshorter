package urls

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, urls *URLs) error
}