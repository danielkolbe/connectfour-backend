package game

type Color int

const(
	NONE Color = iota
	BLUE
	RED
)

type Board [7][6] Color

func NewBoard() Board {
	return [7][6]Color{};
}

func (b *Board) AddChip(column int, color Color){
	for i := len(b)-1; i>=0; i-- {
		if(NONE == b[i][column]) {
			b[i][column] = color
			break
		}
	}
}