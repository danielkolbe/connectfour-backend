package game

type Color int
const nRows int = 6;
const nCol int = 7;

const(
	NONE Color = iota
	BLUE
	RED
)

type Board [nRows][nCol] Color

func NewBoard() Board {
	return [nRows][nCol]Color{};
}

func (b *Board) AddChip(column int, color Color){
	for i := len(b)-1; i>=0; i-- {
		if(NONE == b[i][column]) {
			b[i][column] = color
			break
		}
	}
}