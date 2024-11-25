package nokocore

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"io"
	"strconv"
	"strings"
	"sync"
)

var ErrDataTypeInvalid = errors.New("invalid data type")

func KeepVoid(_ ...any) {
	// do nothing...
}

func Default[T any]() T {
	var temp T
	return temp
}

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Copy[T any](values []T) []T {
	size := len(values)
	temp := make([]T, size)
	copy(temp, values)
	return temp
}

func ToInt(value any) int64 {
	if value != nil {
		if val, ok := value.(IntConvertible); ok {
			return val.ToInt()
		}

		switch value.(type) {
		case bool:
			if value.(bool) {
				return 1
			}

			return 0

		case int, int8, int16, int32, int64:
			switch value.(type) {
			case int:
				return int64(value.(int))

			case int8:
				return int64(value.(int8))

			case int16:
				return int64(value.(int16))

			case int32:
				return int64(value.(int32))

			case int64:
				return value.(int64)

			default:
				return 0
			}

		case uint, uint8, uint16, uint32, uint64:
			switch value.(type) {
			case uint:
				return int64(value.(uint))

			case uint8:
				return int64(value.(uint8))

			case uint16:
				return int64(value.(uint16))

			case uint32:
				return int64(value.(uint32))

			case uint64:
				return int64(value.(uint64))

			default:
				return 0
			}

		case uintptr:
			return int64(value.(uintptr))

		case float32, float64:
			switch value.(type) {
			case float32:
				return int64(value.(float32))

			case float64:
				return int64(value.(float64))

			default:
				return 0
			}

		case complex64, complex128:
			switch value.(type) {
			case complex64:
				return int64(real(value.(complex64)))

			case complex128:
				return int64(real(value.(complex128)))

			default:
				return 0

			}

		case string:
			return Unwrap(strconv.ParseInt(value.(string), 10, 64))

		default:
			return ToIntReflect(value)
		}
	}

	return 0
}

func ToFloat(value any) float64 {
	if value != nil {
		if val, ok := value.(FloatConvertible); ok {
			return val.ToFloat()
		}

		switch value.(type) {
		case bool:
			if value.(bool) {
				return 1
			}

			return 0

		case int, int8, int16, int32, int64:
			switch value.(type) {
			case int:
				return float64(value.(int))

			case int8:
				return float64(value.(int8))

			case int16:
				return float64(value.(int16))

			case int32:
				return float64(value.(int32))

			case int64:
				return float64(value.(int64))

			default:
				return 0
			}

		case uint, uint8, uint16, uint32, uint64:
			switch value.(type) {
			case uint:
				return float64(value.(uint))

			case uint8:
				return float64(value.(uint8))

			case uint16:
				return float64(value.(uint16))

			case uint32:
				return float64(value.(uint32))

			case uint64:
				return float64(value.(uint64))

			default:
				return 0
			}

		case uintptr:
			return float64(value.(uintptr))

		case float32, float64:
			switch value.(type) {
			case float32:
				return float64(value.(float32))

			case float64:
				return value.(float64)

			default:
				return 0
			}

		case complex64, complex128:
			switch value.(type) {
			case complex64:
				return float64(real(value.(complex64)))

			case complex128:
				return real(value.(complex128))

			default:
				return 0

			}

		case string:
			return Unwrap(strconv.ParseFloat(value.(string), 64))

		default:
			return ToFloatReflect(value)
		}
	}

	return 0
}

