package nokocore

import (
	"errors"
	"math/big"
	"math/rand"
	"time"
)

var ErrValueIsEmpty = errors.New("value is empty")

func Abs(values ...int) (int, error) {
	size := len(values)
	if 0 < size {
		temp := 0
		for i, value := range values {
			KeepVoid(i)

			if value < 0 {
				temp -= value
			} else {
				temp += value
			}
		}
		return temp, nil
	}
	return Default[int](), ErrValueIsEmpty
}

func Abs32(values ...int32) (int32, error) {
	size := len(values)
	if 0 < size {
		temp := int32(0)
		for _, value := range values {
			if value < 0 {
				temp -= value
			} else {
				temp += value
			}
		}
		return temp, nil
	}
	return Default[int32](), ErrValueIsEmpty
}

func Abs64(values ...int64) (int64, error) {
	size := len(values)
	if 0 < size {
		temp := int64(0)
		for _, value := range values {
			if value < 0 {
				temp -= value
			} else {
				temp += value
			}
		}
		return temp, nil
	}
	return Default[int64](), ErrValueIsEmpty
}

func Min(values ...int) (int, error) {
	size := len(values)
	if 0 < size {
		minVal := values[0]
		for i := 1; i < size; i++ {
			if minVal > values[i] {
				minVal = values[i]
			}
		}
		return minVal, nil
	}
	return Default[int](), ErrValueIsEmpty
}

func Min32(values ...int32) (int32, error) {
	size := len(values)
	if 0 < size {
		minVal := values[0]
		for i := 1; i < size; i++ {
			if minVal > values[i] {
				minVal = values[i]
			}
		}
		return minVal, nil
	}
	return Default[int32](), ErrValueIsEmpty
}

func Min64(values ...int64) (int64, error) {
	size := len(values)
	if 0 < size {
		minVal := values[0]
		for i := 1; i < size; i++ {
			if minVal > values[i] {
				minVal = values[i]
			}
		}
		return minVal, nil
	}
	return Default[int64](), ErrValueIsEmpty
}

func Max(values ...int) (int, error) {
	size := len(values)
	if 0 < size {
		maxVal := values[0]
		for i := 1; i < size; i++ {
			if maxVal < values[i] {
				maxVal = values[i]
			}
		}
		return maxVal, nil
	}
	return Default[int](), ErrValueIsEmpty
}

func Max32(values ...int32) (int32, error) {
	size := len(values)
	if 0 < size {
		maxVal := values[0]
		for i := 1; i < size; i++ {
			if maxVal < values[i] {
				maxVal = values[i]
			}
		}
		return maxVal, nil
	}
	return Default[int32](), ErrValueIsEmpty
}

func Max64(values ...int64) (int64, error) {
	size := len(values)
	if 0 < size {
		maxVal := values[0]
		for i := 1; i < size; i++ {
			if maxVal < values[i] {
				maxVal = values[i]
			}
		}
		return maxVal, nil
	}
	return Default[int64](), ErrValueIsEmpty
}

type MinMaxResultImpl interface {
	Min() int
	Max() int
}

type MinMaxResult struct {
	min int
	max int
}

func NewMinMaxResult(min, max int) MinMaxResultImpl {
	return &MinMaxResult{
		min: min,
		max: max,
	}
}

func (m *MinMaxResult) Min() int {
	return m.min
}

func (m *MinMaxResult) Max() int {
	return m.max
}

func MinMax(values ...int) (MinMaxResultImpl, error) {
	size := len(values)
	if 0 < size {
		minVal := values[0]
		maxVal := values[0]
		for i := 1; i < size; i++ {
			if minVal > values[i] {
				minVal = values[i]
			}
			if maxVal < values[i] {
				maxVal = values[i]
			}
		}
		return NewMinMaxResult(minVal, maxVal), nil
	}
	return nil, ErrValueIsEmpty
}

type MinMaxResult32Impl interface {
	Min32() int32
	Max32() int32
}

type MinMaxResult32 struct {
	min int32
	max int32
}

func NewMinMaxResult32(min, max int32) MinMaxResult32Impl {
	return &MinMaxResult32{
		min: min,
		max: max,
	}
}

func (m *MinMaxResult32) Min32() int32 {
	return m.min
}

func (m *MinMaxResult32) Max32() int32 {
	return m.max
}

func MinMax32(values ...int32) (MinMaxResult32Impl, error) {
	size := len(values)
	if 0 < size {
		minVal := values[0]
		maxVal := values[0]
		for i := 1; i < size; i++ {
			if minVal > values[i] {
				minVal = values[i]
			}
			if maxVal < values[i] {
				maxVal = values[i]
			}
		}
		return NewMinMaxResult32(minVal, maxVal), nil
	}
	return nil, ErrValueIsEmpty
}

type MinMaxResult64Impl interface {
	Min64() int64
	Max64() int64
}

type MinMaxResult64 struct {
	min int64
	max int64
}

func NewMinMaxResult64(min, max int64) MinMaxResult64Impl {
	return &MinMaxResult64{
		min: min,
		max: max,
	}
}

