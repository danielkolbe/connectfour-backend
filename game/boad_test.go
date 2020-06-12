package game;

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()
	for i, row := range b {
		for j, color := range row {
			if NONE != color {
				t.Errorf("A newly created board should be blank but the color of field %v,%v was %v", i, j, Color(1))
			}
		}
	}
}

func TestAddChip(t *testing.T) {
	b := NewBoard()
	b.AddChip(4,2)
	b.AddChip(4,1)
	if RED != b[5][4] {
		t.Errorf("Color of field 6,4 of the board should be red but was %v", Color(b[5][4]))
	}
	if BLUE != b[4][4] {
		t.Errorf("Color of field 6,4 of the board should be blue but was %v", Color(b[4][4]))
	}
}