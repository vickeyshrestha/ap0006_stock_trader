package main

import (
	"log"
	"stockzilla/services/mongo-engine/internal/applicationDriver"
	"stockzilla/services/mongo-engine/internal/initialConfig"
)

func main() {
	config, err := initialConfig.NewConfiguration()
	if err != nil {
		log.Println(err)
	}
	applicationDriver.Start(config)
}
