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

func TestFindfindwinner(t *testing.T) {
	// Arrange
	b := Board{Fields: [nRows][nCols]color{
		{none, none, none, blue, none, none, none},
		{none, none, none, none, blue, none, none},
		{none, none, none, none, none, blue, none},
		{none, none, none, none, none, none, blue},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, blue, findwinner(&b), "should return blue (diagonal win)")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, blue, blue, blue, blue},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, blue, findwinner(&b), "should return blue (horizontal win)")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, red},
		{none, none, none, none, none, none, red},
		{none, none, none, none, none, none, red},
		{none, none, none, none, none, none, red},
	},
	}
	// Act & Assert
	require.Equal(t, red, findwinner(&b), "should return red (vertical win)")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, blue, blue, blue, none, none},
		{red, none, none, none, none, none, red},
		{none, red, none, none, none, none, red},
		{none, none, red, none, none, none, red},
	},
	}
	// Act & Assert
	require.Equal(t, none, findwinner(&b), "should return none (no win)")

}

func TestWinDiagonal(t *testing.T) {
	// Arrange
	b := Board{Fields: [nRows][nCols]color{
		{none, none, none, blue, none, none, none},
		{none, none, none, none, blue, none, none},
		{none, none, none, none, none, blue, none},
		{none, none, none, none, none, none, blue},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, blue, winDiagonal(&b), "should return blue")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, blue, none, none, none},
		{none, red, none, none, blue, none, none},
		{none, none, red, none, none, blue, none},
		{none, none, none, red, none, none, none},
		{none, none, none, none, red, none, none},
		{none, none, none, none, none, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, red, winDiagonal(&b), "should return red")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, blue, none, none, none},
		{none, none, none, none, blue, none, none},
		{none, none, red, none, none, blue, none},
		{none, none, none, red, none, none, none},
		{none, none, none, none, red, none, none},
		{none, none, none, none, none, red, none},
	},
	}
	// Act & Assert
	require.Equal(t, red, winDiagonal(&b), "should return red")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{blue, none, red, none, none, none, none},
		{none, blue, none, red, none, none, none},
		{none, none, blue, none, red, none, none},
		{none, none, none, blue, none, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, blue, winDiagonal(&b), "should return blue")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, blue, blue, blue, blue},
		{red, none, none, none, blue, none, none},
		{red, none, red, none, none, blue, none},
		{red, blue, none, red, none, none, none},
		{red, none, blue, none, red, none, none},
		{red, none, none, blue, none, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, none, winDiagonal(&b), "should return none")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, red, none, red, none, none},
		{none, none, none, red, none, none, none},
		{none, none, red, none, red, none, none},
		{none, red, none, blue, none, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, red, winDiagonal(&b), "should return red")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, blue, none},
		{none, none, none, none, blue, none, none},
		{none, none, red, blue, red, none, none},
		{none, none, blue, none, none, none, none},
		{none, none, red, none, red, none, none},
		{none, red, none, blue, none, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, blue, winDiagonal(&b), "should return blue")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, red, none, red, none},
		{none, none, red, none, blue, none, none},
		{none, red, red, blue, red, none, none},
		{red, none, blue, none, none, none, none},
		{none, none, red, none, red, none, none},
		{none, red, none, blue, none, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, red, winDiagonal(&b), "should return red")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, red, red, none, blue},
		{red, none, none, none, blue, blue, none},
		{red, none, red, none, blue, blue, none},
		{red, blue, none, blue, none, none, none},
		{red, none, blue, none, red, none, none},
		{red, none, none, blue, none, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, blue, winDiagonal(&b), "should return blue")
}

func TestWinHorizontal(t *testing.T) {
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
	require.Equal(t, red, winHorizontal(&b), "should return red")

	// Act & Assert
	require.Equal(t, red, winHorizontal(&b), "should return red")

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
	require.Equal(t, blue, winHorizontal(&b), "should return blue")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, red, none, blue, none, none, none},
		{none, red, none, none, blue, none, none},
		{none, red, none, none, none, blue, none},
		{none, red, none, none, none, none, blue},
		{none, red, blue, none, blue, none, none},
		{none, red, red, red, blue, red, none},
	},
	}
	// Act & Assert
	require.Equal(t, none, winHorizontal(&b), "should return none")
}

func TestWinVertical(t *testing.T) {
	// Arrange
	b := Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, blue, none, none, none, none, none},
		{none, blue, none, none, none, none, none},
		{none, blue, none, none, none, none, none},
		{none, blue, none, none, none, none, none},
		{none, red, red, red, red, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, blue, winVertical(&b), "should return blue")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{blue, none, none, none, none, none, none},
		{blue, none, none, none, none, none, none},
		{blue, none, none, none, none, none, none},
		{blue, none, none, none, none, none, none},
		{red, red, red, red, red, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, blue, winVertical(&b), "should return blue")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, blue},
		{none, none, none, none, none, none, blue},
		{none, none, none, none, none, none, blue},
		{none, none, none, none, none, none, blue},
		{none, none, none, none, none, none, blue},
		{none, red, red, red, red, none, red},
	},
	}
	// Act & Assert
	require.Equal(t, blue, winVertical(&b), "should return blue")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, blue, none, none, none},
		{none, none, none, none, blue, none, none},
		{none, red, none, none, none, blue, none},
		{none, red, none, none, none, none, blue},
		{none, red, blue, none, blue, none, none},
		{none, red, red, red, blue, red, none},
	},
	}
	// Act & Assert
	require.Equal(t, red, winVertical(&b), "should return red")

	// Arrange
	b = Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, blue, none, none, none, none, none},
		{none, red, none, none, none, none, none},
		{none, red, blue, none, blue, none, none},
		{none, red, red, red, red, none, none},
	},
	}
	// Act & Assert
	require.Equal(t, none, winVertical(&b), "should return none")
}