func ToDecimal(value any) decimal.Decimal {
	if value != nil {
		if val, ok := value.(DecimalConvertible); ok {
			return val.ToDecimal()
		}

		switch value.(type) {
		case bool:
			if value.(bool) {
				return decimal.NewFromInt(1)
			}

			return decimal.NewFromInt(0)

		case int, int8, int16, int32, int64:
			switch value.(type) {
			case int:
				return decimal.NewFromInt(int64(value.(int)))

			case int8:
				return decimal.NewFromInt(int64(value.(int8)))

			case int16:
				return decimal.NewFromInt(int64(value.(int16)))

			case int32:
				return decimal.NewFromInt(int64(value.(int32)))

			case int64:
				return decimal.NewFromInt(value.(int64))

			default:
				return decimal.NewFromInt(0)
			}

		case uint, uint8, uint16, uint32, uint64:
			switch value.(type) {
			case uint:
				return decimal.NewFromUint64(uint64(value.(uint)))

			case uint8:
				return decimal.NewFromUint64(uint64(value.(uint8)))

			case uint16:
				return decimal.NewFromUint64(uint64(value.(uint16)))

			case uint32:
				return decimal.NewFromUint64(uint64(value.(uint32)))

			case uint64:
				return decimal.NewFromUint64(value.(uint64))

			default:
				return decimal.NewFromInt(0)
			}

		case uintptr:
			return decimal.NewFromUint64(uint64(value.(uintptr)))

		case float32, float64:
			switch value.(type) {
			case float32:
				return decimal.NewFromFloat32(value.(float32))

			case float64:
				return decimal.NewFromFloat(value.(float64))

			default:
				return decimal.NewFromInt(0)
			}

		case complex64, complex128:
			switch value.(type) {
			case complex64:
				return decimal.NewFromFloat32(real(value.(complex64)))

			case complex128:
				return decimal.NewFromFloat(real(value.(complex128)))

			default:
				return decimal.NewFromInt(0)
			}

		case string:
			return Unwrap(decimal.NewFromString(value.(string)))

		default:
			return ToDecimalReflect(value)
		}
	}

	return decimal.NewFromInt(0)
}

func ToString(value any) string {
	if value != nil {
		if val, ok := value.(StringConvertible); ok {
			return val.ToString()
		}

		switch value.(type) {
		case bool:
			return fmt.Sprintf("%t", value)

		case int, int8, int16, int32, int64:
			return fmt.Sprintf("%d", value)

		case uint, uint8, uint16, uint32, uint64:
			return fmt.Sprintf("%d", value)

		case uintptr:
			return fmt.Sprintf("0x%016x", value)

		case float32, float64:
			return fmt.Sprintf("%f", value)

		case complex64, complex128:
			return fmt.Sprintf("%f", value)

		case string:
			return value.(string)

		default:
			return ToStringReflect(value)
		}
	}

	return "<nil>"
}

type ErrOrOkImpl interface {
	bool | any
}

func IsOk[T ErrOrOkImpl](eOk T) bool {
	var ok bool
	var bOk bool
	KeepVoid(ok, bOk)

	if bOk, ok = CastBool(eOk); ok {
		return bOk
	}

	return false
}

func IsErr[T ErrOrOkImpl](eOk T) bool {
	var ok bool
	var err error
	KeepVoid(ok, err)

	if err, ok = CastErr(eOk); ok {
		return err != nil
	}

	return false
}

func CastErr[T ErrOrOkImpl](eOk T) (error, bool) {
	var ok bool
	var err error
	KeepVoid(ok, err)

	if IsNone(eOk) {
		return nil, true
	}

	if err, ok = Cast[error](eOk); ok {
		return err, true
	}

	return nil, false
}

func IsValid(value any) bool {
	// determine if the given value is valid.
	// this method uses reflection to find out the validity of the value.
	// a value is considered valid if it is non-nil and has a non-zero value.
	return IsValidReflect(value)
}

func IsNone(value any) bool {
	// check if the given value is neither valid nor has a non-zero value.
	// this method leverages reflection to determine if the value is nil or a zero value.
	return !IsValid(value)
}

func IsNoneOrEmpty(value any) bool {
	if IsValid(value) {
		val := ToString(value)
		switch val {
		case "<nil>":
			return true

		case "<array>":
			return GetSizeReflect(value) == 0

		case "<object>":
			// TODO: not implemented yet
			return false

		default:
			return val == ""
		}
	}

	return true
}

func IsNoneOrEmptyWhiteSpace(value any) bool {
	if IsNoneOrEmpty(value) {
		return true
	}
	temp := strings.TrimSpace(ToString(value))
	if temp != "" {
		for i, character := range temp {
			KeepVoid(i)

			// maybe some character is not white space
			if !strings.Contains(WhiteSpace, string(character)) {
				return false
			}
		}
	}
	return true
}

func Unwrap[T any, E ErrOrOkImpl](result T, eOk E) T {
	var ok bool
	var err error
	KeepVoid(ok, err)

	if IsNone(eOk) {
		return result
	}

	if err, ok = CastErr(eOk); !ok {
		if !IsOk(eOk) {
			panic("invalid data type")
		}

		return result
	}

	NoErr(err)
	return result
}

