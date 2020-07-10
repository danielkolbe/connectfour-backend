package api

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGameID(t *testing.T) {
	// Arrange
	cookie := &http.Cookie{Name: "gameID", Value: "324234-555"}
	req, _ := http.NewRequest("", "", nil)
    req.AddCookie(cookie)
	rr := httptest.NewRecorder()
	// Act
	id := gameID(rr, req)
	// Assert
	require.Equal(t, "324234-555", id, "Should return gameID extracted from cookie.")
	
	// Arrange 
	req, _ = http.NewRequest("", "", nil)
	// Regex for uuid
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	// Act
	id = gameID(rr, req)
	// Assert
	require.Equal(t, true, r.MatchString(id), "Should return a generated uuid if no cookie given.")
}
