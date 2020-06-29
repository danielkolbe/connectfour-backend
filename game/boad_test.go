package game;

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()
	for i, row := range b.fields {
		for j, color := range row {
			require.Equal(t, color, none, fmt.Sprintf("A newly created board should be blank but the color of field %v,%v was %v", i, j, color)) 
		}
	}
}
func TestErrorHandlingFullColumn(t *testing.T) {
	b := NewBoard()
	// Fill column 0
	for i:=0 ; i < nRows; i++ {
		error := b.addChip(0)
		require.Equal(t, nil, error, fmt.Sprintf("Error should be nil but was: `%v`.", error))
	}
	// Column 0 is full already.
	error := b.addChip(0)
	require.NotEqual(t, nil, error, "No error returned.")
}
func TestAddChip(t *testing.T) {
	b := NewBoard()
	b.addChip(4)
	b.addChip(4)
	require.Equal(t, red, b.fields[5][4], fmt.Sprintf("Color of field 5,4 of the board should be red but was %v", b.fields[5][4]))
	require.Equal(t, blue, b.fields[4][4], fmt.Sprintf("Color of field 4,4 of the board should be blue but was %v", b.fields[4][4]))
}

func TestNextColor(t *testing.T) {
	b := NewBoard()
	require.Equal(t, red, b.nextColor, fmt.Sprintf("Next color must be red but was %v", b.nextColor))
	b.addChip(4)
	require.Equal(t, blue, b.nextColor, fmt.Sprintf("Next color must be blue but was %v", b.nextColor))
	b.addChip(3)
	require.Equal(t, red, b.nextColor, fmt.Sprintf("Next color must be red but was %v", b.nextColor))
	b.addChip(3)
	require.Equal(t, blue, b.nextColor, fmt.Sprintf("Next color must be blue but was %v", b.nextColor))
}		