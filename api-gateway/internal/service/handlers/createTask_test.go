package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/service/handlers"
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/service/url"
)

func TestCreateTaskHandler(t *testing.T) {
	req, err := http.NewRequest("POST", url.TaskURL, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	service := &handlers.Service{}
	handler := http.HandlerFunc(service.CreateTaskHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}
