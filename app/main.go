package main

import (
	"fmt"
	"github.com/danielkolbe/connectfour/game"
)

func main() {
	board := game.NewBoard()
	err := board.AddChip(2)
	if nil != err {
		fmt.Println(err)
	}
	for _, row := range board.Fields {
		fmt.Printf("%v\n", row)
	}
}
