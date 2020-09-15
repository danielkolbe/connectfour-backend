package game

import (
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

type TestGameDb struct {
	games *map[string]*Board
	mutex sync.Mutex
}

var testGameDb = TestGameDb{games: &gameDb}

func TestTurn(t *testing.T) {
	// Arrange
	testGameDb.mutex.Lock()
	defer func() {
		gameDb = map[string]*Board{}
		testGameDb.mutex.Unlock()
	}()
	c := CFour{}
	// Act & Assert
	error := c.Turn(3, "id_1")
	require.NotEqual(t, nil, error, "should return an BoardDoesNotExistError")
	require.Equal(t, NewBoardDoesNotExistError("id_1").Error(), error.Error(), "should return an BoardDoesNotExistError")

	// Arrange
	(*testGameDb.games)["id_1"] = newBoard()
	// Act
	error = c.Turn(3, "id_1")
	// Assert
	require.Equal(t, nil, error, "should return an nil-error if board exists")
	require.Equal(t,
		"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n r n n n \n", (*testGameDb.games)["id_1"].String(), "should call addChip if board exists")
}

func TestTurnAI(t *testing.T) {
	// Arrange
	testGameDb.mutex.Lock()
	defer func() {
		gameDb = map[string]*Board{}
		testGameDb.mutex.Unlock()
	}()
	c := CFour{}
	// Act & Assert
	column, error := c.TurnAI("id_1", MC{})
	require.NotEqual(t, nil, error, "should return an BoardDoesNotExistError")
	require.Equal(t, NewBoardDoesNotExistError("id_1").Error(), error.Error(), "should return an BoardDoesNotExistError")
	require.Equal(t, -1, column, "should return -1 for column value if board does not exist")

	// Arange
	(*testGameDb.games)["id_1"] = &Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, red, none, none, none, none},
		{red, blue, blue, blue, blue, none, red},
	}, winner: none, NextColor: red,
	}
	// Act
	column, error = c.TurnAI("id_1", MC{})
	// Assert
	require.Equal(t, NewMatchIsOverError("match has already a winner"), error, "should forward error from NextTurn method")
	require.Equal(t, -1, column, "should return -1 for column value if error returned by NextTurn method")

	// Arange
	(*testGameDb.games)["id_1"] = &Board{Fields: [nRows][nCols]color{
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, none, none, none, none, none},
		{none, none, red, none, none, none, none},
		{blue, blue, blue, none, none, none, none},
	}, winner: none, NextColor: red,
	}
	// Act
	column, error = c.TurnAI("id_1", MC{})
	// Assert
	require.Equal(t,
		"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n n n n n n \n"+
			"n n r n n n n \n"+
			"b b b r n n n \n", (*testGameDb.games)["id_1"].String(), "should add chip to column recommended by ai")
	require.Equal(t, 3, column, "should the column recommended by ai")
}

func TestBoard(t *testing.T) {
	// Arrange
	testGameDb.mutex.Lock()
	defer func() {
		gameDb = map[string]*Board{}
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
		gameDb = map[string]*Board{}
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
	require.Equal(t, "n", color, "should return none for winning color on new board")
}
func TestReset(t *testing.T) {
	// Arrange
	testGameDb.mutex.Lock()
	defer func() {
		gameDb = map[string]*Board{}
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
