package healthCheck

import (
	"encoding/json"
	engine "godzilla/services/mongo-engine"
	"net/http"
	"time"
)

type Service struct {
	config engine.InitialConfig
}

func NewHealthService(config engine.InitialConfig) (engine.HealthHandler, error) {
	return &Service{config: config}, nil
}

func (s *Service) HealthCheck(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	responseByte, _ := json.Marshal(HealthEndpoint{
		Application:  "Mongo Engine",
		Version:      s.config.GetApplicationBinary(),
		HealthStatus: "200 OK",
		Message:      "Up and running for " + time.Since(s.config.GetAppStartupTime()).String(),
	})
	_, _ = writer.Write(responseByte)
}
