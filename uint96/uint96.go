package uint96

import (
	"encoding/binary"
	"math/big"
	"math/bits"
)

type uint96 struct {
	Lo, Mid, Hi uint32
}

func New(lo, mid, hi uint32) uint96 {
	return uint96{lo, mid, hi}
}

func FromUInt32(value uint32) uint96 {
	return uint96{
		Lo:  value,
		Mid: 0,
		Hi:  0,
	}
}

func FromUInt64(value uint64) uint96 {

	uint64LoMask := (uint64)(0x00000000ffffffff)
	uint64HiMask := (uint64)(0xffffffff00000000)

	low := (uint32)(uint64LoMask & value)
	high := (uint32)((uint64HiMask & value) >> 32)

	return uint96{
		Lo:  low,
		Mid: high,
		Hi:  0,
	}
}

func FromBytes(b []byte) uint96 {
	return New(
		binary.LittleEndian.Uint32(b[:4]),
		binary.LittleEndian.Uint32(b[4:8]),
		binary.LittleEndian.Uint32(b[8:12]),
	)
}

func FromBig(i *big.Int) (u uint96) {

	mask := uint96{
		Lo:  0xffffffff,
		Mid: 0xffffffff,
		Hi:  0xffffffff,
	}

	i = i.And(i, mask.Big())

	ret := FromUInt64(i.Uint64())
	ret.Hi = uint32(i.Rsh(i, 64).Uint64())

	return ret
}

// To bytes in little-first order
func (u uint96) ToBytes() []byte {

	ret := make([]byte, 12)
	binary.LittleEndian.PutUint32(ret[:4], u.Lo)
	binary.LittleEndian.PutUint32(ret[4:8], u.Mid)
	binary.LittleEndian.PutUint32(ret[8:], u.Hi)

	return ret
}

// To bytes in big-first order
func (u uint96) ToBytesBE() []byte {

	ret := make([]byte, 12)
	binary.BigEndian.PutUint32(ret[:4], u.Hi)
	binary.BigEndian.PutUint32(ret[4:8], u.Mid)
	binary.BigEndian.PutUint32(ret[8:], u.Lo)

	return ret
}

func (u uint96) IsZero() bool {
	return u == uint96{}
}

func (u uint96) Equals(value uint96) bool {
	return u == value
}

// Cmp compares u and v and returns:
//
//	-1 if u < v
//	 0 if u == v
//	+1 if u > v
func (u uint96) Cmp(value uint96) int {
	if u.Equals(value) {
		return 0
	}

	if u.Hi < value.Hi {
		return -1
	}

	if u.Hi == value.Hi {

		if u.Mid < value.Mid {
			return -1
		}

		if u.Mid == value.Mid && u.Lo < value.Lo {
			return -1
		}
	}

	return 1
}

// And returns u&v.
func (u uint96) And(value uint96) uint96 {
	return uint96{
		Lo:  u.Lo & value.Lo,
		Mid: u.Mid & value.Mid,
		Hi:  u.Hi & value.Hi,
	}
}

// And returns u|v.
func (u uint96) Or(value uint96) uint96 {
	return uint96{
		Lo:  u.Lo | value.Lo,
		Mid: u.Mid | value.Mid,
		Hi:  u.Hi | value.Hi,
	}
}

// And returns u^v.
func (u uint96) Xor(value uint96) uint96 {
	return uint96{
		Lo:  u.Lo ^ value.Lo,
		Mid: u.Mid ^ value.Mid,
		Hi:  u.Hi ^ value.Hi,
	}
}

// Lsh returns u<<n.
func (u uint96) Lsh(n uint) uint96 {
	if n > 64 {

		return uint96{
			Lo:  0,
			Mid: 0,
			Hi:  u.Lo << (n - 64),
		}
	}

	if n <= 64 && n > 32 {
		return uint96{
			Lo:  0,
			Mid: u.Lo << (n - 32),
			Hi:  (u.Mid << (n - 32)) | (u.Lo >> (64 - n)),
		}
	}

	return uint96{
		Lo:  u.Lo << n,
		Mid: (u.Mid << n) | (u.Lo >> (32 - n)),
		Hi:  (u.Hi << n) | (u.Mid >> (32 - n)),
	}
}

