package applicationDriver

import (
	"github.com/gorilla/mux"
	"gopkg.in/tylerb/graceful.v1"
	engine "maxzilla/services/mongo-engine"
	"net/http"
	"time"
)

type Service struct {
	mongo  engine.MongoAdapter
	health engine.HealthHandler
}

func NewService(mongoServer engine.MongoAdapter, healthServer engine.HealthHandler) *Service {
	return &Service{
		mongo:  mongoServer,
		health: healthServer,
	}
}

func (s *Service) Routes(request *mux.Router) *graceful.Server {
	// example: http://localhost:8085/health
	request.HandleFunc(engine.HealthCheck, s.health.HealthCheck).Methods("GET")

	// example: http://localhost:8085/getallconfigs
	request.HandleFunc(engine.GetAllConfigsFromDatabase, s.mongo.GetClientConfigAll).Methods("GET")

	// example http://localhost:8085/getconfig?app=testApplication&bin=0.0.2&site=dev
	request.HandleFunc(engine.GetSingleConfig, s.mongo.GetClientConfigBasedOnAppNameAndBinaryVersionAndSite).Methods("GET")

	// example http://localhost:8085/insertnew
	request.HandleFunc(engine.InsertConfig, s.mongo.InsertNewConfig).Methods("POST")

	// example http://localhost:8085/delete?app=testApplication&bin=0.0.2&site=dev
	request.HandleFunc(engine.DeleteConfig, s.mongo.DeleteRecordUsingID).Methods("DELETE")

	server := &graceful.Server{
		Timeout: 30 * time.Second,
		Server: &http.Server{
			Addr:    ":8085",
			Handler: request,
		},
	}
	return server
}
