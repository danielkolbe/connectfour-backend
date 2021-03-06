package api

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	// Arrange
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/reset", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusMethodNotAllowed, rr.Code, "should return http status 405 if method not equal to PATCH")

	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("PATCH", "/reset", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusNotFound, rr.Code, "should return http status 404")

	// Arrange
	rr = httptest.NewRecorder()
	body := struct{ Column int }{4}
	bytesBody, _ := json.Marshal(body)
	req, _ = http.NewRequest("PATCH", "/turn", bytes.NewReader(bytesBody))
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusNotFound, rr.Code, "should return http status 404")

	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/board", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusOK, rr.Code, "should return http status 200")

	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/turn?column=1", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusMethodNotAllowed, rr.Code, "should return http status 405 if method not equal to POST")

	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/board", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusOK, rr.Code, "should return http status 200")

	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/board", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusMethodNotAllowed, rr.Code, "should return http status 405 if method not equal to GET")

	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/unknown", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusNotFound, rr.Code, "should return http status 404 if path does not exist")

	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/winner", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusNotFound, rr.Code, "should return http status 404 if board not found/no cookie set")

	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/winner", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusMethodNotAllowed, rr.Code, "should return http status 405 if method not equal to GET")

	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("PATCH", "/ai/montecarlo", bytes.NewReader(nil))
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusNotFound, rr.Code, "should return http status 404")

	// Arrange
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/ai/montecarlo", nil)
	// Act
	NewRouter().ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusMethodNotAllowed, rr.Code, "should return http status 405 if method not equal to PATCH")
}
