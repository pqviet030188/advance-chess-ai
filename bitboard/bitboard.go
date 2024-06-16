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

func ToSquare(row uint8, col uint8) uint8 {
	return row*SIZE + col
}

func shiftDeltaForLRTB(square uint8) int {
	row, col := ToRowCol(square)
	if row == col {
		return 0
	}

	index := SIZE - 1 - row

	if row < col {
		index = SIZE - 1 - col
	}

	newSquareRow, newSquareCol := ToRowCol(square + index*(SIZE+1))

	if newSquareRow == newSquareCol {
		return 0
	}

	if newSquareRow > newSquareCol {
		return (int(newSquareRow) - int(newSquareCol))
	}

	return -(int(newSquareCol) - int(newSquareRow))
}

func (b *Bitboard) Shift(delta int) *Bitboard {
	if delta == 0 {
		return b
	}

	if delta < 0 {
		res := b.Uint96.Rsh(uint(-delta))

		return &Bitboard{
			Uint96: &res,
		}
	}

	res := b.Uint96.Lsh(uint(delta))

	return &Bitboard{
		Uint96: &res,
	}
}
func (b *Bitboard) ShiftToMoveCalPositionForLRTB(square uint8) *Bitboard {

	delta := shiftDeltaForLRTB(square)
	return b.Shift(delta)
}

func (b *Bitboard) ReverseShiftToMoveCalPositionForLRTB(orgSquare uint8) *Bitboard {

	delta := shiftDeltaForLRTB(orgSquare)
	return b.Shift(-delta)
}

func shiftDeltaForLRBT(square uint8) int {

	row, col := ToRowCol(square)
	if row+col == SIZE-1 {
		return 0
	}

	// below the line
	if row+col < SIZE-1 {
		return int((SIZE - 1) - (row + col))
	}
	return -int((row + col) - (SIZE - 1))
}

func (b *Bitboard) ShiftToMoveCalPositionForLRBT(square uint8) *Bitboard {

	delta := shiftDeltaForLRBT(square)
	return b.Shift(delta)
}

func (b *Bitboard) ReverseShiftToMoveCalPositionForLRBT(orgSquare uint8) *Bitboard {

	delta := shiftDeltaForLRBT(orgSquare)
	return b.Shift(-delta)
}

func shiftDeltaForHor(square uint8) int {

	row, _ := ToRowCol(square)
	delta := row * SIZE
	return -(int(delta))
}

func shiftDeltaForVer(square uint8) int {

	_, col := ToRowCol(square)
	return -(int(col))
}

func (b *Bitboard) ShiftToMoveCalPositionForHor(square uint8) *Bitboard {

	delta := shiftDeltaForHor(square)
	return b.Shift(delta)
}

func (b *Bitboard) ReverseShiftToMoveCalPositionForHor(orgSquare uint8) *Bitboard {

	delta := shiftDeltaForHor(orgSquare)
	return b.Shift(-delta)
}

func (b *Bitboard) ShiftToMoveCalPositionForVer(square uint8) *Bitboard {

	delta := shiftDeltaForVer(square)
	return b.Shift(delta)
}

func (b *Bitboard) ReverseShiftToMoveCalPositionForVer(orgSquare uint8) *Bitboard {
	delta := shiftDeltaForVer(orgSquare)
	return b.Shift(-delta)
}