// Rsh returns u>>n.
func (u uint96) Rsh(n uint) uint96 {
	if n > 64 {

		return uint96{
			Hi:  0,
			Mid: 0,
			Lo:  u.Hi >> (n - 64),
		}
	}

	if n <= 64 && n > 32 {
		return uint96{
			Hi:  0,
			Mid: u.Hi >> (n - 32),
			Lo:  (u.Mid >> (n - 32)) | (u.Hi << (64 - n)),
		}
	}

	return uint96{
		Hi:  u.Hi >> n,
		Mid: (u.Mid >> n) | (u.Hi << (32 - n)),
		Lo:  (u.Lo >> n) | (u.Mid << (32 - n)),
	}
}

func (u uint96) Reverse() uint96 {
	return uint96{
		Lo:  bits.Reverse32(u.Hi),
		Mid: bits.Reverse32(u.Mid),
		Hi:  bits.Reverse32(u.Lo),
	}
}

func (u uint96) Copy() uint96 {
	return uint96{
		Lo:  u.Lo,
		Mid: u.Mid,
		Hi:  u.Hi,
	}
}

func (u *uint96) SetBit(i uint, b uint) {

	if b == 1 {
		if i >= 64 {
			v := uint32(1) << (i - 64)
			u.Hi = u.Hi | v
		} else if i < 64 && i >= 32 {
			v := uint32(1) << (i - 32)
			u.Mid = u.Mid | v
		} else {
			v := uint32(1) << (i)
			u.Lo = u.Lo | v
		}
	} else if b == 0 {
		if i >= 64 {
			v := ((uint32(0xffffffff) << (i - 64 + 1)) | (uint32(1) << (i - 64))) - 1
			u.Hi = u.Hi & v
		} else if i < 64 && i >= 32 {
			v := ((uint32(0xffffffff) << (i - 32 + 1)) | (uint32(1) << (i - 32))) - 1
			u.Mid = u.Mid & v
		} else {
			v := ((uint32(0xffffffff) << (i + 1)) | (uint32(1) << (i))) - 1
			u.Lo = u.Lo & v
		}
	}
}

// TrailingZeros returns the number of trailing zero bits in u; the result is
// 128 for u == 0.
func (u uint96) TrailingZeros() int {
	if u.Lo > 0 {
		return bits.TrailingZeros32(u.Lo)
	}

	if u.Mid > 0 {
		return bits.TrailingZeros32(u.Mid) + 32
	}

	return bits.TrailingZeros32(u.Hi) + 64
}

// OnesCount returns the number of one bits ("population count") in u.
func (u uint96) OnesCount() int {
	return bits.OnesCount32(u.Hi) + bits.OnesCount32(u.Mid) + bits.OnesCount32(u.Lo)
}

// ZerosCount returns the number of one bits ("population count") in u.
func (u uint96) ZerosCount() int {
	return 96 - u.OnesCount()
}

func (u uint96) Add(value uint96) (uint96, uint32) {
	lo, locarry := bits.Add32(u.Lo, value.Lo, 0)
	mid, midcarry := bits.Add32(u.Mid, value.Mid, locarry)
	hi, hicarry := bits.Add32(u.Hi, value.Hi, midcarry)

	ret := uint96{
		Lo:  lo,
		Mid: mid,
		Hi:  hi,
	}

	return ret, hicarry
}

// Sub returns u-v.
func (u uint96) Sub(value uint96) (uint96, uint32) {
	lo, loborrow := bits.Sub32(u.Lo, value.Lo, 0)
	mid, midborrow := bits.Sub32(u.Mid, value.Mid, loborrow)
	hi, hiborrow := bits.Sub32(u.Hi, value.Hi, midborrow)

	ret := uint96{
		Lo:  lo,
		Mid: mid,
		Hi:  hi,
	}

	return ret, hiborrow
}

