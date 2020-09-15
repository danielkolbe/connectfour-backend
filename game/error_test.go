package game

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestErrors(t *testing.T) {
	// Act & Assert
	require.Equal(t, "message", NewMatchIsOverError("message").Error(), "should contain the correct error message")
	// Act & Assert
	require.Equal(t, "board with gameID <id_1> does not exist", NewBoardDoesNotExistError("id_1").Error(), "should contain the correct error message")
	// Act & Assert
	require.Equal(t, "column 2 is full", NewColumnIsFullError(2).Error(), "should contain the correct error message")
	// Act & Assert
	require.Equal(t, "column 1 is out of bounds: 0-6", NewColumnIsOutOfBoundsError(1).Error(), "should contain the correct error message")
}
