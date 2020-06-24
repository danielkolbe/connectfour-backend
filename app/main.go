package main

import (
	"fmt"
	"github.com/danielkolbe/connectfour/game"
)

func main() {
	board := game.NewBoard()
	err := board.AddChip(2, 1)
	if nil != err {
		fmt.Println(err)
	}
	for _, row := range board {
		fmt.Printf("%v\n", row)
	}
}
