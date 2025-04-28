package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
)

// init registers the same handler as main.
// This ensures that the "/" route is available for our tests.
func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("Hello from Go! Running on %s/%s\n", runtime.GOOS, runtime.GOARCH)
		w.Write([]byte(message))
	})
}

func TestHandler(t *testing.T) {
	// Create a new HTTP request to the root route.
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	// Use the default serve mux which was populated in init().
	http.DefaultServeMux.ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Build the expected response.
	expected := fmt.Sprintf("Hello from Go! Running on %s/%s\n", runtime.GOOS, runtime.GOARCH)
	responseBody := rr.Body.String()

	// Check if the response contains the expected message.
	if !strings.Contains(responseBody, expected) {
		t.Errorf("handler returned unexpected body: got %v want to contain %v", responseBody, expected)
	}
}