func (u uint96) Mul(value uint96) (uint96, uint96) {
	hi1, lo := bits.Mul32(u.Lo, value.Lo)

	// cal mid
	hi2, lo2 := bits.Mul32(u.Mid, value.Lo)
	hi3, lo3 := bits.Mul32(u.Lo, value.Mid)
	mid, midc0 := bits.Add32(hi1, lo2, 0)
	mid, midc1 := bits.Add32(mid, lo3, 0)

	// cal high
	hi4, lo4 := bits.Mul32(u.Mid, value.Mid)
	hi5, lo5 := bits.Mul32(u.Hi, value.Lo)
	hi6, lo6 := bits.Mul32(u.Lo, value.Hi)

	hi, hic0 := bits.Add32(midc0, midc1, 0)
	hi, hic1 := bits.Add32(hi, hi2, 0)
	hi, hic2 := bits.Add32(hi, hi3, 0)
	hi, hic3 := bits.Add32(hi, lo4, 0)
	hi, hic4 := bits.Add32(hi, lo5, 0)
	hi, hic5 := bits.Add32(hi, lo6, 0)

	// cal overflow low
	hi7, lo7 := bits.Mul32(u.Hi, value.Mid)
	hi8, lo8 := bits.Mul32(u.Mid, value.Hi)

	loov, loovc1 := bits.Add32(hic0, hic1, 0)
	loov, loovc2 := bits.Add32(loov, hic2, 0)
	loov, loovc3 := bits.Add32(loov, hic3, 0)
	loov, loovc4 := bits.Add32(loov, hic4, 0)
	loov, loovc5 := bits.Add32(loov, hic5, 0)
	loov, loovc6 := bits.Add32(loov, hi4, 0)
	loov, loovc7 := bits.Add32(loov, hi5, 0)
	loov, loovc8 := bits.Add32(loov, hi6, 0)
	loov, loovc9 := bits.Add32(loov, lo7, 0)
	loov, loovc10 := bits.Add32(loov, lo8, 0)

	hi9, lo9 := bits.Mul32(u.Hi, value.Hi)
	midov, midovc1 := bits.Add32(loovc1, loovc2, 0)
	midov, midovc2 := bits.Add32(midov, loovc3, 0)
	midov, midovc3 := bits.Add32(midov, loovc4, 0)
	midov, midovc4 := bits.Add32(midov, loovc5, 0)
	midov, midovc5 := bits.Add32(midov, loovc6, 0)
	midov, midovc6 := bits.Add32(midov, loovc7, 0)
	midov, midovc7 := bits.Add32(midov, loovc8, 0)
	midov, midovc8 := bits.Add32(midov, loovc9, 0)
	midov, midovc9 := bits.Add32(midov, loovc10, 0)
	midov, midovc10 := bits.Add32(midov, hi7, 0)
	midov, midovc11 := bits.Add32(midov, hi8, 0)
	midov, midovc12 := bits.Add32(midov, lo9, 0)

	hiov, _ := bits.Add32(midovc1, midovc2, 0)
	hiov, _ = bits.Add32(hiov, midovc3, 0)
	hiov, _ = bits.Add32(hiov, midovc4, 0)
	hiov, _ = bits.Add32(hiov, midovc5, 0)
	hiov, _ = bits.Add32(hiov, midovc6, 0)
	hiov, _ = bits.Add32(hiov, midovc7, 0)
	hiov, _ = bits.Add32(hiov, midovc8, 0)
	hiov, _ = bits.Add32(hiov, midovc9, 0)
	hiov, _ = bits.Add32(hiov, midovc10, 0)
	hiov, _ = bits.Add32(hiov, midovc11, 0)
	hiov, _ = bits.Add32(hiov, midovc12, 0)
	hiov, _ = bits.Add32(hiov, hi9, 0)

	return uint96{
			Lo:  lo,
			Mid: mid,
			Hi:  hi,
		}, uint96{
			Lo:  loov,
			Mid: midov,
			Hi:  hiov,
		}
}

// convert to big
func (u uint96) Big() *big.Int {
	i := new(big.Int).SetUint64(uint64(u.Hi))
	i = i.Lsh(i, 64)

	mid := new(big.Int).SetUint64(uint64(u.Mid))
	mid = mid.Lsh(mid, 32)
	i = i.Or(i, mid)

	i = i.Or(i, new(big.Int).SetUint64(uint64(u.Lo)))
	return i
}
