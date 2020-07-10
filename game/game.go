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
}

// Turn calls Board.addChip with the given column
// on the Board with the given gameID if such a Board
// already exists. If not a new Board with given gameID
// wil be created and addChip be called on it.
func (c CFour) Turn(column int, gameID string) error {
	var b *Board
	if bo, ok := gameDb[gameID]; ok {
		b = bo
	} else {
		b = newBoard()
		gameDb[gameID] = b
	}

	return b.addChip(column)
}

// Board returns a copy of the game board that belongs to 
// the given game id. If such a board does not exist a new board
// with this id will be instantiated and returned.
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

