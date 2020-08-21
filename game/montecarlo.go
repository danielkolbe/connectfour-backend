package game

import (
	"math/rand"
	"time"
)

var randomInt func (n int) int = rand.Intn

func empiricalLikelihoodOfWinning(b *Board, rep int) float64 {
	player := b.nextColor
	count := 0
	for i :=1; i<=rep; i++  {
		if player == randomMatch(*b) {
			count ++
		}
	}

	return float64(count)/float64(rep);
}

func randomMatch(b Board) color {
	for none == b.win() {
		rColumn := randomColumn(&b)
		if -1 == rColumn {
			return none
		}
		b.addChip(rColumn)
	}

	return b.winner;
}
 
func randomColumn(b *Board) int {
	rand.Seed(time.Now().UTC().UnixNano())
	fColumns := b.freeColumns()
	if 0 == len(fColumns) {
		return -1
	}
	randIndex := randomInt(len(fColumns))
	
	return fColumns[randIndex];
}
