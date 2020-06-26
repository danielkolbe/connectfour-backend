package game

import (
	"fmt"
)

type color int
type Board struct {
	Fields [nRows][nCol] color
	nextColor color
}	

const nRows int = 6
const nCol int = 7

const(
	NONE color = iota
	BLUE
	RED
)

func (c color) String() string {
	if(c > 2 || c < 0) {
		return "Unknown"
	}
	return []string{"NONE", "BLUE", "RED"}[c]
}

func NewBoard() Board {
	return Board{Fields: [nRows][nCol]color{}, nextColor: RED}
}

func (b *Board) AddChip(column int) error {
	
	if NONE != b.Fields[0][column]  {
		return fmt.Errorf("Column %v is full", column)
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

 