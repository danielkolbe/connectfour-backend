package api

import (
	"net/http"

	"github.com/danielkolbe/connectfour/game"
	"github.com/gorilla/mux"
)

// NewRouter returns a new router instance with registered http handlers.
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/turn", newTurnHandler(game.CFour{})).Methods("POST").Schemes("http")
	r.Handle("/favicon.ico", http.NotFoundHandler())
	return r
}