package healthCheck_test

import (
	"net/http"
	"net/http/httptest"
	"stockzilla/services/mongo-engine/internal/healthCheck"
	"stockzilla/services/mongo-engine/mocks"
	"testing"
	"time"
)

func TestHealthCheckHandler(t *testing.T) {

	mockInitialConfig := &mocks.InitialConfig{}
	mockInitialConfig.On("GetApplicationBinary").Return("0.0.1")
	mockInitialConfig.On("GetAppStartupTime").Return(time.Now())
	service, err := healthCheck.NewHealthService(mockInitialConfig)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.HealthCheck)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
