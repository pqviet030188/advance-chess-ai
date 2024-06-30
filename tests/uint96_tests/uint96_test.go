package uint96_tests

import (
	"math/big"
	mrand "math/rand/v2"
	"testing"

	. "github.com/pqviet030188/advance-chess-ai/uint96"
	"github.com/pqviet030188/advance-chess-ai/utilities"
)

func zeroUInt96() Uint96 {
	randBuf := make([]byte, 12)
	return FromBytes(randBuf)
}

func TestZero(t *testing.T) {
	result := zeroUInt96()
	expected := FromUInt32(0)

	if result != expected {
		t.Errorf("Expected to be zero, Result was incorrect, got: %b, want: %b.", expected, result)
	}
}

func TestEqual(t *testing.T) {
	resultBytes := *Randomise12Byte()
	expectedBytes := make([]byte, len(resultBytes))
	copy(expectedBytes, resultBytes)

	result := FromBytes(resultBytes)
	expected := FromBytes(expectedBytes)

	if !result.Equals(expected) {
		t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", expected, result)
	}

	expectedBytes[0] = expectedBytes[0] + 1
	expected = FromBytes(expectedBytes)
	if result.Equals(expected) {
		t.Errorf("Expected to be not equals, Result was incorrect, got: %b, want: %b.", expected, result)
	}
}

func TestAnd(t *testing.T) {
	for i := 0; i < 100; i++ {
		num1 := RandUInt96()
		num2 := RandUInt96()
		num := num1.And(num2).Big()

		num1Big := num1.Big()
		res := num1Big.And(num1Big, num2.Big())

		if num.Cmp(res) != 0 {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", num, res)
		}
	}
}

func TestOr(t *testing.T) {
	for i := 0; i < 100; i++ {
		num1 := RandUInt96()
		num2 := RandUInt96()

		num := num1.Or(num2).Big()

		num1Big := num1.Big()
		res := num1Big.Or(num1Big, num2.Big())

		if num.Cmp(res) != 0 {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", num, res)
		}
	}
}

func TestXor(t *testing.T) {
	for i := 0; i < 100; i++ {
		num1 := RandUInt96()
		num2 := RandUInt96()
		num := num1.Xor(num2).Big()

		num1Big := num1.Big()
		res := num1Big.Xor(num1Big, num2.Big())

		if num.Cmp(res) != 0 {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", num, res)
		}
	}
}

func lsh(bitshift uint, number Uint96, t *testing.T) {

	mask := Uint96{
		Lo:  0xffffffff,
		Mid: 0xffffffff,
		Hi:  0xffffffff,
	}

	// shift 70 bits
	leftNum := number.Lsh(bitshift)
	leftNumBig := leftNum.Big()

	numBig := number.Big()
	leftNumBigRes := numBig.Lsh(numBig, bitshift)
	leftNumBigRes = leftNumBigRes.And(leftNumBigRes, mask.Big())

	if leftNumBig.Cmp(leftNumBigRes) != 0 {
		t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", leftNumBig, leftNumBigRes)
	}
}

func rsh(bitshift uint, number Uint96, t *testing.T) {

	mask := Uint96{
		Lo:  0xffffffff,
		Mid: 0xffffffff,
		Hi:  0xffffffff,
	}

	// shift 70 bits
	rightNum := number.Rsh(bitshift)
	rightNumBig := rightNum.Big()

	numBig := number.Big()
	rightNumBigRes := numBig.Rsh(numBig, bitshift)
	rightNumBigRes = rightNumBigRes.And(rightNumBigRes, mask.Big())

	if rightNumBig.Cmp(rightNumBigRes) != 0 {
		t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", rightNumBig, rightNumBigRes)
	}
}

