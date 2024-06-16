package bitboard

func (board *Bitboard) HorizontalMoveOnly(square uint8, factMask *FactBoardDictionary, horizontalDict *BoardDictionary) *Bitboard {

	// get horizontal mask
	numberMask, ok := factMask.Get(square, HORIZONTAL_MASK)
	if !ok {
		panic("Cannot find horizontal mask")
	}

	mask := Bitboard{
		Uint96: numberMask,
	}

	// extract horizontal line
	horNumber := board.And(*mask.Uint96)
	horBoard := Bitboard{
		Uint96: &horNumber,
	}

	// move the line to bottom
	horCalBoard, shiftedSquare := horBoard.ShiftToMoveCalPositionForHor(square)

	// get sliding moves via dictionary lookup
	horSlidingNumber, ok := horizontalDict.Get(horCalBoard.Uint96, shiftedSquare)
	if !ok {
		panic("Cannot find horizontal sliding board")
	}

	horSlidingBoard := &Bitboard{
		Uint96: horSlidingNumber,
	}

	// reverse to original spot
	horSlidingBoard = horSlidingBoard.ReverseShiftToMoveCalPositionForHor(square)
	horSlidingBoardNumber := horSlidingBoard.And(*mask.Uint96)
	horSlidingBoard = &Bitboard{
		Uint96: &horSlidingBoardNumber,
	}
	return horSlidingBoard
}

func (board *Bitboard) VerticalMoveOnly(square uint8, factMask *FactBoardDictionary, verticalDict *BoardDictionary) *Bitboard {

	// get vertical mask
	numberMask, ok := factMask.Get(square, VERTICAL_MASK)
	if !ok {
		panic("Cannot find vertical mask")
	}

	mask := Bitboard{
		Uint96: numberMask,
	}

	// extract vertical line
	verNumber := board.And(*mask.Uint96)
	verBoard := Bitboard{
		Uint96: &verNumber,
	}

	// move the column to right
	verCalBoard, shiftedSquare := verBoard.ShiftToMoveCalPositionForVer(square)

	// row, _ := ToRowCol(square)

	// get sliding moves via dictionary lookup
	verSlidingNumber, ok := verticalDict.Get(verCalBoard.Uint96, shiftedSquare)
	if !ok {
		panic("Cannot find vertical sliding board")
	}

	verSlidingBoard := &Bitboard{
		Uint96: verSlidingNumber,
	}

	// reverse to original spot
	verSlidingBoard = verSlidingBoard.ReverseShiftToMoveCalPositionForVer(square)
	verSlidingBoardNumber := verSlidingBoard.And(*mask.Uint96)
	verSlidingBoard = &Bitboard{
		Uint96: &verSlidingBoardNumber,
	}
	return verSlidingBoard
}

func (board *Bitboard) LRTBMoveOnly(square uint8, factMask *FactBoardDictionary, lrtbDict *BoardDictionary) *Bitboard {
	// get lrtb mask
	numberMask, ok := factMask.Get(square, LRTB_MASK)
	if !ok {
		panic("Cannot find lrtb mask")
	}

	mask := Bitboard{
		Uint96: numberMask,
	}

	// extract lrtb line
	lrtbNumber := board.And(*mask.Uint96)
	lrtbBoard := Bitboard{
		Uint96: &lrtbNumber,
	}

	// move the column to right
	calBoard, shiftedSquare := lrtbBoard.ShiftToMoveCalPositionForLRTB(square)

	// get sliding moves via dictionary lookup
	lrtbSlidingNumber, ok := lrtbDict.Get(calBoard.Uint96, shiftedSquare)
	if !ok {
		panic("Cannot find lrtb sliding board")
	}

	lrtbSlidingBoard := &Bitboard{
		Uint96: lrtbSlidingNumber,
	}

	// reverse to original spot
	lrtbSlidingBoard = lrtbSlidingBoard.ReverseShiftToMoveCalPositionForLRTB(square)
	lrtbSlidingBoardNumber := lrtbSlidingBoard.And(*mask.Uint96)
	lrtbSlidingBoard = &Bitboard{
		Uint96: &lrtbSlidingBoardNumber,
	}

	return lrtbSlidingBoard
}

func (board *Bitboard) LRBTMoveOnly(square uint8, factMask *FactBoardDictionary, lrbtDict *BoardDictionary) *Bitboard {
	// get lrbt mask
	numberMask, ok := factMask.Get(square, LRBT_MASK)
	if !ok {
		panic("Cannot find lrbt mask")
	}

	mask := Bitboard{
		Uint96: numberMask,
	}

	// extract lrbt line
	lrbtNumber := board.And(*mask.Uint96)
	lrbtBoard := Bitboard{
		Uint96: &lrbtNumber,
	}

	// move the column to right
	calBoard, shiftedSquare := lrbtBoard.ShiftToMoveCalPositionForLRBT(square)

	// row, _ := ToRowCol(square)

	// get sliding moves via dictionary lookup
	lrbtSlidingNumber, ok := lrbtDict.Get(calBoard.Uint96, shiftedSquare)
	if !ok {
		panic("Cannot find lrbt sliding board")
	}

	lrbtSlidingBoard := &Bitboard{
		Uint96: lrbtSlidingNumber,
	}

	// reverse to original spot
	lrbtSlidingBoard = lrbtSlidingBoard.ReverseShiftToMoveCalPositionForLRBT(square)
	lrbtSlidingBoardNumber := lrbtSlidingBoard.And(*mask.Uint96)
	lrbtSlidingBoard = &Bitboard{
		Uint96: &lrbtSlidingBoardNumber,
	}

	return lrbtSlidingBoard
}

