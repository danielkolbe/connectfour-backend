package game

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestTurn(t *testing.T) {
	require.Equal(t, 0, len(gameDb), "initial size of game database should be zero")
	Turn(3, "id_1")	
	require.Equal(t, 1, len(gameDb), "size of game database should be 1")
	Turn(3, "id_1")
	require.Equal(t, 1, len(gameDb), "size of game database should be 1")
	Turn(3, "id_2")
	require.Equal(t, 2, len(gameDb), "size of game database should be 2")
}
