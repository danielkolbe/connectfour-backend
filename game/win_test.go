package game

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHasFour(t *testing.T) {
	// Act & Assert
	require.Equal(t, red, hasFour([]color{none, red, blue, red, red, red, red}), "should return red")
	// Act & Assert
	require.Equal(t, blue, hasFour([]color{blue, blue, blue, blue, red, none}), "should return blue")
	// Act & Assert
	require.Equal(t, none, hasFour([]color{red, red, red, none, red}), "should return none")
}

func TestWinH(t *testing.T) {
	// Arrange
	b := Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, blue, none, none, none, none, none},
		{none, blue, none, none, none, none, none},
		{none, blue, none, none, none, none, none},
		{none, red, red, red, red, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, red, winH(&b), "should return red")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, blue, blue, blue, blue, none, none},
		{none, none, none, none, none, none, none},
		{none, red, none, none, none, none, none},
		{none, red, none, none, none, none, none},
		{none, red, blue, none, blue, none, none},
		{none, red, red, red, blue, red, none},
	},
	}
	// Act & Assert
	require.Equal(t, blue, winH(&b), "should return blue")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, red, none, none, none, none, none},
		{none, red, none, none, none, none, none},
		{none, red, none, none, none, none, none},
		{none, red, none, none, none, none, none},
		{none, red, blue, none, blue, none, none},
		{none, red, red, red, blue, red, none},
	},
	}
	// Act & Assert
	require.Equal(t, none, winH(&b), "should return none")
}
