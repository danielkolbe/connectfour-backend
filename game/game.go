package game

var gameDb = map[string]*board{}

type GameService interface {
	Turn(int, string) error
}

type CFour struct {} 

func (c CFour) Turn(column int, gameId string) error {
	var b *board
	if bo, ok := gameDb[gameId]; ok {
		b = bo
	} else {
		b = NewBoard()
		gameDb[gameId] = b
	}

	return b.addChip(column)	
}