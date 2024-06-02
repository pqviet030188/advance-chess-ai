package board_dictionary_tests

import (
	"testing"

	. "github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/uint96"
)

func TestNewFromString(t *testing.T) {
	board := NewBitboardFromStr(`
		111111111
		111111111
		111111111
		111111111
		111111111
		111111111
		111111111
		111111111
		111111111
	`)

	expected := uint96.Uint96{
		Hi:  0x0001ffff,
		Mid: 0xffffffff,
		Lo:  0xffffffff,
	}
	if !board.Uint96.Equals(expected) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *board.Uint96, expected)
	}

	board2 := NewBitboardFromStr(`
		000000001
		000000001
		000000001
		000000001
		000000001
		000000001
		000000001
		000000001
		000000001
	`)

	expected2 := uint96.Uint96{
		Hi:  0x00000100,
		Mid: 0x80402010,
		Lo:  0x08040201,
	}
	if !board2.Uint96.Equals(expected2) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *board2.Uint96, expected2)
	}

	board3 := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	expected3 := uint96.Uint96{
		Hi:  0x00000000,
		Mid: 0x00000000,
		Lo:  0x00000000,
	}
	if !board3.Uint96.Equals(expected3) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *board3.Uint96, expected3)
	}
}
