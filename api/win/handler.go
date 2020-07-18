package win

import (
	"fmt"
	"github.com/danielkolbe/connectfour/game"
	"github.com/danielkolbe/connectfour/logger"
	"net/http"
)

// Handler implements the http.Handler interface.
type Handler struct {
	gameService game.Service
	gameID      func(w http.ResponseWriter, req *http.Request) string
}

// NewHandler returns a new Handler instance.
// If NewHandler is used, the returned handler will
// be wrapped so that any panic that is escalated to the
// handler will be turned into an http 500 response
func NewHandler(gameService game.Service, gameID func(w http.ResponseWriter, req *http.Request) string) http.Handler {
	return panicRecover(Handler{gameService, gameID})
}

// ServerHTTP takes an incoming (Get) request for the winning
// color of the related game board.
// Steps:
// 1) Extract gameID from cookie if present, else create and set cookie
// 2) Calls gameService.Winner with the gameID and return result or
//    convert error into matching http response if any
func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	gameID := h.gameID(w, req)
	logger.Logger.Debugf("Retrieving winner for game %v", gameID)
	color, err := h.gameService.Winner(gameID)
	if nil != err {
		switch t := err.(type) {
		case *game.BoardDoesNotExistError:
			logger.Logger.Error(t)
			err = fmt.Errorf("no board created, please perform GET request on /board first")
			w.WriteHeader(http.StatusNotFound)			
		default:
			logger.Logger.Error(t)
			err = fmt.Errorf("sorry for that")
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, color)
}

// panicRecover wraps an http.Handler instance so that any panic that
// is escalated to the handler will be turned into an http 500 response
func panicRecover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				logger.Logger.Error(r)
				http.Error(w, fmt.Errorf("sorry for that").Error(), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, req)
	})
}
