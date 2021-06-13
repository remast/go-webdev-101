package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// 1. Create test request
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 2. Create a ResponseRecorder (which satisfies http.ResponseWriter)
	recorder := httptest.NewRecorder()

	// 3. Invoke handler
	healthHandler(recorder, req)

	// 4. Check the status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v expected %v", status, http.StatusOK)
	}
}
