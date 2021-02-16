package stocktrader

import (
	"context"
	sql "github.com/vickeyshrestha/sharing-services/drivers/postgres"
)

type postgresClient struct {
	postgresDriver sql.Driver
}

func (p postgresClient) GetStatus(ctx context.Context) {
	panic("implement me")
}
