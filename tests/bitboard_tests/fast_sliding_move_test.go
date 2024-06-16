package bitboard_tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/pqviet030188/advance-chess-ai/bitboard"
)

func TestHorizontalMask(t *testing.T) {
	horizontalLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/horizontalsm")
	factMask := bitboard.NewFactBoardDictionaryFromFile("../../artifacts/factmask")
	board := bitboard.NewBitboardFromStr(`
		000000000
		000010100
		001000000
		000000100
		101001011
		001000000
		000000100
		000010000
		001000000
	`)
	expected := bitboard.NewBitboardFromStr(`
		000000000
		000010100
		001000000
		000000100
		001110110
		001000000
		000000100
		000010000
		001000000
	`)

	start := time.Now()
	result := board.HorizontalMove(bitboard.F5, factMask, horizontalLookup)
	duration := time.Since(start)

	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		000000000
		000010100
		001000000
		000000100
		101001011
		001000000
		000000100
		000010000
		101001000
	`)
	expected = bitboard.NewBitboardFromStr(`
		000000000
		000010100
		001000000
		000000100
		101001011
		001000000
		000000100
		000010000
		110111000
	`)

	start = time.Now()
	result = board.HorizontalMove(bitboard.C1, factMask, horizontalLookup)
	duration = time.Since(start)

	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		101001000
		000010100
		001000000
		000000100
		101001011
		001000000
		000000100
		000010000
		101001000
	`)
	expected = bitboard.NewBitboardFromStr(`
		110111000
		000010100
		001000000
		000000100
		101001011
		001000000
		000000100
		000010000
		101001000
	`)

	start = time.Now()
	result = board.HorizontalMove(bitboard.C9, factMask, horizontalLookup)
	duration = time.Since(start)

	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}
}

func TestVerticalMask(t *testing.T) {
	verticalLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/verticalsm")
	factMask := bitboard.NewFactBoardDictionaryFromFile("../../artifacts/factmask")
	board := bitboard.NewBitboardFromStr(`
		000000100
		000010000
		001000000
		000000000
		101001111
		001000000
		000000000
		000010000
		001000100
	`)
	expected := bitboard.NewBitboardFromStr(`
		000000100
		000010100
		001000100
		000000100
		101001011
		001000100
		000000100
		000010100
		001000100
	`)

	start := time.Now()
	result := board.VerticalMove(bitboard.G5, factMask, verticalLookup)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		// fmt.Printf("%s\n%s\n%s\n", board.Rep(), expected.Rep(), result.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		000000100
		000010000
		001000001
		000000000
		101001111
		001000000
		000000000
		000010000
		001000101
	`)
	expected = bitboard.NewBitboardFromStr(`
		000000100
		000010000
		001000001
		000000001
		101001110
		001000001
		000000001
		000010001
		001000101
	`)

	start = time.Now()
	result = board.VerticalMove(bitboard.I5, factMask, verticalLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		// fmt.Printf("%s\n%s\n%s\n", board.Rep(), expected.Rep(), result.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		100000100
		000010000
		001000001
		000000000
		101001111
		001000000
		000000000
		000010000
		101000101
	`)
	expected = bitboard.NewBitboardFromStr(`
		100000100
		100010000
		101000001
		100000000
		001001111
		101000000
		100000000
		100010000
		101000101
	`)

	start = time.Now()
	result = board.VerticalMove(bitboard.A5, factMask, verticalLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		// fmt.Printf("%s\n%s\n%s\n", board.Rep(), expected.Rep(), result.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}
}

