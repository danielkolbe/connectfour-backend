package game;

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()
	for i, row := range b.fields {
		for j, color := range row {
			if none != color {
				t.Errorf("A newly created board should be blank but the color of field %v,%v was %v", i, j, color)
			}
		}
	}
}
func TestError(t *testing.T) {
	b := NewBoard()
	for i:=0 ; i < nRows; i++ {
		error := b.addChip(0)
		if nil != error {
			t.Errorf("Error should be nil but was: `%v`.", error)
		}
	}
	error := b.addChip(0)
	if nil == error {
		t.Errorf("No error returned.")
	}
}
func TestAddChip(t *testing.T) {
	b := NewBoard()
	b.addChip(4)
	b.addChip(4)
	if red != b.fields[5][4] {
		t.Errorf("Color of field 6,4 of the board should be red but was %v", b.fields[5][4])
	}
	if blue != b.fields[4][4] {
		t.Errorf("Color of field 6,4 of the board should be blue but was %v", b.fields[4][4])
	}
}

func TestNextColor(t *testing.T) {
	b := NewBoard()
	if red != b.nextColor {
		t.Errorf("Next color must be %v but was %v", none, b.nextColor)
	}
	b.addChip(4)
	if blue !=  b.nextColor {
		t.Errorf("Next color must be %v but was %v", blue, b.nextColor)
	}
	b.addChip(3)
	if red !=  b.nextColor {
		t.Errorf("Next color must be %v but was %v", red, b.nextColor)
	}
	b.addChip(3)
	if blue !=  b.nextColor {
		t.Errorf("Next color must be %v but was %v", blue, b.nextColor)
	}
}		