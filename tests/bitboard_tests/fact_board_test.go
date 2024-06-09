package bitboard_tests

import (
	"testing"

	. "github.com/pqviet030188/advance-chess-ai/bitboard"
)

func TestGenerateHorizontalMask(t *testing.T) {
	exp := NewBitboardFromStr(`
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

	for i := range uint8(1) {
		square := i

		copy := exp.Copy()

		board := GenerateHorizontalMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		111111111
		000000000
		000000000
		000000000
		000000000
	`)

	for i := range SIZE {
		square := SIZE*4 + i

		copy := exp.Copy()

		board := GenerateHorizontalMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		111111111
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	for i := range SIZE {
		square := SIZE*8 + i

		copy := exp.Copy()

		board := GenerateHorizontalMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}
}

func TestGenerateVerticalMask(t *testing.T) {
	exp := NewBitboardFromStr(`
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

	for i := range uint8(1) {
		square := i * SIZE

		copy := exp.Copy()

		board := GenerateVerticalMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		000010000
		000010000
		000010000
		000010000
		000010000
		000010000
		000010000
		000010000
		000010000
	`)

	for i := range SIZE {
		square := i*SIZE + 4

		copy := exp.Copy()

		board := GenerateVerticalMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		100000000
		100000000
		100000000
		100000000
		100000000
		100000000
		100000000
		100000000
		100000000
	`)

	for i := range SIZE {
		square := i*SIZE + (SIZE - 1)

		copy := exp.Copy()

		board := GenerateVerticalMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}
}

func TestGenerateLRTBMask(t *testing.T) {
	exp := NewBitboardFromStr(`
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

	for i := range uint8(1) {
		square := i*SIZE + i

		copy := exp.Copy()

		board := GenerateLRTBMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		100000000
		010000000
		001000000
		000100000
		000010000
	`)

	for i := range uint8(5) {
		square := i*SIZE + i + 4

		copy := exp.Copy()

		board := GenerateLRTBMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		000010000
		000001000
		000000100
		000000010
		000000001
		000000000
		000000000
		000000000
		000000000
	`)

	for i := range uint8(5) {
		square := 4*SIZE + i*SIZE + i

		copy := exp.Copy()

		board := GenerateLRTBMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		100000000
	`)

	for range uint8(1) {
		square := SIZE - 1

		copy := exp.Copy()

		board := GenerateLRTBMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		000000001
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	for range uint8(1) {
		square := SIZE * (SIZE - 1)

		copy := exp.Copy()

		board := GenerateLRTBMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}
}

func TestGenerateLRBTMask(t *testing.T) {
	exp := NewBitboardFromStr(`
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

	for i := range uint8(SIZE) {
		square := i*SIZE + SIZE - 1 - i

		copy := exp.Copy()

		board := GenerateLRBTMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000001
		000000010
		000000100
		000001000
		000010000
	`)

	for i := range uint8(5) {
		square := i*SIZE + 4 - i

		copy := exp.Copy()

		board := GenerateLRBTMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		000010000
		000100000
		001000000
		010000000
		100000000
		000000000
		000000000
		000000000
		000000000
	`)

	for i := range uint8(5) {
		square := 4*SIZE + SIZE*i + SIZE - 1 - i

		copy := exp.Copy()

		board := GenerateLRBTMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000001
	`)

	for range uint8(1) {
		square := uint8(0)

		copy := exp.Copy()

		board := GenerateLRBTMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}

	exp = NewBitboardFromStr(`
		100000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	for range uint8(1) {
		square := SIZE*(SIZE-1) + SIZE - 1

		copy := exp.Copy()

		board := GenerateLRBTMask(square)

		if !copy.Uint96.Equals(*board.Uint96) {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
		}
	}
}
