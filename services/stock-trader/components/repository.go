package stocktrader

import (
	"context"
	sql "github.com/vickeyshrestha/sharing-services/drivers/postgres"
	pb "github.com/vickeyshrestha/sharing-services/protobuf/stock_trader"
	"os"
	"strconv"
)

type RepositoryClient interface {
	GetStatus(ctx context.Context) (*pb.Status, error)
}

func NewRepositoryClient(databaseUserName, databasePassword, databaseName string, config Configuration) (RepositoryClient, error) {
	var dbHostName = os.Getenv("dbHost")
	var dbPort = 0
	var err error

	if os.Getenv("dbPort") != "" {
		dbPort, err = strconv.Atoi(os.Getenv("dbPort"))
		if err != nil {
			return nil, err
		}
	}

	if len(dbHostName) == 0 {
		dbHostName = config.DatabaseHost
	}
	if dbPort == 0 {
		dbPort = config.DatabasePort
	}

	pstgrsDriver, err := sql.NewPostgresDbConnection(dbHostName, databaseUserName, databasePassword, databaseName, dbPort)
	if err != nil {
		return nil, err
	}
	return &postgresClient{postgresDriver: *pstgrsDriver}, nil
}
