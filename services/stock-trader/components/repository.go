package stocktrader

import (
	"context"
	sql "github.com/vickeyshrestha/sharing-services/drivers/postgres"
)

type RepositoryClient interface {
	GetStatus(ctx context.Context)
}

func NewRepositoryClient(databaseUserName, databasePassword, databaseName string) (RepositoryClient, error) {

	pstgrsDriver, err := sql.NewPostgresDbConnection(ApplicationConfiguration.DatabaseHost, databaseUserName, databasePassword, databaseName, ApplicationConfiguration.databasePort)
	if err != nil {
		return nil, err
	}
	return &postgresClient{postgresDriver: pstgrsDriver}, nil
}
