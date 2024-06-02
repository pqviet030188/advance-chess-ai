package bitboard

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/pqviet030188/advance-chess-ai/uint96"
	"github.com/pqviet030188/advance-chess-ai/utilities"
)

type Bitboard struct {
	*uint96.Uint96
}

func NewBitboardFromStr(board string) *Bitboard {
	convertToBit := func(value string) uint8 {
		i, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			panic(err)
		}
		return uint8(i)
	}

	ret := uint96.FromUInt32(0)

	littleToHighSlices := utilities.SliceMap(
		strings.Split(strings.ReplaceAll(
			strings.ReplaceAll(strings.ReplaceAll(board, "\n", ""), " ", ""), "\t", ""), ""), convertToBit)

	slices.Reverse(littleToHighSlices)

	for i := range len(littleToHighSlices) {
		ret.SetBit(uint8(i), littleToHighSlices[i])
	}

	return &Bitboard{
		Uint96: &ret,
	}
}

func RandBitboard() *Bitboard {
	random := uint96.RandUInt96()
	mBitboard := random.To81Bitboard()
	return &Bitboard{
		Uint96: &mBitboard,
	}
}

func (b *Bitboard) StrArr() []string {
	rows := []string{
		// bottom
		"000000000",
		"000000000",
		"000000000",
		"000000000",
		"000000000",
		"000000000",
		"000000000",
		"000000000",
		"000000000",
		//top
	}

	for i := 0; i < int(SIZE); i++ {
		rows[i] = "000000000"
	}

	flag := 1
	for i := 0; i < int(SIZE)*int(SIZE); i++ {
		// do something
		rowIndex := i / int(SIZE)
		bitColIndex := i % int(SIZE)

		row := rows[rowIndex]

		before := ""
		if int(SIZE)-bitColIndex-1 >= 0 && int(SIZE)-bitColIndex-1 <= int(SIZE) {
			before = row[0 : int(SIZE)-bitColIndex-1]
		}

		after := ""
		if int(SIZE)-bitColIndex >= 0 && int(SIZE)-bitColIndex < int(SIZE) {
			after = row[int(SIZE)-bitColIndex:]
		}

		rows[rowIndex] = fmt.Sprintf("%s%1b%s",
			before, b.GetBit(uint8(i)), after)

		// shift bit
		flag <<= 1
	}

	slices.Reverse(rows)
	return rows
}

func (b *Bitboard) Rep() string {
	rep := b.StrArr()
	ret := ""
	for i := 0; i < int(SIZE); i++ {
		ret += fmt.Sprintf("%-5d", (int(SIZE) - i))

		for j := 0; j < int(SIZE); j++ {
			ret += fmt.Sprintf("%s ", string(rep[i][j]))
		}

		ret += fmt.Sprintf("%s\n", "")
	}
	ret += "\n"
	ret += fmt.Sprintf("%-5s", "")
	ret += fmt.Sprintf("%s\n", "A B C D E F G H I ")

	return ret
}

func (b *Bitboard) CoorStr() string {
	var ret string = ""
	for i := 0; i < int(SIZE); i++ {
		ret += fmt.Sprintf("I%d\nH%d\nG%d\nF%d\nE%d\nD%d\nC%d\nB%d\nA%d\n",
			i+1, i+1, i+1, i+1, i+1, i+1, i+1, i+1, i+1)
	}

	return ret
}

func (b *Bitboard) SetBit(square uint8, bit uint8) {
	b.Uint96.SetBit(square, bit)
}
