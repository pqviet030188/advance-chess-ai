package bitboard

import (
	"math"

	"github.com/pqviet030188/advance-chess-ai/uint96"
)

func GenerateSlidingMoveDictionary() (horizontal *BoardDictionary, vertical *BoardDictionary, lrtb *BoardDictionary, lrbt *BoardDictionary) {
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

func GenerateSlidingMoveDictionaryFile(horizontalPath string, verticalPath string, lrtbPath string, lrbtPath string) {
	horizontal, vertical, lrtb, lrbt := GenerateSlidingMoveDictionary()

	horizontal.ToFile(horizontalPath)
	vertical.ToFile(verticalPath)
	lrtb.ToFile(lrtbPath)
	lrbt.ToFile(lrbtPath)
}
