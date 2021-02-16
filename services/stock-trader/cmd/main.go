package main

import (
	"context"
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

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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
	startService(repository, e)

}

func startService(repository service.RepositoryClient, e *echo.Echo) {
	s := service.NewStockTraderService(repository)

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

	// TODO: implement service here
	ctx := context.Background()
	s.GetStatus(ctx, nil)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
