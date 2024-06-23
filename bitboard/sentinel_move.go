package bitboard

import "fmt"

func SentinelMoves(
	square uint8,
	factMask *FactBoardDictionary,
) *Bitboard {
	movesNumber, ok := factMask.Get(square, SENTINEL_MASK)
	if movesNumber == nil || !ok {
		panic(fmt.Sprintf("Cannot get zombie near attack mask square: %d", square))
	}

	moves := &Bitboard{
		Uint96: movesNumber,
	}

	return moves
}
