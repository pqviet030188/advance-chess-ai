package pgeneral

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

	// get all moves
	allMoves := bitboard.Nearby(square, model.FactMask)
	if allMoves == nil {
		panic(fmt.Sprintf("Cannot look up nearby moves for nearby square: %d", square))
	}

	// get spaces
	spaces := model.Everything.Not()

	// moves only consider moving to empty spaces
	moveNumber := allMoves.And(spaces)
	move = &bitboard.Bitboard{
		Uint96: &moveNumber,
	}

	// attack considers enemies only
	attackNumber := allMoves.And(
		*model.GetEnemyPieces(side).Uint96)
	attack = &bitboard.Bitboard{
		Uint96: &attackNumber,
	}

	attack = model.RefineAttacksWithEnemySentinelProtection(side, attack)
	return move, attack
}
