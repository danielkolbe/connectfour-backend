package api

import (
	"net/http"
	uuid "github.com/satori/go.uuid"
)

// GameID extracts the gameId from cookie if present,
// else create and write cookie to response writer
func gameID(w http.ResponseWriter, req *http.Request) string {
	c, err := req.Cookie("gameID")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "gameID",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c.Value
}
