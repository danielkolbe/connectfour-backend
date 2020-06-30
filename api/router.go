package api

import (
	"net/http"

	"github.com/danielkolbe/connectfour/game"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/turn", TurnHandler{game.CFour{}}).Methods("POST").Schemes("http")
	r.Handle("/favicon.ico", http.NotFoundHandler())
	return r
}