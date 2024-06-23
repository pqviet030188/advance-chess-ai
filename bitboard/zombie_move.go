package bitboard

import "fmt"

func ZombieAttacks(
	square uint8,
	side uint8,
	factMask *FactBoardDictionary,
) *Bitboard {
	var attacks *Bitboard = nil

	if side == NEAR {
		attacksNumber, ok := factMask.Get(square, ZOMBIE_NEAR_ATTACK_MASK)
		if attacksNumber == nil || !ok {
			panic(fmt.Sprintf("Cannot get zombie near attack mask square: %d", square))
		}

		attacks = &Bitboard{
			Uint96: attacksNumber,
		}
	} else if side == FAR {
		attacksNumber, ok := factMask.Get(square, ZOMBIE_FAR_ATTACK_MASK)
		if attacksNumber == nil || !ok {
			panic(fmt.Sprintf("Cannot get zombie near attack mask square: %d", square))
		}

		attacks = &Bitboard{
			Uint96: attacksNumber,
		}
	}

	return attacks
}

func ZombieMoves(
	square uint8,
	side uint8,
	factMask *FactBoardDictionary,
) *Bitboard {
	var moves *Bitboard = nil

	if side == NEAR {
		movesNumber, ok := factMask.Get(square, ZOMBIE_NEAR_MOVE_MASK)
		if movesNumber == nil || !ok {
			panic(fmt.Sprintf("Cannot get zombie near attack mask square: %d", square))
		}

		moves = &Bitboard{
			Uint96: movesNumber,
		}
	} else if side == FAR {
		movesNumber, ok := factMask.Get(square, ZOMBIE_FAR_MOVE_MASK)
		if movesNumber == nil || !ok {
			panic(fmt.Sprintf("Cannot get zombie near attack mask square: %d", square))
		}

		moves = &Bitboard{
			Uint96: movesNumber,
		}
	}

	return moves
}