func Unwrap2[T1, T2 any, E ErrOrOkImpl](value1 T1, value2 T2, eOk E) (T1, T2) {
	var ok bool
	var err error
	KeepVoid(ok, err)

	if IsNone(eOk) {
		return value1, value2
	}

	if err, ok = CastErr(eOk); !ok {
		if !IsOk(eOk) {
			panic("invalid data type")
		}

		return value1, value2
	}

	NoErr(err)

	return value1, value2
}

func Unwrap3[T1, T2, T3 any, E ErrOrOkImpl](value1 T1, value2 T2, value3 T3, eOk E) (T1, T2, T3) {
	var ok bool
	var err error
	KeepVoid(ok, err)

	if IsNone(eOk) {
		return value1, value2, value3
	}

	if err, ok = CastErr(eOk); !ok {
		if !IsOk(eOk) {
			panic("invalid data type")
		}

		return value1, value2, value3
	}

	NoErr(err)

	return value1, value2, value3
}

func HandlePanic(handler func(error)) {
	if err := recover(); err != nil {
		switch val := err.(type) {
		case error:
			fmt.Printf("[RECOVERY] Panic: %s\n", val.Error())
			if handler != nil {
				handler(val)
			}
		case string:
			fmt.Printf("[RECOVERY] Panic: %s\n", val)
			if handler != nil {
				handler(errors.New(val))
			}
		default:
			fmt.Println("[RECOVERY] Panic.")
			if handler != nil {
				handler(errors.New("recovery"))
			}
		}
	}
}

type MapImpl[T any] interface {
	Len() int
	Keys() []string
	Values() []T
	HasKey(key string) bool
	ContainKeys(keys ...string) bool
	GetVal(key string) T
	Get(key string) any
	SetVal(key string, value T) bool
	Set(key string, value any) bool
	DelVal(key string) bool
	Del(key string) bool
}

type MapAnyImpl interface {
	MapImpl[any]
}

type Map[T any] map[string]T
type MapAny = Map[any]

func NewMap[T any]() MapImpl[T] {
	return make(Map[T])
}

func NewMapAny() MapAnyImpl {
	return NewMap[any]()
}

func (m Map[T]) GetNameType() string {
	return "Map"
}

func (m Map[T]) Len() int {
	return len(m)
}

func (m Map[T]) Keys() []string {
	size := len(m)
	keys := make([]string, 0, size)
	for k, v := range m {
		KeepVoid(k, v)

		keys = append(keys, k)
	}

	return keys
}

func (m Map[T]) Values() []T {
	size := len(m)
	values := make([]T, 0, size)
	for k, v := range m {
		KeepVoid(k, v)

		values = append(values, v)
	}

	return values
}

func (m Map[T]) HasKey(key string) bool {
	var ok bool
	var value T
	KeepVoid(ok, value)
	if value, ok = m[key]; !ok {
		return false
	}
	return true
}

func (m Map[T]) ContainKeys(keys ...string) bool {
	for i, key := range keys {
		KeepVoid(i, key)

		if !m.HasKey(key) {
			return false
		}
	}
	return true
}

func (m Map[T]) GetVal(key string) T {
	return GetMapValue(m, key)
}

func (m Map[T]) Get(key string) any {
	return GetMapValueWithSuperKeyReflect(m, key)
}

func (m Map[T]) SetVal(key string, value T) bool {
	return SetMapValue(m, key, value)
}

func (m Map[T]) Set(key string, value any) bool {
	return SetMapValueWithSuperKeyReflect(m, key, value)
}

func (m Map[T]) DelVal(key string) bool {
	return DeleteMapValue(m, key)
}

func (m Map[T]) Del(key string) bool {
	return DeleteMapValueWithSuperKeyReflect(m, key)
}

// TODO: not implemented yet

type ArrayImpl[T any] interface {
	At(index int) T
	Set(index int, value T) bool
	Del(index int) bool
	Concat(arrays ...ArrayImpl[T]) ArrayImpl[T]
	Append(values ...T) ArrayImpl[T]
	Len() int
}

type Array[T any] []T
type ArrayAny = Array[any]
type ArrayStr = Array[string]

