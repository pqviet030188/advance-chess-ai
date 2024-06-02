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
		Lo:  0x0001ffff,
		Mid: 0xffffffff,
		Hi:  0xffffffff,
	}
	if !board.Uint96.Equals(expected) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *board.Uint96, expected)
	}
}
