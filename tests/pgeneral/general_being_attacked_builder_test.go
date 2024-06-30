package pgeneral_test

import (
	"testing"

	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
)

func SetupWithBuilder(wall string, nearGeneral string, farGeneral string, farBuilder string) *gamemodel.GameModel {
	side := bitboard.NEAR

	model := gamemodel.SetupWithFacts(side)
	model.NearGeneral = bitboard.NewBitboardFromStr(nearGeneral)
	model.FarGeneral = bitboard.NewBitboardFromStr(farGeneral)
	model.FarBuilder = bitboard.NewBitboardFromStr(farBuilder)
	model.Wall = bitboard.NewBitboardFromStr(wall)

	gamemodel.Update(model, side)
	return model
}

func TestGeneralBeingAttackedByBuilder1(t *testing.T) {
	model := SetupWithBuilder(`
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

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByBuilder2(t *testing.T) {
	model := SetupWithBuilder(`
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
		000000000
		000100000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByBuilder3(t *testing.T) {
	model := SetupWithBuilder(`
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
		000000000
		000001000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByBuilder4(t *testing.T) {
	model := SetupWithBuilder(`
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
		000000000
		000000100
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}

func TestGeneralBeingAttackedByBuilderSentinelProtected(t *testing.T) {
	model := SetupWithBuilder(`
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

func TestGeneralBeingAttackedByBuilder5(t *testing.T) {
	model := SetupWithBuilder(`
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

func TestGeneralBeingAttackedByBuilder6(t *testing.T) {
	model := SetupWithBuilder(`
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
