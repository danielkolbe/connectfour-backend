package game

import (
	"sync"
	"testing"
	"github.com/stretchr/testify/require"
)

type TestGameDb struct{
	games *map[string]*Board
	mutex sync.Mutex
}

var testGameDb = TestGameDb{games: &gameDb}

func TestTurn(t *testing.T) {
	// Arrange
	testGameDb.mutex.Lock()
	defer func() {
		gameDb = map[string] *Board{}
		testGameDb.mutex.Unlock()
	}()	
	c := CFour{}
	// Assert
	require.Equal(t, 0, len(*testGameDb.games), "initial size of game database should be zero")
	// Act & Assert
	c.Turn(3, "id_1")
	require.Equal(t, 1, len(*testGameDb.games), "size of game database should be 1")
	// Act & Assert
	c.Turn(3, "id_1")
	require.Equal(t, 1, len(*testGameDb.games), "size of game database should be 1")
	// Act & Assert
	c.Turn(3, "id_2")
	require.Equal(t, 2, len(*testGameDb.games), "size of game database should be 2")
}

func TestBoard(t *testing.T) {
	// Arrange
	testGameDb.mutex.Lock()
	defer func() {
		gameDb = map[string] *Board{}
		testGameDb.mutex.Unlock()
	}()
	c := CFour{}
	// Assert
	require.Equal(t, 0, len(*testGameDb.games), "initial size of game database should be zero")
	// Act & Assert
	b := c.Board("id_2")
	require.Equal(t, *newBoard(), b, "should create a new board after requesting a non-existing board")
	require.Equal(t, 1, len(*testGameDb.games), "size of game database should be increased by 1 after requesting a non-existing board")
	// Act & Assert
	b = c.Board("id_3")	
	require.Equal(t, 2, len(*testGameDb.games), "size of game database should be increased by 1 after requesting a non-existing board")
    // Act & Assert
	gameDb["id_2"].addChip(0)
	b = c.Board("id_2")	
    require.Equal(t, red, b.Fields[5][0], "should return the right board (the one with red chip added)")
}

func TestWinner(t *testing.T) {
	// Arrange
	testGameDb.mutex.Lock()
	defer func() {
		gameDb = map[string] *Board{}
		testGameDb.mutex.Unlock()
	}()
	c := CFour{}
	// Act
	color, error := c.Winner("unknown")
	// Assert
	require.NotEqual(t, nil, error, "should return an non-nil error if board does not exist")
	require.Equal(t, NewBoardDoesNotExistError("unknown").Error(), error.Error(), "should return a NewBoardDoesNotExistError if board does not exit")
	require.Equal(t, "", color, "should return an empty string for color if board does not exit")
	
	// Act & Assert
	c.Board("id_1")
	color, error = c.Winner("id_1")
	require.Equal(t, nil, error, "should return an nil error if board does exist")
	require.Equal(t, "n",color, "should return none for winning color on new board")
}
