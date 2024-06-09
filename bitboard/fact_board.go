package bitboard

import (
	"github.com/pqviet030188/advance-chess-ai/uint96"
)

func GenerateHorizontalMask(square uint8) *Bitboard {

	ret := uint96.FromUInt32(0)
	row := int(square / 9)
	col := int(square % 9)

	ret.SetBit(square, 1)
	for r, c := row, col+1; c < int(SIZE); c++ {

		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	for r, c := row, col-1; c >= 0; c-- {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func GenerateVerticalMask(square uint8) *Bitboard {

	ret := uint96.FromUInt32(0)
	row := int(square / 9)
	col := int(square % 9)

	ret.SetBit(square, 1)
	for r, c := row+1, col; r < int(SIZE); r++ {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	for r, c := row-1, col; r >= 0; r-- {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func GenerateLRTBMask(square uint8) *Bitboard {

	ret := uint96.FromUInt32(0)
	row := int(square / 9)
	col := int(square % 9)

	ret.SetBit(square, 1)
	for r, c := row+1, col+1; r < int(SIZE) && c < int(SIZE); r, c = r+1, c+1 {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func GenerateLRBTMask(square uint8) *Bitboard {

	ret := uint96.FromUInt32(0)
	row := int(square / 9)
	col := int(square % 9)

	ret.SetBit(square, 1)
	for r, c := row-1, col+1; r >= 0 && c < int(SIZE); r, c = r-1, c+1 {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	for r, c := row+1, col-1; r < int(SIZE) && c >= 0; r, c = r+1, c-1 {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func GenerateFactMask() *FactBoardDictionary {
	fact := NewFactBoardDictionary()

	for i := range SIZE * SIZE {
		square := uint8(i)
		mask := GenerateHorizontalMask(square).Uint96
		notMask := mask.Not()
		fact.Put(square, HORIZONTAL_MASK, mask)
		fact.Put(square, NOT_HORIZONTAL_MASK, &notMask)
	}

	for i := range SIZE * SIZE {
		square := uint8(i)
		mask := GenerateVerticalMask(square).Uint96
		notMask := mask.Not()
		fact.Put(square, VERTICAL_MASK, mask)
		fact.Put(square, NOT_VERTICAL_MASK, &notMask)
	}

	for i := range SIZE * SIZE {
		square := uint8(i)
		mask := GenerateLRTBMask(square).Uint96
		notMask := mask.Not()
		fact.Put(square, LRTB_MASK, mask)
		fact.Put(square, NOT_LRTB_MASK, &notMask)
	}

	for i := range SIZE * SIZE {
		square := uint8(i)
		mask := GenerateLRBTMask(square).Uint96
		notMask := mask.Not()
		fact.Put(square, LRBT_MASK, mask)
		fact.Put(square, NOT_LRBT_MASK, &notMask)
	}

	for i := range SIZE * SIZE {
		square := uint8(i)
		lrbt := GenerateLRBTMask(square).Uint96
		lrtb := GenerateLRTBMask(square).Uint96
		horizontal := GenerateHorizontalMask(square).Uint96
		vertical := GenerateVerticalMask(square).Uint96
		mask := lrbt.Or(*lrtb).Or(*horizontal).Or(*vertical)
		notMask := mask.Not()

		fact.Put(square, DIRECTIONAL_MASK, &mask)
		fact.Put(square, NOT_DIRECTIONAL_MASK, &notMask)
	}

	for i := range SIZE * SIZE {
		square := uint8(i)
		horizontal := GenerateHorizontalMask(square).Uint96
		vertical := GenerateVerticalMask(square).Uint96
		mask := horizontal.Or(*vertical)
		notMask := mask.Not()

		fact.Put(square, MINER_MASK, &mask)
		fact.Put(square, NOT_MINER_MASK, &notMask)
	}

	return fact
}

func GenerateFactMaskFile(path string) {
	mask := GenerateFactMask()
	mask.ToFile(path)
}
