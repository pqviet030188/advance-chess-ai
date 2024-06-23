package gamemodel

import (
	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/uint96"
)

type GameModel struct {
	Everything *bitboard.Bitboard
	NearPieces *bitboard.Bitboard
	FarPieces  *bitboard.Bitboard

	Wall         *bitboard.Bitboard
	NearSentinel *bitboard.Bitboard
	FarSentinel  *bitboard.Bitboard

	FactMask       *bitboard.FactBoardDictionary
	LrtbDict       *bitboard.BoardDictionary
	LrbtDict       *bitboard.BoardDictionary
	HorizontalDict *bitboard.BoardDictionary
	VerticalDict   *bitboard.BoardDictionary
}

func (model *GameModel) GetNearPieces(update bool) *bitboard.Bitboard {
	zero := uint96.FromUInt32(0)
	nearPiecesNumber := model.NearSentinel.Or(zero)
	nearPieces := &bitboard.Bitboard{
		Uint96: &nearPiecesNumber,
	}
	if update {
		model.NearPieces = nearPieces
	}

	return nearPieces
}

func (model *GameModel) GetFarPieces(update bool) *bitboard.Bitboard {
	zero := uint96.FromUInt32(0)
	farPiecesNumber := model.FarSentinel.Or(zero)
	farPieces := &bitboard.Bitboard{
		Uint96: &farPiecesNumber,
	}
	if update {
		model.FarPieces = farPieces
	}

	return farPieces
}

func (model *GameModel) GetEverything(updateSides bool, update bool) *bitboard.Bitboard {
	nearPieces := model.NearPieces
	farPieces := model.FarPieces

	if updateSides {
		nearPieces = model.GetNearPieces(update)
		farPieces = model.GetNearPieces(update)
	}

	everythingNumber := nearPieces.Or(*farPieces.Uint96).Or(*model.Wall.Uint96)
	everything := &bitboard.Bitboard{
		Uint96: &everythingNumber,
	}

	if update {
		model.Everything = everything
	}

	return everything
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

func (model *GameModel) GetSentinel(side uint8) *bitboard.Bitboard {
	if side == NEAR {
		return model.NearSentinel
	}

	if side == FAR {
		return model.FarSentinel
	}

	panic("Side is not valid")
}

func (model *GameModel) GetEnemySentinel(side uint8) *bitboard.Bitboard {
	if side == NEAR {
		return model.FarSentinel
	}

	if side == FAR {
		return model.NearSentinel
	}

	panic("Side is not valid")
}

func (model *GameModel) GetEnemyProtection(side uint8) *bitboard.Bitboard {
	enemySentinel := model.GetEnemySentinel(side)
	return bitboard.SentinelProtection(enemySentinel, model.FactMask)
}

func (model *GameModel) GeProtection(side uint8) *bitboard.Bitboard {
	sentinel := model.GetSentinel(side)
	return bitboard.SentinelProtection(sentinel, model.FactMask)
}

func (model *GameModel) RefineAttacksWithEnemySentinelProtection(side uint8, attack *bitboard.Bitboard) *bitboard.Bitboard {
	enemyProtection := model.GetEnemyProtection(side)

	// remove enemy protected blocks
	newAttackNumber := attack.And(enemyProtection.Not())
	return &bitboard.Bitboard{
		Uint96: &newAttackNumber,
	}
}
