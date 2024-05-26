package bitboard

import (
	"fmt"
	"slices"

	"github.com/pqviet030188/advance-chess-ai/uint96"
)

type Bitboard struct {
	*uint96.Uint96
}

func RandBitboard() Bitboard {
	random := uint96.RandUInt96()
	mBitboard := random.To81Bitboard()
	return Bitboard{
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

	for i := 0; i < 9; i++ {
		rows[i] = "000000000"
	}

	flag := 1
	for i := 0; i < 81; i++ {
		// do something
		rowIndex := i / 9
		bitColIndex := i % 9

		row := rows[rowIndex]

		before := ""
		if 9-bitColIndex-1 >= 0 && 9-bitColIndex-1 <= 9 {
			before = row[0 : 9-bitColIndex-1]
		}

		after := ""
		if 9-bitColIndex >= 0 && 9-bitColIndex < 9 {
			after = row[9-bitColIndex:]
		}

		rows[rowIndex] = fmt.Sprintf("%s%1b%s",
			before, b.GetBit(uint(i)), after)

		// shift bit
		flag <<= 1
	}

	slices.Reverse(rows)
	return rows
}

func (b *Bitboard) Rep() string {
	rep := b.StrArr()
	ret := ""
	for i := 0; i < 9; i++ {
		ret += fmt.Sprintf("%-5d", (9 - i))

		for j := 0; j < 9; j++ {
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
	for i := 0; i < 9; i++ {
		ret += fmt.Sprintf("I%d\nH%d\nG%d\nF%d\nE%d\nD%d\nC%d\nB%d\nA%d\n",
			i+1, i+1, i+1, i+1, i+1, i+1, i+1, i+1, i+1)
	}

	return ret
}

func (b *Bitboard) SetBit(square uint, bit uint) {
	b.Uint96.SetBit(square, bit)
}
