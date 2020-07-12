package game

import (
	"fmt"
	"strings"
	"sync"
)

// Board represents a connect four game Board and its current state
// (all added chips and the color of the next chip) .
type Board struct {
	Fields    [nRows][nCols]color
	nextColor color
	mutex sync.Mutex
}

type color int

const nRows int = 6
const nCols int = 7

const (
	none color = iota
	blue
	red
)

func (b Board) String() string {
	var str strings.Builder
	for row := 0; row < nRows; row++ {
		for column := 0; column < nCols; column++ {
			str.WriteString(b.Fields[row][column].String())
			str.WriteString(" ")
		}    
		str.WriteString("\n")
	}
	return str.String()
}

func (c color) String() string {
	if c > 2 || c < 0 {
		return "Unknown"
	}
	return []string{"n", "b", "r"}[c]
}

// newBoard returns a pointer to a new Board instance.
// The nextColor field will be pre-set to red.
func newBoard() *Board {
	return &Board{Fields: [nRows][nCols]color{}, nextColor: red}
}

// addChip adds a new chip to the Board inserting
// it at the specified column. If the column is
// full or out of bounds an error will be returned. 
func (b *Board) addChip(column int) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if none != b.win() {
		return NewMatchIsOverError()
	}
	if nCols-1 < column || 0 > column {
		return fmt.Errorf("column %v is out of bounds: 0-%v", column, nCols-1)
	}
	if none != b.Fields[0][column] {
		return NewColumnIsFullError(column)
	}
	for row := len(b.Fields) - 1; row >= 0; row-- {
		if none == b.Fields[row][column] {
			b.Fields[row][column] = b.nextColor
			break
		}
	}
	b.nextColor = (2 - (b.nextColor - 1))
	
	return nil
}

// win returns the winning color or none if
// the board has no winner yet
func (b *Board) win() color {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return win(b);
} 