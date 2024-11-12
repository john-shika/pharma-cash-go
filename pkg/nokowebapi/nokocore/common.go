package nokocore

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"sync"
)

var ErrDataTypeInvalid = errors.New("invalid data type")

var EmptyString string

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

func CopyStack[T any](value []T) []T {
	temp := make([]T, len(value))
	copy(temp, value)
	return temp
}

func ToString(value any) string {
	if value == nil {
		return "<null>"
	}

	if val, ok := value.(StringableImpl); ok {
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
		return fmt.Sprint(value)
	}
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

func IsNone(value any) bool {
	return value == nil
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

type StackCollectionImpl[T any] interface {
	*T | []T
}

type MapCollectionImpl[T any] interface {
	StackCollectionImpl[T] | map[string]T
}

type MapImpl[T any] interface {
	Len() int
	Keys() []string
	Values() []T
	HasKey(key string) bool
	ContainKeys(keys ...string) bool
	Get(key string) any
	Set(key string, value any) bool
	delVal(key string) bool
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
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Map[T]) Values() []T {
	values := make([]T, 0, len(m))
	for i, v := range m {
		KeepVoid(i, v)

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

func (m Map[T]) getVal(key string) T {
	var ok bool
	var value T
	KeepVoid(ok, value)

	if value, ok = m[key]; !ok {
		return Default[T]()
	}
	return value
}

func (m Map[T]) Get(key string) any {
	tokens := strings.Split(strings.Trim(key, "."), ".")
	if len(tokens) == 0 {
		panic("invalid key")
	}

	var ok bool
	var err error
	var idx int
	var temp any
	KeepVoid(ok, idx, temp)

	temp = m
	for i, token := range tokens {
		KeepVoid(i)

		// find by key name
		if idx, err = strconv.Atoi(token); err != nil {
			temp = Unwrap(Cast[MapAny](temp)).getVal(token)
			continue
		}

		// find by index
		temp = Unwrap(Cast[ArrayAny](temp)).At(idx)
	}

	return temp
}

func (m Map[T]) setVal(key string, value T) bool {
	m[key] = value
	return true
}

func (m Map[T]) Set(key string, value any) bool {
	tokens := strings.Split(strings.Trim(key, "."), ".")
	if len(tokens) == 0 {
		panic("invalid key")
	}

	var ok bool
	var err error
	var idx int
	var temp any
	KeepVoid(ok, idx, temp)

	n := len(tokens)

	temp = m
	for i, token := range tokens {
		KeepVoid(i)

		// find by key name
		if idx, err = strconv.Atoi(token); err != nil {
			if n == i+1 {
				return Unwrap(Cast[MapAny](temp)).setVal(token, value)
			}
			temp = Unwrap(Cast[MapAny](temp)).getVal(token)
			continue
		}

		// find by index
		if n == i+1 {
			return Unwrap(Cast[ArrayAny](temp)).Set(idx, value)
		}
		temp = Unwrap(Cast[ArrayAny](temp)).At(idx)
	}

	return false
}

func (m Map[T]) delVal(key string) bool {
	var ok bool
	var temp T
	KeepVoid(ok, temp)

	if temp, ok = m[key]; !ok {
		return false
	}

	delete(m, key)
	return true
}

func (m Map[T]) Del(key string) bool {
	tokens := strings.Split(strings.Trim(key, "."), ".")
	if len(tokens) == 0 {
		panic("invalid key")
	}

	var ok bool
	var err error
	var idx int
	var temp any
	KeepVoid(ok, idx, temp)

	n := len(tokens)

	temp = m
	for i, token := range tokens {
		KeepVoid(i)

		// find by key name
		if idx, err = strconv.Atoi(token); err != nil {
			if n == i+1 {
				return Unwrap(Cast[MapAny](temp)).delVal(token)
			}
			temp = Unwrap(Cast[MapAny](temp)).getVal(token)
			continue
		}

		// find by index
		if n == i+1 {
			return Unwrap(Cast[ArrayAny](temp)).Del(idx)
		}
		temp = Unwrap(Cast[ArrayAny](temp)).At(idx)
	}

	return false
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
	fmt.Println("[WARN] array deletion for slices is not supported. Use a dynamic array implementation instead.")
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

func IsNoneOrEmpty(value string) bool {
	return value == "" || value == "\x00"
}

func IsNoneOrEmptyWhiteSpace(value string) bool {
	if IsNoneOrEmpty(value) {
		return true
	}
	temp := strings.TrimSpace(value)
	return temp == "" ||
		temp == "\x00" ||
		temp == "\xA0" ||
		temp == "\t" ||
		temp == "\r" ||
		temp == "\n" ||
		temp == "\r\n"
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
	var err error
	var temp uuid.UUID
	KeepVoid(err, temp)

	if temp, err = uuid.NewV7(); err != nil {
		panic(err)
	}

	return temp
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
