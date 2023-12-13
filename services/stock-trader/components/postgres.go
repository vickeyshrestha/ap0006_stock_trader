package stocktrader

import (
	"context"
	sql "github.com/vickeyshrestha/sharing-services/drivers/postgres"
	pb "github.com/vickeyshrestha/sharing-services/protobuf/stock_trader"
)

type postgresClient struct {
	postgresDriver sql.Driver
}

func (p postgresClient) GetActions(ctx context.Context) (*pb.ActionsResponse, error) {
	// TODO: add the actions here
	return &pb.ActionsResponse{}, nil
}

func (p postgresClient) GetStatus(ctx context.Context) (*pb.Status, error) {
	return &pb.Status{ServiceName: serviceName, Message: statusAlive}, nil
}
