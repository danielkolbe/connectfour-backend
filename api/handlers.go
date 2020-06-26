package handlers

import (
	"fmt"
	"net/http"
	"github.com/danielkolbe/connectfour/game"
	"strconv"
)

var b game.Board = game.NewBoard()

func Turn(w http.ResponseWriter, req *http.Request) {
	c, err := strconv.Atoi(req.FormValue("column"))
	if(nil != err) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Please provide a query parameter 'column' as integer greater or equal 0")
	}
	err = b.AddChip(c)
	if(nil != err) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}
	for _, row := range b.Fields {
		fmt.Printf("%v\n", row)
	}
}