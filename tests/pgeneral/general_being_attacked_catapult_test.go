package pgeneral_test

import (
	"testing"

	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
)

func SetupWithCatapult(wall string, nearGeneral string, farGeneral string, farCatapult string) *gamemodel.GameModel {
	side := bitboard.NEAR

	model := gamemodel.SetupWithFacts(side)
	model.NearGeneral = bitboard.NewBitboardFromStr(nearGeneral)
	model.FarGeneral = bitboard.NewBitboardFromStr(farGeneral)
	model.FarCatapult = bitboard.NewBitboardFromStr(farCatapult)
	model.Wall = bitboard.NewBitboardFromStr(wall)

	gamemodel.Update(model, side)
	return model
}

func TestGeneralBeingAttackedByCatapult1(t *testing.T) {
	model := SetupWithCatapult(`
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
		001000000
		000000000
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByCatapult2(t *testing.T) {
	model := SetupWithCatapult(`
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
		000000100
		000000000
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByCatapult3(t *testing.T) {
	model := SetupWithCatapult(`
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
		000010000
		000000000
		000000000
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByCatapult4(t *testing.T) {
	model := SetupWithCatapult(`
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
		000000100
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
		000010000
		000000000
	`)

	gamemodel.Update(model, bitboard.NEAR)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}