func NewArray[T any](size int) ArrayImpl[T] {
	return make(Array[T], size)
}

func (a Array[T]) GetNameType() string {
	return "Array"
}

func (a Array[T]) At(index int) T {
	return a[index]
}

func (a Array[T]) Set(index int, value T) bool {
	a[index] = value
	return true
}

func (a Array[T]) Del(index int) bool {
	fmt.Println("[WARN] Array element deletion unsupported. Use a DeleteArrayItem instead.")
	KeepVoid(index)
	return false
}

func (a Array[T]) Concat(arrays ...ArrayImpl[T]) ArrayImpl[T] {
	var temp Array[T]
	KeepVoid(temp)

	copy(temp, a)
	for i, array := range arrays {
		KeepVoid(i)

		var ok bool
		var slice []T
		KeepVoid(ok, slice)

		if slice, ok = Cast[[]T](array); !ok {
			for j := 0; j < array.Len(); j++ {
				temp = append(temp, array.At(j))
			}
			continue
		}

		temp = append(temp, slice...)
	}

	return temp
}

func (a Array[T]) Append(values ...T) ArrayImpl[T] {
	temp := append(a, values...)
	return temp
}

func (a Array[T]) Len() int {
	return len(a)
}

func Cast[T any](value any) (T, bool) {
	temp, ok := value.(T)
	return temp, ok
}

func CastAny(value any) (any, bool) {
	return value, true
}

func CastBool(value any) (bool, bool) {
	return Cast[bool](value)
}

func CastInt8(value any) (int8, bool) {
	return Cast[int8](value)
}

func CastUint8(value any) (uint8, bool) {
	return Cast[uint8](value)
}

func CastInt16(value any) (int16, bool) {
	return Cast[int16](value)
}

func CastUint16(value any) (uint16, bool) {
	return Cast[uint16](value)
}

func CastInt(value any) (int, bool) {
	return Cast[int](value)
}

func CastUint(value any) (uint, bool) {
	return Cast[uint](value)
}

func CastInt32(value any) (int32, bool) {
	return Cast[int32](value)
}

func CastUint32(value any) (uint32, bool) {
	return Cast[uint32](value)
}

func CastInt64(value any) (int64, bool) {
	return Cast[int64](value)
}

func CastUint64(value any) (uint64, bool) {
	return Cast[uint64](value)
}

func CastFloat32(value any) (float32, bool) {
	return Cast[float32](value)
}

func CastFloat64(value any) (float64, bool) {
	return Cast[float64](value)
}

func CastString(value any) (string, bool) {
	return Cast[string](value)
}

