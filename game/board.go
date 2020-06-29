package game

import (
	"fmt"
)

type color int
type board struct {
	fields    [nRows][nCol]color
	nextColor color
}

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

func NewBoard() *board {
	return &board{fields: [nRows][nCol]color{}, nextColor: red}
}

func (b *board) addChip(column int) error {

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
