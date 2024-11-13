package nokocore

import (
	"errors"
	"fmt"
	"iter"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

func EqualsReflect(value, other any) bool {
	return reflect.DeepEqual(value, other)
}

func PassTypeIndirectReflect(value any) reflect.Type {
	typ := GetTypeReflect(value)
	switch typ.Kind() {
	case reflect.Pointer:
		return PassTypeIndirectReflect(typ.Elem())
	default:
		return typ
	}
}

func TypeEqualsReflect(value, other any) bool {
	return PassTypeIndirectReflect(value) == PassTypeIndirectReflect(other)
}

func GetTypeReflect(value any) reflect.Type {
	var val reflect.Value
	switch value.(type) {
	case reflect.Type:
		return value.(reflect.Type)
	case reflect.Value:
		val = value.(reflect.Value)
	default:
		val = reflect.ValueOf(value)
	}
	return val.Type()
}

func GetValueReflect(value any) reflect.Value {
	var val reflect.Value
	switch value.(type) {
	case reflect.Type:
		panic("should not pass reflect type")
	case reflect.Value:
		val = value.(reflect.Value)
	default:
		val = reflect.ValueOf(value)
	}
	return val
}

func GetKindReflect(value any) reflect.Kind {
	return PassValueIndirectReflect(value).Kind()
}

func IsValidReflect(value any) bool {
	val := GetValueReflect(value)
	// must be not zero value
	if val.IsValid() {
		switch val.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
			// chan, func, interface, map, pointer, slice
			// will be considered as nullable value
			return !val.IsNil()
		default:
			// not chan, func, interface, map, pointer, slice
			// will be considered as notnull value
			return true
		}
	}
	// zero or null value
	return false
}

func PassValueIndirectReflect(value any) reflect.Value {
	val := GetValueReflect(value)

	if !val.IsValid() {
		return val
	}

	switch val.Kind() {
	case reflect.Interface, reflect.Pointer:
		return PassValueIndirectReflect(val.Elem())
	default:
		return val
	}
}

func IsCountableReflect[T any](value T) bool {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		panic("invalid value")
	}
	switch val.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Chan:
		return true
	default:
		return false
	}
}

func GetSizeReflect(value any) int {
	val := PassValueIndirectReflect(value)
	if !IsCountableReflect(val) {
		panic("value is not countable")
	}
	return val.Len()
}

func IsNoneOrEmptyReflect(value any) bool {
	val := GetValueReflect(value)
	if !val.IsValid() {
		return true
	}
	if IsCountableReflect(val) {
		return GetSizeReflect(val) == 0
	}
	return false
}

func IsStringableReflect(value any) bool {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		panic("invalid value")
	}
	switch val.Kind() {
	case reflect.String:
		return true
	default:
		return false
	}
}

func IsExportedFieldReflect(value any) bool {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return false
	}
	return val.CanInterface()
}

func IsExportedFieldAtReflect(value any, index int) bool {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return false
	}
	return IsExportedFieldReflect(val.Field(index))
}

func IsExportedFieldByNameReflect(value any, name string) bool {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return false
	}
	return IsExportedFieldReflect(val.FieldByName(name))
}

func ToStringReflect(value any) string {
	if value == nil {
		return "<null>"
	}

	if val, ok := value.(StringableImpl); ok {
		return val.ToString()
	}

	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return "<null>"
	}

	switch val.Kind() {
	case reflect.Bool:
		v := val.Bool()
		return strconv.FormatBool(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := val.Int()
		return strconv.FormatInt(v, 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := val.Uint()
		return strconv.FormatUint(v, 10)
	case reflect.Uintptr:
		v := Unwrap(Cast[uintptr](val.Interface()))
		return ToString(v)
	case reflect.Float32, reflect.Float64:
		v := val.Float()
		return strconv.FormatFloat(v, 'f', -1, 64)
	case reflect.Complex64, reflect.Complex128:
		v := val.Complex()
		return strconv.FormatComplex(v, 'f', -1, 128)
	case reflect.String:
		return val.String()
	case reflect.Struct:
		if IsTimeUtcISO8601(val) {
			return strconv.Quote(ToTimeUtcStringISO8601(val))
		}
		return "<struct>"
	case reflect.Array, reflect.Slice:
		return "<array>"
	case reflect.Map:
		return "<map>"
	default:
		return fmt.Sprint(val.Interface())
	}
}

func CastPassValueIndirect[T any](value any) T {
	var ok bool
	var val any
	var temp T
	KeepVoid(ok, val, temp)

	if val = PassValueIndirect(value); val == nil {
		return temp
	}

	if temp, ok = Cast[T](val); !ok {
		panic("invalid data type")
	}

	return temp
}

func PassValueIndirect(value any) any {
	val := PassValueIndirectReflect(value)

	if !val.IsValid() {
		return nil
	}

	return val.Interface()
}

func GetNameTypeReflect(value any) string {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		panic("invalid value")
	}
	method := val.MethodByName("GetNameType")
	if method.IsValid() {
		results := method.Call(nil)
		if len(results) != 1 || !results[0].IsValid() {
			panic("invalid results")
		}
		name := Unwrap(CastString(results[0].Interface()))
		return name
	}

	return val.Type().Name()
}

