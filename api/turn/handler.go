package turn

import (
	"encoding/json"
	"fmt"
	"github.com/danielkolbe/connectfour/game"
	"github.com/danielkolbe/connectfour/logger"
	"io/ioutil"
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

// ServerHTTP takes an incoming (POST) request which is required to have
// a body that can be unmarshalled to struct {Column int}. Performs the
// next turn on the related game board.
// Steps:
// 1) Extract gameID from cookie if present, else create and set cookie
// 2) Parse column number from request
// 3) Calls gameService.Turn with the given column number and return error if any
func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	gameID := h.gameID(w, req)
	c, err := parseColumn(req)
	logger.Logger.Debugf("Column param parsed from request (with gameID %v): %v", gameID, c)
	if nil != err {
		logger.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "missing or negative column property in post body")
		return
	}
	logger.Logger.Debugf("Adding next chip to column %v for game %v", c, gameID)
	if err := h.gameService.Turn(c, gameID); nil != err {
		switch t := err.(type) {
		case *game.ColumnIsFullError:
			logger.Logger.Error(t)
			w.WriteHeader(http.StatusConflict)
		case *game.MatchIsOverError:
			logger.Logger.Error(t)
			w.WriteHeader(http.StatusConflict)
		case *game.BoardDoesNotExistError:
			logger.Logger.Error(t)
			w.WriteHeader(http.StatusNotFound)
		case *game.ColumnIsOutOfBoundsError:
			logger.Logger.Error(t)
			w.WriteHeader(http.StatusBadRequest)				
		default:
			logger.Logger.Error(t)
			err = fmt.Errorf("sorry for that")
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprint(w, err)
	}
}

// parseColumn extracts the column number from the given request body
// If the body cannot be unmarshalled to struct {column int}
// an error will be returned likewice if column < 0
func parseColumn(req *http.Request) (int, error) {
	body, err := ioutil.ReadAll(req.Body)
	if nil != err {
		return -1, err
	}
	t := struct{ Column int }{-1}
	err = json.Unmarshal(body, &t)
	if nil != err {
		return -1, err
	}
	if 0 > t.Column {
		return -1, fmt.Errorf("could not parse column or value of column < 0")
	}

	return t.Column, nil
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
