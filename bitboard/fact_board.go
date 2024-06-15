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

func GenerateNearbyMask(square uint8) *Bitboard {

	ret := uint96.FromUInt32(0)
	row := int(square / 9)
	col := int(square % 9)

	r, c := row+1, col+1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+1, col
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+1, col-1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row, col+1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	// r, c = row, col
	// if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
	// 	ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	// }

	r, c = row, col-1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-1, col+1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-1, col
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-1, col-1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func GenerateZombieNearMask(square uint8) *Bitboard {

	ret := uint96.FromUInt32(0)
	row := int(square / 9)
	col := int(square % 9)

	// r, c := row, col
	// if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
	// 	ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	// }

	r, c := row+1, col+1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+2, col+2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+1, col
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+2, col
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+1, col-1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+2, col-2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func GenerateZombieNearMoveMask(square uint8) *Bitboard {

	ret := uint96.FromUInt32(0)
	row := int(square / 9)
	col := int(square % 9)

	// r, c := row, col
	// if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
	// 	ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	// }

	r, c := row+1, col+1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+1, col
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+1, col-1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func GenerateZombieFarMask(square uint8) *Bitboard {

	ret := uint96.FromUInt32(0)
	row := int(square / 9)
	col := int(square % 9)

	// r, c := row, col
	// if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
	// 	ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	// }

	r, c := row-1, col
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-2, col
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-1, col+1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-2, col+2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-1, col-1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-2, col-2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func GenerateZombieFarMoveMask(square uint8) *Bitboard {

	ret := uint96.FromUInt32(0)
	row := int(square / 9)
	col := int(square % 9)

	// r, c := row, col
	// if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
	// 	ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	// }

	r, c := row-1, col
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-1, col+1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-1, col-1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func GenerateSentinelMask(square uint8) *Bitboard {

	ret := uint96.FromUInt32(0)
	row := int(square / 9)
	col := int(square % 9)

	// r, c := row, col
	// if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
	// 	ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	// }

	r, c := row+2, col+1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+2, col-1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+1, col+2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+1, col-2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-2, col+1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-2, col-1
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-1, col+2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-1, col-2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func GenerateCatapultMask(square uint8) *Bitboard {

	ret := uint96.FromUInt32(0)
	row := int(square / 9)
	col := int(square % 9)

	// r, c := row, col
	// if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
	// 	ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	// }

	r, c := row+3, col
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+2, col+2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row+2, col-2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row, col+3
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row, col-3
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-2, col+2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-2, col-2
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	r, c = row-3, col
	if r >= 0 && r < int(SIZE) && c >= 0 && c < int(SIZE) {
		ret = ret.Or(uint96.FromUInt32(1).Lsh(uint(r*9 + c)))
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func GenerateFactMaskDictionary() *FactBoardDictionary {
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

	for i := range SIZE * SIZE {
		square := uint8(i)
		mask := GenerateNearbyMask(square).Uint96
		notMask := mask.Not()
		fact.Put(square, NEARBY_MASK, mask)
		fact.Put(square, NOT_NEARBY_MASK, &notMask)
	}

	for i := range SIZE * SIZE {
		square := uint8(i)
		zombieFarMask := GenerateZombieFarMask(square).Uint96
		notZombieFarMask := zombieFarMask.Not()

		zombieNearMask := GenerateZombieNearMask(square).Uint96
		notZombieNearMask := zombieNearMask.Not()

		zombieFarMoveMask := GenerateZombieFarMoveMask(square).Uint96
		notZombieFarMoveMask := zombieFarMoveMask.Not()

		zombieNearMoveMask := GenerateZombieNearMoveMask(square).Uint96
		notZombieNearMoveMask := zombieNearMoveMask.Not()

		fact.Put(square, ZOMBIE_FAR_ATTACK_MASK, zombieFarMask)
		fact.Put(square, NOT_ZOMBIE_FAR_ATTACK_MASK, &notZombieFarMask)

		fact.Put(square, ZOMBIE_NEAR_ATTACK_MASK, zombieNearMask)
		fact.Put(square, NOT_ZOMBIE_NEAR_ATTACK_MASK, &notZombieNearMask)

		fact.Put(square, ZOMBIE_FAR_MOVE_MASK, zombieFarMoveMask)
		fact.Put(square, NOT_ZOMBIE_FAR_MOVE_MASK, &notZombieFarMoveMask)

		fact.Put(square, ZOMBIE_NEAR_MOVE_MASK, zombieNearMoveMask)
		fact.Put(square, NOT_ZOMBIE_NEAR_MOVE_MASK, &notZombieNearMoveMask)
	}

	for i := range SIZE * SIZE {
		square := uint8(i)
		mask := GenerateSentinelMask(square).Uint96
		notMask := mask.Not()
		fact.Put(square, SENTINEL_MASK, mask)
		fact.Put(square, NOT_SENTINEL_MASK, &notMask)
	}

	for i := range SIZE * SIZE {
		square := uint8(i)
		mask := GenerateCatapultMask(square).Uint96
		notMask := mask.Not()
		fact.Put(square, CATAPULT_MASK, mask)
		fact.Put(square, NOT_CATAPULT_MASK, &notMask)
	}

	return fact
}

func GenerateFactMaskDictionaryFile(path string) {
	mask := GenerateFactMaskDictionary()
	mask.ToFile(path)
}
