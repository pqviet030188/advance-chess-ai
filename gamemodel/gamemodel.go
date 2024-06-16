package gamemodel

import "github.com/pqviet030188/advance-chess-ai/bitboard"

type GameModel struct {
	Everything *bitboard.Bitboard
	NearPieces *bitboard.Bitboard
	FarPieces  *bitboard.Bitboard

	Wall *bitboard.Bitboard

	NearSentinelSquares []uint8
	FarSentinelSquares  []uint8

	FactMask       *bitboard.FactBoardDictionary
	LrtbDict       *bitboard.BoardDictionary
	LrbtDict       *bitboard.BoardDictionary
	HorizontalDict *bitboard.BoardDictionary
	VerticalDict   *bitboard.BoardDictionary
}

func (model *GameModel) GetEnemyPieces(side uint8) *bitboard.Bitboard {
	if side == NEAR {
		return model.FarPieces
	}

	if side == FAR {
		return model.NearPieces
	}

	panic("Side is not valid")
}

func (model *GameModel) GetEnemySentinelSquares(side uint8) []uint8 {
	if side == NEAR {
		return model.FarSentinelSquares
	}

	if side == FAR {
		return model.NearSentinelSquares
	}

	panic("Side is not valid")
}

func (model *GameModel) GetSentinelSquares(side uint8) []uint8 {
	if side == NEAR {
		return model.NearSentinelSquares
	}

	if side == FAR {
		return model.FarSentinelSquares
	}

	panic("Side is not valid")
}

func (model *GameModel) GetEnemyProtection(side uint8) *bitboard.Bitboard {
	enemySentinelSquares := model.GetEnemySentinelSquares(side)
	return bitboard.SentinelProtection(enemySentinelSquares, model.FactMask)
}

func (model *GameModel) GeProtection(side uint8) *bitboard.Bitboard {
	sentinelSquares := model.GetSentinelSquares(side)
	return bitboard.SentinelProtection(sentinelSquares, model.FactMask)
}

func (model *GameModel) RefineAttacksWithEnemySentinelProtection(side uint8, attack *bitboard.Bitboard) *bitboard.Bitboard {
	enemySentinelSquares := model.GetEnemySentinelSquares(side)
	enemyProtection := bitboard.SentinelProtection(enemySentinelSquares, model.FactMask)

	// remove enemy protected blocks
	newAttackNumber := attack.And(enemyProtection.Not())
	return &bitboard.Bitboard{
		Uint96: &newAttackNumber,
	}
}
