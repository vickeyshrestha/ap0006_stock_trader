package stocktrader

import (
	"context"
)

type RepositoryClient interface {
	GetStatus(ctx context.Context)
}
