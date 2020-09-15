package game

var gameDb = map[string]*Board{}

// CFour implements the game.Service interface.
type CFour struct{}

// MC is used to implement the game.AI interface
// using a MonteCarlo algorithm
type MC struct{}

// Service provides a collection of functions
// that can perform actions on a game.Board.
type Service interface {
	// TurnAI performs next turn on the gameboard with the given id
	// using an artificial intelligence algorithm. Returns column
	// where the next chip was inserted or an error.
	TurnAI(gameID string, ai AI) (int, error)
	// Turn performs the next turn on the gameboard with the given id
	// on the column with the given number.
	Turn(int, string) error
	// Board returns the game board for the given id or a new one that
	// can be referenced with the id on later calls.
	Board(string) Board
	// Winner returns the winning color for the board
	// with the given gameID.
	Winner(string) (string, error)
	// Reset sets the board with given gameID back to its initial state.
	Reset(string) error
}

// AI provides a collection of functions
// used to compute artifical intelligence turns
// based on a given game board.
type AI interface {
	// NextTurn returns the index of the column
	// that has been computed as the next AI turn.
	NextTurn(b *Board) (int, error)
}

// TurnAI calls the NextTurn function on the given instance of
// AI passing the Board with the given gameID.
// It calls addChip with the column index returned by NextTurn and
// returns the column index. It returns an BoardDoesNotExistError
// if no board matching the given gameID exists, a specific error if
// the board is not in an legal state for such an operation.
func (c CFour) TurnAI(gameID string, ai AI) (int, error) {
	b, ok := gameDb[gameID]
	if !ok {
		return -1, NewBoardDoesNotExistError(gameID)
	}
	column, err := ai.NextTurn(b)
	if nil != err {
		return -1, err
	}

	return column, b.addChip(column)
}

// Turn calls Board.addChip with the given column
// on the Board with the given gameID. It returns an
// BoardDoesNotExistError if no such board exists,
// a specific error if the board is not in an legal
// state for such an operation.
func (c CFour) Turn(column int, gameID string) error {
	if b, ok := gameDb[gameID]; ok {
		return b.addChip(column)
	}
	return NewBoardDoesNotExistError(gameID)
}

// Board returns a copy of the game board that belongs to
// the given game id. If such a board does not exist a new board
// with that id will be instantiated and returned.
func (c CFour) Board(gameID string) Board {
	var b *Board
	if bo, ok := gameDb[gameID]; ok {
		b = bo
	} else {
		b = newBoard()
		gameDb[gameID] = b
	}
	return *b
}

// Winner returns the winning color for the board
// with the given gameID. It returns an
// BoardDoesNotExistError if no such board exists.
func (c CFour) Winner(gameID string) (string, error) {
	if board, ok := gameDb[gameID]; ok {
		return board.win().String(), nil
	}
	return "", NewBoardDoesNotExistError(gameID)
}

// Reset sets the board with given gameID back to its initial state.
// It returns an BoardDoesNotExistError if no such board exists.
func (c CFour) Reset(gameID string) error {
	if _, ok := gameDb[gameID]; ok {
		gameDb[gameID] = newBoard()
		return nil
	}
	return NewBoardDoesNotExistError(gameID)
}
