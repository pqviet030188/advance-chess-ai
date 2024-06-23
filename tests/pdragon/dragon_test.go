package pdragon_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
	dragon "github.com/pqviet030188/advance-chess-ai/pdragon"
)

func setup(
	nearPieces string,
	farPieces string,

	wall string,
	nearSetinel string,
	farSentinel string,
) *gamemodel.GameModel {
	horizontalLookup := NewBoardDictionaryFromFile("../../artifacts/horizontalsm")
	verticalLookup := NewBoardDictionaryFromFile("../../artifacts/verticalsm")
	lrtbLookup := NewBoardDictionaryFromFile("../../artifacts/lrtbsm")
	lrbtLookup := NewBoardDictionaryFromFile("../../artifacts/lrbtsm")
	factMask := NewFactBoardDictionaryFromFile("../../artifacts/factmask")

	nearPiecesBoard := NewBitboardFromStr(nearPieces)
	farPiecesBoard := NewBitboardFromStr(farPieces)
	wallBoard := NewBitboardFromStr(wall)

	nearSentinelBoard := NewBitboardFromStr(nearSetinel)
	farSentinelBoard := NewBitboardFromStr(farSentinel)

	model := &gamemodel.GameModel{
		NearPieces:     nearPiecesBoard,
		FarPieces:      farPiecesBoard,
		Wall:           wallBoard,
		FactMask:       factMask,
		LrtbDict:       lrtbLookup,
		LrbtDict:       lrbtLookup,
		HorizontalDict: horizontalLookup,
		VerticalDict:   verticalLookup,
		NearSentinel:   nearSentinelBoard,
		FarSentinel:    farSentinelBoard,
	}

	model.GetEverything(false, true)

	return model
}

func TestDragonMoveAttack(t *testing.T) {
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

	model := setup(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := dragon.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		001000000
		101010000
		011100000
		010111000
		011100000
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
	// fmt.Printf("attack:\n%s\n", attack.Rep())
}

func TestDragonMoveAttackWithNearbyAttack(t *testing.T) {
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

	model := setup(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := dragon.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		100000000
		010000000
		010111000
		011100000
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
	// fmt.Printf("attack:\n%s\n", attack.Rep())
}

func TestDragonMoveAttackWithEnemySentinelProtection(t *testing.T) {
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
		000100000
		000000000
		000000000
		000000000
	`

	model := setup(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := dragon.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		100000000
		011000000
		010111000
		011100000
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
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestDragonMoveAttackWithFalseEnemySentinelProtection(t *testing.T) {
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
		000100000
		000000000
		000000000
		000000000
		000000000
	`

	model := setup(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := dragon.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		100000000
		011000000
		010111000
		011100000
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
	// fmt.Printf("attack:\n%s\n", attack.Rep())
	// fmt.Printf("attackExpected:\n%s\n", attackExpected.Rep())
}

func TestDragonMoveAttackWithCrossAttackEnemy(t *testing.T) {
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

	model := setup(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := dragon.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		100000000
		011100000
		010111000
		011100000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		001010000
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

func TestDragonMoveAttackWithCrossAttackSentinelProtection(t *testing.T) {
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

	model := setup(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := dragon.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		100000000
		011100000
		010111000
		011100000
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

func TestDragonMoveAttackWithWalls(t *testing.T) {
	wall := `
		000000000
		001000000
		000000000
		000000100
		000000000
		000000000
		000100000
		100001010
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
		000101000
		000000000
		000000000
		000000000
		000000000
	`

	model := setup(nearPieces, farPieces, wall, nearSentinel, farSentinel)

	start := time.Now()
	move, attack := dragon.GenerateMoves(C2, NEAR, model)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	moveExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		100000000
		011000000
		010110000
		011100000
	`)

	attackExpected := NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		001000000
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
