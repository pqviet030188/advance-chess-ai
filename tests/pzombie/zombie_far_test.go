package pzombie_test

import (
	"testing"

	. "github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
	zombie "github.com/pqviet030188/advance-chess-ai/pzombie"
)

func TestZombieFarMoveAttack(t *testing.T) {
	wall := `
		000000000
		000000000
		000100000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearPieces := `
		000000000
		000000000
		001000000
		100000000
		000001000
		000000000
		000000000
		000000000
		000000000
	`

	farPieces := `
		000000000
		011110000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	farSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	move, attack := zombie.GenerateMoves(C8, FAR, model)
	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		010000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		001000000
		100000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}
}

func TestZombieFarMoveAttack2(t *testing.T) {
	wall := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearPieces := `
		000000000
		000000000
		000000000
		000100001
		010000000
		000000000
		000000000
		000000000
		000000000
	`

	farPieces := `
		000000000
		010000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	farSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	move, attack := zombie.GenerateMoves(B8, FAR, model)
	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		111000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000100000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}
}

func TestZombieFarMoveAttack3(t *testing.T) {
	wall := `
		000000000
		000000000
		100000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearPieces := `
		000000000
		000000000
		000000001
		001000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	farPieces := `
		000000000
		100000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	farSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	move, attack := zombie.GenerateMoves(A8, FAR, model)
	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		010000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		001000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}
}

func TestZombieFarMoveAttack4(t *testing.T) {
	wall := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearPieces := `
		000000000
		000000000
		000000000
		000001010
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	farPieces := `
		000000000
		000000010
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000010
		000000000
		000000000
		000000000
		000000000
	`

	farSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	move, attack := zombie.GenerateMoves(H8, FAR, model)
	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000111
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000001000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}
}

func TestZombieFarMoveAttack5(t *testing.T) {
	wall := `
		000000000
		000000000
		000000001
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearPieces := `
		000000000
		000000000
		000000000
		000000100
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	farPieces := `
		000000000
		000000001
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	farSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	move, attack := zombie.GenerateMoves(I8, FAR, model)
	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000010
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000100
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}
}

func TestZombieFarMoveAttack6(t *testing.T) {
	wall := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearPieces := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`

	nearSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	farSentinel := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	move, attack := zombie.GenerateMoves(E1, FAR, model)
	moveExpected := NewBitboardFromStr(`
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

	attackExpected := NewBitboardFromStr(`
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

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}
}
