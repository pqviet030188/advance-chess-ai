package pzombie

import (
	"fmt"

	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
)

func GenerateMoves(
	square uint8,
	side uint8,
	model *gamemodel.GameModel,
) (move *bitboard.Bitboard, attack *bitboard.Bitboard) {
	// get all moves and attacks
	allMoves := bitboard.ZombieMoves(square, side, model.FactMask)
	allAttacks := bitboard.ZombieAttacks(square, side, model.FactMask)

	if allMoves == nil {
		panic(fmt.Sprintf("Cannot look up moves for zombie, square: %d", square))
	}

	if allAttacks == nil {
		panic(fmt.Sprintf("Cannot look up attack moves for zombie, square: %d", square))
	}

	// get spaces
	spaces := model.Everything.Not()

	// moves only consider moving to empty spaces
	moveNumber := allMoves.And(spaces)
	move = &bitboard.Bitboard{
		Uint96: &moveNumber,
	}

	// get all directional moves
	directionalMoves := model.Everything.DirectionalMoveOnly(square, model.FactMask,
		model.LrtbDict, model.LrbtDict,
		model.HorizontalDict, model.VerticalDict)

	// attack remove nearby squares
	// and consider enemies only
	attackNumber := directionalMoves.And(*allAttacks.Uint96).And(
		*model.GetEnemyPieces(side).Uint96)
	attack = &bitboard.Bitboard{
		Uint96: &attackNumber,
	}

	attack = model.RefineAttacksWithEnemySentinelProtection(side, attack)
	return move, attack
}
