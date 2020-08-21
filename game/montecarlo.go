package game

import (
	"math/rand"
	"time"
)

var randomInt func (n int) int = rand.Intn

func randomColumn(b *Board) int {
	rand.Seed(time.Now().UTC().UnixNano())
	fColumns := b.freeColumns()
	if 0 == len(fColumns) {
		return -1
	}
	randIndex := randomInt(len(fColumns))
	
	return fColumns[randIndex];
}
