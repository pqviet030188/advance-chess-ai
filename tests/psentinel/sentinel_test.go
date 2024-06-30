package psentinel_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
	sentinel "github.com/pqviet030188/advance-chess-ai/psentinel"
)

func TestSentinelMoveAttack(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		001111100
		000000000
		000000000
		000000000
		100000010
		000000000
	`

	nearPieces := `
		000000000
		001010000
		000100000
		000000000
		000000000
		000000000
		000000000
		001000010
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101100
		000011000
		000000000
		000000000
		000000100
		000000000
	`

	nearSentinel := `
		000000000
		000010000
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
		000100000
		000010000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := sentinel.GenerateMoves(E5, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000001000
		000000000
		000000000
		001000100
		000101000
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

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestSentinelMoveAttackLeft1(t *testing.T) {
	wall := `
		000000000
		110000000
		101000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearPieces := `
		100000000
		000000000
		000100000
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
		010000000
		001000000
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
		001000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := sentinel.GenerateMoves(A9, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		001000000
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
		010000000
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

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestSentinelMoveAttackLeft2(t *testing.T) {
	wall := `
		000000000
		000000000
		000000000
		110000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearPieces := `
		000000000
		000000000
		000100000
		001000000
		100000000
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
		001100000
		010000000
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
		000100000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := sentinel.GenerateMoves(A5, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

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
		000000000
		000000000
		000000000
		010000000
		000000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestSentinelMoveAttackLeft3(t *testing.T) {
	wall := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		110000000
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
		100000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		001100000
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
		000100000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := sentinel.GenerateMoves(A1, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		010000000
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

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestSentinelMoveAttackRight1(t *testing.T) {
	wall := `
		000000000
		000000011
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
		000000010
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	farPieces := `
		000000000
		000001100
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
		000001000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := sentinel.GenerateMoves(I9, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

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

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestSentinelMoveAttackRight2(t *testing.T) {
	wall := `
		000000000
		000000000
		000000010
		000000010
		000000010
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
		000000001
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
		000000100
		000000010
		000000010
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
		000000010
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := sentinel.GenerateMoves(I5, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
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

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000100
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

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestSentinelMoveAttackRight3(t *testing.T) {
	wall := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000010
		000000010
		000000010
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
		000000001
	`

	farPieces := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000100
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

	start := time.Now()
	move, attack := sentinel.GenerateMoves(I1, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

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
		000000100
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestSentinelMoveAttackTop(t *testing.T) {
	wall := `
		000000000
		000111000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	nearPieces := `
		000010000
		000000000
		000100000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	farPieces := `
		000000000
		001000100
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
		000000100
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := sentinel.GenerateMoves(E9, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000001000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		001000000
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

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestSentinelMoveAttackBottom(t *testing.T) {
	wall := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000111000
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
		001000000
		000010000
	`

	farPieces := `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000011000
		000000100
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
		000001000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := sentinel.GenerateMoves(E1, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000100000
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
		000000100
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}
