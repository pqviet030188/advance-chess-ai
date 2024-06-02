package bitboard

func (board *Bitboard) CalculateHorizontalSlidingMoves(square uint8) *Bitboard {

	// 0->8 bottom to top
	// rowIndex := square / 9

	// // 0->8 right to left
	_, colIndex := ToRowCol(square)

	res := board.Uint96.Copy()
	res.SetBit(square, 0)

	// modify res for sliding moves

	// to left
	setZero := false

	if SIZE-colIndex-1 > 0 {
		for li := range SIZE - colIndex - 1 {

			bitIndex := li + 1 + square
			bit := res.GetBit(bitIndex)

			if !setZero {
				res.SetBit(bitIndex, uint8(1))
			} else {
				res.SetBit(bitIndex, uint8(0))
			}

			if bit == 1 && !setZero {
				setZero = true
			}
		}
	}

	// to right
	setZero = false
	if colIndex > 0 {
		for ri := range colIndex {

			bitIndex := square - (ri + 1)
			bit := res.GetBit(bitIndex)

			if !setZero {
				res.SetBit(bitIndex, uint8(1))
			} else {
				res.SetBit(bitIndex, uint8(0))
			}

			if bit == 1 && !setZero {
				setZero = true
			}
		}
	}

	return &Bitboard{
		Uint96: &res,
	}
}

func (board *Bitboard) CalculateVerticalSlidingMoves(square uint8) *Bitboard {
	// 0->8 bottom to top
	rowIndex, _ := ToRowCol(square)

	// // 0->8 right to left
	// colIndex := square % SIZE

	res := board.Uint96.Copy()
	res.SetBit(square, 0)

	// modify res for sliding moves

	// to top
	setZero := false
	if SIZE-rowIndex-1 > 0 {
		for ui := range SIZE - rowIndex - 1 {

			bitIndex := square + SIZE*(ui+1)
			bit := res.GetBit(bitIndex)

			if !setZero {
				res.SetBit(bitIndex, uint8(1))
			} else {
				res.SetBit(bitIndex, uint8(0))
			}

			if bit == 1 && !setZero {
				setZero = true
			}
		}
	}

	// to bottom
	setZero = false
	if rowIndex > 0 {
		for bi := range rowIndex {

			bitIndex := square - SIZE*(bi+1)
			bit := res.GetBit(bitIndex)
			if !setZero {
				res.SetBit(bitIndex, uint8(1))
			} else {
				res.SetBit(bitIndex, uint8(0))
			}

			if bit == 1 && !setZero {
				setZero = true
			}
		}
	}

	// to right
	return &Bitboard{
		Uint96: &res,
	}
}

func (board *Bitboard) CalculateLRTBDiagSlidingMoves(square uint8) *Bitboard {
	// 0->8 bottom to top
	rowIndex, _ := ToRowCol(square)

	// // 0->8 right to left
	// colIndex := square % SIZE

	res := board.Uint96.Copy()
	res.SetBit(square, 0)

	// modify res for sliding moves

	// to top
	setZero := false
	if SIZE-rowIndex-1 > 0 {
		prevSquare := square
		for range SIZE - rowIndex - 1 {

			bitIndex := prevSquare + SIZE + 1

			if bitIndex >= SIZE*SIZE {
				break
			}

			prevSquare = bitIndex

			bit := res.GetBit(bitIndex)

			if !setZero {
				res.SetBit(bitIndex, uint8(1))
			} else {
				res.SetBit(bitIndex, uint8(0))
			}

			if bit == 1 && !setZero {
				setZero = true
			}
		}
	}

	// to bottom
	setZero = false
	if rowIndex > 0 {
		prevSquare := square
		for range rowIndex {

			bitIndex := prevSquare - SIZE - 1

			if prevSquare < SIZE+1 {
				break
			}

			prevSquare = bitIndex

			bit := res.GetBit(bitIndex)
			if !setZero {
				res.SetBit(bitIndex, uint8(1))
			} else {
				res.SetBit(bitIndex, uint8(0))
			}

			if bit == 1 && !setZero {
				setZero = true
			}
		}
	}

	return &Bitboard{
		Uint96: &res,
	}
}

func (board *Bitboard) CalculateLRBTDiagSlidingMoves(square uint8) *Bitboard {
	// 0->8 bottom to top
	rowIndex, _ := ToRowCol(square)

	// // 0->8 right to left
	// colIndex := square % SIZE

	res := board.Uint96.Copy()
	res.SetBit(square, 0)

	// modify res for sliding moves

	// to top
	setZero := false
	if SIZE-rowIndex-1 > 0 {
		prevSquare := square
		for range SIZE - rowIndex - 1 {

			bitIndex := prevSquare + SIZE - 1

			if bitIndex >= SIZE*SIZE {
				break
			}

			prevSquare = bitIndex

			bit := res.GetBit(bitIndex)

			if !setZero {
				res.SetBit(bitIndex, uint8(1))
			} else {
				res.SetBit(bitIndex, uint8(0))
			}

			if bit == 1 && !setZero {
				setZero = true
			}
		}
	}

	// to bottom
	setZero = false
	if rowIndex > 0 {
		prevSquare := square
		for range rowIndex {

			bitIndex := prevSquare - SIZE + 1

			if prevSquare < SIZE-1 {
				break
			}

			prevSquare = bitIndex

			bit := res.GetBit(bitIndex)
			if !setZero {
				res.SetBit(bitIndex, uint8(1))
			} else {
				res.SetBit(bitIndex, uint8(0))
			}

			if bit == 1 && !setZero {
				setZero = true
			}
		}
	}

	return &Bitboard{
		Uint96: &res,
	}
}
