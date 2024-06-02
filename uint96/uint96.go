package uint96

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	"math/bits"
)

type Uint96 struct {
	Lo, Mid, Hi uint32
}

func New(lo, mid, hi uint32) Uint96 {
	return Uint96{lo, mid, hi}
}

func Randomise12Byte() *[]byte {
	randBuf := make([]byte, 12)
	rand.Read(randBuf)
	return &randBuf
}

func RandUInt96() Uint96 {
	randBuf := Randomise12Byte()
	return FromBytes(*randBuf)
}

func FromUInt32(value uint32) Uint96 {
	return Uint96{
		Lo:  value,
		Mid: 0,
		Hi:  0,
	}
}

func FromUInt64(value uint64) Uint96 {

	uint64LoMask := (uint64)(0x00000000ffffffff)
	uint64HiMask := (uint64)(0xffffffff00000000)

	low := (uint32)(uint64LoMask & value)
	high := (uint32)((uint64HiMask & value) >> 32)

	return Uint96{
		Lo:  low,
		Mid: high,
		Hi:  0,
	}
}

func FromBytes(b []byte) Uint96 {
	return New(
		binary.LittleEndian.Uint32(b[:4]),
		binary.LittleEndian.Uint32(b[4:8]),
		binary.LittleEndian.Uint32(b[8:12]),
	)
}