func TestLsh(t *testing.T) {
	for i := 0; i < 100; i++ {

		num := RandUInt96()

		// shift 96 bits
		lsh(96, num, t)

		// shift 95 bits
		lsh(95, num, t)

		// shift 70 bits
		lsh(70, num, t)

		// shift 64 bits
		lsh(64, num, t)

		// shift 60 bits
		lsh(60, num, t)

		// shift 32 bits
		lsh(32, num, t)

		// shift 30, 20 bits
		lsh(30, num, t)
		lsh(20, num, t)

		// shift 1 bit
		lsh(1, num, t)

		// shift 0 bit
		lsh(0, num, t)
	}
}

func TestRsh(t *testing.T) {
	for i := 0; i < 100; i++ {

		num := RandUInt96()

		// shift 96 bits
		rsh(96, num, t)

		// shift 95 bits
		rsh(95, num, t)

		// shift 70 bits
		rsh(70, num, t)

		// shift 64 bits
		rsh(64, num, t)

		// shift 60 bits
		rsh(60, num, t)

		// shift 32 bits
		rsh(32, num, t)

		// shift 30, 20 bits
		rsh(30, num, t)
		rsh(20, num, t)

		// shift 1 bit
		rsh(1, num, t)

		// shift 0 bit
		rsh(0, num, t)
	}
}

func TestAdd(t *testing.T) {
	for i := 0; i < 100; i++ {

		num1 := RandUInt96()
		num2 := RandUInt96()
		sum, _ := num1.Add(num2)
		sumBig := sum.Big()

		mask := Uint96{
			Lo:  0xffffffff,
			Mid: 0xffffffff,
			Hi:  0xffffffff,
		}

		num1Big := num1.Big()
		num2Big := num2.Big()

		sumBigRes := num1Big.Add(num1Big, num2Big)
		sumBigRes = sumBigRes.And(sumBigRes, mask.Big())

		if sumBig.Cmp(sumBigRes) != 0 {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", sumBig, sumBigRes)
		}
	}
}

func TestSub(t *testing.T) {
	for i := 0; i < 100; i++ {

		num1 := RandUInt96()
		num2 := RandUInt96()

		sub, _ := num1.Sub(num2)
		subBig := sub.Big()

		mask := Uint96{
			Lo:  0xffffffff,
			Mid: 0xffffffff,
			Hi:  0xffffffff,
		}

		num1Big := num1.Big()
		num2Big := num2.Big()

		subBigRes := num1Big.Sub(num1Big, num2Big)
		subBigRes = subBigRes.And(subBigRes, mask.Big())

		if subBig.Cmp(subBigRes) != 0 {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", subBig, subBigRes)
		}
	}
}

func TestReverse(t *testing.T) {
	for i := 0; i < 100; i++ {

		num := RandUInt96()
		reverse := num.Reverse()
		start := reverse.Reverse()
		startBig := start.Big()

		resBig := num.Big()

		if startBig.Cmp(resBig) != 0 {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", startBig, resBig)
		}
	}
}

func TestFromBig(t *testing.T) {
	for i := 0; i < 100; i++ {

		num := RandUInt96()
		resBig := num.Big()
		num2 := FromBig(resBig)

		if num.Cmp(num2) != 0 {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", num2, num)
		}
	}
}

func TestFromBytes(t *testing.T) {
	for i := 0; i < 100; i++ {

		num := RandUInt96()
		bytes := num.ToBytes()
		num2 := FromBytes(bytes)

		if num.Cmp(num2) != 0 {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", num2, num)
		}
	}
}

func TestCopy(t *testing.T) {
	for i := 0; i < 100; i++ {

		num := RandUInt96()
		new := num.Copy()

		if (&num) == (&new) {
			t.Errorf("Expected to be different, Result was incorrect, got: %t, want: %t.", (&num) == (&new), false)
		}

		if num.Cmp(new) != 0 {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", new, num)
		}
	}
}

func setBit(i uint8, b uint8, number Uint96, t *testing.T) {

	numBig := number.Big()
	number.SetBit(i, b)

	res := number.Big()
	numBigRes := numBig.SetBit(numBig, int(i), uint(b))

	if res.Cmp(numBigRes) != 0 {
		t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", res, numBigRes)
	}
}

