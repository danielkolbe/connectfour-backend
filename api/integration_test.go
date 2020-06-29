package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	// Arrange
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/turn?column=1", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// require
	require.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code" )
	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/turn?column=1", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// require
	require.Equal(t, http.StatusMethodNotAllowed, rr.Code, "handler returned wrong status code" )
	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/unknown?column=1", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// require
	require.Equal(t, http.StatusNotFound, rr.Code, "handler returned wrong status code" )
	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/turn?unknown=1", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// require
	require.Equal(t, http.StatusBadRequest, rr.Code, "handler returned wrong status code" )
}
