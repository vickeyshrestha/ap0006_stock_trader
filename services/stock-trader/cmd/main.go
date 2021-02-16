package main

import (
	"fmt"
	service "github/godzilla/services/stock-trader/components"
	"os"
	"time"
)

func main() {

	databaseUserName := os.Getenv("dbUser")
	databasePassword := os.Getenv("dbPassword")
	databaseName := os.Getenv("dbName")
	var repository service.RepositoryClient
	var err error
	var index = 1
	for {
		repository, err = service.NewRepositoryClient(databaseUserName, databasePassword, databaseName)
		if err != nil {
			fmt.Println(fmt.Sprintf(service.ErrConnectingToDb, err.Error(), index))
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println(service.InfoSuccessfulConnectionDb)
			break
		}
		index++
	}
	startGrpc(repository)

}

func startGrpc(repository service.RepositoryClient) {
	_ = service.NewStockTraderService(repository)
}
