package stocktrader

import (
	"context"
	pb "github.com/vickeyshrestha/sharing-services/protobuf/stock_trader"
)

type RepositoryClient interface {
	GetStatus(ctx context.Context)
}

func NewRepositoryClient(databaseUserName, databasePassword, databaseName string) (RepositoryClient, error) {

	return &postgresClient{}, nil
}
