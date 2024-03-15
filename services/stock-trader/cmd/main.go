package main

import (
	"context"
	"encoding/json"
	"fmt"
	grpcRunTime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/vickeyshrestha/sharing-services/protobuf/stock_trader"
	service "github/stockzilla/services/stock-trader/components"
	"google.golang.org/grpc"
	grpcHealth "google.golang.org/grpc/health"
	grpcHealthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/handlers"

	_ "github.com/jnewmano/grpc-json-proxy/codec"
)

func main() {

	databaseUserName := os.Getenv("dbUser")
	databasePassword := os.Getenv("dbPassword")
	databaseName := os.Getenv("dbName")
	natsUrl := os.Getenv("natsUrl")
	configFileFullPath := "config.json"
	//configFileFullPath := "I:\\go\\src\\stockzilla\\services\\stock-trader\\resources\\config.json" // while using windows for dev purpose only
	configuration, err := readConfigJson(configFileFullPath)
	if err != nil {
		fmt.Println(err)
		return
	}

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

	startGrpcServer(repository, configuration)
	startHttpAgent(configuration)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Starting Subscriber")
		service.Subscriber(natsUrl)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Starting Dequeue")
		service.Dequeue()
	}()

	// TODO: Third Go routine that would calculate the Mean and inserts into the table
	wg.Wait()

	runtime.Goexit()

}

/*
starts grpc service
*/
func startGrpcServer(repository service.RepositoryClient, config service.Configuration) {
	listener, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		panic(err)
	}
	stockTraderService := service.NewStockTraderService(repository)
	var server = grpc.NewServer()
	pb.RegisterStockTraderServer(server, stockTraderService)
	healthServer := grpcHealth.NewServer()
	grpcHealthv1.RegisterHealthServer(server, healthServer)
	go func() {
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()
}

/*
following method acts as an http agent for grpc protocol
*/
func startHttpAgent(config service.Configuration) {
	var DialOptions []grpc.DialOption
	DialOptions = append(DialOptions, grpc.WithInsecure())

	go func() {
		ctx := context.Background()
		ctxWithCancel, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := grpcRunTime.NewServeMux(grpcRunTime.WithMarshalerOption(grpcRunTime.MIMEWildcard, &grpcRunTime.JSONPb{}))
		_ = pb.RegisterStockTraderHandlerFromEndpoint(ctxWithCancel, mux, config.GrpcPort, DialOptions)
		fmt.Println("starting http and listening to port", config.Httpport)
		if err := http.ListenAndServe(config.Httpport, handlers.CORS(handlers.AllowedHeaders([]string{"grpc-metadata-token", "content-type"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(mux)); err != nil {
			panic(err)
		}
	}()
}

/*
reads configuration from Config.json file
*/
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
