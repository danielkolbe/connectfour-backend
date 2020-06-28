package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	// Arrange
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/turn?column=1", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/turn?column=1", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	if status := rr.Code; status != http.StatusMethodNotAllowed {
        t.Errorf("Handler returned wrong status code: got %v want %v", status,  http.StatusMethodNotAllowed)
	}
	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/unknown?column=1", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status,   http.StatusNotFound)
	}
}
