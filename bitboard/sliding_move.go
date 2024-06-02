package bitboard

func (board *Bitboard) CalculateHorizontalSlidingMoves(square uint8) *Bitboard {

	// 0->8 bottom to top
	// rowIndex := square / 9

	// // 0->8 right to left
	colIndex := square % SIZE

	res := board.Uint96.Copy()

	// modify res for sliding moves

	// to left
	if SIZE-colIndex-1 > 0 {
		for li := range SIZE - colIndex - 1 {

			bitIndex := li + 1 + square
			res.SetBit(bitIndex, uint8(1))

			bit := res.GetBit(bitIndex)
			if bit == 1 {
				break
			}
		}
	}

	// to right
	if colIndex > 0 {
		for ri := range colIndex {

			bitIndex := square - (ri + 1)
			res.SetBit(bitIndex, uint8(1))

			bit := res.GetBit(bitIndex)
			if bit == 1 {
				break
			}
		}
	}

	// to right
	return &Bitboard{
		Uint96: &res,
	}
}

func (board *Bitboard) CalculateVerticalSlidingMoves(square uint8) *Bitboard {
	// 0->8 bottom to top
	// rowIndex := square / 9

	// // 0->8 right to left
	colIndex := square % SIZE

	res := board.Uint96.Copy()

	// modify res for sliding moves

	// to left
	if SIZE-colIndex-1 > 0 {
		for ui := range SIZE - colIndex - 1 {

			bitIndex := square + SIZE*(ui+1)
			res.SetBit(bitIndex, uint8(1))

			bit := res.GetBit(bitIndex)
			if bit == 1 {
				break
			}
		}
	}

	// to right
	if colIndex > 0 {
		for bi := range colIndex {

			bitIndex := square - SIZE*(bi+1)
			res.SetBit(bitIndex, uint8(1))

			bit := res.GetBit(bitIndex)
			if bit == 1 {
				break
			}
		}
	}

	// to right
	return &Bitboard{
		Uint96: &res,
	}
}
