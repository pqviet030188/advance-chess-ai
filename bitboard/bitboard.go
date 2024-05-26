package bitboard

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pqviet030188/advance-chess-ai/uint96"
)

type bitboard struct {
	*uint96.Uint96
}

type Blah struct {
	Lo, Mid, Hi uint32
}

func RandBitboard() bitboard {
	random := uint96.RandUInt96()
	mBitboard := random.To81Bitboard()
	return bitboard{
		Uint96: &mBitboard,
	}
}

func (b *bitboard) Str() string {
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

		flag <<= 1
	}

	slices.Reverse(rows)
	return strings.Join(rows, "\n")
}
