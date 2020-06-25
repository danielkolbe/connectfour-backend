package game

import (
	"fmt"
)

type Color int
type Board struct {
	Fields [nRows][nCol] Color
	nextColor Color
}	
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
	return Board{Fields: [nRows][nCol]Color{}, nextColor: RED}
}

func (b *Board) AddChip(column int) error {
	
	if NONE != b.Fields[0][column]  {
		return fmt.Errorf("Selected column %v is full", column)
	}
	
	for row := len(b.Fields)-1; row>=0; row-- {
		if(NONE == b.Fields[row][column]) {
			b.Fields[row][column] = b.nextColor
			break
		}
	}
	b.nextColor = (2-(b.nextColor-1))
	return nil
}

 