package main

import (
	"fmt"
	"github.com/vickeyshrestha/sharing-services/drivers/nats"
	"log"
	"math/rand"
	"os"
	stocks "stockzilla/simulators/stock/components"
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
		mNasdaq["AAPL"] = rand.Float64()
		mNasdaq["ADP"] = rand.Float64()
		mNasdaq["AMAT"] = rand.Float64()
		mNasdaq["AMZN"] = rand.Float64()
		mNasdaq["AXON"] = rand.Float64()
		sNasdaq := stocks.Stock{
			Companies: mNasdaq,
		}

		mNYSE["BABA"] = rand.Float64()
		mNYSE["BAC"] = rand.Float64()
		mNYSE["BB"] = rand.Float64()
		mNYSE["BBAI"] = rand.Float64()
		mNYSE["BBD"] = rand.Float64()
		sNYSE := stocks.Stock{
			Companies: mNYSE,
		}

		mLondonSE["CCL"] = rand.Float64()
		mLondonSE["CCR"] = rand.Float64()
		mLondonSE["CGT"] = rand.Float64()
		mLondonSE["CHG"] = rand.Float64()
		mLondonSE["CHRY"] = rand.Float64()
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
		time.Sleep(time.Second * 5)
		i = i + 1

		stocks.ClearMap(mNasdaq)
		stocks.ClearMap(mNYSE)
		stocks.ClearMap(mLondonSE)
	}

}
