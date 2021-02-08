package main

import (
	"fmt"
	service "github/maxzilla/services/stock-trader/components"
)

func main() {

	// TODO: Main code goes here
	fmt.Println("Hello stock traders")
}

func startGrpc() {
	service.NewStockTraderService()
}
