package stocktrader

import (
	"context"
	sql "github.com/vickeyshrestha/sharing-services/drivers/postgres"
	pb "github.com/vickeyshrestha/sharing-services/protobuf/stock_trader"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"strconv"
)

type postgresClient struct {
	postgresDriver sql.Driver
}

func (p postgresClient) GetActions(ctx context.Context) (*pb.ActionsResponse, error) {
	// TODO: add the actions here. For time being, let us only query the most recent record. This will change later.

	var uid string
	var symbol string
	var company string
	var actionType string
	var currentValue string
	var addedTimestamp string
	var isDeleted string

	getRecentAction := `SELECT * FROM  
						stocktrader.actions s1 WHERE added_timestamp = (SELECT MAX(added_timestamp) FROM stocktrader.actions s2 WHERE s1.uid = s2.uid)
						ORDER BY uid, added_timestamp;`

	rows, err := p.postgresDriver.Sql.Query(getRecentAction)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&uid, &symbol, &company, &actionType, &currentValue, &addedTimestamp, &isDeleted)
	}

	isDeletedBool, err := strconv.ParseBool(isDeleted)
	currentValueFloat64, _ := strconv.ParseFloat(currentValue, 64)
	uidInt64, _ := strconv.ParseInt(uid, 10, 64)
	return &pb.ActionsResponse{
		Uid:            wrapperspb.Int64(uidInt64),
		Symbol:         wrapperspb.String(symbol),
		Company:        wrapperspb.String(company),
		ActionType:     wrapperspb.String(actionType),
		CurrentValue:   wrapperspb.Double(currentValueFloat64),
		AddedTimestamp: wrapperspb.String(addedTimestamp),
		IsDeleted:      wrapperspb.Bool(isDeletedBool),
	}, nil
}

func (p postgresClient) GetStatus(ctx context.Context) (*pb.Status, error) {
	return &pb.Status{ServiceName: serviceName, Message: statusAlive}, nil
}
