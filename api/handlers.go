package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/danielkolbe/connectfour/game"
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
	c, err := parseColumn(req)
	if(nil != err) {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Please provide a query parameter 'column' as integer greater or equal 0")
		return
	}
	if(nil != h.gameService.Turn(c, gameId)) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}
}

func parseColumn(req *http.Request) (int, error) {
	body, err := ioutil.ReadAll(req.Body)
	if nil != err {
		return -1, err
	}
	var t struct {Column int}
	err = json.Unmarshal(body, &t)
	fmt.Printf("%+v", t)
	if nil != err {
		return -1, err
	}
	
	return t.Column, nil
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