func TestSetBitOne(t *testing.T) {
	for i := 0; i < 100; i++ {

		num := RandUInt96()
		setBit(95, 1, num, t)
		setBit(70, 1, num, t)
		setBit(64, 1, num, t)
		setBit(60, 1, num, t)
		setBit(32, 1, num, t)
		setBit(30, 1, num, t)
		setBit(20, 1, num, t)
		setBit(1, 1, num, t)
		setBit(0, 1, num, t)
	}
}

func TestSetBitZero(t *testing.T) {
	for i := 0; i < 100; i++ {

		num := RandUInt96()
		setBit(95, 0, num, t)
		setBit(70, 0, num, t)
		setBit(64, 0, num, t)
		setBit(60, 0, num, t)
		setBit(32, 0, num, t)
		setBit(30, 0, num, t)
		setBit(20, 0, num, t)
		setBit(1, 0, num, t)
		setBit(0, 0, num, t)
	}
}

func TestCountOnes(t *testing.T) {
	val := Uint96{
		Lo:  0x0000ffff,
		Mid: 0x00000000,
		Hi:  0x00000000,
	}

	if val.OnesCount() != 16 {
		t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", val.OnesCount(), 16)
	}

	val = Uint96{
		Lo:  0x0000ffff,
		Mid: 0x000ff000,
		Hi:  0x00000000,
	}

	if val.OnesCount() != 24 {
		t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", val.OnesCount(), 24)
	}

	val = Uint96{
		Lo:  0x0000ffff,
		Mid: 0x000ff000,
		Hi:  0x00fff00000,
	}

	if val.OnesCount() != (9 * 4) {
		t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", val.OnesCount(), (9 * 4))
	}
}

func TestCountZeros(t *testing.T) {
	val := Uint96{
		Lo:  0x0000ffff,
		Mid: 0x00000000,
		Hi:  0x00000000,
	}

	if val.ZerosCount() != (96 - 16) {
		t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", val.ZerosCount(), (96 - 16))
	}

	val = Uint96{
		Lo:  0x0000ffff,
		Mid: 0x000ff000,
		Hi:  0x00000000,
	}

	if val.ZerosCount() != (96 - 24) {
		t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", val.ZerosCount(), (96 - 24))
	}

	val = Uint96{
		Lo:  0x0000ffff,
		Mid: 0x000ff000,
		Hi:  0x00fff00000,
	}

	if val.ZerosCount() != (96 - (9 * 4)) {
		t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", val.ZerosCount(), (96 - (9 * 4)))
	}
}

func TestCountTrailingZeros(t *testing.T) {
	val := Uint96{
		Lo:  0x0000ffff,
		Mid: 0x00000000,
		Hi:  0x00000000,
	}

	if val.TrailingZeros() != 0 {
		t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", val.TrailingZeros(), 0)
	}

	val = Uint96{
		Lo:  0x0000fff0,
		Mid: 0x000ff000,
		Hi:  0x00000000,
	}

	if val.TrailingZeros() != 4 {
		t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", val.TrailingZeros(), 4)
	}

	val = Uint96{
		Lo:  0x00000000,
		Mid: 0x000ff000,
		Hi:  0x00000000,
	}

	if val.TrailingZeros() != 32+(4*3) {
		t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", val.TrailingZeros(), 32+(4*3))
	}

	val = Uint96{
		Lo:  0x00000000,
		Mid: 0x00000000,
		Hi:  0x00ff0000,
	}

	if val.TrailingZeros() != 64+(4*4) {
		t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", val.TrailingZeros(), 64+(4*4))
	}
}

