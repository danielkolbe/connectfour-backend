package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	// Test http 200
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/turn?column=1", nil)
	NewRouter().ServeHTTP(rr, req)
	require.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code" )
	// Test http 405
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/turn?column=1", nil)	
	NewRouter().ServeHTTP(rr, req)
	require.Equal(t, http.StatusMethodNotAllowed, rr.Code, "handler returned wrong status code" )
	// Test http 404
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/unknown?column=1", nil)
	NewRouter().ServeHTTP(rr, req)
	require.Equal(t, http.StatusNotFound, rr.Code, "handler returned wrong status code" )
	// Test http 400
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/turn?unknown=1", nil)
	NewRouter().ServeHTTP(rr, req)
	require.Equal(t, http.StatusBadRequest, rr.Code, "handler returned wrong status code" )
}
