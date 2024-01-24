package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	apiKey   = "Replace_me_with_ALPHA_VANTAGE_API_KEY"
	symbol   = "NVDA" // Stock symbol for Nvidia Inc. (you can change it to any stock symbol)
	interval = 5      // Time interval between API requests in seconds
)

type StockQuote struct {
	Symbol    string `json:"01. symbol"`
	Price     string `json:"05. price"`
	Timestamp string `json:"07. latest trading day"`
}

func fetchStockQuote() (*StockQuote, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", symbol, apiKey)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data map[string]StockQuote
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	quote, ok := data["Global Quote"]
	if !ok {
		return nil, fmt.Errorf("Failed to get stock quote")
	}

	return &quote, nil
}

func main() {
	if apiKey == "Replace_me_with_ALPHA_VANTAGE_API_KEY" {
		fmt.Println("Please provide your Alpha Vantage API key")
		os.Exit(1)
	}

	ticker := time.NewTicker(time.Second * time.Duration(interval))

	for range ticker.C {
		quote, err := fetchStockQuote()
		if err != nil {
			log.Println("Error fetching stock quote:", err)
			continue
		}

		fmt.Printf("Symbol: %s, Price: %s, Timestamp: %s\n", quote.Symbol, quote.Price, quote.Timestamp)
	}
}