func FromBig(i *big.Int) (u Uint96) {

	mask := Uint96{
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
func (u Uint96) ToBytes() []byte {

	ret := make([]byte, 12)
	binary.LittleEndian.PutUint32(ret[:4], u.Lo)
	binary.LittleEndian.PutUint32(ret[4:8], u.Mid)
	binary.LittleEndian.PutUint32(ret[8:], u.Hi)

	return ret
}

// To bytes in big-first order
func (u Uint96) ToBytesBE() []byte {

	ret := make([]byte, 12)
	binary.BigEndian.PutUint32(ret[:4], u.Hi)
	binary.BigEndian.PutUint32(ret[4:8], u.Mid)
	binary.BigEndian.PutUint32(ret[8:], u.Lo)

	return ret
}

func (u Uint96) IsZero() bool {
	return u == Uint96{}
}

func (u Uint96) Equals(value Uint96) bool {
	return u == value
}

// Cmp compares u and v and returns:
//
//	-1 if u < v
//	 0 if u == v
//	+1 if u > v
func (u Uint96) Cmp(value Uint96) int {
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
func (u Uint96) And(value Uint96) Uint96 {
	return Uint96{
		Lo:  u.Lo & value.Lo,
		Mid: u.Mid & value.Mid,
		Hi:  u.Hi & value.Hi,
	}
}

// And returns u|v.
func (u Uint96) Or(value Uint96) Uint96 {
	return Uint96{
		Lo:  u.Lo | value.Lo,
		Mid: u.Mid | value.Mid,
		Hi:  u.Hi | value.Hi,
	}
}

// And returns u^v.
func (u Uint96) Xor(value Uint96) Uint96 {
	return Uint96{
		Lo:  u.Lo ^ value.Lo,
		Mid: u.Mid ^ value.Mid,
		Hi:  u.Hi ^ value.Hi,
	}
}

// Lsh returns u<<n.
func (u Uint96) Lsh(n uint) Uint96 {
	if n > 64 {

		return Uint96{
			Lo:  0,
			Mid: 0,
			Hi:  u.Lo << (n - 64),
		}
	}

	if n <= 64 && n > 32 {
		return Uint96{
			Lo:  0,
			Mid: u.Lo << (n - 32),
			Hi:  (u.Mid << (n - 32)) | (u.Lo >> (64 - n)),
		}
	}

	return Uint96{
		Lo:  u.Lo << n,
		Mid: (u.Mid << n) | (u.Lo >> (32 - n)),
		Hi:  (u.Hi << n) | (u.Mid >> (32 - n)),
	}
}

// Rsh returns u>>n.
func (u Uint96) Rsh(n uint) Uint96 {
	if n > 64 {

		return Uint96{
			Hi:  0,
			Mid: 0,
			Lo:  u.Hi >> (n - 64),
		}
	}

	if n <= 64 && n > 32 {
		return Uint96{
			Hi:  0,
			Mid: u.Hi >> (n - 32),
			Lo:  (u.Mid >> (n - 32)) | (u.Hi << (64 - n)),
		}
	}

	return Uint96{
		Hi:  u.Hi >> n,
		Mid: (u.Mid >> n) | (u.Hi << (32 - n)),
		Lo:  (u.Lo >> n) | (u.Mid << (32 - n)),
	}
}

func (u Uint96) Reverse() Uint96 {
	return Uint96{
		Lo:  bits.Reverse32(u.Hi),
		Mid: bits.Reverse32(u.Mid),
		Hi:  bits.Reverse32(u.Lo),
	}
}

func (u *Uint96) Copy() Uint96 {
	return Uint96{
		Lo:  u.Lo,
		Mid: u.Mid,
		Hi:  u.Hi,
	}
}

func (u *Uint96) SetBit(i uint8, b uint8) {

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

func (u *Uint96) GetBit(i uint8) uint8 {
	if i < 32 {
		return (uint8)(u.Lo>>i) & 1
	}

	if i >= 32 && i < 64 {
		return (uint8)(u.Mid>>(i-32)) & 1
	}

	return (uint8)(u.Hi>>(i-64)) & 1
}

// TrailingZeros returns the number of trailing zero bits in u; the result is
// 128 for u == 0.
func (u Uint96) TrailingZeros() int {
	if u.Lo > 0 {
		return bits.TrailingZeros32(u.Lo)
	}

	if u.Mid > 0 {
		return bits.TrailingZeros32(u.Mid) + 32
	}

	return bits.TrailingZeros32(u.Hi) + 64
}

func Lsh128(hi uint64, lo uint64, n uint) (rhi uint64, rlo uint64) {
	if n > 64 {
		rlo = 0
		rhi = lo << (n - 64)
	} else {
		rlo = lo << n
		rhi = hi<<n | lo>>(64-n)
	}
	return
}

func Rsh128(hi uint64, lo uint64, n uint) (rhi uint64, rlo uint64) {
	if n > 64 {
		rlo = hi >> (n - 64)
		rhi = 0
	} else {
		rlo = lo>>n | hi<<(64-n)
		rhi = hi >> n
	}
	return
}

func (u Uint96) QuoRem(value Uint96) (Uint96, Uint96) {

	u64Hi := uint64(u.Hi)
	u64Lo := ((uint64)(u.Mid) << 32) | uint64(u.Lo)

	v64Hi := uint64(value.Hi)
	v64Lo := ((uint64)(value.Mid) << 32) | uint64(value.Lo)

	if v64Hi == 0 {
		var res64Hi uint64
		var res64Lo uint64
		var r uint64

		if u64Hi < v64Lo {
			res64Lo, r = bits.Div64(u64Hi, u64Lo, v64Lo)
		} else {
			res64Hi, r = bits.Div64(0, u64Hi, v64Lo)
			res64Lo, r = bits.Div64(r, u64Lo, v64Lo)
		}

		res32Mid := (uint32)(((uint64)(0xffffffff00000000) & res64Lo) >> 32)
		res32Low := (uint32)((uint64)(0x00000000ffffffff) & res64Lo)
		res32Hi := (uint32)((uint64)(0x00000000ffffffff) & res64Hi)
		res := Uint96{
			Lo:  res32Low,
			Mid: res32Mid,
			Hi:  res32Hi,
		}

		r96 := FromUInt64(r)

		return res, r96
	}

	// generate a "trial quotient," guaranteed to be within 1 of the actual
	// quotient, then adjust.
	n := uint(bits.LeadingZeros64(v64Hi))
	v1Hi, _ := Lsh128(v64Hi, v64Lo, n)
	u1Hi, u1Lo := Rsh128(u64Hi, u64Lo, 1)
	tq, _ := bits.Div64(u1Hi, u1Lo, v1Hi)
	tq >>= 63 - n
	if tq != 0 {
		tq--
	}
	q := FromUInt64(tq)
	// calculate remainder using trial quotient, then adjust if remainder is
	// greater than divisor
	_, mulLo := value.Mul(q)
	r, _ := u.Sub(mulLo)
	if r.Cmp(value) >= 0 {
		q, _ = q.Add(Uint96{
			Lo:  1,
			Hi:  0,
			Mid: 0,
		})
		r, _ = r.Sub(value)
	}

	return q, r
}

func (u Uint96) To81Bitboard() Uint96 {
	return Uint96{
		Lo:  u.Lo,
		Mid: u.Mid,
		Hi:  u.Hi & (uint32)(0x0001ffff),
	}
}

func (u Uint96) Str() string {
	return fmt.Sprintf("%32b\n%32b\n%32b",
		u.Hi, u.Mid, u.Lo)
}

// OnesCount returns the number of one bits ("population count") in u.
func (u Uint96) OnesCount() int {
	return bits.OnesCount32(u.Hi) + bits.OnesCount32(u.Mid) + bits.OnesCount32(u.Lo)
}

// ZerosCount returns the number of one bits ("population count") in u.
func (u Uint96) ZerosCount() int {
	return 96 - u.OnesCount()
}

func (u Uint96) Add(value Uint96) (Uint96, uint32) {
	lo, locarry := bits.Add32(u.Lo, value.Lo, 0)
	mid, midcarry := bits.Add32(u.Mid, value.Mid, locarry)
	hi, hicarry := bits.Add32(u.Hi, value.Hi, midcarry)

	ret := Uint96{
		Lo:  lo,
		Mid: mid,
		Hi:  hi,
	}

	return ret, hicarry
}

// Sub returns u-v.
func (u Uint96) Sub(value Uint96) (Uint96, uint32) {
	lo, loborrow := bits.Sub32(u.Lo, value.Lo, 0)
	mid, midborrow := bits.Sub32(u.Mid, value.Mid, loborrow)
	hi, hiborrow := bits.Sub32(u.Hi, value.Hi, midborrow)

	ret := Uint96{
		Lo:  lo,
		Mid: mid,
		Hi:  hi,
	}

	return ret, hiborrow
}

func (u Uint96) Not() Uint96 {
	return Uint96{
		Lo:  ^u.Lo,
		Mid: ^u.Mid,
		Hi:  ^u.Hi,
	}
}

func (u Uint96) Mul(value Uint96) (Uint96, Uint96) {
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

	return Uint96{
			Lo:  loov,
			Mid: midov,
			Hi:  hiov,
		}, Uint96{
			Lo:  lo,
			Mid: mid,
			Hi:  hi,
		}
}

// convert to big
func (u Uint96) Big() *big.Int {
	i := new(big.Int).SetUint64(uint64(u.Hi))
	i = i.Lsh(i, 64)

	mid := new(big.Int).SetUint64(uint64(u.Mid))
	mid = mid.Lsh(mid, 32)
	i = i.Or(i, mid)

	i = i.Or(i, new(big.Int).SetUint64(uint64(u.Lo)))
	return i
}
