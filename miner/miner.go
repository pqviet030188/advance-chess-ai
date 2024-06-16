package miner

import (
	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
)

func GenerateMoves(
	square uint8,
	side uint8,
	model *gamemodel.GameModel,
) (move *bitboard.Bitboard, attack *bitboard.Bitboard) {

	// get all moves
	horizontalMoves := model.Everything.HorizontalMoveOnly(square, model.FactMask, model.HorizontalDict)
	verticalMoves := model.Everything.VerticalMoveOnly(square, model.FactMask, model.VerticalDict)
	allMoves := horizontalMoves.Or(*verticalMoves.Uint96)

	// get spaces
	spaces := model.Everything.Not()

	// moves only consider moving to empty spaces
	moveNumber := allMoves.And(spaces)
	move = &bitboard.Bitboard{
		Uint96: &moveNumber,
	}

	// attack remove nearby squares
	// and consider enemies only
	attackNumber := allMoves.And(
		*model.GetEnemyPieces(side).Uint96)

	attack = &bitboard.Bitboard{
		Uint96: &attackNumber,
	}
	return move, attack
}
