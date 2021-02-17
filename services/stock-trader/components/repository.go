package stocktrader

import (
	"context"
	sql "github.com/vickeyshrestha/sharing-services/drivers/postgres"
)

type RepositoryClient interface {
	GetStatus(ctx context.Context)
}

func NewRepositoryClient(databaseUserName, databasePassword, databaseName string, config Configuration) (RepositoryClient, error) {

	pstgrsDriver, err := sql.NewPostgresDbConnection(config.DatabaseHost, databaseUserName, databasePassword, databaseName, config.DatabasePort)
	if err != nil {
		return nil, err
	}
	return &postgresClient{postgresDriver: pstgrsDriver}, nil
}
