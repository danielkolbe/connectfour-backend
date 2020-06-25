package game;

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()
	for i, row := range b.Fields {
		for j, color := range row {
			if NONE != color {
				t.Errorf("A newly created board should be blank but the color of field %v,%v was %v", i, j, Color(1))
			}
		}
	}
}
func TestError(t *testing.T) {
	b := NewBoard()
	for i:=0 ; i < nRows; i++ {
		error := b.AddChip(0)
		if nil != error {
			t.Errorf("Error should be nil but was: `%v`.", error)
		}
	}
	error := b.AddChip(0)
	if nil == error {
		t.Errorf("No error returned.")
	}
}
func TestAddChip(t *testing.T) {
	b := NewBoard()
	b.AddChip(4)
	b.AddChip(4)
	if RED != b.Fields[5][4] {
		t.Errorf("Color of field 6,4 of the board should be red but was %v", Color(b.Fields[5][4]))
	}
	if BLUE != b.Fields[4][4] {
		t.Errorf("Color of field 6,4 of the board should be blue but was %v", Color(b.Fields[4][4]))
	}
}

func TestNextColor(t *testing.T) {
	b := NewBoard()
	if RED != b.nextColor {
		t.Errorf("Next color must be %v but was %v", NONE, b.nextColor)
	}
	b.AddChip(4)
	if BLUE !=  b.nextColor {
		t.Errorf("Next color must be %v but was %v", BLUE, b.nextColor)
	}
	b.AddChip(3)
	if RED !=  b.nextColor {
		t.Errorf("Next color must be %v but was %v", RED, b.nextColor)
	}
	b.AddChip(3)
	if BLUE !=  b.nextColor {
		t.Errorf("Next color must be %v but was %v", BLUE, b.nextColor)
	}
}		