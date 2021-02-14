package main

import (
	"godzilla/services/mongo-engine/internal/applicationDriver"
	"godzilla/services/mongo-engine/internal/initialConfig"
	"log"
)

func main() {
	config, err := initialConfig.NewConfiguration()
	if err != nil {
		log.Println(err)
	}
	applicationDriver.Start(config)
}