func makePtrReflect(value any) (reflect.Value, bool) {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		panic("invalid value")
	}

	if val.CanAddr() {
		return reflect.NewAt(val.Type(), unsafe.Pointer(val.UnsafeAddr())), true
	}

	return reflect.Value{}, false
}

func MakePointerReflect(value any) reflect.Value {
	var ok bool
	var ptr reflect.Value
	KeepVoid(ok, ptr)

	val := GetValueReflect(value)
	if !val.IsValid() {
		panic("invalid value")
	}

	if ptr, ok = makePtrReflect(value); ok {
		return ptr
	}

	switch val.Kind() {
	case reflect.Interface:
		fmt.Println("[WARN] Value is an interface value. The result is a pointer to the value stored in the interface.")
		return MakePointerReflect(val.Elem())
	case reflect.Pointer:
		return val
	case reflect.Func:
		fmt.Println("[WARN] The returned pointer is an underlying code pointer, but not necessarily enough to identify a single function uniquely.")
		return reflect.NewAt(val.Type(), val.UnsafePointer())
	case reflect.Slice:
		fmt.Println("[WARN] The returned pointer is to the first element of the slice.")
		return reflect.NewAt(val.Type(), val.UnsafePointer())
	case reflect.String:
		fmt.Println("[WARN] The returned pointer is to the first element of the underlying bytes of string.")
		return reflect.NewAt(val.Type(), val.UnsafePointer())
	case reflect.Chan, reflect.Map, reflect.UnsafePointer:
		return reflect.NewAt(val.Type(), val.UnsafePointer())
	//case reflect.Slice, reflect.Array, reflect.Struct:
	//	if ptr, ok = makePtrReflect(value); !ok {
	//		panic("value is not addressable")
	//	}
	//	return ptr
	default:

		panic("use a generic pointer instead")
	}
}

func GetArrayItem[T any](array []T, index int) T {
	if index >= 0 && index < len(array) {
		return array[index]
	}

	return Default[T]()
}

func GetArrayItemReflect(array any, index int) reflect.Value {
	val := PassValueIndirectReflect(array)
	//ref := MakePointerReflect(val).Elem()

	if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
		panic("invalid data type")
	}

	if index >= 0 && index < val.Len() {
		//return ref.Index(index)
		return val.Index(index)
	}

	return reflect.Value{}
}

func SetArrayItem[T any](array []T, index int, value T) bool {
	if index >= 0 && index < len(array) {
		array[index] = value
		return true
	}

	return false
}

func SetArrayItemReflect(array any, index int, value any) bool {
	val := PassValueIndirectReflect(array)
	//ref := MakePointerReflect(val).Elem()

	if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
		panic("invalid data type")
	}

	if index >= 0 && index < val.Len() {
		//ref.Index(index).Set(GetValueReflect(value))
		val.Index(index).Set(GetValueReflect(value))
		return true
	}

	return false
}

func DeleteArrayItem[T any](arrayPtr *[]T, index int) bool {
	array := *arrayPtr
	j := index + 1
	if j <= len(array) {
		*arrayPtr = append(array[:index], array[j:]...)
		return true
	}

	*arrayPtr = array[:index]
	return true
}

func DeleteArrayItemReflect(array any, index int) bool {
	val := PassValueIndirectReflect(array)
	//ref := MakePointerReflect(val).Elem()

	if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
		panic("invalid data type")
	}

	if !val.CanSet() {
		panic("value is not addressable")
	}

	start := index + 1
	if start <= val.Len() {

		val.Set(reflect.AppendSlice(val.Slice(0, index), val.Slice(start, val.Len())))
		return true
	}

	val.Set(val.Slice(0, index))
	return true
}

