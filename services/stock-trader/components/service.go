package stocktrader

import (
	"context"
	pb "github.com/vickeyshrestha/sharing-services/protobuf/stock_trader"
	"google.golang.org/protobuf/types/known/emptypb"
)

type service struct {
	repositoryClient RepositoryClient
}

func (s service) GetStatus(ctx context.Context, empty *emptypb.Empty) (*pb.Status, error) {
	response, err := s.repositoryClient.GetStatus(ctx)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s service) GetActions(ctx context.Context, empty *emptypb.Empty) (*pb.ActionsResponse, error) {
	response, err := s.repositoryClient.GetActions(ctx)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewStockTraderService(repoClt RepositoryClient) pb.StockTraderServer {
	s := service{repositoryClient: repoClt}
	return &s
}
