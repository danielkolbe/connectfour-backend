package ai

import (
	"encoding/json"
	"fmt"
	"github.com/danielkolbe/connectfour/game"
	"github.com/danielkolbe/connectfour/logger"
	"net/http"
)

// Handler implements the http.Handler interface.
type Handler struct {
	gameService game.Service
	ai          game.AI
	gameID      func(w http.ResponseWriter, req *http.Request) string
}

type AITurn struct {
	Column int
}

// NewHandler returns a new Handler instance.
// If NewHandler is used, the returned handler will
// be wrapped so that any panic that is escalated to the
// handler will be turned into an http 500 response
func NewHandler(gameService game.Service, ai game.AI, gameID func(w http.ResponseWriter, req *http.Request) string) http.Handler {
	return panicRecover(Handler{gameService, ai, gameID})
}

// ServerHTTP takes an incoming (PATCH) request to perform on
// artifical intelligence move on the related board. Returns
// the index of the column where the next chip was inserted.
// Steps:
// 1) Extract gameID from cookie if present, else create and set cookie
// 2) Calls gameService.TurnAI with the given gameID the implementation
//    of the game.AI interface the handler was instatiated with.
// 3) Convert error into matching http response if any or return json
//    of format {Column: int} to indicate which ai move was rendered
func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	gameID := h.gameID(w, req)
	logger.Logger.Debugf("Performing ai turn on board with gameID %v", gameID)
	col, err := h.gameService.TurnAI(gameID, h.ai)
	if nil != err {
		switch t := err.(type) {
		case *game.MatchIsOverError:
			logger.Logger.Error(t)
			w.WriteHeader(http.StatusConflict)
		case *game.BoardDoesNotExistError:
			err = fmt.Errorf("no board created, please perform a GET request on /board first")
			logger.Logger.Error(t)
			w.WriteHeader(http.StatusNotFound)
		default:
			logger.Logger.Error(t)
			err = fmt.Errorf("sorry for that")
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprint(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AITurn{col})
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
