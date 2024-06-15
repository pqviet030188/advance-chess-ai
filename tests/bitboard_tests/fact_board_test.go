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

func TestGenerateNearbyMask(t *testing.T) {
	exp := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000111000
		000101000
		000111000
		000000000
		000000000
		000000000
	`)

	copy := exp.Copy()
	board := GenerateNearbyMask(E5)
	if !copy.Uint96.Equals(*board.Uint96) {
		// fmt.Printf("%s %s \n", copy.Rep(), board.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000011
		000000010
	`)

	copy = exp.Copy()
	board = GenerateNearbyMask(I1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		010000000
		110000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateNearbyMask(A9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000010
		000000011
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateNearbyMask(I9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		110000000
		010000000
	`)

	copy = exp.Copy()
	board = GenerateNearbyMask(A1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}
}

func TestGenerateZombieNearMask(t *testing.T) {
	exp := NewBitboardFromStr(`
		000000000
		000000000
		001010100
		000111000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy := exp.Copy()
	board := GenerateZombieNearMask(E5)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000101
		000000011
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieNearMask(I1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
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
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieNearMask(A9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
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
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieNearMask(I9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		101000000
		110000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieNearMask(A1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}
}

func TestGenerateZombieNearMoveMask(t *testing.T) {
	exp := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000111000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy := exp.Copy()
	board := GenerateZombieNearMoveMask(E5)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000011
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieNearMoveMask(I1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
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
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieNearMoveMask(A9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
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
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieNearMoveMask(I9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		110000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieNearMoveMask(A1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}
}

func TestGenerateZombieFarMask(t *testing.T) {
	exp := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000111000
		001010100
		000000000
		000000000
	`)

	copy := exp.Copy()
	board := GenerateZombieFarMask(E5)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
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
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieFarMask(I1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		110000000
		101000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieFarMask(A9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000011
		000000101
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieFarMask(I9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
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
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieFarMask(A1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}
}

func TestGenerateZombieFarMoveMask(t *testing.T) {
	exp := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000111000
		000000000
		000000000
		000000000
	`)

	copy := exp.Copy()
	board := GenerateZombieFarMoveMask(E5)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
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
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieFarMoveMask(I1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		110000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieFarMoveMask(A9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000011
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieFarMoveMask(I9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
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
		000000000
	`)

	copy = exp.Copy()
	board = GenerateZombieFarMoveMask(A1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}
}

func TestGenerateSentinelMask(t *testing.T) {
	exp := NewBitboardFromStr(`
		000000000
		000000000
		000101000
		001000100
		000000000
		001000100
		000101000
		000000000
		000000000
	`)

	copy := exp.Copy()
	board := GenerateSentinelMask(E5)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000010
		000000100
		000000000
	`)

	copy = exp.Copy()
	board = GenerateSentinelMask(I1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		001000000
		010000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateSentinelMask(A9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000100
		000000010
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateSentinelMask(I9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		010000000
		001000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateSentinelMask(A1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}
}

func TestGenerateCatapultMask(t *testing.T) {
	exp := NewBitboardFromStr(`
		000000000
		000010000
		001000100
		000000000
		010000010
		000000000
		001000100
		000010000
		000000000
	`)

	copy := exp.Copy()
	board := GenerateCatapultMask(E5)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000001
		000000100
		000000000
		000001000
	`)

	copy = exp.Copy()
	board = GenerateCatapultMask(I1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000100000
		000000000
		001000000
		100000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateCatapultMask(A9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000001000
		000000000
		000000100
		000000001
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateCatapultMask(I9)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		100000000
		001000000
		000000000
		000100000
	`)

	copy = exp.Copy()
	board = GenerateCatapultMask(A1)
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}
}

func TestGenerateUintMask(t *testing.T) {
	exp := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000011111
		111111111
		111111111
		111111111
	`)

	copy := exp.Copy()
	board := GenerateLowUintMask()
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		000000000
		000000001
		111111111
		111111111
		111111111
		111100000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()
	board = GenerateMidUintMask()
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}

	exp = NewBitboardFromStr(`
		111111111
		111111110
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	copy = exp.Copy()

	DIM := SIZE * SIZE
	for i := DIM; i < 96; i++ {
		copy.SetBit(i, 1)
	}

	board = GenerateHiUintMask()
	if !copy.Uint96.Equals(*board.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", board.Uint96, copy.Uint96)
	}
}
