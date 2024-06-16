package dragon

import "github.com/pqviet030188/advance-chess-ai/bitboard"

func GenerateMoves(bitboard *bitboard.Bitboard, square uint8) (move *bitboard.Bitboard, attack *bitboard.Bitboard) {

}

// func AttackMask(square uint) bitboard.Bitboard {

// 	ret := uint96.FromUInt32(0)

// 	row := square / 9
// 	col := square % 9

// 	for r, c := row+1, col+1; r <= 7 && c <= 7; r, c = r+1, c+1 {
// 		ret = ret.Or(uint96.FromUInt32(1).Lsh(r*9 + c))
// 	}

// 	for r, c := row-1, col+1; r >= 1 && c <= 7; r, c = r-1, c+1 {
// 		ret = ret.Or(uint96.FromUInt32(1).Lsh(r*9 + c))
// 	}

// 	for r, c := row+1, col-1; r <= 7 && c >= 1; r, c = r+1, c-1 {
// 		ret = ret.Or(uint96.FromUInt32(1).Lsh(r*9 + c))
// 	}

// 	for r, c := row-1, col-1; r >= 1 && c >= 1; r, c = r-1, c-1 {
// 		ret = ret.Or(uint96.FromUInt32(1).Lsh(r*9 + c))
// 	}

// 	for r, c := row+1, col; r <= 7; r++ {
// 		ret = ret.Or(uint96.FromUInt32(1).Lsh(r*9 + c))
// 	}

// 	for r, c := row-1, col; r >= 1; r-- {
// 		ret = ret.Or(uint96.FromUInt32(1).Lsh(r*9 + c))
// 	}

// 	for r, c := row, col+1; c <= 7; c++ {
// 		ret = ret.Or(uint96.FromUInt32(1).Lsh(r*9 + c))
// 	}

// 	for r, c := row, col-1; c >= 1; c-- {
// 		ret = ret.Or(uint96.FromUInt32(1).Lsh(r*9 + c))
// 	}

// 	return bitboard.Bitboard{
// 		Uint96: &ret,
// 	}
// }

// calculate move
// do & near
