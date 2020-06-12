package game;

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()
	b.AddChip(1, RED)
	for i, row := range b {
		for j, color := range row {
			if NONE != color {
				t.Errorf("A newly created board should be blank but the color of field %v,%v was %v", i, j, Color(1))
			}
		}
	}
}