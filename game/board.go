package game

type Color int

const(
	NONE Color = iota
	BLUE
	RED
)

type Board struct {
	COLORS [7][6] Color
}

func (b Board) addChip(row int, color Color){

}