func CastPtr[T any](value any) (*T, bool) {
	temp, ok := value.(*T)
	return temp, ok
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64EncodeURLSafe(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func Base64DecodeURLSafe(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}

func NewUUID() uuid.UUID {
	return Unwrap(uuid.NewV7())
}

// TODO: not implemented yet

type ActionBindingImpl[T any] interface {
	Func() func()
	Bind(this T) ActionBindingImpl[T]
	Apply(this T)
	Call()
}

type ActionSingleParamBindingImpl[T, P any] interface {
	Func() func(arg P)
	Bind(this T) ActionSingleParamBindingImpl[T, P]
	Apply(this T, arg P)
	Call(arg P)
}

type ActionAnyParamsBindingImpl[T any] interface {
	Func() func(args ...any)
	Bind(this T) ActionAnyParamsBindingImpl[T]
	Apply(this T, args ...any)
	Call(args ...any)
}

type ActionParamsBindingImpl[T, P any] interface {
	Func() func(args ...P)
	Bind(this T) ActionParamsBindingImpl[T, P]
	Apply(this T, args ...P)
	Call(args ...P)
}

type ActionReturnBindingImpl[T, R any] interface {
	Func() func() R
	Bind(this T) ActionReturnBindingImpl[T, R]
	Apply(this T) R
	Call() R
}

type ActionSingleParamReturnBindingImpl[T, V, R any] interface {
	Func() func(arg V) R
	Bind(this T) ActionSingleParamReturnBindingImpl[T, V, R]
	Apply(this T, arg V) R
	Call(arg V) R
}

type ActionAnyParamsReturnBindingImpl[T, R any] interface {
	Func() func(args ...any) R
	Bind(this T) ActionAnyParamsReturnBindingImpl[T, R]
	Apply(this T, args ...any) R
	Call(args ...any) R
}

type ActionParamsReturnBindingImpl[T, V, R any] interface {
	Func() func(args ...V) R
	Bind(this T) ActionParamsReturnBindingImpl[T, V, R]
	Apply(this T, args ...V) R
	Call(args ...V) R
}

// TODO: not implemented yet

type ActionImpl interface {
	Func() func()
	Call()
}

type Action func()

func NewAction(action func()) ActionImpl {
	return Action(action)
}

func (a Action) Func() func() {
	return func() {
		a.Call()
	}
}

func (a Action) Call() {
	a()
}

type LockerImpl interface {
	Lock(action Action)
	TryLock(action Action) bool
	IsLocked() bool
}

type Locker struct {
	mutex  *sync.Mutex
	locked bool
}

func NewLocker() LockerImpl {
	return &Locker{
		mutex:  &sync.Mutex{},
		locked: false,
	}
}

func (lock *Locker) Release() {
	if !lock.locked {
		return
	}
	lock.mutex.Unlock()
	lock.locked = false
}

func (lock *Locker) Lock(action Action) {
	defer lock.Release()
	lock.mutex.Lock()
	lock.locked = true
	action.Call()
}

func (lock *Locker) TryLock(action Action) bool {
	if !lock.mutex.TryLock() {
		return false
	}
	defer lock.Release()
	lock.locked = true
	action.Call()
	return true
}

func (lock *Locker) IsLocked() bool {
	return lock.locked
}

type StateImpl[T any] interface {
	Get() T
	Set(value T)
}

type State[T any] struct {
	mutex *sync.RWMutex
	data  T
}

func NewState[T any](value T) StateImpl[T] {
	return &State[T]{
		mutex: &sync.RWMutex{},
		data:  value,
	}
}

func (state *State[T]) Get() T {
	defer state.mutex.RUnlock()
	state.mutex.RLock()
	return state.data
}

func (state *State[T]) Set(value T) {
	defer state.mutex.Unlock()
	state.mutex.Lock()
	state.data = value
}

type StateActionImpl interface {
	Set(action Action)
	Func() func()
	Call()
}

type StateAction struct {
	state StateImpl[Action]
}

func NewStateAction(action Action) StateActionImpl {
	return &StateAction{
		state: NewState(action),
	}
}

func (a *StateAction) Set(action Action) {
	a.state.Set(action)
}

func (a *StateAction) Func() func() {
	return func() {
		a.Call()
	}
}

func (a *StateAction) Call() {
	a.state.Get().Call()
}

// single params

type ActionSingleParamImpl[P any] interface {
	Func() func(arg P)
	Call(arg P)
}

type ActionSingleParam[P any] func(arg P)

func NewActionSingleParam[P any](action func(arg P)) ActionSingleParamImpl[P] {
	return ActionSingleParam[P](action)
}

func (a ActionSingleParam[P]) Func() func(arg P) {
	return func(arg P) {
		a.Call(arg)
	}
}

func (a ActionSingleParam[P]) Call(arg P) {
	a(arg)
}

type StateActionSingleParamImpl[P any] interface {
	Set(action ActionSingleParam[P])
	Func() func(arg P)
	Call(arg P)
}

type StateActionSingleParam[P any] struct {
	state StateImpl[ActionSingleParam[P]]
}

func NewStateActionSingleParam[P any](action ActionSingleParam[P]) StateActionSingleParamImpl[P] {
	return &StateActionSingleParam[P]{
		state: NewState(action),
	}
}

func (a *StateActionSingleParam[P]) Set(action ActionSingleParam[P]) {
	a.state.Set(action)
}

func (a *StateActionSingleParam[P]) Func() func(arg P) {
	return func(arg P) {
		a.Call(arg)
	}
}

func (a *StateActionSingleParam[P]) Call(arg P) {
	a.state.Get().Call(arg)
}

// double params

type ActionDoubleParamsImpl[P1, P2 any] interface {
	Func() func(p1 P1, p2 P2)
	Call(p1 P1, p2 P2)
}

type ActionDoubleParams[P1, P2 any] func(p1 P1, p2 P2)

func NewActionDoubleParams[P1, P2 any](action func(p1 P1, p2 P2)) ActionDoubleParamsImpl[P1, P2] {
	return ActionDoubleParams[P1, P2](action)
}

func (a ActionDoubleParams[P1, P2]) Func() func(p1 P1, p2 P2) {
	return func(p1 P1, p2 P2) {
		a.Call(p1, p2)
	}
}

func (a ActionDoubleParams[P1, P2]) Call(p1 P1, p2 P2) {
	a(p1, p2)
}

type StateActionDoubleParamsImpl[P1, P2 any] interface {
	Set(action ActionDoubleParams[P1, P2])
	Func() func(p1 P1, p2 P2)
	Call(p1 P1, p2 P2)
}

type StateActionDoubleParams[P1, P2 any] struct {
	state StateImpl[ActionDoubleParams[P1, P2]]
}

func NewStateActionDoubleParams[P1, P2 any](action ActionDoubleParams[P1, P2]) StateActionDoubleParamsImpl[P1, P2] {
	return &StateActionDoubleParams[P1, P2]{
		state: NewState(action),
	}
}

func (a *StateActionDoubleParams[P1, P2]) Set(action ActionDoubleParams[P1, P2]) {
	a.state.Set(action)
}

func (a *StateActionDoubleParams[P1, P2]) Func() func(p1 P1, p2 P2) {
	return func(p1 P1, p2 P2) {
		a.Call(p1, p2)
	}
}

func (a *StateActionDoubleParams[P1, P2]) Call(p1 P1, p2 P2) {
	a.state.Get().Call(p1, p2)
}

type ActionAnyParamsImpl interface {
	Func() func(args ...any)
	Call(args ...any)
}

type ActionAnyParams func(args ...any)

func NewActionAnyParams(action func(args ...any)) ActionAnyParamsImpl {
	return ActionAnyParams(action)
}

func (a ActionAnyParams) Func() func(args ...any) {
	return func(args ...any) {
		a.Call(args...)
	}
}

func (a ActionAnyParams) Call(args ...any) {
	a(args...)
}

type StateActionAnyParamsImpl interface {
	Set(action ActionAnyParams)
	Func() func(args ...any)
	Call(args ...any)
}

type StateActionAnyParams struct {
	state StateImpl[ActionAnyParams]
}

func NewStateActionAnyParams(action ActionAnyParams) StateActionAnyParamsImpl {
	return &StateActionAnyParams{
		state: NewState(action),
	}
}

func (a *StateActionAnyParams) Set(action ActionAnyParams) {
	a.state.Set(action)
}

func (a *StateActionAnyParams) Func() func(args ...any) {
	return func(args ...any) {
		a.Call(args...)
	}
}

func (a *StateActionAnyParams) Call(args ...any) {
	a.state.Get().Call(args...)
}

type ActionParamsImpl[P any] interface {
	Func() func(args ...P)
	Call(args ...P)
}

type ActionParams[P any] func(args ...P)

func NewActionParams[P any](action func(args ...P)) ActionParamsImpl[P] {
	return ActionParams[P](action)
}

func (a ActionParams[P]) Func() func(args ...P) {
	return func(args ...P) {
		a.Call(args...)
	}
}

func (a ActionParams[params]) Call(args ...params) {
	a(args...)
}

type StateActionParamsImpl[P any] interface {
	Set(action ActionParams[P])
	Func() func(args ...P)
	Call(args ...P)
}

type StateActionParams[P any] struct {
	state StateImpl[ActionParams[P]]
}

func NewStateActionParams[P any](action ActionParams[P]) StateActionParamsImpl[P] {
	return &StateActionParams[P]{
		state: NewState(action),
	}
}

func (a *StateActionParams[P]) Set(action ActionParams[P]) {
	a.state.Set(action)
}

func (a *StateActionParams[P]) Func() func(args ...P) {
	return func(args ...P) {
		a.Call(args...)
	}
}

func (a *StateActionParams[P]) Call(args ...P) {
	a.state.Get().Call(args...)
}

type ActionReturnImpl[R any] interface {
	Func() func() R
	Call() R
}

type ActionReturn[R any] func() R

func NewActionReturn[R any](action func() R) ActionReturnImpl[R] {
	return ActionReturn[R](action)
}

func (a ActionReturn[R]) Func() func() R {
	return func() R {
		return a.Call()
	}
}

func (a ActionReturn[R]) Call() R {
	return a()
}

type StateActionReturnImpl[R any] interface {
	Set(action ActionReturn[R])
	Func() func() R
	Call() R
}

type StateActionReturn[R any] struct {
	state StateImpl[ActionReturn[R]]
}

func NewStateActionReturn[R any](action ActionReturn[R]) StateActionReturnImpl[R] {
	return &StateActionReturn[R]{
		state: NewState(action),
	}
}

func (a *StateActionReturn[R]) Set(action ActionReturn[R]) {
	a.state.Set(action)
}

func (a *StateActionReturn[R]) Func() func() R {
	return func() R {
		return a.Call()
	}
}

func (a *StateActionReturn[R]) Call() R {
	return a.state.Get().Call()
}

// single params return

type ActionSingleParamReturnImpl[V, R any] interface {
	Func() func(arg V) R
	Call(arg V) R
}

type ActionSingleParamReturn[V, R any] func(arg V) R

func NewActionSingleParamReturn[V, R any](action func(arg V) R) ActionSingleParamReturnImpl[V, R] {
	return ActionSingleParamReturn[V, R](action)
}

func (a ActionSingleParamReturn[V, R]) Func() func(arg V) R {
	return func(arg V) R {
		return a.Call(arg)
	}
}

func (a ActionSingleParamReturn[V, R]) Call(arg V) R {
	return a(arg)
}

type StateActionSingleParamReturnImpl[V, R any] interface {
	Set(action ActionSingleParamReturn[V, R])
	Func() func(arg V) R
	Call(arg V) R
}

type StateActionSingleParamReturn[V, R any] struct {
	state StateImpl[ActionSingleParamReturn[V, R]]
}

func NewStateActionSingleParamReturn[V, R any](action ActionSingleParamReturn[V, R]) StateActionSingleParamReturnImpl[V, R] {
	return &StateActionSingleParamReturn[V, R]{
		state: NewState(action),
	}
}

func (a *StateActionSingleParamReturn[V, R]) Set(action ActionSingleParamReturn[V, R]) {
	a.state.Set(action)
}

func (a *StateActionSingleParamReturn[V, R]) Func() func(arg V) R {
	return func(arg V) R {
		return a.Call(arg)
	}
}

func (a *StateActionSingleParamReturn[V, R]) Call(arg V) R {
	return a.state.Get().Call(arg)
}

// double params return

type ActionDoubleParamsReturnImpl[V1, V2, R any] interface {
	Func() func(v1 V1, v2 V2) R
	Call(v1 V1, v2 V2) R
}

type ActionDoubleParamsReturn[V1, V2, R any] func(v1 V1, v2 V2) R

func NewActionDoubleParamsReturn[V1, V2, R any](action func(v1 V1, v2 V2) R) ActionDoubleParamsReturnImpl[V1, V2, R] {
	return ActionDoubleParamsReturn[V1, V2, R](action)
}

func (a ActionDoubleParamsReturn[V1, V2, R]) Func() func(v1 V1, v2 V2) R {
	return func(v1 V1, v2 V2) R {
		return a.Call(v1, v2)
	}
}

func (a ActionDoubleParamsReturn[V1, V2, R]) Call(v1 V1, v2 V2) R {
	return a(v1, v2)
}

type StateActionDoubleParamsReturnImpl[V1, V2, R any] interface {
	Set(action ActionDoubleParamsReturn[V1, V2, R])
	Func() func(v1 V1, v2 V2) R
	Call(v1 V1, v2 V2) R
}

type StateActionDoubleParamsReturn[V1, V2, R any] struct {
	state StateImpl[ActionDoubleParamsReturn[V1, V2, R]]
}

func NewStateActionDoubleParamsReturn[V1, V2, R any](action ActionDoubleParamsReturn[V1, V2, R]) StateActionDoubleParamsReturnImpl[V1, V2, R] {
	return &StateActionDoubleParamsReturn[V1, V2, R]{
		state: NewState(action),
	}
}

func (a *StateActionDoubleParamsReturn[V1, V2, R]) Set(action ActionDoubleParamsReturn[V1, V2, R]) {
	a.state.Set(action)
}

func (a *StateActionDoubleParamsReturn[V1, V2, R]) Func() func(v1 V1, v2 V2) R {
	return func(v1 V1, v2 V2) R {
		return a.Call(v1, v2)
	}
}

func (a *StateActionDoubleParamsReturn[V1, V2, R]) Call(v1 V1, v2 V2) R {
	return a.state.Get().Call(v1, v2)
}

type ActionAnyParamsReturnImpl[R any] interface {
	Func() func(args ...any) R
	Call(args ...any) R
}

type ActionAnyParamsReturn[R any] func(args ...any) R

func NewActionAnyParamsReturn[R any](action func(args ...any) R) ActionAnyParamsReturnImpl[R] {
	return ActionAnyParamsReturn[R](action)
}

func (a ActionAnyParamsReturn[R]) Func() func(args ...any) R {
	return func(args ...any) R {
		return a.Call(args...)
	}
}

func (a ActionAnyParamsReturn[R]) Call(args ...any) R {
	return a(args...)
}

type StateActionAnyParamsReturnImpl[R any] interface {
	Set(action ActionAnyParamsReturn[R])
	Func() func(args ...any) R
	Call(args ...any) R
}

type StateActionAnyParamsReturn[R any] struct {
	state StateImpl[ActionAnyParamsReturn[R]]
}

func NewStateActionAnyParamsReturn[R any](action ActionAnyParamsReturn[R]) StateActionAnyParamsReturnImpl[R] {
	return &StateActionAnyParamsReturn[R]{
		state: NewState(action),
	}
}

func (a *StateActionAnyParamsReturn[R]) Set(action ActionAnyParamsReturn[R]) {
	a.state.Set(action)
}

func (a *StateActionAnyParamsReturn[R]) Func() func(args ...any) R {
	return func(args ...any) R {
		return a.Call(args...)
	}
}

func (a *StateActionAnyParamsReturn[R]) Call(args ...any) R {
	return a.state.Get().Call(args...)
}

type ActionParamsReturnImpl[P, R any] interface {
	Func() func(args ...P) R
	Call(args ...P) R
}

type ActionParamsReturn[P, R any] func(args ...P) R

func NewActionParamsReturn[P, R any](action func(args ...P) R) ActionParamsReturnImpl[P, R] {
	return ActionParamsReturn[P, R](action)
}

func (a ActionParamsReturn[P, R]) Func() func(args ...P) R {
	return func(args ...P) R {
		return a.Call(args...)
	}
}

func (a ActionParamsReturn[P, R]) Call(args ...P) R {
	return a(args...)
}

type StateActionParamsReturnImpl[P, R any] interface {
	Set(action ActionParamsReturn[P, R])
	Func() func(args ...P) R
	Call(args ...P) R
}

type StateActionParamsReturn[P, R any] struct {
	state StateImpl[ActionParamsReturn[P, R]]
}

func NewStateActionParamsReturn[P, R any](action ActionParamsReturn[P, R]) StateActionParamsReturnImpl[P, R] {
	return &StateActionParamsReturn[P, R]{
		state: NewState(action),
	}
}

func (a *StateActionParamsReturn[P, R]) Set(action ActionParamsReturn[P, R]) {
	a.state.Set(action)
}

func (a *StateActionParamsReturn[P, R]) Func() func(args ...P) R {
	return func(args ...P) R {
		return a.Call(args...)
	}
}

func (a *StateActionParamsReturn[P, R]) Call(args ...P) R {
	return a.state.Get().Call(args...)
}

type SafeReader struct {
	mutex  *sync.Mutex
	reader io.Reader
}

func NewSafeReader(writer io.Reader) *SafeReader {
	return &SafeReader{
		reader: writer,
		mutex:  &sync.Mutex{},
	}
}

func (sw *SafeReader) Read(p []byte) (n int, err error) {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()
	return sw.reader.Read(p)
}

type SafeWriter struct {
	mutex  *sync.Mutex
	writer io.Writer
}

func NewSafeWriter(writer io.Writer) *SafeWriter {
	return &SafeWriter{
		writer: writer,
		mutex:  &sync.Mutex{},
	}
}

func (sw *SafeWriter) Write(p []byte) (n int, err error) {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()
	return sw.writer.Write(p)
}