func GetMapValue[K comparable, V any](m map[K]V, key K) V {
	if val, ok := m[key]; ok {
		return val
	}

	return Default[V]()
}

func GetMapValueReflect(m any, key string) reflect.Value {
	val := PassValueIndirectReflect(m)
	//ref := MakePointerReflect(val).Elem()

	if val.Kind() != reflect.Map {
		panic("invalid data type")
	}

	//return ref.MapIndex(GetValueReflect(key))
	return val.MapIndex(GetValueReflect(key))
}

func SetMapValue[K comparable, V any](m map[K]V, key K, value V) bool {
	m[key] = value
	return true
}

func SetMapValueReflect(m any, key string, value any) bool {
	val := PassValueIndirectReflect(m)
	//ref := MakePointerReflect(val).Elem()

	if val.Kind() != reflect.Map {
		panic("invalid data type")
	}

	//ref.SetMapIndex(GetValueReflect(key), GetValueReflect(value))
	val.SetMapIndex(GetValueReflect(key), GetValueReflect(value))
	return true
}

func DeleteMapValue[K comparable, V any](m map[K]V, key K) bool {
	delete(m, key)
	return true
}

func DeleteMapValueReflect(m any, key string) bool {
	val := PassValueIndirectReflect(m)
	//ref := MakePointerReflect(val).Elem()

	if val.Kind() != reflect.Map {
		panic("invalid data type")
	}

	//ref.SetMapIndex(GetValueReflect(key), reflect.Value{})
	val.SetMapIndex(GetValueReflect(key), reflect.Value{})
	return true
}

func GetBoolValueReflect(value any) bool {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return false
	}

	switch val.Kind() {
	case reflect.Bool:
		return val.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return val.Float() != 0
	case reflect.Complex64, reflect.Complex128:
		return val.Complex() != 0
	case reflect.String:
		return ParseEnvToBool(val.String())
	default:
		return false
	}
}

func GetIntValueReflect(value any) int64 {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return 0
	}

	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			return 1
		}

		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int64(val.Uint())
	case reflect.Float32, reflect.Float64:
		return int64(val.Float())
	case reflect.Complex64, reflect.Complex128:
		return int64(real(val.Complex()))
	case reflect.String:
		return ParseEnvToInt(val.String())
	default:
		return 0
	}
}

func GetUintValueReflect(value any) uint64 {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return 0
	}

	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			return 1
		}

		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint()
	case reflect.Float32, reflect.Float64:
		return uint64(val.Float())
	case reflect.Complex64, reflect.Complex128:
		return uint64(real(val.Complex()))
	case reflect.String:
		return ParseEnvToUint(val.String())
	default:
		return 0
	}
}

func GetFloatValueReflect(value any) float64 {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return 0
	}

	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			return 1
		}

		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(val.Uint())
	case reflect.Float32, reflect.Float64:
		return val.Float()
	case reflect.Complex64, reflect.Complex128:
		return real(val.Complex())
	case reflect.String:
		return ParseEnvToFloat(val.String())
	default:
		return 0
	}
}

func GetComplexValueReflect(value any) complex128 {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return 0
	}

	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			return 1
		}

		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return complex(float64(val.Int()), 0)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return complex(float64(val.Uint()), 0)
	case reflect.Float32, reflect.Float64:
		return complex(val.Float(), 0)
	case reflect.Complex64, reflect.Complex128:
		return val.Complex()
	case reflect.String:
		return ParseEnvToComplex(val.String())
	default:
		return 0
	}
}

func GetStringValueReflect(value any) string {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return ""
	}

	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			return "true"
		}

		return "false"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return fmt.Sprintf("%d", val.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", val.Float())
	case reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%f", real(val.Complex()))
	case reflect.String:
		return val.String()
	default:
		return ""
	}
}

func GetArrayValueReflect(value any) []any {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return nil
	}

	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		temp := make([]any, val.Len())
		for i := 0; i < val.Len(); i++ {
			temp[i] = val.Index(i).Interface()
		}

		return temp

	default:
		return nil
	}
}

