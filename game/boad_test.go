package game

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewBoard(t *testing.T) {
	// Arrange
	b := newBoard()
	for i, row := range b.fields {
		for j, color := range row {
			// Assert
			require.Equal(t, color, none, fmt.Sprintf("A newly created board should be blank but the color of field %v,%v was %v", i, j, color))
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
	require.NotEqual(t, nil, error, "No error returned.")
}
func TestAddChip(t *testing.T) {
	// Arrange
	b := newBoard()
	// Act
	err := b.addChip(4)
	// Assert
	require.Equal(t, nil, err, fmt.Sprintf("Expected error to be nil but was: %v", err))
	require.Equal(t, red, b.fields[5][4], fmt.Sprintf("Color of field 5,4 of the board should be red but was %v", b.fields[5][4]))

	// Act
	err = b.addChip(4)
	// Assert
	require.Equal(t, nil, err, fmt.Sprintf("Expected error to be nil but was: %v", err))
	require.Equal(t, blue, b.fields[4][4], fmt.Sprintf("Color of field 4,4 of the board should be blue but was %v", b.fields[4][4]))

	// Act
	err = b.addChip(nCol)
	// Assert
	require.NotEqual(t, nil, err, fmt.Sprintf("Should return an error if column number exceeds upper limit"))

	// Act
	err = b.addChip(-1)
	// Assert
	require.NotEqual(t, nil, err, fmt.Sprintf("Should return an error if column number is negative"))
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
