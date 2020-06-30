package game

import "fmt"

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
	fmt.Printf("gameid: %v\n", gameId)
	for _, row := range b.fields {
		fmt.Printf("%v\n", row)
	}
	
	return b.addChip(column)	
}