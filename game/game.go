package game

var gameDb = map[string]*Board{}

// CFour implements the game.Service interface.
type CFour struct{}

// Service provides a collection of function
// that can perform actions on a game.Board.
type Service interface {
	// Perform the next turn on the gameboard with the given id
	// on the column with the given number.
	Turn(int, string) error
	// Returns the game board for the given id or a new one that
	// can be referenced with the id on later calls.
	Board(string) Board
	// Winner returns the winning color for the board 
    // with the given gameID. 
	Winner(string) (string, error)
    // Reset sets the board with given gameID back to its initial state.
	Reset(string) (error)
}

// AI provides a collection of functions
// used to compute artifical intelligence turns
// based on a given game board.
type AI interface {
	NextTurn(b *Board) int
}

// Turn calls Board.addChip with the given column
// on the Board with the given gameID. It returns an 
// BoardDoesNotExistError if no such board exists.
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
		return  nil
	} 
	return NewBoardDoesNotExistError(gameID)
}
