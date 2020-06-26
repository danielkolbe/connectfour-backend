package game

import (
	"testing"
)

func TestTurn(t *testing.T) {
	if 0 != len(gameDb) {
		t.Errorf("Initial size of game database should be zero but was: %v", len(gameDb))
	}
	Turn(3, "id_1")	
	
	if 1 != len(gameDb) {
		t.Errorf("Size of game database should be 1 but was: %v", len(gameDb))
	}
	Turn(3, "id_1")
	
	if 1 != len(gameDb) {
		t.Errorf("Size of game database should be 1 but was: %v", len(gameDb))
	}
	Turn(3, "id_2")
	
	if 2 != len(gameDb) {
		t.Errorf("Size of game database should be 2 but was: %v", len(gameDb))
	}
}
