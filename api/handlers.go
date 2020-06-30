package api

import (
	"fmt"
	"github.com/danielkolbe/connectfour/game"
	"net/http"
	"strconv"
	"github.com/satori/go.uuid"
)

type TurnHandler struct {
	gameService game.GameService
}

func NewTurnHandler(gameService game.GameService) TurnHandler {
	return TurnHandler{gameService}
}

func (h TurnHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	gameId := gameId(w, req)
	c, err := strconv.Atoi(req.FormValue("column"))
	if(nil != err) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Please provide a query parameter 'column' as integer greater or equal 0")
		return
	}
	if(nil != h.gameService.Turn(c, gameId)) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}
}

func gameId(w http.ResponseWriter, req *http.Request) string {
	c, err := req.Cookie("gameId")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "gameId",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c.Value
}