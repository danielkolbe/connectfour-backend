package game

import (
	"math/rand"
	"time"
)

// used for to enable mocking in unit tests
var randomInt func (n int) int = rand.Intn

// number of random matches that will be played
// for each available column
const rep = 500

type result struct {
    column int
    probability float64
}

// NextTurn performs a Monte Carlo algorithm to determine
// which of the remaining non-full columns of the given board
// the next chip should be inserted in order the maximize the chances
// to win for the side that chip belongs to. Returns an error if the given
// board is not in an legal state for turn computation.
// Steps:
// 1) Choose a column that is not completely filled yet and insert the next chip
// 2) Take the board with the inserted next chip and play rep many random games
//    all beginning with the game constellation of that board
// 3) Determine the ratio of won matches (won by the side that owns the chip inserted in step 1)
// 4) Repeat step 2 and 3 for all remaining non-full columns
// 5) Return the index of the column with the best ratio (highest empirical likelihood of winning) 
func (mc MC) NextTurn(b *Board) (int, error) {
	fColumns := b.freeColumns()
	if 0 == len(fColumns) {
		return -1, NewMatchIsOverError("board is already full")
	}
	channel := make(chan result, len(fColumns))
	for _, column := range(fColumns) {
		nb := newBoard()
		copy(nb.Fields[:], b.Fields[:])
		err := nb.addChip(column)
		if nil != err {
			return -1, err
		}
		go func(c chan result, b *Board, column int) {
			c <- result{probability: 1- empiricalLikelihoodOfWinning(b, rep), column: column}
		}(channel, nb, column)
	}
	bestResult := result{column: -1, probability: -1.0}
	for i := 1; i <= len(fColumns); i++ {
		result := <- channel
		if result.probability > bestResult.probability {
			bestResult = result
		}	
	}
	close(channel)

	return bestResult.column, nil;
}

// empiricalLikelihoodOfWinning performs rep many random matches
// all with the given board as initial board. Returns the (empirical)
// likelihood of a victory of the player who's turn is next on the
// given board. 
func empiricalLikelihoodOfWinning(b *Board, rep int) float64 {
	player := b.NextColor
	count := 0
	for i :=1; i<=rep; i++ {
		if player == randomMatch(*b) {
			count ++
		}
	}
	return float64(count)/float64(rep);
}

// performs a random match, using the given
// board as starting point, by choosing the 
// column to add the next chip to randomly among
// the remaining columns. Returns the color of the winner, 
// or none if its a draw.
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
