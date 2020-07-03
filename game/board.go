package game

import (
	"fmt"
)

// board represents a connect four game board and its current state
// (all added chips and the color of the next chip) .
type board struct {
	fields    [nRows][nCol]color
	nextColor color
}

type color int

const nRows int = 6
const nCol int = 7

const (
	none color = iota
	blue
	red
)

func (c color) String() string {
	if c > 2 || c < 0 {
		return "Unknown"
	}
	return []string{"none", "blue", "red"}[c]
}

// newBoard returns a new board instance. 
// The nextColor field will be preset to red
func newBoard() *board {
	return &board{fields: [nRows][nCol]color{}, nextColor: red}
}

// addChip adds a new chip to the board inserting
// it at the specified column. If the column is
// full or out of bounds an error will be returned. 
func (b *board) addChip(column int) error {

	if nCol-1 < column || 0 > column {
		return fmt.Errorf("column %v is out of bounds: 0-%v", column, nCol-1)
	}

	if none != b.fields[0][column] {
		return fmt.Errorf("column %v is full", column)
	}

	for row := len(b.fields) - 1; row >= 0; row-- {
		if none == b.fields[row][column] {
			b.fields[row][column] = b.nextColor
			break
		}
	}
	b.nextColor = (2 - (b.nextColor - 1))
	return nil
}
