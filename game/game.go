package game

var gameDb = map[string]*board{}

// CFour implements the game.GameService interface.
type CFour struct{}

// Service provides a collection of function
// that can perform actions on a game.board.
type Service interface {
	Turn(int, string) error
}

// Turn calls board.addChip with the given column
// on the board with the given gameID if such a board
// already exists. If not a new board with given gameID
// wil be created and addChip be called on it.
func (c CFour) Turn(column int, gameID string) error {
	var b *board
	if bo, ok := gameDb[gameID]; ok {
		b = bo
	} else {
		b = newBoard()
		gameDb[gameID] = b
	}

	return b.addChip(column)
}
