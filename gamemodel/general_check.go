package gamemodel

import (
	"github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/uint96"
)

func (model *GameModel) IsGeneralBeingAttacked(square uint8, side uint8) bool {
	general := model.GetGeneral(side)
	// if general.GetBit(square) == uint8(0) {
	// 	return false
	// }

	zero := uint96.FromUInt32(0)
	protection := model.GetProtection(side)

	// general being protected
	if !protection.And(*general.Uint96).Equals(zero) {
		return false
	}

	// --- check attacks of builder, general ---
	nearby := bitboard.Nearby(square, model.FactMask)
	enemyBuilder := model.GetEnemyBuilder(side)

	// nearby includes enemy builder
	if !nearby.And(*enemyBuilder.Uint96).Equals(zero) {
		return true
	}

	enemyGeneral := model.GetEnemyGeneral(side)
	// nearby includes enemy general
	if !nearby.And(*enemyGeneral.Uint96).Equals(zero) {
		return true
	}

	// --- check attacks of zombie ---
	horSlidingBoard := model.Everything.HorizontalMoveOnly(square, model.FactMask, model.HorizontalDict)
	verSlidingBoard := model.Everything.VerticalMoveOnly(square, model.FactMask, model.VerticalDict)
	horVerSlidingMoves := horSlidingBoard.Or(*verSlidingBoard.Uint96)
	lrbtSlidingBoard := model.Everything.LRBTMoveOnly(square, model.FactMask, model.LrbtDict)
	lrtbSlidingBoard := model.Everything.LRTBMoveOnly(square, model.FactMask, model.LrtbDict)
	directionalSlidingMoves := horVerSlidingMoves.Or(*lrbtSlidingBoard.Uint96).Or(*lrtbSlidingBoard.Uint96)

	// temp := &bitboard.Bitboard{
	// 	Uint96: &directionalSlidingMoves,
	// }
	// fmt.Printf("everything:\n%s\n", model.Everything.Rep())
	// fmt.Printf("move:\n%s\n", temp.Rep())

	// pretend that zombie is at current general spot, see if the zombie
	// can attack enemy zombie
	zombieMaskType := bitboard.ZOMBIE_FAR_ATTACK_MASK
	if side == bitboard.NEAR {
		zombieMaskType = bitboard.ZOMBIE_NEAR_ATTACK_MASK
	}

	enemyZombieAttackMask, ok := model.FactMask.Get(square, zombieMaskType)
	if !ok {
		panic("Should return zombie mask for general check")
	}

	enemyZombie := model.GetEnemyZombie(side)

	// enemy zombie attack mask includes enemy zombies (indicating that enemy zombies can attack us)
	if !enemyZombieAttackMask.And(directionalSlidingMoves).And(*enemyZombie.Uint96).Equals(zero) {
		return true
	}

	// --- check attacks of miner and dragon ---
	enemyMiner := model.GetEnemyMiner(side)

	// vertical and horizontal sliding moves include enemy miners
	if !horVerSlidingMoves.And(*enemyMiner.Uint96).Equals(zero) {
		return true
	}

	enemyDragon := model.GetEnemyDragon(side)
	notnearbyMask, ok := model.FactMask.Get(square, bitboard.NOT_NEARBY_MASK)

	if !ok {
		panic("Should be able to get not nearby mask in general check")
	}

	// directional sliding moves except nearby ranges includes enemy dragons
	if !directionalSlidingMoves.And(*notnearbyMask).And(*enemyDragon.Uint96).Equals(zero) {
		return true
	}

	// --- check attacks of sentinel ---
	enemySentinel := model.GetEnemySentinel(side)
	sentinelMask, ok := model.FactMask.Get(square, bitboard.SENTINEL_MASK)
	if !ok {
		panic("Should be able to get sentinel mask in general check")
	}

	if !sentinelMask.And(*enemySentinel.Uint96).Equals(zero) {
		return true
	}

	// directional sliding moves except nearby ranges includes enemy dragons
	if !directionalSlidingMoves.And(*notnearbyMask).And(*enemyDragon.Uint96).Equals(zero) {
		return true
	}

	// --- check attacks of catapult ---
	enemyCatapult := model.GetEnemyCatapult(side)
	catapultMask, ok := model.FactMask.Get(square, bitboard.CATAPULT_MASK)
	if !ok {
		panic("Should be able to get catapult mask in general check")
	}

	if !catapultMask.And(*enemyCatapult.Uint96).Equals(zero) {
		return true
	}

	return false
}
