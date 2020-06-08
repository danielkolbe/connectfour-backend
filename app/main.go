package main

import (
	"fmt"
	"github.com/danielkolbe/connectfour/game"
)

func main() {
	board := game.Board{ COLORS: [7][6]game.Color{{game.RED},{game.BLUE, game.RED}}}
	for _,row := range board.COLORS {
		fmt.Printf("%v\n", row)
	}
}
