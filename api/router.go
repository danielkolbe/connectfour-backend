package api

import (
	"net/http"
	"github.com/danielkolbe/connectfour/api/board"
	"github.com/danielkolbe/connectfour/api/turn"
	"github.com/danielkolbe/connectfour/api/win"
	"github.com/danielkolbe/connectfour/api/reset"
	"github.com/danielkolbe/connectfour/game"
	"github.com/gorilla/mux"
)

// NewRouter returns a new router instance with registered http handlers.
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/turn",  turn.NewHandler(game.CFour{}, gameID)).Methods("PATCH").Schemes("http")
	r.Handle("/reset", reset.NewHandler(game.CFour{}, gameID)).Methods("PATCH").Schemes("http")
	r.Handle("/board", board.NewHandler(game.CFour{}, gameID)).Methods("GET").Schemes("http")
	r.Handle("/winner", win.NewHandler(game.CFour{}, gameID)).Methods("GET").Schemes("http")
	r.Handle("/favicon.ico", http.NotFoundHandler())
	return r
}