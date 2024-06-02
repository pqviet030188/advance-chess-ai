package board_dictionary_tests

import (
	"testing"

	. "github.com/pqviet030188/advance-chess-ai/bitboard"
)

func TestHorizontalSlidingMoves(t *testing.T) {

	board := NewBitboardFromStr(`
		000000000
		000001000
		000000000
		101000010
		011010100
		000011100
		000000000
		000000000
		000000000
	`)

	expected := NewBitboardFromStr(`
		000000000
		000001000
		000000000
		001110110
		011010100
		000011100
		000000000
		000000000
		000000000
	`)

	res := board.CalculateHorizontalSlidingMoves(48)

	boardValue := res.Uint96

	if !boardValue.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", boardValue, *expected.Uint96)
	}
}
