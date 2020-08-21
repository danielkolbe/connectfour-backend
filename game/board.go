package game

import (
	"strings"
	"sync"
)

// Board represents a connect four game Board and its current state
// (all added chips and the color of the next chip) .
type Board struct {
	Fields    [nRows][nCols]color
	nextColor color
	winner color
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
	return &Board{Fields: [nRows][nCols]color{}, nextColor: red, winner: none}
}

// addChip adds a new chip to the Board inserting
// it at the specified column. If the column is
// full or out of bounds or the match has already a winner 
// an customn error will be returned. 
func (b *Board) addChip(column int) error {
	winner := b.win()
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if none != winner {
		return NewMatchIsOverError()
	}
	if nCols-1 < column || 0 > column {
		return NewColumnIsOutOfBoundsError(column)
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
	if winner := b.winner; none != winner {
		return winner
	}
	b.winner = findwinner(b)

	return b.winner;
} 

// freeColumns returns indices of the columns of the
// board that are not already full
func (b *Board) freeColumns() []int {
	c := []int{}
	for i := 0; i < nCols; i++ {
		if none == b.Fields[0][i] {
			c = append(c, i)
		}
	}  
	return c
}	