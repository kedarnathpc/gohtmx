package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	// Create a new request for the root endpoint
	req1, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Write a "test successful" response
	h1 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Test successful")
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler functions
	h1(rr, req1)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	// TODO: Add more assertions to validate the response body and behavior of the handlers
}