func GetMapStrAnyValueReflect(value any) map[string]any {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return nil
	}

	switch val.Kind() {
	case reflect.Map:
		temp := make(map[string]any)
		mapIter := val.MapRange()
		for mapIter.Next() {
			if keyName := GetStringValueReflect(mapIter.Key()); keyName != "" {
				SetMapValue(temp, keyName, mapIter.Value().Interface())
				continue
			}
			panic("invalid key name")
		}
		return temp
	case reflect.Struct:
		temp := make(map[string]any)
		options := &ForEachStructFieldsOptions{
			Validation: false,
		}

		// for each struct field
		NoErr(ForEachStructFieldsReflect(val.Interface(), options, func(name string, sFieldX StructFieldExpandedImpl) error {
			if mapStrAny := GetMapStrAnyValueReflect(sFieldX.GetValue()); mapStrAny != nil {
				SetMapValue(temp, name, any(mapStrAny))
				return nil
			}

			// set map item as any
			SetMapValue(temp, name, sFieldX.GetValue().Interface())
			return nil
		}))
		return temp
	default:
		return nil
	}
}

func GetMapValueWithSuperKeyReflect(m any, key string) any {
	tokens := strings.Split(strings.Trim(key, "."), ".")
	if len(tokens) == 0 {
		panic("invalid key")
	}

	var ok bool
	var err error
	var idx int
	KeepVoid(ok, idx)

	temp := m

	for i, token := range tokens {
		KeepVoid(i)

		// find by key name
		if idx, err = strconv.Atoi(token); err != nil {
			temp = GetMapValueReflect(temp, token).Interface()
			continue
		}

		// find by index
		temp = GetArrayItemReflect(temp, idx)
	}

	return temp
}

func SetMapValueWithSuperKeyReflect(m any, key string, value any) bool {
	tokens := strings.Split(strings.Trim(key, "."), ".")
	if len(tokens) == 0 {
		panic("invalid key")
	}

	var ok bool
	var err error
	var idx int
	KeepVoid(ok, idx)

	n := len(tokens)
	val := PassValueIndirectReflect(value)
	temp := m

	for i, token := range tokens {
		KeepVoid(i)

		// find by key name
		if idx, err = strconv.Atoi(token); err != nil {
			if n == i+1 {
				SetMapValueReflect(temp, token, val)
				return true
			}
			temp = GetMapValueReflect(temp, token).Interface()
			continue
		}

		// find by index
		if n == i+1 {
			SetArrayItemReflect(temp, idx, value)
			return true
		}
		temp = GetArrayItemReflect(temp, idx).Interface()
	}

	return false
}

func DeleteMapValueWithSuperKeyReflect(m any, key string) bool {
	tokens := strings.Split(strings.Trim(key, "."), ".")
	if len(tokens) == 0 {
		panic("invalid key")
	}

	var ok bool
	var err error
	var idx int
	KeepVoid(ok, idx)

	n := len(tokens)
	temp := m

	for i, token := range tokens {
		KeepVoid(i)

		// find by key name
		if idx, err = strconv.Atoi(token); err != nil {
			if n == i+1 {
				DeleteMapValueReflect(temp, token)
				return true
			}
			temp = GetMapValueReflect(temp, token).Interface()
			continue
		}

		// find by index
		if n == i+1 {
			DeleteArrayItemReflect(temp, idx)
			return true
		}
		temp = GetArrayItemReflect(temp, idx).Interface()
	}

	return false
}

func SetValueReflect(field any, value any) error {
	vField := PassValueIndirectReflect(field)
	val := PassValueIndirectReflect(value)

	if !vField.IsValid() {
		return errors.New("field is not valid")
	}

	if !vField.CanSet() {
		return errors.New("value is not addressable")
	}

	if !val.IsValid() {
		vField.SetZero()
		return nil
	}

	switch vField.Kind() {
	case reflect.Bool:
		vField.SetBool(GetBoolValueReflect(val))
		return nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vField.SetInt(GetIntValueReflect(val))
		return nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		vField.SetUint(GetUintValueReflect(val))
		return nil

	case reflect.Float32, reflect.Float64:
		vField.SetFloat(GetFloatValueReflect(val))
		return nil

	case reflect.Complex64, reflect.Complex128:
		vField.SetComplex(GetComplexValueReflect(val))
		return nil

	case reflect.String:
		vField.SetString(GetStringValueReflect(val))
		return nil

	case reflect.Array, reflect.Slice:
		temp := GetValueReflect(GetArrayValueReflect(val))
		if temp.Kind() != reflect.Array && temp.Kind() != reflect.Slice {
			return errors.New("value type does not match field type")
		}

		// Convert an interface array to slice
		// Initialize the slice with the same length as the original array
		vField.Set(reflect.MakeSlice(vField.Type(), temp.Len(), temp.Len()))

		// Iterate through the original array and set each element in the new slice,
		// converting it to the appropriate type
		for i := 0; i < temp.Len(); i++ {
			if err := SetValueReflect(vField.Index(i), temp.Index(i)); err != nil {
				return err
			}
		}

		return nil

	case reflect.Map:
		// map[string]any
		temp := GetValueReflect(GetMapStrAnyValueReflect(val))
		if !TypeEqualsReflect(vField, temp) {
			return errors.New("value type does not match field type")
		}

		vField.Set(temp)
		return nil

	case reflect.Struct:
		if TypeEqualsReflect(vField, val) {
			vField.Set(val)
			return nil
		}

		temp := GetMapStrAnyValueReflect(val)

		if temp != nil {
			options := &ForEachStructFieldsOptions{
				Validation: false,
			}

			return ForEachStructFieldsReflect(vField, options, func(name string, sFieldX StructFieldExpandedImpl) error {
				return SetValueReflect(sFieldX.GetValue(), GetMapValue(temp, name))
			})
		}

		return errors.New("invalid data type")

	default:
		// interface
		vField.Set(val)
		return nil
	}
}

