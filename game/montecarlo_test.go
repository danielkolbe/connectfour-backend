package game

import (
	"github.com/stretchr/testify/require"
	"testing"
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
	}, winner: none, NextColor: red,
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
	}, winner: none, NextColor: blue,
	}
	// Act & Assert
	result = empiricalLikelihoodOfWinning(b, 5000)
	require.Greater(t, 1.0-result, 0.51, "winning likelihood of red should be greater than 51% if red starts with a chip at the side")
	require.Less(t, 1.0-result, 0.65, "winning likelihood of red should be less than 65% if red starts with a chip at the side")

	// Arrange
	b = &Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, red, none, none, none},
	}, winner: none, NextColor: blue,
	}
	// Act & Assert
	result = empiricalLikelihoodOfWinning(b, 5000)
	require.Greater(t, 1.0-result, 0.92, "winning likelihood of red should be greater than 92% if red starts with a chip in the middle")
	require.Less(t, 1.0-result, 0.96, "winning likelihood of red should be less than 96% if red starts with a chip in the middle")
}

func TestNextTurn(t *testing.T) {
	// Arrange
	b := newBoard()
	// Act
	column, err := MC{}.NextTurn(b)
	// Assert
	require.Equal(t, 3, column, "should return column 3 if empty board given (middle column is the best option for first turn)")
	require.Equal(t, nil, err, "error should be nil")
	// Arrange
	b = &Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, red},
		{none, none, none, none, none, none, red},
		{none, none, none, none, none, none, red},
	}, winner: none, NextColor: blue,
	}
	// Act
	column, err = MC{}.NextTurn(b)
	// Assert
	require.Equal(t, 6, column, "should return turn that prevents oppenents immediate victory")
	require.Equal(t, nil, err, "error should be nil")
	// Arrange
	b = &Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{blue, blue, blue, none, none, none, none},
	}, winner: none, NextColor: blue,
	}
	// Act
	column, err = MC{}.NextTurn(b)
	// Assert
	require.Equal(t, 3, column, "should return column so that ai immediately wins")
	require.Equal(t, nil, err, "error should be nil")

	// Arrange
	b = &Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, red, none, none, none, none},
		{none, none, blue, blue, none, none, none},
	}, winner: none, NextColor: red,
	}
	// Act
	column, err = MC{}.NextTurn(b)
	// Assert
	require.Equal(t, 4, column, "should return column 4 to prevent blue players victory in two turns and choose the column that is closer to the middle (than column 1)")
	require.Equal(t, nil, err, "error should be nil")
	
	// Arrange
	b = &Board{Fields: [nRows][nCols]color{
		{red, red, red, none, red, red, red},
		{red, red, blue, blue, red, red, red},
		{blue, red, blue, blue, red, red, red},
		{red, blue, red, red, blue, blue, blue},
		{blue, red, red, blue, red, red, red},
		{blue, red, blue, blue, red, red, red},
	}, winner: none, NextColor: red,
	}
	// Act
	column, err = MC{}.NextTurn(b)
	// Assert
	require.Equal(t, 3, column, "should return last non-full column")
	require.Equal(t, nil, err, "error should be nil")
	
	// Arrange
	b = &Board{Fields: [nRows][nCols]color{
		{red, red, red, blue, red, red, red},
		{red, red, blue, blue, red, red, red},
		{blue, red, blue, blue, red, red, red},
		{red, blue, red, red, blue, blue, blue},
		{blue, red, red, blue, red, red, red},
		{blue, red, blue, blue, red, red, red},
	}, winner: none, NextColor: red,
	}
	// Act
	column, err = MC{}.NextTurn(b)
	// Assert
	require.Equal(t, -1, column, "should return -1 if board is already full")
	require.Equal(t, NewMatchIsOverError("board is already full"), err, "should return a NewMatchIsOverError if the board is already full")

	// Arrange
	b = &Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, red, none, none, none, none},
		{red, blue, blue, blue, blue, none, red},
	}, winner: none, NextColor: red,
	}
	// Act
	column, err = MC{}.NextTurn(b)
	// Assert
	require.Equal(t, -1, column, "should return -1 if the board has a winner already")
	require.Equal(t, NewMatchIsOverError("match has already a winner"), err, "should return a NewMatchIsOverError if the board has a winner already")
}
	