package stocktrader

import (
	"context"
	pb "github.com/vickeyshrestha/sharing-services/protobuf/stock_trader"
	"google.golang.org/protobuf/types/known/emptypb"
)

type service struct {
	// TODO: service components goes here
}

func (s service) GetStatus(ctx context.Context, empty *emptypb.Empty) (*pb.Status, error) {
	panic("implement me")
}

func (s service) GetActions(ctx context.Context, empty *emptypb.Empty) (*pb.ActionsResponse, error) {
	panic("implement me")
}

func NewStockTraderService() pb.StockTraderServer {
	s := service{}
	return &s
}
