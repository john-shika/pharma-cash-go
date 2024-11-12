package nokocore

import (
	"math/big"
	"math/rand"
)

// windows types like
// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types
// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#basic-types

type CHAR = int8
type SHORT = int16
type WCHAR = uint16
type USHORT = uint16
type WORD = uint16
type DWORD = uint
type DWORD32 = uint32
type DWORD64 = uint64
type LONG = int
type ULONG = uint
type LONG64 = int64
type ULONG64 = uint64
type FLOAT = float32
type DOUBLE = float64

type Ordering int

const (
	OrderingLess Ordering = iota - 1
	OrderingEqual
	OrderingGreater
)

type ComparableImpl[T any] interface {
	CompareTo(other T) Ordering
	Equals(other T) bool
}

type EquatableImpl[T any] interface {
	Equals(other T) bool
}

type StringableImpl interface {
	ToString() string
}

type CountableImpl interface {
	Len() int
}

type HashableImpl interface {
	HashCode() int
}

type JsonableImpl interface {
	ToJson() string
}

type NameableImpl interface {
	GetNameType() string
}

type BigFloatImpl interface {
	SetPrec(prec uint) *big.Float
	SetMode(mode big.RoundingMode) *big.Float
	Prec() uint
	MinPrec() uint
	Mode() big.RoundingMode
	Acc() big.Accuracy
	Sign() int
	MantExp(mant *big.Float) (exp int)
	SetMantExp(mant *big.Float, exp int) *big.Float
	Signbit() bool
	IsInf() bool
	IsInt() bool
	SetUint64(x uint64) *big.Float
	SetInt64(x int64) *big.Float
	SetFloat64(x float64) *big.Float
	SetInt(x *big.Int) *big.Float
	SetRat(x *big.Rat) *big.Float
	SetInf(signbit bool) *big.Float
	Set(x *big.Float) *big.Float
	Copy(x *big.Float) *big.Float
	Uint64() (uint64, big.Accuracy)
	Int64() (int64, big.Accuracy)
	Float32() (float32, big.Accuracy)
	Float64() (float64, big.Accuracy)
	Int(z *big.Int) (*big.Int, big.Accuracy)
	Rat(z *big.Rat) (*big.Rat, big.Accuracy)
	Abs(x *big.Float) *big.Float
	Neg(x *big.Float) *big.Float
	Add(x, y *big.Float) *big.Float
	Sub(x, y *big.Float) *big.Float
	Mul(x, y *big.Float) *big.Float
	Quo(x, y *big.Float) *big.Float
	Cmp(y *big.Float) int
}

type BigIntImpl interface {
	Sign() int
	SetInt64(x int64) *big.Int
	SetUint64(x uint64) *big.Int
	Set(x *big.Int) *big.Int
	Bits() []big.Word
	SetBits(abs []big.Word) *big.Int
	Abs(x *big.Int) *big.Int
	Neg(x *big.Int) *big.Int
	Add(x, y *big.Int) *big.Int
	Sub(x, y *big.Int) *big.Int
	Mul(x, y *big.Int) *big.Int
	MulRange(a, b int64) *big.Int
	Binomial(n, k int64) *big.Int
	Quo(x, y *big.Int) *big.Int
	Rem(x, y *big.Int) *big.Int
	QuoRem(x, y, r *big.Int) (*big.Int, *big.Int)
	Div(x, y *big.Int) *big.Int
	Mod(x, y *big.Int) *big.Int
	DivMod(x, y, m *big.Int) (*big.Int, *big.Int)
	Cmp(y *big.Int) (r int)
	CmpAbs(y *big.Int) int
	Int64() int64
	Uint64() uint64
	IsInt64() bool
	IsUint64() bool
	Float64() (float64, big.Accuracy)
	SetString(s string, base int) (*big.Int, bool)
	SetBytes(buf []byte) *big.Int
	Bytes() []byte
	FillBytes(buf []byte) []byte
	BitLen() int
	TrailingZeroBits() uint
	Exp(x, y, m *big.Int) *big.Int
	GCD(x, y, a, b *big.Int) *big.Int
	Rand(rnd *rand.Rand, n *big.Int) *big.Int
	ModInverse(g, n *big.Int) *big.Int
	ModSqrt(x, p *big.Int) *big.Int
	Lsh(x *big.Int, n uint) *big.Int
	Rsh(x *big.Int, n uint) *big.Int
	Bit(i int) uint
	SetBit(x *big.Int, i int, b uint) *big.Int
	And(x, y *big.Int) *big.Int
	AndNot(x, y *big.Int) *big.Int
	Or(x, y *big.Int) *big.Int
	Xor(x, y *big.Int) *big.Int
	Not(x *big.Int) *big.Int
	Sqrt(x *big.Int) *big.Int
}

type RandomImpl interface {
	Seed(seed int64)
	Int63() int64
	Uint32() uint32
	Uint64() uint64
	Int31() int32
	Int() int
	Int63n(n int64) int64
	Int31n(n int32) int32
	Intn(n int) int
	Float64() float64
	Float32() float32
	Perm(n int) []int
	Shuffle(n int, swap func(i, j int))
	Read(p []byte) (n int, err error)
}
