package game

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestHasFour(t *testing.T) {
	// Act & Assert
	require.Equal(t, red, hasFour(&[]color{none, red, blue, red, red, red, red}), "should return red")
	// Act & Assert
	require.Equal(t, blue, hasFour(&[]color{blue, blue, blue, blue, red, none}), "should return blue")
	// Act & Assert
	require.Equal(t, none, hasFour(&[]color{red, red, red, none, red}), "should return none")
}
