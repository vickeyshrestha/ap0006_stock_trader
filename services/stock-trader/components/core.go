package stocktrader

import (
	"fmt"
	nats "github.com/vickeyshrestha/sharing-services/drivers/nats"
	"log"
)

type TimeSeries struct {
	ExchangeName string
	Stocks       Stock
}

type Stock struct {
	Companies map[string]float64
}

/*
	core will start the business logic
*/
func Subscriber(natsUrl string) {
	encodedConnection, err := nats.NewNatsConnectionClient(natsUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer encodedConnection.Close()

	tsChanRecv := make(chan *TimeSeries)
	encodedConnection.BindRecvChan(SimNatsTopic, tsChanRecv)

	for {
		req := <-tsChanRecv
		log.Printf("Received request -> Exchange name: %s | Stocks: %v", req.ExchangeName, req.Stocks)
	}
}
