package bitboard_tests

import (
	"fmt"
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

func TestShiftToMoveCalPositionForLRTB(t *testing.T) {
	zero := uint96.FromUInt32(0)
	expected := Bitboard{
		Uint96: &zero,
	}

	index := uint8(0)
	for square := int(SIZE) - 1; square >= 0; square = square - 1 {
		mask := GenerateLRTBMask(uint8(square))
		expected.SetBit(index, 1)
		result, _ := mask.ShiftToMoveCalPositionForLRTB(uint8(square))
		reversedResult := result.ReverseShiftToMoveCalPositionForLRTB(uint8(square))

		if !result.Uint96.Equals(*expected.Uint96) {
			fmt.Printf("%s\n%s\n%s\n%d\n", mask.Rep(), result.Rep(), expected.Rep(), square)
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *result.Uint96, *expected.Uint96)
		}

		if !reversedResult.Uint96.Equals(*mask.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *reversedResult.Uint96, *mask.Uint96)
		}

		index = index + SIZE + 1
	}

	zero = uint96.FromUInt32(0)
	expected = *NewBitboardFromStr(`
		100000000
		010000000
		001000000
		000100000
		000010000
		000001000
		000000100
		000000010
		000000001
	`)
	index = uint8(0)
	for vsquare := 0; vsquare < int(SIZE); vsquare = vsquare + 1 {
		square := ToSquare(uint8(vsquare), 0)
		mask := GenerateLRTBMask(square)

		if int(index-SIZE-1) >= 0 {
			expected.SetBit(index-SIZE-1, 0)
		}

		result, _ := mask.ShiftToMoveCalPositionForLRTB(square)
		reversedResult := result.ReverseShiftToMoveCalPositionForLRTB(square)

		if !result.Uint96.Equals(*expected.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x", *result.Uint96, *expected.Uint96)
		}

		if !reversedResult.Uint96.Equals(*mask.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *reversedResult.Uint96, *mask.Uint96)
		}

		index = index + SIZE + 1
	}
}

func TestShiftToMoveCalPositionForLRBT(t *testing.T) {
	// zero := uint96.FromUInt32(0)
	expected := *NewBitboardFromStr(`
		000000001
		000000010
		000000100
		000001000
		000010000
		000100000
		001000000
		010000000
		100000000
	`)
	index := -1
	for col := int(SIZE - 1); col >= 0; col = col - 1 {
		square := ToSquare(uint8(0), uint8(col))
		mask := GenerateLRBTMask(uint8(square))

		if index >= 0 {
			bitSquare := ToSquare(SIZE-1-uint8(index), uint8(index))
			expected.SetBit(bitSquare, 0)
		}

		result, _ := mask.ShiftToMoveCalPositionForLRBT(square)
		reversedResult := result.ReverseShiftToMoveCalPositionForLRBT(square)

		if !result.Uint96.Equals(*expected.Uint96) {
			// fmt.Printf("%s\n%s\n%s\n%d", mask.Rep(), result.Rep(), expected.Rep(), square)
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x", *result.Uint96, *expected.Uint96)
		}

		if !reversedResult.Uint96.Equals(*mask.Uint96) {
			// fmt.Printf("board:\n%s\nresult:\n%s\nexpected:\n%s\nreverse:\n%s\n%d", mask.Rep(), result.Rep(), expected.Rep(), reversedResult.Rep(), square)
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *reversedResult.Uint96, *mask.Uint96)
		}

		index = index + 1
	}

	expected = *NewBitboardFromStr(`
		000000001
		000000010
		000000100
		000001000
		000010000
		000100000
		001000000
		010000000
		100000000
	`)
	index = -1
	for col := 0; col < int(SIZE); col = col + 1 {
		square := ToSquare(uint8(SIZE-1), uint8(col))
		mask := GenerateLRBTMask(uint8(square))

		if index >= 0 {
			bitSquare := ToSquare(uint8(index), uint8(SIZE)-1-uint8(index))
			expected.SetBit(bitSquare, 0)
		}

		result, _ := mask.ShiftToMoveCalPositionForLRBT(square)
		reversedResult := result.ReverseShiftToMoveCalPositionForLRBT(square)

		if !result.Uint96.Equals(*expected.Uint96) {
			// fmt.Printf("%s\n%s\n%s\n%d", mask.Rep(), result.Rep(), expected.Rep(), square)
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x", *result.Uint96, *expected.Uint96)
		}

		if !reversedResult.Uint96.Equals(*mask.Uint96) {
			// fmt.Printf("board:\n%s\nresult:\n%s\nexpected:\n%s\nreverse:\n%s\n%d", mask.Rep(), result.Rep(), expected.Rep(), reversedResult.Rep(), square)
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *reversedResult.Uint96, *mask.Uint96)
		}

		index = index + 1
	}
}

func TestShiftToMoveCalPositionForHor(t *testing.T) {
	expected := *NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		111111111
	`)

	for row := 0; row < int(SIZE); row = row + 1 {
		square := ToSquare(uint8(row), uint8(0))
		mask := GenerateHorizontalMask(uint8(square))

		result, _ := mask.ShiftToMoveCalPositionForHor(square)
		reversedResult := result.ReverseShiftToMoveCalPositionForHor(square)

		if !result.Uint96.Equals(*expected.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *result.Uint96, *expected.Uint96)
		}

		if !reversedResult.Uint96.Equals(*mask.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *reversedResult.Uint96, *mask.Uint96)
		}
	}
}

func TestShiftToMoveCalPositionForVer(t *testing.T) {
	expected := *NewBitboardFromStr(`
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

	for col := 0; col < int(SIZE); col = col + 1 {
		square := ToSquare(uint8(SIZE-1), uint8(col))
		mask := GenerateVerticalMask(uint8(square))

		result, _ := mask.ShiftToMoveCalPositionForVer(square)
		reversedResult := result.ReverseShiftToMoveCalPositionForVer(square)

		if !result.Uint96.Equals(*expected.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *result.Uint96, *expected.Uint96)
		}

		if !reversedResult.Uint96.Equals(*mask.Uint96) {
			// fmt.Printf("board:\n%s\nresult:\n%s\nexpected:\n%s\nreverse:\n%s\n%d", mask.Rep(), result.Rep(), expected.Rep(), reversedResult.Rep(), square)
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *reversedResult.Uint96, *mask.Uint96)
		}
	}
}
