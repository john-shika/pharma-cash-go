package bigfloat

import (
	"fmt"
	"math/big"
	"strconv"
)

type Impl interface {
	Scan(s fmt.ScanState, ch rune) error
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(data []byte) (err error)
	MarshalText() (text []byte, err error)
	UnmarshalText(text []byte) (err error)
	MarshalJSON() (data []byte, err error)
	UnmarshalJSON(data []byte) (err error)
	String() string

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

type BigFloat struct {
	Value *big.Float
}

func New(value float64) *BigFloat {
	return &BigFloat{
		Value: new(big.Float).SetFloat64(value),
	}
}

func (b *BigFloat) Scan(s fmt.ScanState, ch rune) error {
	return b.Value.Scan(s, ch)
}

func (b *BigFloat) MarshalBinary() (data []byte, err error) {
	return b.Value.GobEncode()
}

func (b *BigFloat) UnmarshalBinary(data []byte) (err error) {
	return b.Value.GobDecode(data)
}

func (b *BigFloat) MarshalText() (text []byte, err error) {
	return b.Value.MarshalText()
}

func (b *BigFloat) UnmarshalText(text []byte) (err error) {
	return b.Value.UnmarshalText(text)
}

func (b *BigFloat) MarshalJSON() (data []byte, err error) {
	var temp []byte
	if temp, err = b.Value.MarshalText(); err != nil {
		return nil, err
	}
	return []byte(strconv.Quote(string(temp))), nil
}

func (b *BigFloat) UnmarshalJSON(data []byte) (err error) {
	var temp string
	if temp, err = strconv.Unquote(string(data)); err != nil {
		return err
	}
	return b.Value.UnmarshalText([]byte(temp))
}

func (b *BigFloat) String() string {
	return b.Value.String()
}

func (b *BigFloat) SetPrec(prec uint) *BigFloat {
	b.Value.SetPrec(prec)
	return b
}

func (b *BigFloat) SetMode(mode big.RoundingMode) *BigFloat {
	b.Value.SetMode(mode)
	return b
}

func (b *BigFloat) SetFloat64(value float64) *BigFloat {
	b.Value.SetFloat64(value)
	return b
}

func (b *BigFloat) SetInt64(value int64) *BigFloat {
	b.Value.SetInt64(value)
	return b
}

func (b *BigFloat) SetInt(value *big.Int) *BigFloat {
	b.Value.SetInt(value)
	return b
}

func (b *BigFloat) SetUint64(value uint64) *BigFloat {
	b.Value.SetUint64(value)
	return b
}

// TODO: continued, not implemented yet
