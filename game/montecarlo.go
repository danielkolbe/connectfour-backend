package game

import (
	"fmt"
	"math/rand"
	"time"
)

var randomInt func (n int) int = rand.Intn
const rep = 500

type result struct {
    column int
    probability float64
}

func NextTurn(b *Board) int{
	fColumns := b.freeColumns()
	channel := make(chan result, len(fColumns))
	for column := range(fColumns) {
		c := newBoard()
		copy(c.Fields[:], b.Fields[:])
		c.addChip(column)
		go func(c chan result, b *Board, column int) {
			c <-result{probability: 1-empiricalLikelihoodOfWinning(b, rep), column: column}
		}(channel, c, column)
	}
	bestResult := result{column: -1, probability: -1.0}
	for i := 1; i <= len(fColumns); i++ {
		result := <- channel
		if result.probability > bestResult.probability {
			bestResult = result
		}	
	}
	close(channel)

	return bestResult.column;
}

func empiricalLikelihoodOfWinning(b *Board, rep int) float64 {
	player := b.nextColor
	count := 0
	for i :=1; i<=rep; i++ {
		if player == randomMatch(*b) {
			count ++
		}
	}
	fmt.Println(float64(count)/float64(rep))
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
