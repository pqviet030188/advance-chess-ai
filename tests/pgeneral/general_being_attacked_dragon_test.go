package pgeneral_test

import (
	"testing"

	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
)

func SetupWithDragon(wall string, nearGeneral string, farGeneral string, farDragon string) *gamemodel.GameModel {
	side := bitboard.NEAR

	model := gamemodel.SetupWithFacts(side)
	model.NearGeneral = bitboard.NewBitboardFromStr(nearGeneral)
	model.FarGeneral = bitboard.NewBitboardFromStr(farGeneral)
	model.FarDragon = bitboard.NewBitboardFromStr(farDragon)
	model.Wall = bitboard.NewBitboardFromStr(wall)

	gamemodel.Update(model, side)
	return model
}

func TestGeneralBeingAttackedByDragon1(t *testing.T) {
	model := SetupWithDragon(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`, `
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
		000000000
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByDragon2(t *testing.T) {
	model := SetupWithDragon(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`, `
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		001000000
		000000000
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByDragon3(t *testing.T) {
	model := SetupWithDragon(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`, `
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000100
		000000000
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByDragonWithWallBlockers(t *testing.T) {
	model := SetupWithDragon(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000111000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`, `
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
		000000000
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}

func TestGeneralBeingAttackedByDragonWithAllyBlockers(t *testing.T) {
	model := SetupWithDragon(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`, `
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
		000000000
		000000000
	`)

	model.NearBuilder = bitboard.NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000111000
		000000000
	`)

	gamemodel.Update(model, bitboard.NEAR)
	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}

func TestGeneralBeingAttackedByDragonWithProtection(t *testing.T) {
	model := SetupWithDragon(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`, `
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
		000000000
		000000000
	`)

	model.NearSentinel = bitboard.NewBitboardFromStr(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000001000
	`)

	gamemodel.Update(model, bitboard.NEAR)
	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}

func TestGeneralBeingAttackedByDragonInNearby(t *testing.T) {
	model := SetupWithDragon(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
	`, `
		000010000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
	`, `
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

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}
