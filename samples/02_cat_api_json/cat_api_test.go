package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCatsAPIHandler(t *testing.T) {
	// 1. Create test request
	req, err := http.NewRequest("GET", "/api/cats", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 2. Create a ResponseRecorder (which satisfies http.ResponseWriter)
	recorder := httptest.NewRecorder()

	// 3. Invoke handler
	catAPIHandler(recorder, req)

	// 4. Check the status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Wrong status code: got %v expected %v", recorder.Code, http.StatusOK)
	}
}
