package bitboard

import "github.com/pqviet030188/advance-chess-ai/uint96"

func SentinelProtection(
	sentinelSquares []uint8,
	factMask *FactBoardDictionary,
) *Bitboard {

	result := uint96.FromUInt32(0)

	for _, square := range sentinelSquares {
		number, ok := factMask.Get(square, NEARBY_MASK)
		if !ok {
			panic("Cannot find nearby mask")
		}

		result = result.Or(*number)
	}

	return &Bitboard{
		Uint96: &result,
	}
}

func Nearby(
	square uint8,
	factMask *FactBoardDictionary,
) *Bitboard {

	number, ok := factMask.Get(square, NEARBY_MASK)
	if !ok {
		panic("Cannot find nearby mask")
	}

	return &Bitboard{
		Uint96: number,
	}
}
