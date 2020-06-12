package main

import (
	"fmt"
	"github.com/danielkolbe/connectfour/game"
)

func main() {
	board := game.Board{}
	/*for _,row := range board.COLORS {
		fmt.Printf("%v\n", row)
	}*/
	board.AddChip(2, game.RED)
	for _,row := range board {
		fmt.Printf("%v\n", row)
	}
}