func TestLRTBMask(t *testing.T) {
	lrtbLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/lrtbsm")
	factMask := bitboard.NewFactBoardDictionaryFromFile("../../artifacts/factmask")
	board := bitboard.NewBitboardFromStr(`
		001000100
		000010000
		001010000
		000000000
		101001111
		001000000
		000000001
		000010000
		001000100
	`)
	expected := bitboard.NewBitboardFromStr(`
		000000100
		000010000
		001010000
		000001000
		101001011
		001000010
		000000001
		000010000
		001000100
	`)

	start := time.Now()
	result := board.LRTBMove(bitboard.G5, factMask, lrtbLookup)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		// fmt.Printf("%s\n%s\n%s\n", board.Rep(), expected.Rep(), result.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		001000100
		000010000
		001010000
		100000000
		101001111
		001000000
		000000001
		000000000
		001001100
	`)
	expected = bitboard.NewBitboardFromStr(`
		001000100
		000010000
		001010000
		100000000
		111001111
		000000000
		000100001
		000010000
		001001100
	`)

	start = time.Now()
	result = board.LRTBMove(bitboard.C4, factMask, lrtbLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		101000100
		000010000
		001010000
		100000000
		101001111
		001000000
		000001101
		000000000
		001001101
	`)
	expected = bitboard.NewBitboardFromStr(`
		101000100
		010010000
		000010000
		100100000
		101011111
		001001000
		000001101
		000000000
		001001100
	`)

	start = time.Now()
	result = board.LRTBMove(bitboard.C7, factMask, lrtbLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		001000101
		000010000
		001010000
		100000000
		101001111
		001000000
		000000001
		000000000
		101001100
	`)
	expected1 := bitboard.NewBitboardFromStr(`
		001000100
		000010000
		001010000
		100000000
		101001111
		001000000
		000000001
		000000000
		101001100
	`)

	expected2 := bitboard.NewBitboardFromStr(`
		001000101
		000010000
		001010000
		100000000
		101001111
		001000000
		000000001
		000000000
		001001100
	`)

	start = time.Now()
	result = board.LRTBMove(bitboard.I9, factMask, lrtbLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected1.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected1.Uint96)
	}

	start = time.Now()
	result = board.LRTBMove(bitboard.A1, factMask, lrtbLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected2.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected2.Uint96)
	}
}

func TestLRBTMask(t *testing.T) {
	lrbtLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/lrbtsm")
	factMask := bitboard.NewFactBoardDictionaryFromFile("../../artifacts/factmask")
	board := bitboard.NewBitboardFromStr(`
		001000100
		000010000
		001010000
		000000000
		101001111
		001000100
		000000001
		000010000
		001000100
	`)
	expected := bitboard.NewBitboardFromStr(`
		001000100
		000010000
		001010000
		000000000
		101001101
		001000100
		000001001
		000000000
		001100100
	`)

	start := time.Now()
	result := board.LRBTMove(bitboard.E2, factMask, lrbtLookup)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		// fmt.Printf("%s\n%s\n%s\n", board.Rep(), expected.Rep(), result.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		001010100
		000010000
		001010000
		110000000
		111001111
		000000000
		000100001
		000010000
		001001100
	`)
	expected = bitboard.NewBitboardFromStr(`
		001010100
		000110000
		000010000
		110000000
		011001111
		000000000
		000100001
		000010000
		001001100
	`)

	start = time.Now()
	result = board.LRBTMove(bitboard.C7, factMask, lrbtLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		101000101
		000010010
		001010000
		100000000
		101011111
		001000000
		000001101
		010000000
		101001101
	`)
	expected = bitboard.NewBitboardFromStr(`
		101000100
		000010010
		001010100
		100001000
		101001111
		001100000
		001001101
		010000000
		001001101
	`)

	start = time.Now()
	result = board.LRBTMove(bitboard.E5, factMask, lrbtLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		101000101
		000010000
		001010000
		100000000
		101001111
		001000000
		000000001
		000000000
		101001101
	`)
	expected1 := bitboard.NewBitboardFromStr(`
		101000101
		000010000
		001010000
		100000000
		101001111
		001000000
		000000001
		000000000
		101001100
	`)

	expected2 := bitboard.NewBitboardFromStr(`
		001000101
		000010000
		001010000
		100000000
		101001111
		001000000
		000000001
		000000000
		101001101
	`)

	start = time.Now()
	result = board.LRBTMove(bitboard.I1, factMask, lrbtLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected1.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected1.Uint96)
	}

	start = time.Now()
	result = board.LRBTMove(bitboard.A9, factMask, lrbtLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected2.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected2.Uint96)
	}
}

func TestDirectionalMask(t *testing.T) {
	lrtbLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/lrtbsm")
	lrbtLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/lrbtsm")
	horizontalLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/horizontalsm")
	verticalLookup := bitboard.NewBoardDictionaryFromFile("../../artifacts/verticalsm")
	factMask := bitboard.NewFactBoardDictionaryFromFile("../../artifacts/factmask")

	board := bitboard.NewBitboardFromStr(`
		101000101
		000010010
		001010000
		000000000
		101011111
		001000100
		000000001
		000010000
		101000100
	`)
	expected := bitboard.NewBitboardFromStr(`
		001000100
		000000010
		001010100
		000111000
		001101000
		001111100
		001010101
		010010010
		101000101
	`)

	start := time.Now()
	result := board.DirectionalMove(bitboard.E5, factMask, lrtbLookup, lrbtLookup, horizontalLookup, verticalLookup)
	duration := time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		fmt.Printf("%s\n%s\n%s\n", board.Rep(), expected.Rep(), result.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		101000101
		000010000
		001010100
		000000000
		101011111
		001000100
		000000001
		000010000
		101000100
	`)
	expected = bitboard.NewBitboardFromStr(`
		000000110
		000010011
		001010101
		000000001
		101001111
		001000100
		000000000
		000010000
		001000100
	`)

	start = time.Now()
	result = board.DirectionalMove(bitboard.I9, factMask, lrtbLookup, lrbtLookup, horizontalLookup, verticalLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		fmt.Printf("%s\n%s\n%s\n", board.Rep(), expected.Rep(), result.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		101000101
		000010000
		001010100
		000000000
		101011111
		001000100
		000000001
		000010000
		101000101
	`)
	expected = bitboard.NewBitboardFromStr(`
		001000100
		000010000
		000010100
		000000000
		101011110
		001001100
		000000101
		000010011
		000000110
	`)

	start = time.Now()
	result = board.DirectionalMove(bitboard.I1, factMask, lrtbLookup, lrbtLookup, horizontalLookup, verticalLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		fmt.Printf("%s\n%s\n%s\n", board.Rep(), expected.Rep(), result.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		101000101
		000010000
		001010100
		000000000
		101011111
		001000100
		000000001
		000010000
		101000101
	`)
	expected = bitboard.NewBitboardFromStr(`
		001000100
		000010000
		001010000
		000000000
		101011111
		101100100
		101000001
		110010000
		011000000
	`)

	start = time.Now()
	result = board.DirectionalMove(bitboard.A1, factMask, lrtbLookup, lrbtLookup, horizontalLookup, verticalLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		fmt.Printf("%s\n%s\n%s\n", board.Rep(), expected.Rep(), result.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		101000101
		000010000
		001010100
		000000000
		101011111
		001000100
		000000001
		000010000
		101000101
	`)
	expected = bitboard.NewBitboardFromStr(`
		011000000
		110010000
		101010100
		100000000
		101001111
		001000100
		000000001
		000010000
		001000100
	`)

	start = time.Now()
	result = board.DirectionalMove(bitboard.A9, factMask, lrtbLookup, lrbtLookup, horizontalLookup, verticalLookup)
	duration = time.Since(start)
	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		fmt.Printf("%s\n%s\n%s\n", board.Rep(), expected.Rep(), result.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}
}
