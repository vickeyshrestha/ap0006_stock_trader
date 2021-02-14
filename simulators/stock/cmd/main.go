package main

import (
	"fmt"
	nats "github.com/vickeyshrestha/sharing-services/drivers/nats"
	stocks "godzilla/simulators/stock/components"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	natsUrl := os.Getenv("natsUrl")
	encodedNatsConnection, err := nats.NewNatsConnectionClient(natsUrl)
	if err != nil {
		fmt.Println(fmt.Sprintf(stocks.ErrFailedToInitializeNatsClient, err))
		return
	}
	defer encodedNatsConnection.Close()

	personChanSend := make(chan *stocks.TimeSeries)
	encodedNatsConnection.BindSendChan("request_subject", personChanSend)

	i := 0
	for {

		s := stocks.Stock{
			CompanyA: rand.Float64(),
			CompanyB: rand.Float64(),
			CompanyC: rand.Float64(),
			CompanyD: rand.Float64(),
			CompanyE: rand.Float64(),
			CompanyF: rand.Float64(),
		}

		// Create instance of type Request with Id set to
		// the current value of i
		req := stocks.TimeSeries{
			ExchangeName: "NASDAQ Stock Exchange",
			Stocks:       s,
		}

		// Just send to the channel! :)
		log.Printf("Sending request %s", req)

		personChanSend <- &req

		// Pause and increment counter
		time.Sleep(time.Second * 1)
		i = i + 1
	}

}
