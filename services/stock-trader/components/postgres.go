package stocktrader

import (
	"context"
	sql "github.com/vickeyshrestha/sharing-services/drivers/postgres"
	pb "github.com/vickeyshrestha/sharing-services/protobuf/stock_trader"
)

type postgresClient struct {
	postgresDriver sql.Driver
}

func (p postgresClient) GetStatus(ctx context.Context) (*pb.Status, error) {
	panic("implement me")
}