type StructFieldExpandedImpl interface {
	GetStructField() reflect.StructField
	GetStructTagExpanded() StructTagExpandedImpl
	GetValue() reflect.Value
	GetName() string
	GetPkgPath() string
	GetType() reflect.Type
	GetTag() reflect.StructTag
	GetOffset() uintptr
	GetIndex() []int
	GetAnonymous() bool
	IsExported() bool
	Seq() iter.Seq[reflect.Value]
	Seq2() iter.Seq2[reflect.Value, reflect.Value]
	Addr() reflect.Value
	Bool() bool
	Bytes() []byte
	CanAddr() bool
	CanSet() bool
	Call(in []reflect.Value) []reflect.Value
	CallSlice(in []reflect.Value) []reflect.Value
	Cap() int
	Close()
	CanComplex() bool
	Complex() complex128
	Elem() reflect.Value
	Field(i int) reflect.Value
	FieldByIndex(index []int) reflect.Value
	FieldByIndexErr(index []int) (reflect.Value, error)
	FieldByName(name string) reflect.Value
	FieldByNameFunc(match func(string) bool) reflect.Value
	CanFloat() bool
	Float() float64
	Index(i int) reflect.Value
	CanInt() bool
	Int() int64
	CanInterface() bool
	Interface() (i any)
	IsNil() bool
	IsValid() bool
	IsZero() bool
	SetZero()
	Kind() reflect.Kind
	Len() int
	MapIndex(key reflect.Value) reflect.Value
	MapKeys() []reflect.Value
	SetIterKey(iter *reflect.MapIter)
	SetIterValue(iter *reflect.MapIter)
	MapRange() *reflect.MapIter
	Method(i int) reflect.Value
	NumMethod() int
	MethodByName(name string) reflect.Value
	NumField() int
	OverflowComplex(x complex128) bool
	OverflowFloat(x float64) bool
	OverflowInt(x int64) bool
	OverflowUint(x uint64) bool
	Pointer() uintptr
	Recv() (x reflect.Value, ok bool)
	Send(x reflect.Value)
	Set(x reflect.Value)
	SetBool(x bool)
	SetBytes(x []byte)
	SetComplex(x complex128)
	SetFloat(x float64)
	SetInt(x int64)
	SetLen(n int)
	SetCap(n int)
	SetMapIndex(key reflect.Value, elem reflect.Value)
	SetUint(x uint64)
	SetPointer(x unsafe.Pointer)
	SetString(x string)
	Slice(i int, j int) reflect.Value
	Slice3(i int, j int, k int) reflect.Value
	String() string
	TryRecv() (x reflect.Value, ok bool)
	TrySend(x reflect.Value) bool
	Type() reflect.Type
	CanUint() bool
	Uint() uint64
	UnsafeAddr() uintptr
	UnsafePointer() unsafe.Pointer
	Grow(n int)
	Clear()
	Convert(t reflect.Type) reflect.Value
	CanConvert(t reflect.Type) bool
	Comparable() bool
	Equal(u reflect.Value) bool
}

type StructFieldExpanded struct {
	structField       reflect.StructField
	structTagExpanded StructTagExpandedImpl
	value             reflect.Value
}

