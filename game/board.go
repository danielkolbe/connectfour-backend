package game

type Color int
type Board [nRows][nCol] Color
const nRows int = 6
const nCol int = 7

const(
	NONE Color = iota
	BLUE
	RED
)

func NewBoard() Board {
	return [nRows][nCol]Color{}
}

func (b *Board) AddChip(column int, color Color){
	for i := len(b)-1; i>=0; i-- {
		if(NONE == b[i][column]) {
			b[i][column] = color
			break
		}
	}
}