func TestMultiply(t *testing.T) {
	for i := 0; i < 100; i++ {

		mask := Uint96{
			Lo:  0xffffffff,
			Mid: 0xffffffff,
			Hi:  0xffffffff,
		}

		num1 := RandUInt96()
		// num1 := uint96{
		// 	Lo:  0xffffffff,
		// 	Mid: 0xffffffff,
		// 	Hi:  0xffffffff,
		// }
		num1Big := num1.Big()

		num2 := RandUInt96()
		// num2 := uint96{
		// 	Lo:  0xffffffff,
		// 	Mid: 0xffffffff,
		// 	Hi:  0xffffffff,
		// }
		num2Big := num2.Big()

		// use new type to cal
		// value, carr := num1.Mul(num2)
		hi, lo := num1.Mul(num2)
		loBig := lo.Big()
		hiBig := hi.Big()

		// use big to cal
		resBig := num1Big.Mul(num1Big, num2Big)
		lowResBig := resBig.And(resBig, mask.Big())

		num1Big2 := num1.Big()
		resBig2 := num1Big2.Mul(num1Big2, num2Big)
		highResBig := resBig2.Rsh(resBig2, 96)

		if loBig.Cmp(lowResBig) != 0 {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", loBig, lowResBig)
		}

		if hiBig.Cmp(highResBig) != 0 {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", hiBig, highResBig)
		}
	}
}

func TestNot(t *testing.T) {
	for i := 0; i < 100; i++ {
		num := RandUInt96()
		not := num.Not()

		bigNum := num.Big()
		notex := FromBig(bigNum.Not(bigNum))

		if not != notex {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", not, notex)
		}
	}
}

func TestQuorem(t *testing.T) {
	for i := 0; i < 100; i++ {
		num := RandUInt96()
		numBig := num.Big()
		num2 := RandUInt96()
		num2Big := num2.Big()

		q, r := numBig.QuoRem(numBig, num2Big, new(big.Int))
		eq, er := FromBig(q), FromBig(r)
		o, or := num.QuoRem(num2)

		if o != eq {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", o, eq)
		}

		if or != er {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", or, er)
		}
	}
}

func TestGetBit(t *testing.T) {
	for i := 0; i < 100; i++ {
		num := RandUInt96()
		numBig := num.Big()
		bitIndex := uint8(mrand.Int32N(96))

		out := num.GetBit(bitIndex)
		want := numBig.Bit(int(bitIndex))

		if out != uint8(want) {
			t.Errorf("Expected to be equals, Result was incorrect, got: %b, want: %b.", out, want)
		}
	}
}

func TestSetBitIndexesLow(t *testing.T) {
	value := Uint96{
		Lo:  0xe0000ee0,
		Mid: 0,
		Hi:  0,
	}

	indexes := value.SetBitIndexes()
	compare := utilities.SliceCompare(indexes, []uint8{5, 6, 7, 9, 10, 11, 29, 30, 31})

	if !compare {
		t.Errorf("Expected to be equals")
	}
}

func TestSetBitIndexesMid(t *testing.T) {
	value := Uint96{
		Lo:  0xe0000ee0,
		Mid: 0xe0000ee0,
		Hi:  0,
	}

	indexes := value.SetBitIndexes()
	compare := utilities.SliceCompare(indexes, []uint8{
		5, 6, 7, 9, 10, 11, 29, 30, 31,
		37, 38, 39, 41, 42, 43, 61, 62, 63,
	})

	if !compare {
		t.Errorf("Expected to be equals")
	}
}

func TestSetBitIndexesHigh(t *testing.T) {
	value := Uint96{
		Lo:  0xe0000ee0,
		Mid: 0xe0000ee0,
		Hi:  0xe0000ee0,
	}

	indexes := value.SetBitIndexes()
	compare := utilities.SliceCompare(indexes, []uint8{
		5, 6, 7, 9, 10, 11, 29, 30, 31,
		37, 38, 39, 41, 42, 43, 61, 62, 63,
		69, 70, 71, 73, 74, 75, 93, 94, 95,
	})

	if !compare {
		t.Errorf("Expected to be equals")
	}
}
