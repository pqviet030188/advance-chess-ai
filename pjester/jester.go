package pjester

import (
	"fmt"

	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/gamemodel"
)

func GenerateMoves(
	square uint8,
	side uint8,
	model *gamemodel.GameModel,
) (move *bitboard.Bitboard, attack *bitboard.Bitboard, swap *bitboard.Bitboard) {
	// get all moves and attacks
	nearbyMoves := bitboard.Nearby(square, model.FactMask)
	if nearbyMoves == nil {
		panic(fmt.Sprintf("Cannot look up nearby moves for jester, square: %d", square))
	}

	// get spaces
	spaces := model.Everything.Not()

	// moves only consider moving to empty spaces
	moveNumber := nearbyMoves.And(spaces)
	move = &bitboard.Bitboard{
		Uint96: &moveNumber,
	}

	// attack remove nearby squares
	// and consider enemies excluding general
	enemyExcludeGeneral := model.GetEnemyPieces(side).Xor(
		*model.GetEnemyGeneral(side).Uint96)
	attackNumber := nearbyMoves.And(enemyExcludeGeneral)
	attack = &bitboard.Bitboard{
		Uint96: &attackNumber,
	}

	// swaps only consider moving to ally spaces
	swapNumber := nearbyMoves.And(*model.GetPieces(side).Uint96)
	swap = &bitboard.Bitboard{
		Uint96: &swapNumber,
	}

	return move, attack, swap
}