func (board *Bitboard) HorizontalMove(square uint8, factMask *FactBoardDictionary, horizontalDict *BoardDictionary) *Bitboard {

	// not
	notNumberMask, ok := factMask.Get(square, NOT_HORIZONTAL_MASK)
	if !ok {
		panic("Cannot find horizontal mask")
	}

	notMask := Bitboard{
		Uint96: notNumberMask,
	}

	notHorNumber := board.And(*notMask.Uint96)
	notHorBoard := Bitboard{
		Uint96: &notHorNumber,
	}

	// hor moves
	horSlidingBoard := board.HorizontalMoveOnly(square, factMask, horizontalDict)

	resultNumber := notHorBoard.Or(*horSlidingBoard.Uint96)
	return &Bitboard{
		Uint96: &resultNumber,
	}
}

func (board *Bitboard) VerticalMove(square uint8, factMask *FactBoardDictionary, verticalDict *BoardDictionary) *Bitboard {

	// not
	notNumberMask, ok := factMask.Get(square, NOT_VERTICAL_MASK)
	if !ok {
		panic("Cannot find vertical mask")
	}

	notMask := Bitboard{
		Uint96: notNumberMask,
	}

	notVerNumber := board.And(*notMask.Uint96)
	notVerBoard := Bitboard{
		Uint96: &notVerNumber,
	}

	// ver moves
	verSlidingBoard := board.VerticalMoveOnly(square, factMask, verticalDict)

	resultNumber := notVerBoard.Or(*verSlidingBoard.Uint96)
	return &Bitboard{
		Uint96: &resultNumber,
	}
}

func (board *Bitboard) LRTBMove(square uint8, factMask *FactBoardDictionary, lrtbDict *BoardDictionary) *Bitboard {

	// not
	notNumberMask, ok := factMask.Get(square, NOT_LRTB_MASK)
	if !ok {
		panic("Cannot find lrtb mask")
	}

	notMask := Bitboard{
		Uint96: notNumberMask,
	}

	notlrtbNumber := board.And(*notMask.Uint96)
	notlrtbBoard := Bitboard{
		Uint96: &notlrtbNumber,
	}

	// lookup move
	lrtbSlidingBoard := board.LRTBMoveOnly(square, factMask, lrtbDict)
	resultNumber := notlrtbBoard.Or(*lrtbSlidingBoard.Uint96)
	return &Bitboard{
		Uint96: &resultNumber,
	}
}

func (board *Bitboard) LRBTMove(square uint8, factMask *FactBoardDictionary, lrbtDict *BoardDictionary) *Bitboard {

	// not
	notNumberMask, ok := factMask.Get(square, NOT_LRBT_MASK)
	if !ok {
		panic("Cannot find lrbt mask")
	}

	notMask := Bitboard{
		Uint96: notNumberMask,
	}

	notlrbtNumber := board.And(*notMask.Uint96)
	notlrbtBoard := Bitboard{
		Uint96: &notlrbtNumber,
	}

	// lookup move
	lrbtSlidingBoard := board.LRBTMoveOnly(square, factMask, lrbtDict)
	resultNumber := notlrbtBoard.Or(*lrbtSlidingBoard.Uint96)
	return &Bitboard{
		Uint96: &resultNumber,
	}
}

func (board *Bitboard) DirectionalMove(square uint8, factMask *FactBoardDictionary,
	lrtbDict *BoardDictionary, lrbtDict *BoardDictionary,
	horizontalDict *BoardDictionary, verticalDict *BoardDictionary) *Bitboard {

	// lookup move
	lrbtSlidingBoard := board.LRBTMoveOnly(square, factMask, lrbtDict)
	lrtbSlidingBoard := board.LRTBMoveOnly(square, factMask, lrtbDict)
	horSlidingBoard := board.HorizontalMoveOnly(square, factMask, horizontalDict)
	verSlidingBoard := board.VerticalMoveOnly(square, factMask, verticalDict)

	slidingNumber := lrbtSlidingBoard.Or(*lrtbSlidingBoard.Uint96).Or(*horSlidingBoard.Uint96).Or(*verSlidingBoard.Uint96)

	// not
	notmaskNumber, ok := factMask.Get(square, NOT_DIRECTIONAL_MASK)
	if !ok {
		panic("Cannot find lrbt mask")
	}
	notNumber := board.And(*notmaskNumber)

	resultNumber := slidingNumber.Or(notNumber)

	return &Bitboard{
		Uint96: &resultNumber,
	}
}
