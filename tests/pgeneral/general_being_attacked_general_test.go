package pgeneral_test

import (
	"testing"

	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
)

func SetupWithGeneral(wall string, nearGeneral string, farGeneral string) *gamemodel.GameModel {
	side := bitboard.NEAR

	model := gamemodel.SetupWithFacts(side)
	model.NearGeneral = bitboard.NewBitboardFromStr(nearGeneral)
	model.FarGeneral = bitboard.NewBitboardFromStr(farGeneral)
	model.Wall = bitboard.NewBitboardFromStr(wall)

	gamemodel.Update(model, side)
	return model
}

func TestGeneralBeingAttackedByGeneral1(t *testing.T) {
	model := SetupWithGeneral(`
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

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByGeneral2(t *testing.T) {
	model := SetupWithGeneral(`
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
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000100000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByGeneral4(t *testing.T) {
	model := SetupWithGeneral(`
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
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000100
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}

func TestGeneralBeingAttackedByGeneralSentinelProtected(t *testing.T) {
	model := SetupWithGeneral(`
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
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000100000
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

func TestGeneralBeingAttackedByGeneral5(t *testing.T) {
	model := SetupWithGeneral(`
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
		000100000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
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
	`)

	model.NearBuilder = bitboard.NewBitboardFromStr(`
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

	gamemodel.Update(model, bitboard.FAR)
	attacked := model.IsGeneralBeingAttacked(bitboard.E9, bitboard.FAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByGeneral6(t *testing.T) {
	model := SetupWithGeneral(`
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
		000100000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
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
	`)

	model.FarSentinel = bitboard.NewBitboardFromStr(`
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

	gamemodel.Update(model, bitboard.FAR)
	attacked := model.IsGeneralBeingAttacked(bitboard.E9, bitboard.FAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}
