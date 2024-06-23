package gamemodel

import "github.com/pqviet030188/advance-chess-ai/bitboard"

func SetupWithStrings(
	nearPieces string,
	farPieces string,

	wall string,
	nearSetinel string,
	farSentinel string,
) *GameModel {
	horizontalLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/horizontalsm")
	verticalLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/verticalsm")
	lrtbLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/lrtbsm")
	lrbtLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/lrbtsm")
	factMask := bitboard.NewFactBoardDictionaryFromFile("../../artifacts/factmask")

	nearPiecesBoard := bitboard.NewBitboardFromStr(nearPieces)
	farPiecesBoard := bitboard.NewBitboardFromStr(farPieces)
	wallBoard := bitboard.NewBitboardFromStr(wall)

	nearSentinelBoard := bitboard.NewBitboardFromStr(nearSetinel)
	farSentinelBoard := bitboard.NewBitboardFromStr(farSentinel)

	model := &GameModel{
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
	model.CalculateSentinelProtection(bitboard.NEAR, true)
	model.CalculateSentinelProtection(bitboard.FAR, true)
	return model
}