func NewStructFieldExpanded(field reflect.StructField, sTagX StructTagExpandedImpl, value reflect.Value) StructFieldExpandedImpl {
	return &StructFieldExpanded{
		structField:       field,
		structTagExpanded: sTagX,
		value:             value,
	}
}

func (s *StructFieldExpanded) GetStructField() reflect.StructField {
	return s.structField
}

func (s *StructFieldExpanded) GetStructTagExpanded() StructTagExpandedImpl {
	return s.structTagExpanded
}

func (s *StructFieldExpanded) GetValue() reflect.Value {
	return s.value
}

func (s *StructFieldExpanded) GetName() string {
	return s.structField.Name
}

func (s *StructFieldExpanded) GetPkgPath() string {
	return s.structField.PkgPath
}

func (s *StructFieldExpanded) GetType() reflect.Type {
	return s.structField.Type
}

func (s *StructFieldExpanded) GetTag() reflect.StructTag {
	return s.structField.Tag
}

func (s *StructFieldExpanded) GetOffset() uintptr {
	return s.structField.Offset
}

func (s *StructFieldExpanded) GetIndex() []int {
	return s.structField.Index
}

func (s *StructFieldExpanded) GetAnonymous() bool {
	return s.structField.Anonymous
}

func (s *StructFieldExpanded) IsExported() bool {
	return s.structField.IsExported()
}

func (s *StructFieldExpanded) Seq() iter.Seq[reflect.Value] {
	return s.value.Seq()
}

func (s *StructFieldExpanded) Seq2() iter.Seq2[reflect.Value, reflect.Value] {
	return s.value.Seq2()
}

func (s *StructFieldExpanded) Addr() reflect.Value {
	return s.value.Addr()
}

func (s *StructFieldExpanded) Bool() bool {
	return s.value.Bool()
}

func (s *StructFieldExpanded) Bytes() []byte {
	return s.value.Bytes()
}

func (s *StructFieldExpanded) CanAddr() bool {
	return s.value.CanAddr()
}

func (s *StructFieldExpanded) CanSet() bool {
	return s.value.CanSet()
}

func (s *StructFieldExpanded) Call(in []reflect.Value) []reflect.Value {
	return s.value.Call(in)
}

func (s *StructFieldExpanded) CallSlice(in []reflect.Value) []reflect.Value {
	return s.value.CallSlice(in)
}

func (s *StructFieldExpanded) Cap() int {
	return s.value.Cap()
}

func (s *StructFieldExpanded) Close() {
	s.value.Close()
}

func (s *StructFieldExpanded) CanComplex() bool {
	return s.value.CanComplex()
}

func (s *StructFieldExpanded) Complex() complex128 {
	return s.value.Complex()
}

func (s *StructFieldExpanded) Elem() reflect.Value {
	return s.value.Elem()
}

func (s *StructFieldExpanded) Field(i int) reflect.Value {
	return s.value.Field(i)
}

func (s *StructFieldExpanded) FieldByIndex(index []int) reflect.Value {
	return s.value.FieldByIndex(index)
}

func (s *StructFieldExpanded) FieldByIndexErr(index []int) (reflect.Value, error) {
	return s.value.FieldByIndexErr(index)
}

func (s *StructFieldExpanded) FieldByName(name string) reflect.Value {
	return s.value.FieldByName(name)
}

func (s *StructFieldExpanded) FieldByNameFunc(match func(string) bool) reflect.Value {
	return s.value.FieldByNameFunc(match)
}

func (s *StructFieldExpanded) CanFloat() bool {
	return s.value.CanFloat()
}

func (s *StructFieldExpanded) Float() float64 {
	return s.value.Float()
}

func (s *StructFieldExpanded) Index(i int) reflect.Value {
	return s.value.Index(i)
}

func (s *StructFieldExpanded) CanInt() bool {
	return s.value.CanInt()
}

func (s *StructFieldExpanded) Int() int64 {
	return s.value.Int()
}

func (s *StructFieldExpanded) CanInterface() bool {
	return s.value.CanInterface()
}

func (s *StructFieldExpanded) Interface() (i any) {
	return s.value.Interface()
}

func (s *StructFieldExpanded) IsNil() bool {
	return s.value.IsNil()
}

func (s *StructFieldExpanded) IsValid() bool {
	return s.value.IsValid()
}

func (s *StructFieldExpanded) IsZero() bool {
	return s.value.IsZero()
}

func (s *StructFieldExpanded) SetZero() {
	s.value.SetZero()
}

