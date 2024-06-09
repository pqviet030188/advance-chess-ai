package bitboard_tests

import (
	"math"
	"testing"

	. "github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/uint96"
)

func generateSlidingMoveDictionary() (horizontal *BoardDictionary, vertical *BoardDictionary, lrtb *BoardDictionary, lrbt *BoardDictionary) {
	bitCount := uint16(9)
	lineRange := (uint16)(math.Pow(2, float64(bitCount))) - 1

	dict := NewBoardDictionary()

	// horizontal line
	occupancy := uint96.FromUInt32(0)
	for occ := range lineRange {

		// build occupancy
		for bitIndex := range bitCount {
			bitValue := (uint8)(occ>>bitIndex) & 1

			occupancy.SetBit(uint8(bitIndex), bitValue)
		}

		// build sliding attack
		for bitIndex := range bitCount {
			square := uint8(bitIndex)

			// copy to build bitboard for occupancy
			boardNumber := occupancy.Copy()
			board := Bitboard{
				Uint96: &boardNumber,
			}

			// calculate sliding moves
			resBoard := board.CalculateHorizontalSlidingMoves(square)
			dict.Put(&occupancy, square, resBoard.Uint96)
		}
	}
	horizontal = dict

	// vertical line
	dict = NewBoardDictionary()
	occupancy = uint96.FromUInt32(0)
	for occ := range lineRange {

		// build occupancy
		for bitIndex := range bitCount {
			bitValue := (uint8)(occ>>bitIndex) & 1

			square := uint8(bitIndex * bitCount)
			occupancy.SetBit(square, bitValue)
		}

		// build sliding attack
		for bitIndex := range bitCount {
			square := uint8(bitIndex * bitCount)

			// copy to build bitboard for occupancy
			boardNumber := occupancy.Copy()
			board := Bitboard{
				Uint96: &boardNumber,
			}

			// calculate sliding moves
			resBoard := board.CalculateVerticalSlidingMoves(square)
			dict.Put(&occupancy, square, resBoard.Uint96)
		}
	}
	vertical = dict

	// lrtb line
	dict = NewBoardDictionary()
	occupancy = uint96.FromUInt32(0)
	for occ := range lineRange {

		// build occupancy
		for bitIndex := range bitCount {
			bitValue := (uint8)(occ>>bitIndex) & 1

			square := uint8(bitIndex*bitCount + bitIndex)
			occupancy.SetBit(square, bitValue)
		}

		// build sliding attack
		for bitIndex := range bitCount {
			square := uint8(bitIndex*bitCount + bitIndex)

			// copy to build bitboard for occupancy
			boardNumber := occupancy.Copy()
			board := Bitboard{
				Uint96: &boardNumber,
			}

			// calculate sliding moves
			resBoard := board.CalculateLRTBDiagSlidingMoves(square)
			dict.Put(&occupancy, square, resBoard.Uint96)
		}
	}
	lrtb = dict

	// lrbt line
	dict = NewBoardDictionary()
	occupancy = uint96.FromUInt32(0)
	for occ := range lineRange {

		// build occupancy
		for bitIndex := range bitCount {
			bitValue := (uint8)(occ>>bitIndex) & 1

			square := uint8(bitIndex*bitCount + (bitCount - 1 - bitIndex))
			occupancy.SetBit(square, bitValue)
		}

		// build sliding attack
		for bitIndex := range bitCount {
			square := uint8(bitIndex*bitCount + (bitCount - 1 - bitIndex))

			// copy to build bitboard for occupancy
			boardNumber := occupancy.Copy()
			board := Bitboard{
				Uint96: &boardNumber,
			}

			// calculate sliding moves
			resBoard := board.CalculateLRBTDiagSlidingMoves(square)
			dict.Put(&occupancy, square, resBoard.Uint96)
		}
	}
	lrbt = dict

	return
}

func compareTwoDict(exp *BoardDictionary, got *BoardDictionary, t *testing.T) {
	expKeys1 := exp.Keys()
	expKeys := &expKeys1
	gotKeys1 := got.Keys()
	gotKeys := &gotKeys1

	if len(*expKeys) != len(*gotKeys) {
		t.Errorf("Expected key count to be the same, Result was incorrect, got: %x, want: %x.", len(*gotKeys), len(*expKeys))
	}

	gotKeyDict := map[BoardDictionaryKey]bool{}
	expKeyDict := map[BoardDictionaryKey]bool{}

	for _, key := range *expKeys {
		bkey := BoardDictionaryKey{
			Occupancy: key.Occupancy,
			Square:    key.Square,
		}

		expKeyDict[bkey] = true
	}

	for _, key := range *gotKeys {
		bkey := BoardDictionaryKey{
			Occupancy: key.Occupancy,
			Square:    key.Square,
		}

		gotKeyDict[bkey] = true
	}

	for key := range expKeyDict {
		if has := gotKeyDict[key]; !has {
			t.Errorf("Expected to contain the key, Result was incorrect, got: %s, %d, %t, want: %s, %d, %t.", key.Occupancy.Str(), key.Square, has,
				key.Occupancy.Str(), key.Square, true)
		}
	}

	for _, key := range expKeys1 {
		expValue, hasKeyInExp := exp.Get(&key.Occupancy, key.Square)
		gotValue, hasKeyInGot := got.Get(&key.Occupancy, key.Square)

		if !hasKeyInExp || !hasKeyInGot || !expValue.Equals(*gotValue) {
			t.Errorf("Expected to be equal, Result was incorrect, got: %s, want: %s.", gotValue.Str(), expValue.Str())
		}
	}
}

func TestGenerateSlidingMoveDictionaryFile(t *testing.T) {
	horizontalPath, verticalPath, lrtbPath, lrbtPath := "../../artifacts/horizontalsm", "../../artifacts/verticalsm", "../../artifacts/lrtbsm", "../../artifacts/lrbtsm"

	GenerateSlidingMoveDictionaryFile(horizontalPath, verticalPath, lrtbPath, lrbtPath)
	horizontalDict := NewBoardDictionaryFromFile(horizontalPath)
	verticalDict := NewBoardDictionaryFromFile(verticalPath)
	lrtbDict := NewBoardDictionaryFromFile(lrtbPath)
	lrbtDict := NewBoardDictionaryFromFile(lrbtPath)

	horizontal, vertical, lrtb, lrbt := generateSlidingMoveDictionary()

	// fmt.Printf("hor count %d \n", horizontalDict.KeyCount())
	// fmt.Printf("ver count %d \n", verticalDict.KeyCount())
	// fmt.Printf("lrtb count %d \n", lrtbDict.KeyCount())
	// fmt.Printf("lrbt count %d \n", lrbtDict.KeyCount())

	compareTwoDict(horizontal, horizontalDict, t)
	compareTwoDict(vertical, verticalDict, t)
	compareTwoDict(lrtb, lrtbDict, t)
	compareTwoDict(lrbt, lrbtDict, t)
}
