package bitboard_tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/pqviet030188/advance-chess-ai/bitboard"
)

func TestSentinelProtectionByBitboard(t *testing.T) {
	factMask := bitboard.NewFactBoardDictionaryFromFile("../../artifacts/factmask")

	board := bitboard.NewBitboardFromStr(`
		100000001
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		100000001
	`)

	expected := bitboard.NewBitboardFromStr(`
		010000010
		100000001
		000000000
		000000000
		000000000
		000000000
		000000000
		100000001
		010000010
	`)

	start := time.Now()
	result := bitboard.SentinelProtection(board, factMask)
	duration := time.Since(start)

	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		fmt.Printf("board:\n%s\nresult:\n%s\nexpected:\n%s", board.Rep(), result.Rep(), expected.Rep())
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		100000001
		100000001
		100000001
		000000000
		000000000
		000000000
		100000001
		100000001
		100000001
	`)

	expected = bitboard.NewBitboardFromStr(`
		110000011
		110000011
		110000011
		100000001
		000000000
		100000001
		110000011
		110000011
		110000011
	`)

	start = time.Now()
	result = bitboard.SentinelProtection(board, factMask)
	duration = time.Since(start)

	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		000000000
		100000001
		100000001
		000000000
		000000000
		000000000
		100000001
		100000001
		000000000
	`)

	expected = bitboard.NewBitboardFromStr(`
		100000001
		110000011
		110000011
		100000001
		000000000
		100000001
		110000011
		110000011
		100000001
	`)

	start = time.Now()
	result = bitboard.SentinelProtection(board, factMask)
	duration = time.Since(start)

	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		011000110
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		011000110
	`)

	expected = bitboard.NewBitboardFromStr(`
		111101111
		011000110
		000000000
		000000000
		000000000
		000000000
		000000000
		011000110
		111101111
	`)

	start = time.Now()
	result = bitboard.SentinelProtection(board, factMask)
	duration = time.Since(start)

	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		111000111
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		000000000
		111000111
	`)

	expected = bitboard.NewBitboardFromStr(`
		111101111
		111000111
		000000000
		000000000
		000000000
		000000000
		000000000
		111000111
		111101111
	`)

	start = time.Now()
	result = bitboard.SentinelProtection(board, factMask)
	duration = time.Since(start)

	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		000000000
		010000010
		000000000
		000000000
		011111110
		000000000
		000000000
		010000010
		000000000
	`)

	expected = bitboard.NewBitboardFromStr(`
		010000010
		101000101
		010000010
		011111110
		111111111
		011111110
		010000010
		101000101
		010000010
	`)

	start = time.Now()
	result = bitboard.SentinelProtection(board, factMask)
	duration = time.Since(start)

	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

	board = bitboard.NewBitboardFromStr(`
		100000001
		010000010
		100000001
		000000000
		011111110
		000000000
		000000000
		010000010
		100000001
	`)

	expected = bitboard.NewBitboardFromStr(`
		010000010
		101000101
		010000010
		111111111
		111111111
		011111110
		010000010
		101000101
		010000010
	`)

	start = time.Now()
	result = bitboard.SentinelProtection(board, factMask)
	duration = time.Since(start)

	fmt.Printf("time taken: %d\n", duration.Nanoseconds())

	if !result.Uint96.Equals(*expected.Uint96) {
		t.Errorf("Expected values to be the same, Result was incorrect, got: %x, want: %x.", result.Uint96, expected.Uint96)
	}

}