func (s *StructFieldExpanded) Kind() reflect.Kind {
	return s.value.Kind()
}

func (s *StructFieldExpanded) Len() int {
	return s.value.Len()
}

func (s *StructFieldExpanded) MapIndex(key reflect.Value) reflect.Value {
	return s.value.MapIndex(key)
}

func (s *StructFieldExpanded) MapKeys() []reflect.Value {
	return s.value.MapKeys()
}

func (s *StructFieldExpanded) SetIterKey(iter *reflect.MapIter) {
	s.value.SetIterKey(iter)
}

func (s *StructFieldExpanded) SetIterValue(iter *reflect.MapIter) {
	s.value.SetIterValue(iter)
}

func (s *StructFieldExpanded) MapRange() *reflect.MapIter {
	return s.value.MapRange()
}

func (s *StructFieldExpanded) Method(i int) reflect.Value {
	return s.value.Method(i)
}

func (s *StructFieldExpanded) NumMethod() int {
	return s.value.NumMethod()
}

func (s *StructFieldExpanded) MethodByName(name string) reflect.Value {
	return s.value.MethodByName(name)
}

func (s *StructFieldExpanded) NumField() int {
	return s.value.NumField()
}

func (s *StructFieldExpanded) OverflowComplex(x complex128) bool {
	return s.value.OverflowComplex(x)
}

func (s *StructFieldExpanded) OverflowFloat(x float64) bool {
	return s.value.OverflowFloat(x)
}

func (s *StructFieldExpanded) OverflowInt(x int64) bool {
	return s.value.OverflowInt(x)
}

func (s *StructFieldExpanded) OverflowUint(x uint64) bool {
	return s.value.OverflowUint(x)
}

func (s *StructFieldExpanded) Pointer() uintptr {
	return s.value.Pointer()
}

func (s *StructFieldExpanded) Recv() (x reflect.Value, ok bool) {
	return s.value.Recv()
}

func (s *StructFieldExpanded) Send(x reflect.Value) {
	s.value.Send(x)
}

func (s *StructFieldExpanded) Set(x reflect.Value) {
	s.value.Set(x)
}

func (s *StructFieldExpanded) SetBool(x bool) {
	s.value.SetBool(x)
}

func (s *StructFieldExpanded) SetBytes(x []byte) {
	s.value.SetBytes(x)
}

func (s *StructFieldExpanded) SetComplex(x complex128) {
	s.value.SetComplex(x)
}

func (s *StructFieldExpanded) SetFloat(x float64) {
	s.value.SetFloat(x)
}

func (s *StructFieldExpanded) SetInt(x int64) {
	s.value.SetInt(x)
}

func (s *StructFieldExpanded) SetLen(n int) {
	s.value.SetLen(n)
}

func (s *StructFieldExpanded) SetCap(n int) {
	s.value.SetCap(n)
}

func (s *StructFieldExpanded) SetMapIndex(key reflect.Value, elem reflect.Value) {
	s.value.SetMapIndex(key, elem)
}

func (s *StructFieldExpanded) SetUint(x uint64) {
	s.value.SetUint(x)
}

func (s *StructFieldExpanded) SetPointer(x unsafe.Pointer) {
	s.value.SetPointer(x)
}

func (s *StructFieldExpanded) SetString(x string) {
	s.value.SetString(x)
}

func (s *StructFieldExpanded) Slice(i int, j int) reflect.Value {
	return s.value.Slice(i, j)
}

func (s *StructFieldExpanded) Slice3(i int, j int, k int) reflect.Value {
	return s.value.Slice3(i, j, k)
}

func (s *StructFieldExpanded) String() string {
	return s.value.String()
}

func (s *StructFieldExpanded) TryRecv() (x reflect.Value, ok bool) {
	return s.value.TryRecv()
}

func (s *StructFieldExpanded) TrySend(x reflect.Value) bool {
	return s.value.TrySend(x)
}

func (s *StructFieldExpanded) Type() reflect.Type {
	return s.value.Type()
}

func (s *StructFieldExpanded) CanUint() bool {
	return s.value.CanUint()
}

func (s *StructFieldExpanded) Uint() uint64 {
	return s.value.Uint()
}

func (s *StructFieldExpanded) UnsafeAddr() uintptr {
	return s.value.UnsafeAddr()
}

func (s *StructFieldExpanded) UnsafePointer() unsafe.Pointer {
	return s.value.UnsafePointer()
}

