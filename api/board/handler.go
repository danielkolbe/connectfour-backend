package board

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/danielkolbe/connectfour/game"
	"github.com/danielkolbe/connectfour/logger"
)

// Handler implements the http.Handler interface.
type Handler struct {
	gameService game.Service
	gameID func(w http.ResponseWriter, req *http.Request) string
}

// NewHandler returns a new Handler instance.
// If NewHandler is used, the returned handler will
// be wrapped so that any panic that is escalated to the
// handler will be turned into an http 500 response
func NewHandler(gameService game.Service,  gameID func(w http.ResponseWriter, req *http.Request) string) http.Handler {
	return panicRecover(Handler{gameService, gameID})
}

// ServerHTTP takes an incoming (GET) request for an game existing
// game board. Reads the gameID from cookie. If the gameID is not present
// or does not match an existing game a new board will be returned.
// Steps:
// 1) Extract gameID from cookie if present, else create and set cookie
// 2) Obtain existing or new board
// 3) Return board as json if content-type is application/json, board as text else
func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	gameID := h.gameID(w, req)
	logger.Logger.Debugf("Retrieving current board for game %v", gameID)
	board := h.gameService.Board(gameID)
	if "application/json" == req.Header.Get("Content-type") {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(board)
		return
	}
	w.Write([]byte(board.String()))
}

// panicRecover wraps an http.Handler instance so that any panic that
// is escalated to the handler will be turned into an http 500 response
func panicRecover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				logger.Logger.Error(r)
				http.Error(w, fmt.Errorf("well that's embarrassing").Error(), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, req)
	})
}
