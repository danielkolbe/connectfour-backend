package game

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestRandomColumn(t *testing.T) {
	randomIntSwap := randomInt
	randomInt = func (n int) int {
		return 0
	}
	defer func(){
		randomInt = randomIntSwap
	}()

	// Arrange
	b := &Board{Fields: [nRows][nCols]color{
		{red, red, none, blue, none, none, none},
		{red, red, none, none, blue, none, none},
		{red, red, none, none, none, blue, none},
		{red, red, none, none, none, none, blue},
		{red, red, blue, none, blue, none, none},
		{red, red, red, red, blue, red, none},
	}, winner: none,
	}
	// Act & Assert
	require.Equal(t, randomColumn(b), 2, "should return index of the first free column if rand returns 0")

	// Arrange
	b = &Board{Fields: [nRows][nCols]color{
		{red, red, red, blue, red, red, red},
		{red, red, red, red, blue, red, red},
		{red, red, red, red, red, blue, red},
		{red, red, red, red, red, red, blue},
		{red, red, blue, red, blue, red, red},
		{red, red, red, red, blue, red,red},
	}, winner: none,
	}
	// Act & Assert
	require.Equal(t, randomColumn(b), -1, "should return -1 if board is full")
}