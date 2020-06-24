package game

import (
	"fmt"
)

type Color int
type Board [nRows][nCol] Color
const nRows int = 6
const nCol int = 7

const(
	NONE Color = iota
	BLUE
	RED
)

func (color Color) String() string {
	if(color > 2 || color < 0) {
		return "Unknown"
	}
	return []string{"NONE", "BLUE", "RED"}[color]
}

func NewBoard() Board {
	return [nRows][nCol]Color{}
}

func (b *Board) AddChip(column int, color Color) error{
	if !(BLUE == color || RED == color) {
		return fmt.Errorf("Expected chip color to be either %v(%d) or %v(%d) but received %v(%d).", BLUE, BLUE, RED, RED, color, color)
	}
	for i := len(b)-1; i>=0; i-- {
		if(NONE == b[i][column]) {
			b[i][column] = color
			break
		}
	}
	return nil
}
