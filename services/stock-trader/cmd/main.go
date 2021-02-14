package main

import (
	"fmt"
	service "github/godzilla/services/stock-trader/components"
)

func main() {

	// TODO: Main code goes here
	fmt.Println("Hello stock traders")
	startGrpc()
}

func startGrpc() {
	service.NewStockTraderService()
}
