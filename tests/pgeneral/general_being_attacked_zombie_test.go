package pgeneral_test

import (
	"testing"

	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
)

func SetupWithZombie(wall string, nearGeneral string, farGeneral string, farZombie string) *gamemodel.GameModel {
	side := bitboard.NEAR
	model := gamemodel.SetupWithFacts(side)
	model.NearGeneral = bitboard.NewBitboardFromStr(nearGeneral)
	model.FarGeneral = bitboard.NewBitboardFromStr(farGeneral)
	model.FarZombie = bitboard.NewBitboardFromStr(farZombie)
	model.Wall = bitboard.NewBitboardFromStr(wall)

	gamemodel.Update(model, side)

	return model
}

func TestGeneralBeingAttackedByZombie1(t *testing.T) {
	model := SetupWithZombie(`
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

func TestGeneralBeingAttackedByZombie2(t *testing.T) {
	model := SetupWithZombie(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000010000
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

func TestGeneralBeingAttackedByZombie3(t *testing.T) {
	model := SetupWithZombie(`
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
		000010000
		000000000
	`)

	gamemodel.Update(model, bitboard.NEAR)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}

}

func TestGeneralBeingAttackedByZombie4(t *testing.T) {
	model := SetupWithZombie(`
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
		000010000
		000000000
		000000000
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}

func TestGeneralBeingAttackedByZombie5(t *testing.T) {
	model := SetupWithZombie(`
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

func TestGeneralBeingAttackedByZombie6(t *testing.T) {
	model := SetupWithZombie(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000100000
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
		001000100
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
		000001000
		000000000
	`)

	gamemodel.Update(model, bitboard.NEAR)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}

func TestGeneralBeingAttackedByZombie7(t *testing.T) {
	model := SetupWithZombie(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000100000
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
		001000100
		000000000
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.E1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByZombieSentinelProtected(t *testing.T) {
	model := SetupWithZombie(`
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000100000
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
		001000100
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

func TestGeneralBeingAttackedByZombie8(t *testing.T) {
	model := SetupWithZombie(`
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
		100000000
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
		100000000
		000000000
		000000001
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.A1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}

func TestGeneralBeingAttackedByZombie9(t *testing.T) {
	model := SetupWithZombie(`
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
		100000000
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
		100000000
		000000001
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.A1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByZombie10(t *testing.T) {
	model := SetupWithZombie(`
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
		100000000
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
		100000000
		000000000
		000000001
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.A1, bitboard.NEAR)

	if attacked {
		t.Errorf("Expected not to be threatened")
	}
}

func TestGeneralBeingAttackedByZombie11(t *testing.T) {
	model := SetupWithZombie(`
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
		100000000
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
		000000001
		000000000
	`)

	attacked := model.IsGeneralBeingAttacked(bitboard.A1, bitboard.NEAR)

	if !attacked {
		t.Errorf("Expected to be threatened")
	}
}

func TestGeneralBeingAttackedByZombie12(t *testing.T) {
	model := SetupWithZombie(`
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

	model.NearZombie = bitboard.NewBitboardFromStr(`
		000000000
		000000000
		001000000
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
