package api

import (
	"encoding/json"
	"fmt"
	"github.com/danielkolbe/boardgames/game"
	"github.com/danielkolbe/boardgames/app/logger"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
)


// turnHandler implements the http.Handler interface.
type turnHandler struct {
	gameService game.Service
}

// newTurnHandler returns a new turnHandler instance.
// If newTurnHandler is used, the returned handler will
// be wrapped so that any panic that is escalated to the
// handler will be turned into an http 500 response
func newTurnHandler(gameService game.Service) http.Handler {
	return panicRecover(turnHandler{gameService})
}

// ServerHTTP takes an incoming (POST) request which is required to have
// a body that can be unmarshalled to struct {Column int}.
// Steps:
// 1) Extract gameID from cookie if present, else create and set cookie
// 2) Parse column number from request
// 3) Calls gameService.Turn with the column number
func (h turnHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	gameID := gameID(w, req)
	c, err := parseColumn(req)
	if nil != err {
		logger.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Please provide a query parameter 'column' as integer greater or equal 0")
		return
	}
	logger.Logger.Debugf("Adding next chip to column %v for game %v", c, gameID)
	if err := h.gameService.Turn(c, gameID);  nil != err {
		logger.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
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

// gameID extracts the gameId from cookie if present,
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