func (s *StructFieldExpanded) Grow(n int) {
	s.value.Grow(n)
}

func (s *StructFieldExpanded) Clear() {
	s.value.Clear()
}

func (s *StructFieldExpanded) Convert(t reflect.Type) reflect.Value {
	return s.value.Convert(t)
}

func (s *StructFieldExpanded) CanConvert(t reflect.Type) bool {
	return s.value.CanConvert(t)
}

func (s *StructFieldExpanded) Comparable() bool {
	return s.value.Comparable()
}

func (s *StructFieldExpanded) Equal(u reflect.Value) bool {
	return s.value.Equal(u)
}

type StructTagExpandedImpl interface {
	GetIgnore() bool
	GetOmitEmpty() bool
	GetRequired() bool
	GetName() string
	Unmatch(value any) bool
	Match(value any) bool
}

type StructTagExpanded struct {
	ignore    bool
	omitEmpty bool
	required  bool
	name      string
}

func (s *StructTagExpanded) GetIgnore() bool {
	return s.ignore
}

func (s *StructTagExpanded) GetOmitEmpty() bool {
	return s.omitEmpty
}

func (s *StructTagExpanded) GetRequired() bool {
	return s.required
}

func (s *StructTagExpanded) GetName() string {
	return s.name
}

func (s *StructTagExpanded) Unmatch(value any) bool {
	if s.ignore {
		return true
	}

	if s.omitEmpty && s.name == "" {
		return true
	}

	if s.required && IsNoneOrEmptyReflect(value) {
		return true
	}

	return false
}

func (s *StructTagExpanded) Match(value any) bool {
	return !s.Unmatch(value)
}

func GetStructTagExpanded(key string, sTag reflect.StructTag) StructTagExpandedImpl {
	switch key {
	case "db", "mapstructure", "json", "yaml":
		if val, ok := sTag.Lookup(key); ok {
			val = strings.TrimSpace(val)
			tokens := strings.Split(val, ",")
			if len(tokens) > 0 {
				var name string
				var ignore, omitEmpty, required bool
				for i, token := range tokens {
					KeepVoid(i)

					switch strings.TrimSpace(token) {
					case "-", "ignore":
						ignore = true
					case "omitempty":
						omitEmpty = true
					case "required":
						required = true
					default:
						name = token
					}
				}

				return &StructTagExpanded{
					ignore:    ignore,
					omitEmpty: omitEmpty,
					required:  required,
					name:      name,
				}
			}
		}
		return nil
	default:
		panic("unsupported struct tag key")
	}
}

type StructFieldAction func(name string, sFieldX StructFieldExpandedImpl) error

func (s StructFieldAction) Call(name string, sFieldX StructFieldExpandedImpl) error {
	return s(name, sFieldX)
}

type ForEachStructFieldsOptions struct {
	// Validation indicates whether the field should be matched.
	// Validation is matching ignored, omitempty, required attributes.
	Validation bool `mapstructure:"validation" json:"validation" yaml:"validation"`
}

func NewForEachStructFieldsOptions() *ForEachStructFieldsOptions {
	return &ForEachStructFieldsOptions{
		Validation: true,
	}
}

func ForEachStructFieldsReflect(value any, options *ForEachStructFieldsOptions, action StructFieldAction) error {
	val := PassValueIndirectReflect(value)
	//ref := MakePointerReflect(val).Elem()

	// tag names with level priority
	sTagNames := []string{
		"mapstructure",
		"json",
		"yaml",
	}

	if val.IsValid() && val.Kind() == reflect.Struct {

		typ := val.Type()
		size := val.NumField()

		for i := 0; i < size; i++ {
			KeepVoid(i)

			vField := val.Field(i)
			sField := typ.Field(i)
			sTag := sField.Tag

			// can't be exported
			if !sField.IsExported() {
				continue
			}

			pName := sField.Name
			pValid := true

			var sTagX StructTagExpandedImpl

			for j, sTagName := range sTagNames {
				KeepVoid(j)

				if sTagX = GetStructTagExpanded(sTagName, sTag); sTagX != nil {
					pName = sTagX.GetName()
					pValid = sTagX.Match(vField)
					break
				}
			}

			if !pValid && options.Validation {
				continue
			}

			sFieldX := NewStructFieldExpanded(sField, sTagX, vField)
			if err := action.Call(pName, sFieldX); err != nil {
				return err
			}
		}
		return nil
	}

	return errors.New("value is not struct")
}
