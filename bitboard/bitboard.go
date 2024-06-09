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
	for i := uint8(0); i < SIZE*SIZE; i++ {
		// do something
		// rowIndex := i / int(SIZE)
		// bitColIndex := i % int(SIZE)
		rowIndex, bitColIndex := ToRowCol(i)

		row := rows[rowIndex]

		before := ""
		if SIZE >= bitColIndex+1 && (SIZE-bitColIndex-1) <= SIZE {
			before = row[0 : SIZE-bitColIndex-1]
		}

		after := ""
		if SIZE >= bitColIndex && SIZE-bitColIndex < SIZE {
			after = row[SIZE-bitColIndex:]
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

func (b *Bitboard) Copy() *Bitboard {
	copy := b.Uint96.Copy()
	return &Bitboard{
		Uint96: &copy,
	}
}

func ToRowCol(square uint8) (row uint8, col uint8) {
	row = square / SIZE
	col = square % SIZE
	return
}

func (b *Bitboard) ShiftToMoveCalPositionForLRTB(square uint8) *Bitboard {

	if square == SIZE*SIZE-1 || square == SIZE-1 {
		res := uint96.FromUInt32(0)

		return &Bitboard{
			Uint96: &res,
		}
	}

	if square == 0 || square == SIZE*SIZE-SIZE {
		return b
	}

	row, col := ToRowCol(square)
	// fmt.Printf("row col %d %d\n", row, col)
	if row == col {
		return b
	}

	index := SIZE - 1 - row
	// if SIZE-1-col < SIZE-1-row {
	if row < col {
		index = SIZE - 1 - col
	}

	newSquareRow, newSquareCol := ToRowCol(square + index*(SIZE+1))

	// index := row
	// delta := newSquareRow - newSquareCol
	// fmt.Printf("nrow ncol %d %d %d %d\n", index, square+index*(SIZE+1), newSquareRow, newSquareCol)

	if newSquareRow == newSquareCol {
		return b
	}

	if newSquareRow > newSquareCol {
		res := b.Uint96.Lsh(uint(newSquareRow - newSquareCol))

		return &Bitboard{
			Uint96: &res,
		}
	}

	res := b.Uint96.Rsh(uint(newSquareCol - newSquareRow))

	return &Bitboard{
		Uint96: &res,
	}
}
