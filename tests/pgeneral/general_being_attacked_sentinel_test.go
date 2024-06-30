package pgeneral_test

import (
	"testing"

	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
)

func SetupWithSentinel(wall string, nearGeneral string, farGeneral string, farSentinel string) *gamemodel.GameModel {
	side := bitboard.NEAR

	model := gamemodel.SetupWithFacts(side)
	model.NearGeneral = bitboard.NewBitboardFromStr(nearGeneral)
	model.FarGeneral = bitboard.NewBitboardFromStr(farGeneral)
	model.FarSentinel = bitboard.NewBitboardFromStr(farSentinel)
	model.Wall = bitboard.NewBitboardFromStr(wall)

	gamemodel.Update(model, side)
	return model
}

func TestGeneralBeingAttackedBySentinel1(t *testing.T) {
	model := SetupWithSentinel(`
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
		000100000
		000000000
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedBySentinel2(t *testing.T) {
	model := SetupWithSentinel(`
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
		000001000
		000000000
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}
