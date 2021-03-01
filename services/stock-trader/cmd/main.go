package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	service "github/godzilla/services/stock-trader/components"
	"net/http"
	"os"
	"time"
)

func main() {

	databaseUserName := os.Getenv("dbUser")
	databasePassword := os.Getenv("dbPassword")
	databaseName := os.Getenv("dbName")
	configFileFullPath := "config.json"
	//configFileFullPath := "C:\\Projects-Golang\\src\\godzilla\\services\\stock-trader\\resources\\config.json" // while using windows for dev purpose only
	configuration, err := readConfigJson(configFileFullPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var repository service.RepositoryClient
	var index = 1
	for {
		repository, err = service.NewRepositoryClient(databaseUserName, databasePassword, databaseName, configuration)
		if err != nil {
			fmt.Println(fmt.Sprintf(service.ErrConnectingToDb, err.Error(), index))
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println(service.InfoSuccessfulConnectionDb)
			break
		}
		index++
	}
	startService(repository, configuration, e)

}

func startService(repository service.RepositoryClient, config service.Configuration, e *echo.Echo) {
	s := service.NewStockTraderService(repository)

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Httpport)))

	// TODO: implement service here
	ctx := context.Background()
	s.GetStatus(ctx, nil)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func readConfigJson(configFilePath string) (service.Configuration, error) {
	configFromJsonFile := service.Configuration{}
	configJsonFile, err := os.Open(configFilePath)
	if err != nil {
		return service.Configuration{}, err
	}
	decoderConfigFile := json.NewDecoder(configJsonFile)
	errDecode := decoderConfigFile.Decode(&configFromJsonFile)
	if errDecode != nil {
		return service.Configuration{}, errDecode
	}
	return configFromJsonFile, nil
}
