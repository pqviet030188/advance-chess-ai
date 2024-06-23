package gamemodel

import (
	"github.com/pqviet030188/advance-chess-ai/bitboard"
)

type GameModel struct {
	Everything *bitboard.Bitboard
	NearPieces *bitboard.Bitboard
	FarPieces  *bitboard.Bitboard

	Wall *bitboard.Bitboard

	NearZombie *bitboard.Bitboard
	FarZombie  *bitboard.Bitboard

	NearBuilder *bitboard.Bitboard
	FarBuilder  *bitboard.Bitboard

	NearDragon *bitboard.Bitboard
	FarDragon  *bitboard.Bitboard

	NearMiner *bitboard.Bitboard
	FarMiner  *bitboard.Bitboard

	NearSentinel *bitboard.Bitboard
	FarSentinel  *bitboard.Bitboard

	NearJester *bitboard.Bitboard
	FarJester  *bitboard.Bitboard

	NearCatapult *bitboard.Bitboard
	FarCatapult  *bitboard.Bitboard

	NearGeneral *bitboard.Bitboard
	FarGeneral  *bitboard.Bitboard

	NearSentinelProtection *bitboard.Bitboard
	FarSentinelProtection  *bitboard.Bitboard

	FactMask       *bitboard.FactBoardDictionary
	LrtbDict       *bitboard.BoardDictionary
	LrbtDict       *bitboard.BoardDictionary
	HorizontalDict *bitboard.BoardDictionary
	VerticalDict   *bitboard.BoardDictionary
}

func (model *GameModel) GetNearPieces(update bool) *bitboard.Bitboard {
	// zero := uint96.FromUInt32(0)
	nearPiecesNumber := model.NearSentinel.Or(*model.NearDragon.Uint96).Or(
		*model.NearMiner.Uint96).Or(*model.NearBuilder.Uint96).Or(
		*model.NearZombie.Uint96).Or(*model.NearGeneral.Uint96).Or(
		*model.NearJester.Uint96).Or(*model.NearCatapult.Uint96)

	nearPieces := &bitboard.Bitboard{
		Uint96: &nearPiecesNumber,
	}
	if update {
		model.NearPieces = nearPieces
	}

	return nearPieces
}

func (model *GameModel) GetFarPieces(update bool) *bitboard.Bitboard {
	// zero := uint96.FromUInt32(0)
	farPiecesNumber := model.FarSentinel.Or(*model.FarDragon.Uint96).Or(
		*model.FarMiner.Uint96).Or(*model.FarBuilder.Uint96).Or(
		*model.FarZombie.Uint96).Or(*model.FarGeneral.Uint96).Or(
		*model.FarJester.Uint96).Or(*model.FarCatapult.Uint96)

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

func (model *GameModel) GetEnemySide(side uint8) uint8 {
	if side == bitboard.NEAR {
		return bitboard.FAR
	}

	if side == bitboard.FAR {
		return bitboard.NEAR
	}

	panic("Side is not valid")
}

func (model *GameModel) GetPieces(side uint8) *bitboard.Bitboard {
	if side == bitboard.NEAR {
		return model.NearPieces
	}

	if side == bitboard.FAR {
		return model.FarPieces
	}

	panic("Side is not valid")
}

func (model *GameModel) GetEnemyPieces(side uint8) *bitboard.Bitboard {
	if side == bitboard.NEAR {
		return model.FarPieces
	}

	if side == bitboard.FAR {
		return model.NearPieces
	}

	panic("Side is not valid")
}

func (model *GameModel) GetSentinel(side uint8) *bitboard.Bitboard {
	if side == bitboard.NEAR {
		return model.NearSentinel
	}

	if side == bitboard.FAR {
		return model.FarSentinel
	}

	panic("Side is not valid")
}

func (model *GameModel) GetEnemySentinel(side uint8) *bitboard.Bitboard {
	if side == bitboard.NEAR {
		return model.FarSentinel
	}

	if side == bitboard.FAR {
		return model.NearSentinel
	}

	panic("Side is not valid")
}

func (model *GameModel) GetGeneral(side uint8) *bitboard.Bitboard {
	if side == bitboard.NEAR {
		return model.NearGeneral
	}

	if side == bitboard.FAR {
		return model.FarGeneral
	}

	panic("Side is not valid")
}

func (model *GameModel) GetEnemyGeneral(side uint8) *bitboard.Bitboard {
	if side == bitboard.NEAR {
		return model.FarGeneral
	}

	if side == bitboard.FAR {
		return model.NearGeneral
	}

	panic("Side is not valid")
}

func (model *GameModel) GetProtection(side uint8) *bitboard.Bitboard {
	if side == bitboard.NEAR {
		return model.NearSentinelProtection
	}

	if side == bitboard.FAR {
		return model.FarSentinelProtection
	}

	panic("Side is not valid")
}

func (model *GameModel) GetEnemyProtection(side uint8) *bitboard.Bitboard {
	if side == bitboard.NEAR {
		return model.FarSentinelProtection
	}

	if side == bitboard.FAR {
		return model.NearSentinelProtection
	}

	panic("Side is not valid")
}

func (model *GameModel) CalculateSentinelProtection(side uint8, update bool) *bitboard.Bitboard {
	sentinel := model.GetSentinel(side)
	protection := bitboard.SentinelProtection(sentinel, model.FactMask)
	if update {
		if side == bitboard.NEAR {
			model.NearSentinelProtection = protection
		}

		if side == bitboard.FAR {
			model.FarSentinelProtection = protection
		}
	}

	return protection
}

func (model *GameModel) CalculateEnemyProtection(side uint8, update bool) *bitboard.Bitboard {
	enemySide := model.GetEnemySide(side)
	return model.CalculateSentinelProtection(enemySide, update)
}

func (model *GameModel) RefineAttacksWithEnemySentinelProtection(side uint8, attack *bitboard.Bitboard) *bitboard.Bitboard {
	enemyProtection := model.GetEnemyProtection(side)

	// remove enemy protected blocks
	newAttackNumber := attack.And(enemyProtection.Not())
	return &bitboard.Bitboard{
		Uint96: &newAttackNumber,
	}
}
