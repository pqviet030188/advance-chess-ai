package bitboard

import (
	"github.com/pqviet030188/advance-chess-ai/uint96"
)

func SentinelProtection(
	sentinel *Bitboard,
	factMask *FactBoardDictionary,
) *Bitboard {

	// a col
	aVerticalMask, _ := factMask.Get(A1, VERTICAL_MASK)
	anotVerticalMask, _ := factMask.Get(A1, NOT_VERTICAL_MASK)

	// i col
	iVerticalMask, _ := factMask.Get(I1, VERTICAL_MASK)
	inotVerticalMask, _ := factMask.Get(I1, NOT_VERTICAL_MASK)

	// row 1
	row1Mask, _ := factMask.Get(A1, HORIZONTAL_MASK)
	row1notMask, _ := factMask.Get(A1, NOT_HORIZONTAL_MASK)

	// row 9
	row9Mask, _ := factMask.Get(A9, HORIZONTAL_MASK)
	row9notMask, _ := factMask.Get(A9, NOT_HORIZONTAL_MASK)

	i1Spot := iVerticalMask.And(*row1Mask)
	i9Spot := iVerticalMask.And(*row9Mask)
	a1Spot := aVerticalMask.And(*row1Mask)
	a9Spot := aVerticalMask.And(*row9Mask)

	iLine := iVerticalMask.Xor(i1Spot).Xor(i9Spot)
	aLine := aVerticalMask.Xor(a1Spot).Xor(a9Spot)
	row1Line := row1Mask.Xor(a1Spot).Xor(i1Spot)
	row9Line := row9Mask.Xor(a9Spot).Xor(i9Spot)

	zero := uint96.FromUInt32(0)
	center := sentinel.And(*anotVerticalMask).And(*inotVerticalMask).
		And(*row1notMask).And(*row9notMask)

	res := uint96.FromUInt32(0)

	if !center.Equals(zero) {
		// << 9, >> 9, << 1, >> 1
		top := center.Lsh(uint(SIZE))
		left := center.Lsh(1)
		bottom := center.Rsh(uint(SIZE))
		right := center.Rsh(1)
		res = top.Or(left).Or(right).Or(bottom).Or(res)
	}

	// has a1
	a1Sentinel := a1Spot.And(*sentinel.Uint96)
	if !a1Sentinel.Equals(zero) {
		// << 9, >> 1
		top := a1Sentinel.Lsh(uint(SIZE))
		right := a1Sentinel.Rsh(1)
		res = top.Or(right).Or(res)
	}

	// has i1
	i1Sentinel := i1Spot.And(*sentinel.Uint96)
	if !i1Sentinel.Equals(zero) {
		// << 9, << 1
		top := i1Sentinel.Lsh(uint(SIZE))
		left := i1Sentinel.Lsh(1)
		res = top.Or(left).Or(res)
	}

	// has a9
	a9Sentinel := a9Spot.And(*sentinel.Uint96)
	if !a9Sentinel.Equals(zero) {
		// >> 9, >> 1
		bottom := a9Sentinel.Rsh(uint(SIZE))
		right := a9Sentinel.Rsh(1)
		res = bottom.Or(right).Or(res)

		// b := Bitboard{
		// 	Uint96: &res,
		// }
		// fmt.Printf("%s\n", b.Rep())
	}

	i9Sentinel := i9Spot.And(*sentinel.Uint96)
	if !i9Sentinel.Equals(zero) {
		// >> 9, << 1
		bottom := i9Sentinel.Rsh(uint(SIZE))
		left := i9Sentinel.Lsh(1)
		res = bottom.Or(left).Or(res)
	}

	// has iline
	iSentinel := iLine.And(*sentinel.Uint96)
	if !iSentinel.Equals(zero) {
		// >> 9, << 9, << 1
		top := iSentinel.Lsh(uint(SIZE))
		left := iSentinel.Lsh(1)
		bottom := iSentinel.Rsh(uint(SIZE))
		res = top.Or(left).Or(bottom).Or(res)
	}

	aSentinel := aLine.And(*sentinel.Uint96)
	if !aSentinel.Equals(zero) {
		// >> 9, << 9, >> 1
		top := aSentinel.Lsh(uint(SIZE))
		bottom := aSentinel.Rsh(uint(SIZE))
		right := aSentinel.Rsh(1)
		res = top.Or(right).Or(bottom).Or(res)
	}

	// has row1line
	row1Sentinel := row1Line.And(*sentinel.Uint96)
	if !row1Sentinel.Equals(zero) {
		// << 9, << 1, >> 1
		top := row1Sentinel.Lsh(uint(SIZE))
		left := row1Sentinel.Lsh(1)
		right := row1Sentinel.Rsh(1)
		res = top.Or(left).Or(right).Or(res)
	}

	row9Sentinel := row9Line.And(*sentinel.Uint96)
	if !row9Sentinel.Equals(zero) {
		// >> 9, << 1, >> 1
		left := row9Sentinel.Lsh(1)
		bottom := row9Sentinel.Rsh(uint(SIZE))
		right := row9Sentinel.Rsh(1)
		res = bottom.Or(left).Or(right).Or(res)
	}

	return &Bitboard{
		Uint96: &res,
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
