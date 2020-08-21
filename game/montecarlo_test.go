package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomColumn(t *testing.T) {
	randomIntSwap := randomInt
	randomInt = func(n int) int {
		return 0
	}
	defer func() {
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
		{red, red, red, red, blue, red, red},
	}, winner: none,
	}
	// Act & Assert
	require.Equal(t, randomColumn(b), -1, "should return -1 if board is full")
}

func TestRandomMatch(t *testing.T) {
	randomIntSwap := randomInt
	randomInt = func(n int) int {
		return n - 1
	}
	defer func() {
		randomInt = randomIntSwap
	}()
	// Arrange
	b := Board{Fields: [nRows][nCols]color{
		{red, red, blue, none, none, none, none},
		{blue, blue, red, none, none, none, none},
		{red, red, blue, none, none, none, none},
		{blue, blue, red, none, none, none, none},
		{red, red, blue, none, none, none, none},
		{blue, blue, blue, none, none, none, none},
	}, winner: none, nextColor: red,
	}
	// Act & Assert
	require.Equal(t, randomMatch(b), red, "should return red")
}

func TestEmpiricalLikelihoodOfWinnig(t *testing.T) {
	// Arrange
	b := newBoard()
	// Act & Assert
	result := empiricalLikelihoodOfWinning(b, 5000)
	require.Greater(t, result, 0.85, "winning likelihood of red should be greater than 85% if red starts with a random chip")
	require.Less(t, result, 0.93, "winning likelihood of red should be less than 93% if red starts with a random chip")

	// Arrange
	b = &Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, red},
	}, winner: none, nextColor: blue,
	}
	// Act & Assert
	result = empiricalLikelihoodOfWinning(b, 5000)
	require.Greater(t, 1.0 - result, 0.51, "winning likelihood of red should be greater than 51% if red starts with a chip at the side")
	require.Less(t, 1.0 - result, 0.65, "winning likelihood of red should be less than 65% if red starts with a chip at the side")

	// Arrange
	b = &Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, red, none, none, none},
	}, winner: none, nextColor: blue,
	}
	// Act & Assert
	result = empiricalLikelihoodOfWinning(b, 5000)
	require.Greater(t, 1.0 - result, 0.92, "winning likelihood of red should be greater than 92% if red starts with a chip in the middle")
	require.Less(t, 1.0 - result, 0.96, "winning likelihood of red should be less than 96% if red starts with a chip in the middle")
}
