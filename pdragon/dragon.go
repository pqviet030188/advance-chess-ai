package dragon

import (
	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
)

func GenerateAllMoves(
	square uint8,
	model *gamemodel.GameModel,
) *bitboard.Bitboard {
	return model.Everything.DirectionalMoveOnly(square, model.FactMask, model.LrtbDict, model.LrbtDict,
		model.HorizontalDict, model.VerticalDict)
}

func GenerateMoves(
	square uint8,
	side uint8,
	model *gamemodel.GameModel,
) (move *bitboard.Bitboard, attack *bitboard.Bitboard) {

	// get all moves
	allMoves := GenerateAllMoves(square, model)

	// get spaces
	spaces := model.Everything.Not()

	// moves only consider moving to empty spaces
	moveNumber := allMoves.And(spaces)
	move = &bitboard.Bitboard{
		Uint96: &moveNumber,
	}

	// attack remove nearby squares
	// and consider enemies only
	nearby := bitboard.Nearby(square, model.FactMask)
	attackNumber := allMoves.And(nearby.Not()).And(
		*model.GetEnemyPieces(side).Uint96)

	attack = &bitboard.Bitboard{
		Uint96: &attackNumber,
	}

	attack = model.RefineAttacksWithEnemySentinelProtection(side, attack)
	return move, attack
}
