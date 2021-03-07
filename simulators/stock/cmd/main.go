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
	natsHost := os.Getenv("natsHost")
	natsPort := os.Getenv("natsPort")
	if natsHost == "" || natsPort == "" {
		fmt.Println(stocks.ErrNatsHostPortRequired)
		return
	}
	natsUrl := fmt.Sprintf("%s:%s", natsHost, natsPort)
	fmt.Println(fmt.Sprintf(stocks.InfoNatsConnection, natsUrl))
	encodedNatsConnection, err := nats.NewNatsConnectionClient(natsUrl)
	if err != nil {
		fmt.Println(fmt.Sprintf(stocks.ErrFailedToInitializeNatsClient, err))
		return
	}
	defer encodedNatsConnection.Close()

	personChanSend := make(chan *stocks.TimeSeries)
	encodedNatsConnection.BindSendChan(stocks.SimNatsTopic, personChanSend)

	var mNasdaq map[string]float64
	mNasdaq = make(map[string]float64)

	var mNYSE map[string]float64
	mNYSE = make(map[string]float64)

	var mLondonSE map[string]float64
	mLondonSE = make(map[string]float64)

	i := 0
	for {
		mNasdaq["CompanyA"] = rand.Float64()
		mNasdaq["CompanyB"] = rand.Float64()
		mNasdaq["CompanyC"] = rand.Float64()
		mNasdaq["CompanyD"] = rand.Float64()
		mNasdaq["CompanyE"] = rand.Float64()
		sNasdaq := stocks.Stock{
			Companies: mNasdaq,
		}

		mNYSE["CompanyA"] = rand.Float64()
		mNYSE["CompanyB"] = rand.Float64()
		mNYSE["CompanyC"] = rand.Float64()
		mNYSE["CompanyD"] = rand.Float64()
		mNYSE["CompanyE"] = rand.Float64()
		sNYSE := stocks.Stock{
			Companies: mNYSE,
		}

		mLondonSE["CompanyA"] = rand.Float64()
		mLondonSE["CompanyB"] = rand.Float64()
		mLondonSE["CompanyC"] = rand.Float64()
		mLondonSE["CompanyD"] = rand.Float64()
		mLondonSE["CompanyE"] = rand.Float64()
		sLondonSE := stocks.Stock{
			Companies: mLondonSE,
		}

		// Create instance of type Request with Id set to
		// the current value of i
		reqNASDAQ := stocks.TimeSeries{
			ExchangeName: "NASDAQ Stock Exchange",
			Stocks:       sNasdaq,
		}
		reqNYSE := stocks.TimeSeries{
			ExchangeName: "NY Stock Exchange",
			Stocks:       sNYSE,
		}
		reqLonSE := stocks.TimeSeries{
			ExchangeName: "London Stock Exchange",
			Stocks:       sLondonSE,
		}

		// Just send to the channel! :)
		log.Printf("Sending request %s, data: %v", reqNASDAQ.ExchangeName, reqNASDAQ.Stocks)
		log.Printf("Sending request %s, data: %v", reqNYSE.ExchangeName, reqNYSE.Stocks)
		log.Printf("Sending request %s, data: %v", reqLonSE.ExchangeName, reqLonSE.Stocks)

		personChanSend <- &reqNASDAQ
		personChanSend <- &reqNYSE
		personChanSend <- &reqLonSE

		// Pause and increment counter
		time.Sleep(time.Second * 1)
		i = i + 1

		stocks.ClearMap(mNasdaq)
		stocks.ClearMap(mNYSE)
		stocks.ClearMap(mLondonSE)
	}

}