func TestDiagonalTopLeftBottomRight(t *testing.T) {
	// Arrange
	b := Board{Fields: [nRows][nCols]color{
		{none, none, none, blue, none, red, red},
		{none, blue, none, none, blue, none, none},
		{none, blue, none, red, none, red, blue},
		{none, blue, none, none, none, none, red},
		{none, blue, none, none, none, red, none},
		{none, red, red, red, red, blue, none},
	},
	}
	// Act & Assert
	require.Equal(t, []color{none, blue, none, none, none, blue}, diagonalTopLeftBottomRight(&b.Fields, 0, 0), "should return the diagonal with given starting point")
	require.Equal(t, []color{none, blue, none, none, red}, diagonalTopLeftBottomRight(&b.Fields, 1, 0), "should return the diagonal with given starting point")
	require.Equal(t, []color{none, blue, none, red}, diagonalTopLeftBottomRight(&b.Fields, 2, 0), "should return the diagonal with given starting point")
	require.Equal(t, []color{none, blue, red}, diagonalTopLeftBottomRight(&b.Fields, 3, 0), "should return the diagonal with given starting point")
	require.Equal(t, []color{none, red}, diagonalTopLeftBottomRight(&b.Fields, 4, 0), "should return the diagonal with given starting point")
	require.Equal(t, []color{none}, diagonalTopLeftBottomRight(&b.Fields, 5, 0), "should return the diagonal with given starting point")

	require.Equal(t, []color{none, none, red, none, red, none}, diagonalTopLeftBottomRight(&b.Fields, 0, 1), "should return the diagonal with given starting point")
	require.Equal(t, []color{none, none, none, none, none}, diagonalTopLeftBottomRight(&b.Fields, 0, 2), "should return the diagonal with given starting point")
	require.Equal(t, []color{blue, blue, red, red}, diagonalTopLeftBottomRight(&b.Fields, 0, 3), "should return the diagonal with given starting point")
	require.Equal(t, []color{none, none, blue}, diagonalTopLeftBottomRight(&b.Fields, 0, 4), "should return the diagonal with given starting point")
	require.Equal(t, []color{red, none}, diagonalTopLeftBottomRight(&b.Fields, 0, 5), "should return the diagonal with given starting point")
	require.Equal(t, []color{red}, diagonalTopLeftBottomRight(&b.Fields, 0, 6), "should return the diagonal with given starting point")

	require.Equal(t, []color{none, none, red}, diagonalTopLeftBottomRight(&b.Fields, 3, 2), "should return the diagonal with given starting point")
}

func TestDiagonalTopRightBottomLeft(t *testing.T) {
	// Arrange
	b := Board{Fields: [nRows][nCols]color{
		{none, none, none, blue, none, red, red},
		{none, blue, none, none, blue, none, none},
		{none, blue, none, red, none, red, blue},
		{none, blue, none, none, none, none, red},
		{none, blue, none, none, none, red, none},
		{none, red, red, red, red, blue, none},
	},
	}
	// Act & Assert
	require.Equal(t, []color{red, none, none, none, none, red}, diagonalTopRightBottomLeft(&b.Fields, 0, 6), "should return the diagonal with given starting point")
	require.Equal(t, []color{none, red, none, none, red}, diagonalTopRightBottomLeft(&b.Fields, 1, 6), "should return the diagonal with given starting point")
	require.Equal(t, []color{blue, none, none, red}, diagonalTopRightBottomLeft(&b.Fields, 2, 6), "should return the diagonal with given starting point")
	require.Equal(t, []color{red, red, red}, diagonalTopRightBottomLeft(&b.Fields, 3, 6), "should return the diagonal with given starting point")
	require.Equal(t, []color{none, blue}, diagonalTopRightBottomLeft(&b.Fields, 4, 6), "should return the diagonal with given starting point")
	require.Equal(t, []color{none}, diagonalTopRightBottomLeft(&b.Fields, 5, 6), "should return the diagonal with given starting point")

	require.Equal(t, []color{red, blue, red, none, blue, none}, diagonalTopRightBottomLeft(&b.Fields, 0, 5), "should return the diagonal with given starting point")
	require.Equal(t, []color{none, none, none, blue, none}, diagonalTopRightBottomLeft(&b.Fields, 0, 4), "should return the diagonal with given starting point")
	require.Equal(t, []color{blue, none, blue, none}, diagonalTopRightBottomLeft(&b.Fields, 0, 3), "should return the diagonal with given starting point")
	require.Equal(t, []color{none, blue, none}, diagonalTopRightBottomLeft(&b.Fields, 0, 2), "should return the diagonal with given starting point")
	require.Equal(t, []color{none, none}, diagonalTopRightBottomLeft(&b.Fields, 0, 1), "should return the diagonal with given starting point")
	require.Equal(t, []color{none}, diagonalTopRightBottomLeft(&b.Fields, 0, 0), "should return the diagonal with given starting point")

	require.Equal(t, []color{none, blue, none}, diagonalTopRightBottomLeft(&b.Fields, 3, 2), "should return the diagonal with given starting point")
}

func BenchmarkFindWinner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findwinner(newBoard())
	}
}
