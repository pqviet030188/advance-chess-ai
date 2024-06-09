package main

import (
	// "fmt"

	"fmt"

	"github.com/pqviet030188/advance-chess-ai/bitboard"
	// "github.com/pqviet030188/advance-chess-ai/dragon"
)

func main() {
	// v1 := uint96.FromUInt64(0xff00ff00ffffffff)
	// v2 := uint96.FromUInt64(0xfff0ff00ffffffff)
	// zero := uint96.FromUInt32(1)
	// fmt.Printf("%32b\n", v1.Lo)
	// fmt.Printf("%32b\n", v1.Mid)
	// fmt.Printf("%32b\n", v1.Hi)
	// fmt.Printf("%t\n", v1.Equals(v2))
	// fmt.Printf("%t\n", zero.IsZero())

	// big := v2.Big()
	// fmt.Printf("0x%96b\n", big)
	// by := byte(0xff)
	// by++

	// square := uint(bitboard.E5)
	// sb := bitboard.I9
	// nbitboard := bitboard.RandBitboard()
	// fmt.Printf("%s\n", nbitboard.Str())
	// fmt.Printf("%s\n", nbitboard.Rep())

	// nbitboard.SetBit(sb, 1)
	// fmt.Printf("%s\n", nbitboard.Rep())

	// mask := dragon.AttackMask(square)
	// fmt.Printf("%s\n", mask.Rep())

	bitboard.GenerateSlidingMoveDictionaryFile("./artifacts/horizontalsm", "./artifacts/verticalsm",
		"./artifacts/lrtbsm", "./artifacts/lrbtsm")

	bitboard.GenerateFactMaskFile("./artifacts/factmask")

	mask := bitboard.NewFactBoardDictionaryFromFile("./artifacts/factmask")
	for i := range bitboard.SIZE * bitboard.SIZE {
		// board, _ := mask.GetBoard(i, bitboard.DIRECTIONAL_MASK)
		// board, _ := mask.GetBoard(i, bitboard.ZOMBIE_NEAR_MOVE_MASK)
		// fmt.Printf("ZOMBIE NEAR\nsquare: %d\nboard:\n%s\n", i, board.Rep())

		// board, _ = mask.GetBoard(i, bitboard.NOT_ZOMBIE_NEAR_MOVE_MASK)
		// fmt.Printf("NOT ZOMBIE NEAR\nsquare: %d\nboard:\n%s\n", i, board.Rep())

		// board, _ = mask.GetBoard(i, bitboard.ZOMBIE_NEAR_ATTACK_MASK)
		// fmt.Printf("ZOMBIE NEAR ATTACK\nsquare: %d\nboard:\n%s\n", i, board.Rep())

		// board, _ = mask.GetBoard(i, bitboard.NOT_ZOMBIE_NEAR_ATTACK_MASK)
		// fmt.Printf("NOT ZOMBIE NEAR ATTACK\nsquare: %d\nboard:\n%s\n", i, board.Rep())

		// board, _ = mask.GetBoard(i, bitboard.ZOMBIE_FAR_MOVE_MASK)
		// fmt.Printf("ZOMBIE FAR\nsquare: %d\nboard:\n%s\n", i, board.Rep())

		// board, _ = mask.GetBoard(i, bitboard.NOT_ZOMBIE_FAR_MOVE_MASK)
		// fmt.Printf("NOT ZOMBIE FAR\nsquare: %d\nboard:\n%s\n", i, board.Rep())

		// board, _ = mask.GetBoard(i, bitboard.ZOMBIE_FAR_ATTACK_MASK)
		// fmt.Printf("ZOMBIE FAR ATTACK\nsquare: %d\nboard:\n%s\n", i, board.Rep())

		// board, _ = mask.GetBoard(i, bitboard.NOT_ZOMBIE_FAR_ATTACK_MASK)
		// fmt.Printf("NOT ZOMBIE FAR ATTACK\nsquare: %d\nboard:\n%s\n", i, board.Rep())

		board, _ := mask.GetBoard(i, bitboard.SENTINEL_MASK)
		fmt.Printf("SENTINEL\nsquare: %d\nboard:\n%s\n", i, board.Rep())

		board, _ = mask.GetBoard(i, bitboard.NOT_SENTINEL_MASK)
		fmt.Printf("NOT SENTINEL\nsquare: %d\nboard:\n%s\n", i, board.Rep())

		board, _ = mask.GetBoard(i, bitboard.CATAPULT_MASK)
		fmt.Printf("CATAPULT\nsquare: %d\nboard:\n%s\n", i, board.Rep())

		board, _ = mask.GetBoard(i, bitboard.NOT_CATAPULT_MASK)
		fmt.Printf("NOT CATAPULT\nsquare: %d\nboard:\n%s\n------------------\n", i, board.Rep())

		var mode string
		fmt.Scanf("%s", &mode)
	}
}
