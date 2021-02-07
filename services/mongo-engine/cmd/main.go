package main

import (
	"log"
	"maxzilla/services/mongo-engine/internal/applicationDriver"
	"maxzilla/services/mongo-engine/internal/initialConfig"
)

func main() {
	config, err := initialConfig.NewConfiguration()
	if err != nil {
		log.Println(err)
	}
	applicationDriver.Start(config)
}
