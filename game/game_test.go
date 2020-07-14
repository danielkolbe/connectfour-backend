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
	// Act & Assert
	error := c.Turn(3, "id_1")
	require.NotEqual(t, nil, error, "should return an BoardDoesNotExistError")
	require.Equal(t, NewBoardDoesNotExistError("id_1").Error(), error.Error(), "should return an BoardDoesNotExistError")
	
	// Act & Assert
	(*testGameDb.games)["id_1"] = newBoard()
	error=c.Turn(3, "id_1")
	require.Equal(t, nil, error, "should return an nil-error if board exists")
	require.Equal(t,
		"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n r n n n \n", (*testGameDb.games)["id_1"].String(), "should call addChip if board exists")
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
func TestReset(t *testing.T) {
	// Arrange
	testGameDb.mutex.Lock()
	defer func() {
		gameDb = map[string] *Board{}
		testGameDb.mutex.Unlock()
	}()
	c := CFour{}
	// Arrange 
	b := newBoard()
	b.addChip(2)
	(*testGameDb.games)["id_1"] = b
	// Act
	error := c.Reset("id_1")
	require.Equal(t, newBoard(), (*testGameDb.games)["id_1"], "should reset board to initial state")
	require.Equal(t, nil, error, "should not return an error if board exists")
	
	// Act
	error = c.Reset("unknown")
	// Assert
	require.Equal(t, NewBoardDoesNotExistError("unknown").Error(), error.Error(), "should return a NewBoardDoesNotExistError if board does not exit")
}
