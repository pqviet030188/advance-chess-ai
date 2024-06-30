package gamemodel

import (
	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/uint96"
)

func SetupWithFacts(side uint8) *GameModel {

	zero := uint96.FromUInt32(0)
	zeroBoard := &bitboard.Bitboard{
		Uint96: &zero,
	}

	horizontalLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/horizontalsm")
	verticalLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/verticalsm")
	lrtbLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/lrtbsm")
	lrbtLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/lrbtsm")
	factMask := bitboard.NewFactBoardDictionaryFromFile("../../artifacts/factmask")

	model := &GameModel{
		Wall:           zeroBoard,
		FactMask:       factMask,
		LrtbDict:       lrtbLookup,
		LrbtDict:       lrbtLookup,
		HorizontalDict: horizontalLookup,
		VerticalDict:   verticalLookup,
		NearBuilder:    zeroBoard,
		FarBuilder:     zeroBoard,
		NearMiner:      zeroBoard,
		FarMiner:       zeroBoard,
		NearSentinel:   zeroBoard,
		FarSentinel:    zeroBoard,
		NearCatapult:   zeroBoard,
		FarCatapult:    zeroBoard,
		NearDragon:     zeroBoard,
		FarDragon:      zeroBoard,
		NearJester:     zeroBoard,
		FarJester:      zeroBoard,
		NearGeneral:    zeroBoard,
		FarGeneral:     zeroBoard,
		NearZombie:     zeroBoard,
		FarZombie:      zeroBoard,
	}
	return model
}

func Update(model *GameModel, side uint8) {
	enemySide := GetEnemySide(side)

	model.CalculateSentinelProtection(side, true)
	model.CalculateSentinelProtection(enemySide, true)
	model.GetEverything(true, true)
}
