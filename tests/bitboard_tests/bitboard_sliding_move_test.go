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

	res := board.CalculateHorizontalSlidingMoves(F6)

	boardValue := res.Uint96

	if !boardValue.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", boardValue, *expected.Uint96)
	}
}

func TestVerticalSlidingMoves(t *testing.T) {

	board := NewBitboardFromStr(`
		000001000
		000001000
		000000000
		101000010
		011010100
		000011100
		000000000
		000001000
		000001000
	`)

	expected := NewBitboardFromStr(`
		000000000
		000001000
		000001000
		101000010
		011011100
		000011100
		000000000
		000000000
		000000000
	`)

	res := board.CalculateVerticalSlidingMoves(F6)

	boardValue := res.Uint96

	if !boardValue.Equals(*expected.Uint96) {
		// fmt.Printf("board: \n%s \n", board.Rep())
		// fmt.Printf("expected: \n%s \n", expected.Rep())
		// fmt.Printf("result: \n%s \n", res.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", boardValue, *expected.Uint96)
	}
}

func TestLRTBDiagSlidingMoves(t *testing.T) {

	board := NewBitboardFromStr(`
		100000000
		010001000
		000001000
		101000010
		011011100
		000010100
		000000100
		000000010
		000000001
	`)

	expected := NewBitboardFromStr(`
		000000000
		010001000
		001001000
		101100010
		011001100
		000011100
		000000100
		000000000
		000000000
	`)

	res := board.CalculateLRTBDiagSlidingMoves(E5)

	boardValue := res.Uint96

	if !boardValue.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", boardValue, *expected.Uint96)
	}
}

func TestLRBTDiagSlidingMoves(t *testing.T) {

	board := NewBitboardFromStr(`
		100000001
		010001000
		000001100
		101000010
		011011100
		000010100
		000000100
		010000010
		100000001
	`)

	expected := NewBitboardFromStr(`
		100000000
		010001000
		000001100
		101001010
		011001100
		000110100
		001000100
		010000010
		000000001
	`)

	res := board.CalculateLRBTDiagSlidingMoves(E5)

	boardValue := res.Uint96

	if !boardValue.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", boardValue, *expected.Uint96)
	}
}
