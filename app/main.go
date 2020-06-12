package main

import (
	"fmt"
	"github.com/danielkolbe/connectfour/game"
)

func main() {
	board := game.NewBoard();
	board.AddChip(2, game.RED)
	for _,row := range board {
		fmt.Printf("%v\n", row)
	}
}
