package applicationDriver

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"os"
	engine "stockzilla/services/mongo-engine"
	"stockzilla/services/mongo-engine/internal/generalUtilities"
	"stockzilla/services/mongo-engine/internal/healthCheck"
	"stockzilla/services/mongo-engine/internal/mongoAdapter"
	"strings"
)

func Start(config engine.InitialConfig) {
	request := mux.NewRouter().StrictSlash(false)

	mongoServer, err := mongoAdapter.NewServer(config)
	if err != nil {
		log.Printf("Cannot connecto to MongoDB. ERROR: %v", err)
		os.Exit(1)
	} else {
		defer mongoServer.Close()

		healthServer, err := healthCheck.NewHealthService(config)
		if err != nil {
			panic(err)
		}

		server := NewService(mongoServer, healthServer)
		server.Routes(request)

		ip, err := generalUtilities.ExternalIP()
		if err != nil {
			fmt.Println(err)
		}

		log.Printf("Application started successfully. Running in ip %v & serving port 8085", ip)
		if strings.EqualFold(*config.GetSSLMode(), "false") {
			log.Printf("Dev mode set to false. Starting application in ssl secured mode")
			errStartingServer := server.Routes(request).ListenAndServeTLS(*config.GetSslCert(), *config.GetSslKey())
			if errStartingServer != nil {
				log.Printf("Failed to start server | Error: %v", errStartingServer)
			}
		} else {
			log.Printf("Starting application in ssl non-secured mode")
			err = server.Routes(request).ListenAndServe()
			if err != nil {
				panic(err)
			}
		}
		log.Printf("Application stopped gracefully")
	}

}
