package pminer_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
	miner "github.com/pqviet030188/advance-chess-ai/pminer"
)

func TestMinerMoveAttack(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		000000000
		100000010
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
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000001000
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
		000101000
		000000000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack, destroy := miner.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		001000000
		001000000
		001000000
		010111000
		001000000
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

	destroyExpected := NewBitboardFromStr(`
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

	if !destroy.Uint96.Equals(*destroyExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *destroy.Uint96, *destroyExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestMinerMoveAttackWithNearbyAttack(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		000000000
		100000010
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
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000001000
		000000000
		001100000
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

	start := time.Now()
	move, attack, destroy := miner.GenerateMoves(C2, NEAR, model)
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
		010111000
		001000000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		001000000
		000000100
		000000000
	`)

	destroyExpected := NewBitboardFromStr(`
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

	if !destroy.Uint96.Equals(*destroyExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *destroy.Uint96, *destroyExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
}

func TestMinerMoveAttackWithEnemySentinelProtection(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		000000000
		100000010
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
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000001000
		001100000
		100100000
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
		000100000
		100000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack, destroy := miner.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		001000000
		010111000
		001000000
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

	destroyExpected := NewBitboardFromStr(`
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

	if !destroy.Uint96.Equals(*destroyExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *destroy.Uint96, *destroyExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestMinerMoveAttackWithFalseEnemySentinelProtection(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		000000000
		100000010
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
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000101000
		001100000
		010100000
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
		000100000
		000000000
		010000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack, destroy := miner.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		001000000
		010111000
		001000000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		001000000
		000000000
		000000100
		000000000
	`)

	destroyExpected := NewBitboardFromStr(`
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

	if !destroy.Uint96.Equals(*destroyExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *destroy.Uint96, *destroyExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestMinerMoveAttackWithCrossAttackEnemy(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		000000000
		100000010
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
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000101000
		001110000
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
		000101000
		000100000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack, destroy := miner.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		001000000
		010111000
		001000000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		001000000
		000000000
		000000100
		000000000
	`)

	destroyExpected := NewBitboardFromStr(`
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

	if !destroy.Uint96.Equals(*destroyExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *destroy.Uint96, *destroyExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestMinerMoveAttackWithCrossAttackSentinelProtection(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		000000000
		100000010
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
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000111000
		001110000
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
		000101000
		000110000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack, destroy := miner.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		001000000
		010111000
		001000000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		001000000
		000000000
		000000100
		000000000
	`)

	destroyExpected := NewBitboardFromStr(`
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

	if !destroy.Uint96.Equals(*destroyExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *destroy.Uint96, *destroyExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestMinerMoveAttackWithWalls(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		001000000
		100000010
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
		000000000
	`

	farPieces := `
		000000000
		000000000
		000000000
		001101000
		000111000
		001110000
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
		000101000
		000110000
		000000000
		000000000
		000000000
		000000000
	`

	model := gamemodel.SetupWithStrings(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack, destroy := miner.GenerateMoves(C2, NEAR, model)
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
		010111000
		001000000
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

	destroyExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		001000000
		100000000
		000000000
	`)

	if !move.Uint96.Equals(*moveExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *move.Uint96, *moveExpected.Uint96)
	}

	if !attack.Uint96.Equals(*attackExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *attack.Uint96, *attackExpected.Uint96)
	}

	if !destroy.Uint96.Equals(*destroyExpected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", *destroy.Uint96, *destroyExpected.Uint96)
	}

	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("near sentinel:\n%s\n", model.NearSentinel.Rep())
	// fmt.Printf("far sentinel:\n%s\n", model.FarSentinel.Rep())
	// fmt.Printf("move:\n%s\n", move.Rep())
	// fmt.Printf("moveExpected:\n%s\n", moveExpected.Rep())
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}
