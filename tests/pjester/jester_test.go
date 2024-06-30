package pjester_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
	jester "github.com/pqviet030188/advance-chess-ai/pjester"
)

func TestJesterMoveAttack(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		000000000
		100100010
		000000000
	`

	nearPieces := `
		000000000
		001010000
		000000000
		000000000
		000000000
		000000000
		011000000
		001000010
		000010000
	`

	farPieces := `
		000010000
		000000000
		000000000
		001101000
		000001000
		000000000
		000100000
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
		000101000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	model.NearGeneral = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`)
	model.FarGeneral = NewBitboardFromStr(`
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	start := time.Now()
	move, attack, swap := jester.GenerateMoves(C2, NEAR, model)
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
		010000000
		011100000
	`)

	attackExpected := NewBitboardFromStr(`
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

	swapExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		011000000
		000000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	if !swap.Uint96.Equals(*swapExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *swap.Uint96, *swapExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
	// fmt.Printf("build:\n%s\n", build.Rep())
	// fmt.Printf("buildExpected:\n%s\n", buildExpected.Rep())
}

func TestJesterMoveAttackWithSentinelProtection(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		000000000
		100100010
		000000000
	`

	nearPieces := `
		000000000
		001010000
		000000000
		000000000
		000000000
		000000000
		011000000
		001000010
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000001000
		000000000
		000110000
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
		000101000
		000000000
		000000000
		000010000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	model.NearGeneral = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`)
	model.FarGeneral = NewBitboardFromStr(`
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	start := time.Now()
	move, attack, swap := jester.GenerateMoves(C2, NEAR, model)
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
		010000000
		011100000
	`)

	attackExpected := NewBitboardFromStr(`
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

	swapExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		011000000
		000000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	if !swap.Uint96.Equals(*swapExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *swap.Uint96, *swapExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
	// fmt.Printf("build:\n%s\n", build.Rep())
	// fmt.Printf("buildExpected:\n%s\n", buildExpected.Rep())
}

func TestJesterMoveAttackWithFalseSentinelProtection(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		000000000
		100100010
		000000000
	`

	nearPieces := `
		000010000
		001010000
		000000000
		000000000
		000000000
		000000000
		011000000
		001000010
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000001000
		000010000
		000110000
		000000100
		000110000
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
		000101000
		000000000
		000010000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	model.NearGeneral = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`)
	model.FarGeneral = NewBitboardFromStr(`
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000100000
	`)

	start := time.Now()
	move, attack, swap := jester.GenerateMoves(C2, NEAR, model)
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
		010000000
		011000000
	`)

	attackExpected := NewBitboardFromStr(`
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

	swapExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		011000000
		000000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	if !swap.Uint96.Equals(*swapExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *swap.Uint96, *swapExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
	// fmt.Printf("build:\n%s\n", build.Rep())
	// fmt.Printf("buildExpected:\n%s\n", buildExpected.Rep())
}

func TestJesterMoveAttackLeft1(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		000000000
		000100010
		000000000
	`

	nearPieces := `
		000000000
		001010000
		000000000
		000000000
		000000000
		000000000
		100000000
		101000010
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000001000
		001010000
		010110000
		000000100
		111000000
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
		000101000
		000000000
		001010000
		000000000
		000000000
		001000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	model.NearGeneral = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`)
	model.FarGeneral = NewBitboardFromStr(`
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		100000000
	`)

	start := time.Now()
	move, attack, swap := jester.GenerateMoves(A2, NEAR, model)
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
		010000000
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
		010000000
	`)

	swapExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		100000000
		000000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	if !swap.Uint96.Equals(*swapExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *swap.Uint96, *swapExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
	// fmt.Printf("swap:\n%s\n", swap.Rep())
	// fmt.Printf("swapExpected:\n%s\n", swapExpected.Rep())
}

func TestJesterMoveAttackLeft2(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		100000000
		000100010
		000000000
	`

	nearPieces := `
		000000000
		001010000
		000000000
		000000000
		000000000
		000000000
		000000000
		101000010
		100000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000001000
		001010000
		010110000
		000000100
		011000000
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
		000101000
		000000000
		001010000
		000000000
		000000000
		001000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	model.NearGeneral = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`)
	model.FarGeneral = NewBitboardFromStr(`
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		010000000
	`)

	start := time.Now()
	move, attack, swap := jester.GenerateMoves(A1, NEAR, model)
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
		010000000
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

	swapExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		100000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	if !swap.Uint96.Equals(*swapExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *swap.Uint96, *swapExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
	// fmt.Printf("build:\n%s\n", build.Rep())
	// fmt.Printf("buildExpected:\n%s\n", buildExpected.Rep())
}

func TestJesterMoveAttackLeft3(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		100000000
		100100010
		000000000
	`

	nearPieces := `
		110000000
		001010000
		000000000
		000000000
		000000000
		000000000
		000000000
		001000010
		100000000
	`

	farPieces := `
		000000000
		100000000
		000000000
		001101000
		000001000
		001010000
		010110000
		000000100
		011000000
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
		000101000
		000000000
		001010000
		000000000
		000000000
		001000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	model.NearGeneral = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`)
	model.FarGeneral = NewBitboardFromStr(`
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	start := time.Now()
	move, attack, swap := jester.GenerateMoves(A9, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		010000000
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
		100000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	swapExpected := NewBitboardFromStr(`
		010000000
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

	if !swap.Uint96.Equals(*swapExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *swap.Uint96, *swapExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
	// fmt.Printf("build:\n%s\n", build.Rep())
	// fmt.Printf("buildExpected:\n%s\n", buildExpected.Rep())
}

func TestJesterMoveAttackRight1(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		100000000
		000100010
		000000000
	`

	nearPieces := `
		000000000
		001010000
		000000000
		000000000
		000000000
		000000000
		000000000
		101000001
		000000010
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000001000
		001010000
		010110001
		000000100
		111000000
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
		000101000
		000000000
		001010000
		000000000
		000000000
		001000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	model.NearGeneral = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`)
	model.FarGeneral = NewBitboardFromStr(`
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	start := time.Now()
	move, attack, swap := jester.GenerateMoves(I2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000010
		000000000
		000000001
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000001
		000000000
		000000000
	`)

	swapExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000010
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	if !swap.Uint96.Equals(*swapExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *swap.Uint96, *swapExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
	// fmt.Printf("build:\n%s\n", build.Rep())
	// fmt.Printf("buildExpected:\n%s\n", buildExpected.Rep())
}

func TestJesterMoveAttackRight2(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		100000000
		100100010
		000000000
	`

	nearPieces := `
		000000000
		001010000
		000000000
		000000000
		000000000
		000000000
		000000000
		001000010
		100000001
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000001000
		001010000
		010110001
		000000101
		011000000
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
		000101000
		000000000
		001010000
		000000001
		000000000
		001000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	model.NearGeneral = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`)
	model.FarGeneral = NewBitboardFromStr(`
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000001
		000000000
	`)

	start := time.Now()
	move, attack, swap := jester.GenerateMoves(I1, NEAR, model)
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
		000000010
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

	swapExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000010
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	if !swap.Uint96.Equals(*swapExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *swap.Uint96, *swapExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
	// fmt.Printf("build:\n%s\n", build.Rep())
	// fmt.Printf("buildExpected:\n%s\n", buildExpected.Rep())
}

func TestJesterMoveAttackRight3(t *testing.T) {
	wall := `
		000000010
		001000000
		000000000
		000000100
		000000000
		000000000
		100000000
		100100010
		000000000
	`

	nearPieces := `
		110000011
		001010000
		000000000
		000000000
		000000000
		000000000
		000000000
		001000010
		100000000
	`

	farPieces := `
		000000000
		100000010
		000000000
		001101000
		000001000
		001010000
		010110000
		000000100
		011000000
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
		000101000
		000000000
		001010000
		000000000
		000000000
		001000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	model.NearGeneral = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`)
	model.FarGeneral = NewBitboardFromStr(`
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	start := time.Now()
	move, attack, swap := jester.GenerateMoves(I9, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000001
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
		000000010
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	swapExpected := NewBitboardFromStr(`
		000000010
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

	if !swap.Uint96.Equals(*swapExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *swap.Uint96, *swapExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
	// fmt.Printf("build:\n%s\n", build.Rep())
	// fmt.Printf("buildExpected:\n%s\n", buildExpected.Rep())
}

func TestJesterMoveAttackTop(t *testing.T) {
	wall := `
		000001010
		001000000
		000000000
		000000100
		000000000
		000000000
		100000000
		100100010
		000000000
	`

	nearPieces := `
		110010001
		001010000
		000000000
		000000000
		000000000
		000000000
		000000000
		001000010
		100000000
	`

	farPieces := `
		000000000
		100101010
		000001000
		001101000
		000001000
		001010000
		010110000
		000000100
		011000000
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
		000001000
		000101000
		000000000
		001010000
		000000000
		000000000
		001000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)
	model.NearGeneral = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`)
	model.FarGeneral = NewBitboardFromStr(`
		000000000
		000001000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	start := time.Now()
	move, attack, swap := jester.GenerateMoves(E9, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000100000
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
		000100000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	swapExpected := NewBitboardFromStr(`
		000000000
		000010000
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

	if !swap.Uint96.Equals(*swapExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *swap.Uint96, *swapExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
	// fmt.Printf("build:\n%s\n", build.Rep())
	// fmt.Printf("buildExpected:\n%s\n", buildExpected.Rep())
}

func TestJesterMoveAttackBottom(t *testing.T) {
	wall := `
		000001010
		001000000
		000000000
		000000100
		000000000
		000000000
		100000000
		100100010
		000000000
	`

	nearPieces := `
		110010001
		001010000
		000000000
		000000000
		000000000
		000000000
		000000000
		001100010
		100010000
	`

	farPieces := `
		000000000
		100101010
		000001000
		001101000
		000001000
		001010000
		010110000
		000011100
		011000000
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
		000001000
		000101000
		000000000
		001010000
		000001000
		000000000
		001000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	model.NearGeneral = NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`)
	model.FarGeneral = NewBitboardFromStr(`
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`)

	start := time.Now()
	move, attack, swap := jester.GenerateMoves(E1, NEAR, model)
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
		000101000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000011000
		000000000
	`)

	swapExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000100000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	if !swap.Uint96.Equals(*swapExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *swap.Uint96, *swapExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
	// fmt.Printf("build:\n%s\n", build.Rep())
	// fmt.Printf("buildExpected:\n%s\n", buildExpected.Rep())
}
