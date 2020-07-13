package game

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewBoard(t *testing.T) {
	// Arrange
	b := newBoard()
	for i, row := range b.Fields {
		for j, color := range row {
			// Assert
			require.Equal(t, color, none, fmt.Sprintf("A newly created Board should be blank but the color of field %v,%v was %v", i, j, color))
		}
	}
}
func TestErrorHandlingFullColumn(t *testing.T) {
	b := newBoard()
	// Fill column 0
	for i := 0; i < nRows; i++ {
		error := b.addChip(0)
		require.Equal(t, nil, error, fmt.Sprintf("Error should be nil but was: `%v`.", error))
	}
	// Column 0 is full already.
	error := b.addChip(0)
	require.NotEqual(t, nil, error, "should return an non-nil error if column is full")
	require.Equal(t, NewColumnIsFullError(0).Error(), error.Error(), "should return an ColumnIsFullError if column is full")
}

func TestErrorHandlingColumnOutOfBounds(t *testing.T) {
	// Arrange
	b := newBoard()
	// Act
	error := b.addChip(nCols)
	// Assert
	require.Equal(t, NewColumnIsOutOfBoundsError(nCols).Error(), error.Error(), "should return an ColumnIsOutOfBoundsError")
}

func TestErrorHandlingMatchIsOver(t *testing.T) {
	// Arrange
	b := newBoard()
	b.winner = red
	// Act
	error := b.addChip(0)
	// Assert
	require.Equal(t, NewMatchIsOverError().Error(), error.Error(), "should return an MatchIsOverError if board has already a winner")
}

func TestAddChip(t *testing.T) {
	// Arrange
	b := newBoard()
	// Act
	err := b.addChip(4)
	// Assert
	require.Equal(t, nil, err, fmt.Sprintf("Expected error to be nil but was: %v", err))
	require.Equal(t, red, b.Fields[5][4], fmt.Sprintf("Color of field 5,4 of the Board should be red but was %v", b.Fields[5][4]))

	// Act
	err = b.addChip(4)
	// Assert
	require.Equal(t, nil, err, fmt.Sprintf("Expected error to be nil but was: %v", err))
	require.Equal(t, blue, b.Fields[4][4], fmt.Sprintf("Color of field 4,4 of the Board should be blue but was %v", b.Fields[4][4]))

	// Act
	err = b.addChip(nCols)
	// Assert
	require.NotEqual(t, nil, err, fmt.Sprintf("Should return an error if column number exceeds upper limit"))

	// Act
	err = b.addChip(-1)
	// Assert
	require.NotEqual(t, nil, err, fmt.Sprintf("Should return an error if column number is negative"))
}

func TestString(t *testing.T) {
	// Arrange
	b := newBoard()
	// Assert
	require.Equal(t,
		"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n", b.String())

	// Arrange
	b.addChip(4)
	b.addChip(2)
	b.addChip(2)
	// Act
	require.Equal(t,
		"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n r n n n n \n"+
			"n n b n r n n \n", b.String())
}

func TestNextColor(t *testing.T) {
	// Arrange
	b := newBoard()
	// Assert
	require.Equal(t, red, b.nextColor, fmt.Sprintf("Next color must be red but was %v", b.nextColor))

	// Act & Assert
	b.addChip(4)
	require.Equal(t, blue, b.nextColor, fmt.Sprintf("Next color must be blue but was %v", b.nextColor))

	// Act & Assert
	b.addChip(3)

	require.Equal(t, red, b.nextColor, fmt.Sprintf("Next color must be red but was %v", b.nextColor))
	// Act & Assert
	b.addChip(3)
	require.Equal(t, blue, b.nextColor, fmt.Sprintf("Next color must be blue but was %v", b.nextColor))
}

func TestWin(t *testing.T) {
	// Arrange
	b := newBoard()
	// Assert
	require.Equal(t, none, b.winner, "Should return none on a newly created board")

	// Arrange
	b = &Board{Fields: [nRows][nCols]color{}, winner: blue}
	// Assert
	require.Equal(t, blue, b.winner, "Should return winner field if winner is already set")
	
	// Arrange
	b = &Board{Fields: [nRows][nCols]color{
		{none, none, none, blue, none, none, none},
		{none, none, none, none, blue, none, none},
		{none, red, none, none, none, blue, none},
		{none, red, none, none, none, none, blue},
		{none, red, blue, none, blue, none, none},
		{none, red, red, red, blue, red, none},
	}, winner: none,
	}
	// Assert
	require.Equal(t, red, b.win(), "Should return winner from win method if field is none")
	require.Equal(t, red, b.winner, "Should set winner field to red")
}
