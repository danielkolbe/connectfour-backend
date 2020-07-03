package game

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTurn(t *testing.T) {
	c := CFour{}
	// Assert
	require.Equal(t, 0, len(gameDb), "initial size of game database should be zero")
	// Act & Assert
	c.Turn(3, "id_1")
	require.Equal(t, 1, len(gameDb), "size of game database should be 1")
	// Act & Assert
	c.Turn(3, "id_1")
	require.Equal(t, 1, len(gameDb), "size of game database should be 1")
	// Act & Assert
	c.Turn(3, "id_2")
	require.Equal(t, 2, len(gameDb), "size of game database should be 2")
}
