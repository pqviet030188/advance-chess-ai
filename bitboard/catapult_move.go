package bitboard

import "fmt"

func CatapultMoves(
	square uint8,
	factMask *FactBoardDictionary,
) *Bitboard {
	moves := Nearby(square, factMask)

	if moves == nil {
		panic(fmt.Sprintf("Cannot get catapult moves, square: %d", square))
	}

	return moves
}

func CatapultAttacks(
	square uint8,
	factMask *FactBoardDictionary,
) *Bitboard {
	attacksNumber, ok := factMask.Get(square, CATAPULT_MASK)
	if attacksNumber == nil || !ok {
		panic(fmt.Sprintf("Cannot get zombie near attack mask square: %d", square))
	}

	attacks := &Bitboard{
		Uint96: attacksNumber,
	}

	return attacks
}
