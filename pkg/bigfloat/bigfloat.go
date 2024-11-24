package bigfloat

import (
	"database/sql/driver"
	"fmt"
	"math/big"
	"nokowebapi/nokocore"
	"strconv"
)

type Impl interface {
	Value() (driver.Value, error)
	Scan(s fmt.ScanState, ch rune) error
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(data []byte) (err error)
	MarshalText() (text []byte, err error)
	UnmarshalText(text []byte) (err error)
	MarshalJSON() (data []byte, err error)
	UnmarshalJSON(data []byte) (err error)
	String() string
	ToString() string

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

	Exp(z *big.Float) *big.Float
	Log(z *big.Float) *big.Float
	Pow(z *big.Float, w *big.Float) *big.Float
}

type BigFloat struct {
	*big.Float
}

func New(value float64) Impl {
	return &BigFloat{
		Float: new(big.Float).SetFloat64(value),
	}
}

func (b BigFloat) Value() (driver.Value, error) {
	return b.Float.MarshalText(), nil
}

func (b *BigFloat) Scan(s fmt.ScanState, ch rune) error {
	return b.Float.Scan(s, ch)
}

func (b BigFloat) MarshalBinary() (data []byte, err error) {
	return b.Float.GobEncode()
}

func (b *BigFloat) UnmarshalBinary(data []byte) (err error) {
	return b.Float.GobDecode(data)
}

func (b BigFloat) MarshalText() (text []byte, err error) {
	return b.Float.MarshalText()
}

func (b *BigFloat) UnmarshalText(text []byte) (err error) {
	return b.Float.UnmarshalText(text)
}

func (b BigFloat) MarshalJSON() (data []byte, err error) {
	var temp []byte
	nokocore.KeepVoid(temp)

	if temp, err = b.Float.MarshalText(); err != nil {
		return nil, err
	}
	return []byte(strconv.Quote(string(temp))), nil
}

func (b *BigFloat) UnmarshalJSON(data []byte) (err error) {
	var temp string
	nokocore.KeepVoid(temp)

	if temp, err = strconv.Unquote(string(data)); err != nil {
		return err
	}
	return b.Float.UnmarshalText([]byte(temp))
}

func (b BigFloat) String() string {
	return b.Float.String()
}

func (b *BigFloat) ToString() string {
	return b.Float.String()
}

func (b *BigFloat) SetPrec(prec uint) *big.Float {
	return b.SetPrec(prec)
}

func (b *BigFloat) SetMode(mode big.RoundingMode) *big.Float {
	return b.SetMode(mode)
}

func (b *BigFloat) Prec() uint {
	return b.Prec()
}

func (b *BigFloat) MinPrec() uint {
	return b.MinPrec()
}

func (b *BigFloat) Mode() big.RoundingMode {
	return b.Mode()
}

func (b *BigFloat) Acc() big.Accuracy {
	return b.Acc()
}

func (b *BigFloat) Sign() int {
	return b.Sign()
}

func (b *BigFloat) MantExp(mant *big.Float) (exp int) {
	return b.MantExp(mant)
}

func (b *BigFloat) SetMantExp(mant *big.Float, exp int) *big.Float {
	return b.SetMantExp(mant, exp)
}

func (b *BigFloat) Signbit() bool {
	return b.Signbit()
}

func (b *BigFloat) IsInf() bool {
	return b.IsInf()
}

func (b *BigFloat) IsInt() bool {
	return b.IsInt()
}

func (b *BigFloat) SetUint64(x uint64) *big.Float {
	return b.SetUint64(x)
}

func (b *BigFloat) SetInt64(x int64) *big.Float {
	return b.SetInt64(x)
}

func (b *BigFloat) SetFloat64(x float64) *big.Float {
	return b.SetFloat64(x)
}

func (b *BigFloat) SetInt(x *big.Int) *big.Float {
	return b.SetInt(x)
}

func (b *BigFloat) SetRat(x *big.Rat) *big.Float {
	return b.SetRat(x)
}

func (b *BigFloat) SetInf(signbit bool) *big.Float {
	return b.SetInf(signbit)
}

func (b *BigFloat) Set(x *big.Float) *big.Float {
	return b.Set(x)
}

func (b *BigFloat) Copy(x *big.Float) *big.Float {
	return b.Copy(x)
}

func (b *BigFloat) Uint64() (uint64, big.Accuracy) {
	return b.Uint64()
}

func (b *BigFloat) Int64() (int64, big.Accuracy) {
	return b.Int64()
}

func (b *BigFloat) Float32() (float32, big.Accuracy) {
	return b.Float32()
}

func (b *BigFloat) Float64() (float64, big.Accuracy) {
	return b.Float64()
}

func (b *BigFloat) Int(z *big.Int) (*big.Int, big.Accuracy) {
	return b.Int(z)
}

func (b *BigFloat) Rat(z *big.Rat) (*big.Rat, big.Accuracy) {
	return b.Rat(z)
}

func (b *BigFloat) Abs(x *big.Float) *big.Float {
	return b.Abs(x)
}

func (b *BigFloat) Neg(x *big.Float) *big.Float {
	return b.Neg(x)
}

func (b *BigFloat) Add(x, y *big.Float) *big.Float {
	return b.Add(x, y)
}

func (b *BigFloat) Sub(x, y *big.Float) *big.Float {
	return b.Sub(x, y)
}

func (b *BigFloat) Mul(x, y *big.Float) *big.Float {
	return b.Mul(x, y)
}

func (b *BigFloat) Quo(x, y *big.Float) *big.Float {
	return b.Quo(x, y)
}

func (b *BigFloat) Cmp(y *big.Float) int {
	return b.Cmp(y)
}

func (b *BigFloat) Exp(z *big.Float) *big.Float {
	return Exp(z)
}

func (b *BigFloat) Log(z *big.Float) *big.Float {
	return Log(z)
}

func (b *BigFloat) Pow(z *big.Float, w *big.Float) *big.Float {
	return Pow(z, w)
}