func (m *MinMaxResult64) Min64() int64 {
	return m.min
}

func (m *MinMaxResult64) Max64() int64 {
	return m.max
}

func MinMax64(values ...int64) (MinMaxResult64Impl, error) {
	size := len(values)
	if 0 < size {
		minVal := values[0]
		maxVal := values[0]
		for i := 1; i < size; i++ {
			if minVal > values[i] {
				minVal = values[i]
			}
			if maxVal < values[i] {
				maxVal = values[i]
			}
		}
		return NewMinMaxResult64(minVal, maxVal), nil
	}
	return nil, ErrValueIsEmpty
}

func MinBigFloat(values ...*big.Float) (*big.Float, error) {
	size := len(values)
	if 0 < size {
		minVal := values[0]
		for i := 1; i < size; i++ {
			if minVal.Cmp(values[i]) > 0 {
				minVal = values[i]
			}
		}
		return minVal, nil
	}
	return nil, ErrValueIsEmpty
}

func MinBigInt(values ...BigIntImpl) (BigIntImpl, error) {
	size := len(values)
	if 0 < size {
		minVal := values[0]
		for i := 1; i < size; i++ {
			if minVal.Cmp(values[i].(*big.Int)) > 0 {
				minVal = values[i]
			}
		}
		return minVal, nil
	}
	return nil, ErrValueIsEmpty
}

func MaxBigFloat(values ...BigFloatImpl) (BigFloatImpl, error) {
	size := len(values)
	if 0 < size {
		maxVal := values[0]
		for i := 1; i < size; i++ {
			if maxVal.Cmp(values[i].(*big.Float)) < 0 {
				maxVal = values[i]
			}
		}
		return maxVal, nil
	}
	return nil, ErrValueIsEmpty
}

func MaxBigInt(values ...BigIntImpl) (BigIntImpl, error) {
	size := len(values)
	if 0 < size {
		maxVal := values[0]
		for i := 1; i < size; i++ {
			if maxVal.Cmp(values[i].(*big.Int)) < 0 {
				maxVal = values[i]
			}
		}
		return maxVal, nil
	}
	return nil, ErrValueIsEmpty
}

type MinMaxBigFloatResultImpl interface {
	Min() BigFloatImpl
	Max() BigFloatImpl
}

type MinMaxBigFloatResult struct {
	min BigFloatImpl
	max BigFloatImpl
}

func NewMinMaxBigFloatResult(min, max BigFloatImpl) MinMaxBigFloatResultImpl {
	return &MinMaxBigFloatResult{
		min: min,
		max: max,
	}
}

func (m *MinMaxBigFloatResult) Min() BigFloatImpl {
	return m.min
}

func (m *MinMaxBigFloatResult) Max() BigFloatImpl {
	return m.max
}

func MinMaxBigFloat(values ...BigFloatImpl) (MinMaxBigFloatResultImpl, error) {
	size := len(values)
	if 0 < size {
		minVal := values[0]
		maxVal := values[0]
		for i := 1; i < size; i++ {
			if minVal.Cmp(values[i].(*big.Float)) > 0 {
				minVal = values[i]
			}
			if maxVal.Cmp(values[i].(*big.Float)) < 0 {
				maxVal = values[i]
			}
		}
		return NewMinMaxBigFloatResult(minVal, maxVal), nil
	}
	return nil, ErrValueIsEmpty
}

type MinMaxBigIntResultImpl interface {
	Min() BigIntImpl
	Max() BigIntImpl
}

type MinMaxBigIntResult struct {
	min BigIntImpl
	max BigIntImpl
}

func NewMinMaxBigIntResult(min, max BigIntImpl) MinMaxBigIntResultImpl {
	return &MinMaxBigIntResult{
		min: min,
		max: max,
	}
}

func (m *MinMaxBigIntResult) Min() BigIntImpl {
	return m.min
}

func (m *MinMaxBigIntResult) Max() BigIntImpl {
	return m.max
}

func MinMaxBigInt(values ...BigIntImpl) (MinMaxBigIntResultImpl, error) {
	size := len(values)
	if 0 < size {
		minVal := values[0]
		maxVal := values[0]
		for i := 1; i < size; i++ {
			if minVal.Cmp(values[i].(*big.Int)) > 0 {
				minVal = values[i]
			}
			if maxVal.Cmp(values[i].(*big.Int)) < 0 {
				maxVal = values[i]
			}
		}
		return NewMinMaxBigIntResult(minVal, maxVal), nil
	}
	return nil, ErrValueIsEmpty
}

func NewRandom() RandomImpl {
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))
	return random
}

func RandomRangeInt(min, max int) int {
	n := max - min
	random := NewRandom()
	return min + random.Intn(n)
}

func RandomRangeInt32(min, max int32) int32 {
	n := max - min
	random := NewRandom()
	return min + random.Int31n(n)
}

func RandomRangeInt64(min, max int64) int64 {
	n := max - min
	random := NewRandom()
	return min + random.Int63n(n)
}
