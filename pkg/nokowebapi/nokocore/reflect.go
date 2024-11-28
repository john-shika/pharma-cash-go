package nokocore

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"iter"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

var ForEachStop = errors.New("for each stop")

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
		break

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
		break

	default:
		val = reflect.ValueOf(value)
	}

	return val
}

func GetKindReflect(value any) reflect.Kind {
	return PassValueIndirectReflect(value).Kind()
}

func IsNoneOrEmptyWhiteSpaceReflect(value any) bool {
	return !IsValidReflect(value) || IsEmptyWhiteSpaceReflect(value)
}

func IsEmptyWhiteSpaceReflect(value any) bool {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return true
	}

	switch val.Kind() {
	case reflect.String:
		return strings.TrimSpace(val.String()) == ""

	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return val.Len() == 0

	default:
		return false
	}
}

func IsValidReflect(value any) bool {
	val := GetValueReflect(value)

	// must be not zero value
	if val.IsValid() {
		switch val.Kind() {
		case reflect.Bool:
			return true

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return true

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return true

		case reflect.Uintptr:
			return val.Uint() != 0

		case reflect.Float32, reflect.Float64:
			return true

		case reflect.Complex64, reflect.Complex128:
			return true

		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
			// chan, func, interface, map, pointer, slice
			// will be considered as nullable value
			return !val.IsNil()

		case reflect.String:
			return val.Len() > 0

		default:
			// not chan, func, interface, map, pointer, slice
			// will be considered as notnull value
			return !val.IsZero()
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
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
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

func ToIntReflect(value any) int64 {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return 0
	}

	method := val.MethodByName("ToInt")
	if method.IsValid() {
		results := method.Call(nil)
		if len(results) != 1 || !results[0].IsValid() {
			panic("invalid results")
		}

		return results[0].Int()
	}

	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			return 1
		}

		return 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(val.Uint())

	case reflect.Uintptr:
		return int64(val.Uint())

	case reflect.Float32, reflect.Float64:
		return int64(val.Float())

	case reflect.Complex64, reflect.Complex128:
		return int64(real(val.Complex()))

	case reflect.String:
		return Unwrap(strconv.ParseInt(val.String(), 10, 64))

	case reflect.Array, reflect.Slice:
		return int64(val.Len())

	case reflect.Map:
		return int64(val.Len())

	case reflect.Struct:
		return 0

	default:
		return 0
	}
}

func ToFloatReflect(value any) float64 {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return 0
	}

	method := val.MethodByName("ToFloat")
	if method.IsValid() {
		results := method.Call(nil)
		if len(results) != 1 || !results[0].IsValid() {
			panic("invalid results")
		}

		return results[0].Float()
	}

	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			return 1
		}

		return 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(val.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(val.Uint())

	case reflect.Uintptr:
		return float64(val.Uint())

	case reflect.Float32, reflect.Float64:
		return val.Float()

	case reflect.Complex64, reflect.Complex128:
		return real(val.Complex())

	case reflect.String:
		return Unwrap(strconv.ParseFloat(val.String(), 64))

	case reflect.Array, reflect.Slice:
		return float64(val.Len())

	case reflect.Map:
		return float64(val.Len())

	case reflect.Struct:
		return 0

	default:
		return 0
	}
}

func ToDecimalReflect(value any) decimal.Decimal {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return decimal.NewFromInt(0)
	}

	method := val.MethodByName("ToDecimal")
	if method.IsValid() {
		results := method.Call(nil)
		if len(results) != 1 || !results[0].IsValid() {
			panic("invalid results")
		}

		return results[0].Interface().(decimal.Decimal)
	}

	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			return decimal.NewFromInt(1)
		}

		return decimal.NewFromInt(0)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return decimal.NewFromInt(val.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return decimal.NewFromUint64(val.Uint())

	case reflect.Uintptr:
		return decimal.NewFromUint64(val.Uint())

	case reflect.Float32, reflect.Float64:
		return decimal.NewFromFloat(val.Float())

	case reflect.Complex64, reflect.Complex128:
		return decimal.NewFromFloat(real(val.Complex()))

	case reflect.String:
		return Unwrap(decimal.NewFromString(val.String()))

	case reflect.Array, reflect.Slice:
		return decimal.NewFromInt(int64(val.Len()))

	case reflect.Map:
		return decimal.NewFromInt(int64(val.Len()))

	case reflect.Struct:
		return decimal.NewFromInt(0)

	default:
		return decimal.NewFromInt(0)
	}
}

func ToStringReflect(value any) string {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return "<nil>"
	}

	method := val.MethodByName("ToString")
	if method.IsValid() {
		results := method.Call(nil)
		if len(results) != 1 || !results[0].IsValid() {
			panic("invalid results")
		}

		return results[0].String()
	}

	switch val.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(val.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(val.Uint(), 10)

	case reflect.Uintptr:
		//return fmt.Sprintf("0x%016x", val.Uint())
		return fmt.Sprintf("%#016x", val.Uint())

	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'f', -1, 64)

	case reflect.Complex64, reflect.Complex128:
		return strconv.FormatComplex(val.Complex(), 'f', -1, 128)

	case reflect.String:
		return val.String()

	case reflect.Array:
		elem := val.Type().Elem()
		switch elem.Kind() {
		case reflect.Uint8:
			size := val.Len()
			temp := make([]byte, size)
			for i := 0; i < size; i++ {
				temp[i] = byte(val.Index(i).Uint())
			}

			return fmt.Sprintf("%s", temp)

		default:
			return "<array>"
		}

	case reflect.Slice:
		elem := val.Type().Elem()
		switch elem.Kind() {
		case reflect.Uint8:
			return fmt.Sprintf("%s", val.Bytes())

		default:
			return "<array>"
		}

	case reflect.Map:
		return "<object>"

	case reflect.Struct:
		if IsTimeUtcISO8601(val) {
			return ToTimeUtcStringISO8601(val)
		}

		if IsURL(val) {
			return ToURLString(val)
		}

		if IsUUID(val) {
			return ToUUIDString(val)
		}

		if IsDecimal(val) {
			return ToDecimalString(val)
		}

		return fmt.Sprintf("%s {}", GetNameTypeReflect(val))

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

		return results[0].String()
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
		fmt.Println("[WARN] Float is an interface value. The result is a pointer to the value stored in the interface.")
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

	//ref.SetMapIndex(GetValueReflect(key), reflect.Float{})
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

func GetByteValueReflect(value any) byte {
	return byte(GetIntValueReflect(value))
}

func GetRuneValueReflect(value any) rune {
	return rune(GetIntValueReflect(value))
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

func GetRealComplexValueReflect(value any) float64 {
	return real(GetComplexValueReflect(value))
}

func GetBytesValueReflect(value any) []byte {
	val := PassValueIndirectReflect(value)

	if !val.IsValid() {
		panic("invalid value")
	}

	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			return []byte{1}
		}

		return []byte{0}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []byte{byte(val.Int())}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return []byte{byte(val.Uint())}

	case reflect.Float32, reflect.Float64:
		return []byte{byte(val.Float())}

	case reflect.Complex64, reflect.Complex128:
		return []byte{byte(real(val.Complex()))}

	case reflect.String:
		return []byte(ToStringReflect(value))

	default:
		panic("invalid data type")
	}
}

func GetRunesValueReflect(value any) []rune {
	val := PassValueIndirectReflect(value)

	if !val.IsValid() {
		panic("invalid value")
	}

	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			return []rune{1}
		}

		return []rune{0}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []rune{rune(val.Int())}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return []rune{rune(val.Uint())}

	case reflect.Float32, reflect.Float64:
		return []rune{rune(val.Float())}

	case reflect.Complex64, reflect.Complex128:
		return []rune{rune(real(val.Complex()))}

	case reflect.String:
		return []rune(ToStringReflect(value))

	default:
		panic("invalid data type")
	}
}

func GetArrayValueReflect(value any) []any {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return nil
	}

	switch val.Kind() {
	case reflect.Array, reflect.Slice, reflect.String:
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
			if keyName := ToStringReflect(mapIter.Key()); keyName != "" {
				SetMapValue(temp, keyName, mapIter.Value().Interface())
				continue
			}
			panic("invalid key name")
		}
		return temp

	case reflect.Struct:
		temp := make(map[string]any)
		options := NewForEachStructFieldsOptions()

		// for each struct field
		NoErr(ForEachStructFieldsReflect(val.Interface(), options, func(name string, sFieldX StructFieldExImpl) error {
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

func getTokens(key string) []string {
	tokens := strings.Split(strings.TrimSpace(key), ".")
	for i, token := range tokens {
		tokens[i] = strings.TrimSpace(token)
	}
	return tokens
}

func GetMapValueWithSuperKeyReflect(m any, key string) any {
	var ok bool
	var err error
	var idx int
	KeepVoid(ok, idx)

	tokens := getTokens(key)
	if len(tokens) == 0 {
		panic("invalid key")
	}

	temp := PassValueIndirectReflect(m)

	for i, token := range tokens {
		KeepVoid(i)

		if !temp.IsValid() {
			return nil
		}

		// find by key name
		if idx, err = strconv.Atoi(token); err != nil {
			temp = GetMapValueReflect(temp, token)
			continue
		}

		// find by index
		temp = GetArrayItemReflect(temp, idx)
	}

	if !temp.IsValid() {
		return nil
	}

	return temp.Interface()
}

func SetMapValueWithSuperKeyReflect(m any, key string, value any) bool {
	var ok bool
	var err error
	var idx int
	KeepVoid(ok, idx)

	tokens := getTokens(key)
	if len(tokens) == 0 {
		panic("invalid key")
	}

	n := len(tokens)
	val := PassValueIndirectReflect(value)
	temp := PassValueIndirectReflect(m)

	for i, token := range tokens {
		KeepVoid(i)

		if !temp.IsValid() {
			return false
		}

		// find by key name
		if idx, err = strconv.Atoi(token); err != nil {
			if n == i+1 {
				SetMapValueReflect(temp, token, val)
				return true
			}
			temp = GetMapValueReflect(temp, token)
			continue
		}

		// find by index
		if n == i+1 {
			SetArrayItemReflect(temp, idx, value)
			return true
		}
		temp = GetArrayItemReflect(temp, idx)
	}

	return false
}

func DeleteMapValueWithSuperKeyReflect(m any, key string) bool {
	var ok bool
	var err error
	var idx int
	KeepVoid(ok, idx)

	tokens := getTokens(key)
	if len(tokens) == 0 {
		panic("invalid key")
	}

	n := len(tokens)
	temp := PassValueIndirectReflect(m)

	for i, token := range tokens {
		KeepVoid(i)

		if !temp.IsValid() {
			return false
		}

		// find by key name
		if idx, err = strconv.Atoi(token); err != nil {
			if n == i+1 {
				DeleteMapValueReflect(temp, token)
				return true
			}
			temp = GetMapValueReflect(temp, token)
			continue
		}

		// find by index
		if n == i+1 {
			DeleteArrayItemReflect(temp, idx)
			return true
		}
		temp = GetArrayItemReflect(temp, idx)
	}

	return false
}

func ParseValueReflect(value any, key string) any {
	var err error
	var idx int
	KeepVoid(err, idx)

	val := PassValueIndirectReflect(value)

	if !val.IsValid() {
		return nil
	}

	switch val.Kind() {
	case reflect.String:
		if idx, err = strconv.Atoi(key); err != nil {
			switch key {
			case "bytes":
				return GetBytesValueReflect(val)

			case "runes":
				return GetRunesValueReflect(val)

			default:
				panic("invalid key")
			}
		}

		if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
			return nil
		}

		if 0 > idx || idx >= val.Len() {
			panic("index out of range")
		}

		return val.Index(idx).Interface()

	case reflect.Array, reflect.Slice:
		if idx, err = strconv.Atoi(key); err != nil {
			panic("invalid key")
		}

		if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
			return nil
		}

		if 0 > idx || idx >= val.Len() {
			panic("index out of range")
		}

		return val.Index(idx).Interface()

	case reflect.Map:
		return GetMapValueReflect(val, key).Interface()

	case reflect.Struct:
		var temp any
		KeepVoid(temp)

		options := NewForEachStructFieldsOptions()

		NoErr(ForEachStructFieldsReflect(val, options, func(name string, sFieldX StructFieldExImpl) error {
			if name != key {
				return nil
			}

			temp = sFieldX.GetValue().Interface()
			return ForEachStop
		}))
		return temp

	default:
		switch key {
		case "bool":
			return GetBoolValueReflect(val)

		case "int":
			return GetIntValueReflect(val)

		case "uint":
			return GetUintValueReflect(val)

		case "float":
			return GetFloatValueReflect(val)

		case "complex":
			return GetComplexValueReflect(val)

		case "real":
			return GetRealComplexValueReflect(val)

		case "string":
			return ToStringReflect(val)

		case "byte":
			return GetByteValueReflect(val)

		case "bytes":
			return GetBytesValueReflect(val)

		case "rune":
			return GetRuneValueReflect(val)

		case "runes":
			return GetRunesValueReflect(val)

		default:
			panic("invalid key")
		}
	}
}

func baseToken(key string) (string, string) {
	tokens := getTokens(key)
	token := ""

	key = ""
	if len(tokens) > 0 {
		key = strings.Join(tokens[1:], ".")
		token = tokens[0]
	}

	// token, full key
	return token, key
}

func GetValueWithSuperKey(data any, key string) any {
	var token string
	KeepVoid(token)

	val := PassValueIndirectReflect(data)

	if !val.IsValid() {
		return nil
	}

	token, key = baseToken(key)

	var temp any
	switch val.Kind() {
	case reflect.Bool:
		temp = GetBoolValueReflect(val)
		break

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		temp = GetIntValueReflect(val)
		break

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		temp = GetUintValueReflect(val)
		break

	case reflect.Float32, reflect.Float64:
		temp = GetFloatValueReflect(val)
		break

	case reflect.Complex64, reflect.Complex128:
		temp = GetComplexValueReflect(val)
		break

	case reflect.String:
		temp = ToStringReflect(val)
		break

	case reflect.Array, reflect.Slice:
		temp = GetArrayValueReflect(val)
		break

	case reflect.Map:
		temp = GetMapStrAnyValueReflect(val)
		break

	case reflect.Struct:
		temp = val.Interface()
		break

	default:
		temp = nil
	}

	if token != "" {
		temp = ParseValueReflect(temp, token)
	}

	if key != "" {
		temp = GetValueWithSuperKey(temp, key)
	}

	return temp
}

func GetValueWithSuperKeyReflect(data any, key string) reflect.Value {
	var err error
	var idx int
	var token string
	KeepVoid(err, idx, token)

	// ensure the field is set immediately, first order
	defaultValueReflect(data)

	val := PassValueIndirectReflect(data)

	if !val.IsValid() {
		return reflect.Value{}
	}

	token, key = baseToken(key)

	var temp reflect.Value
	switch val.Kind() {
	case reflect.Array, reflect.Slice, reflect.String:
		if idx, err = strconv.Atoi(token); err != nil {
			panic("invalid key")
		}

		temp = val.Index(idx)
		break

	case reflect.Map:
		temp = val.MapIndex(GetValueReflect(token))
		break

	case reflect.Struct:
		options := NewForEachStructFieldsOptions()

		// for each struct field
		NoErr(ForEachStructFieldsReflect(val, options, func(name string, sFieldX StructFieldExImpl) error {
			if name != token {
				return nil
			}

			temp = sFieldX.GetValue()
			return ForEachStop
		}))
		break

	default:
		temp = reflect.Value{}
	}

	if key != "" {
		temp = GetValueWithSuperKeyReflect(temp, key)
	}

	return temp
}

func SetValueWithSuperKeyReflect(data any, key string, value any) bool {
	KeepVoid(data, key, value)

	// TODO: implement it

	// super key -> parent key + base key
	// using SetIndex, SetMapIndex, Field.Set

	panic("not implemented yet")
}

func defaultValueReflect(value any) any {
	typ := PassTypeIndirectReflect(value)
	val := GetValueReflect(value)

	if !val.IsValid() {
		panic("invalid value")
	}

	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Slice:
		if val.IsNil() {
			val.Set(reflect.New(typ).Elem())
		}
		return nil

	case reflect.Interface, reflect.Pointer:
		if val.IsNil() {
			val.Set(reflect.New(typ))
		}
		return nil

	default:
		return nil
	}
}

func SetValueReflect(field any, value any) error {
	// ensure the field is set immediately, first order
	defaultValueReflect(field)

	val := PassValueIndirectReflect(value)
	f := PassValueIndirectReflect(field)

	if !f.IsValid() {
		return errors.New("field is not valid")
	}

	if !f.CanSet() {
		return errors.New("value is not addressable")
	}

	if !val.IsValid() {
		f.SetZero()
		return nil
	}

	switch f.Kind() {
	case reflect.Bool:
		f.SetBool(GetBoolValueReflect(val))
		return nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		f.SetInt(GetIntValueReflect(val))
		return nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		f.SetUint(GetUintValueReflect(val))
		return nil

	case reflect.Float32, reflect.Float64:
		f.SetFloat(GetFloatValueReflect(val))
		return nil

	case reflect.Complex64, reflect.Complex128:
		f.SetComplex(GetComplexValueReflect(val))
		return nil

	case reflect.String:
		f.SetString(ToStringReflect(val))
		return nil

	case reflect.Array, reflect.Slice:
		temp := GetValueReflect(GetArrayValueReflect(val))
		if temp.Kind() != reflect.Array && temp.Kind() != reflect.Slice {
			return errors.New("value type does not match field type")
		}

		// Convert an interface array to slice
		// Initialize the slice with the same length as the original array
		f.Set(reflect.MakeSlice(f.Type(), temp.Len(), temp.Len()))

		// Iterate through the original array and set each element in the new slice,
		// converting it to the appropriate type
		for i := 0; i < temp.Len(); i++ {
			if err := SetValueReflect(f.Index(i), temp.Index(i)); err != nil {
				return err
			}
		}

		return nil

	case reflect.Map:
		// map[string]any
		temp := GetValueReflect(GetMapStrAnyValueReflect(val))
		if !TypeEqualsReflect(f, temp) {
			return errors.New("value type does not match field type")
		}

		f.Set(temp)
		return nil

	case reflect.Struct:
		if TypeEqualsReflect(f, val) {
			f.Set(val)
			return nil
		}

		temp := GetMapStrAnyValueReflect(val)

		if temp != nil {
			options := NewForEachStructFieldsOptions()

			return ForEachStructFieldsReflect(f, options, func(name string, sFieldX StructFieldExImpl) error {

				return SetValueReflect(sFieldX.GetValue(), GetMapValue(temp, name))
			})
		}

		return errors.New("invalid data type")

	default:
		f.Set(val)
		return nil
	}
}

type StructFieldExImpl interface {
	GetStructField() reflect.StructField
	GetStructTagEx() StructTagExImpl
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

type StructFieldEx struct {
	structField reflect.StructField
	structTagEx StructTagExImpl
	value       reflect.Value
}

func NewStructFieldEx(field reflect.StructField, sTagX StructTagExImpl, value reflect.Value) StructFieldExImpl {
	return &StructFieldEx{
		structField: field,
		structTagEx: sTagX,
		value:       value,
	}
}

func (s *StructFieldEx) GetStructField() reflect.StructField {
	return s.structField
}

func (s *StructFieldEx) GetStructTagEx() StructTagExImpl {
	return s.structTagEx
}

func (s *StructFieldEx) GetValue() reflect.Value {
	return s.value
}

func (s *StructFieldEx) GetName() string {
	return s.structField.Name
}

func (s *StructFieldEx) GetPkgPath() string {
	return s.structField.PkgPath
}

func (s *StructFieldEx) GetType() reflect.Type {
	return s.structField.Type
}

func (s *StructFieldEx) GetTag() reflect.StructTag {
	return s.structField.Tag
}

func (s *StructFieldEx) GetOffset() uintptr {
	return s.structField.Offset
}

func (s *StructFieldEx) GetIndex() []int {
	return s.structField.Index
}

func (s *StructFieldEx) GetAnonymous() bool {
	return s.structField.Anonymous
}

func (s *StructFieldEx) IsExported() bool {
	return s.structField.IsExported()
}

func (s *StructFieldEx) Seq() iter.Seq[reflect.Value] {
	return s.value.Seq()
}

func (s *StructFieldEx) Seq2() iter.Seq2[reflect.Value, reflect.Value] {
	return s.value.Seq2()
}

func (s *StructFieldEx) Addr() reflect.Value {
	return s.value.Addr()
}

func (s *StructFieldEx) Bool() bool {
	return s.value.Bool()
}

func (s *StructFieldEx) Bytes() []byte {
	return s.value.Bytes()
}

func (s *StructFieldEx) CanAddr() bool {
	return s.value.CanAddr()
}

func (s *StructFieldEx) CanSet() bool {
	return s.value.CanSet()
}

func (s *StructFieldEx) Call(in []reflect.Value) []reflect.Value {
	return s.value.Call(in)
}

func (s *StructFieldEx) CallSlice(in []reflect.Value) []reflect.Value {
	return s.value.CallSlice(in)
}

func (s *StructFieldEx) Cap() int {
	return s.value.Cap()
}

func (s *StructFieldEx) Close() {
	s.value.Close()
}

func (s *StructFieldEx) CanComplex() bool {
	return s.value.CanComplex()
}

func (s *StructFieldEx) Complex() complex128 {
	return s.value.Complex()
}

func (s *StructFieldEx) Elem() reflect.Value {
	return s.value.Elem()
}

func (s *StructFieldEx) Field(i int) reflect.Value {
	return s.value.Field(i)
}

func (s *StructFieldEx) FieldByIndex(index []int) reflect.Value {
	return s.value.FieldByIndex(index)
}

func (s *StructFieldEx) FieldByIndexErr(index []int) (reflect.Value, error) {
	return s.value.FieldByIndexErr(index)
}

func (s *StructFieldEx) FieldByName(name string) reflect.Value {
	return s.value.FieldByName(name)
}

func (s *StructFieldEx) FieldByNameFunc(match func(string) bool) reflect.Value {
	return s.value.FieldByNameFunc(match)
}

func (s *StructFieldEx) CanFloat() bool {
	return s.value.CanFloat()
}

func (s *StructFieldEx) Float() float64 {
	return s.value.Float()
}

func (s *StructFieldEx) Index(i int) reflect.Value {
	return s.value.Index(i)
}

func (s *StructFieldEx) CanInt() bool {
	return s.value.CanInt()
}

func (s *StructFieldEx) Int() int64 {
	return s.value.Int()
}

func (s *StructFieldEx) CanInterface() bool {
	return s.value.CanInterface()
}

func (s *StructFieldEx) Interface() (i any) {
	return s.value.Interface()
}

func (s *StructFieldEx) IsNil() bool {
	return s.value.IsNil()
}

func (s *StructFieldEx) IsValid() bool {
	return s.value.IsValid()
}

func (s *StructFieldEx) IsZero() bool {
	return s.value.IsZero()
}

func (s *StructFieldEx) SetZero() {
	s.value.SetZero()
}

func (s *StructFieldEx) Kind() reflect.Kind {
	return s.value.Kind()
}

func (s *StructFieldEx) Len() int {
	return s.value.Len()
}

func (s *StructFieldEx) MapIndex(key reflect.Value) reflect.Value {
	return s.value.MapIndex(key)
}

func (s *StructFieldEx) MapKeys() []reflect.Value {
	return s.value.MapKeys()
}

func (s *StructFieldEx) SetIterKey(iter *reflect.MapIter) {
	s.value.SetIterKey(iter)
}

func (s *StructFieldEx) SetIterValue(iter *reflect.MapIter) {
	s.value.SetIterValue(iter)
}

func (s *StructFieldEx) MapRange() *reflect.MapIter {
	return s.value.MapRange()
}

func (s *StructFieldEx) Method(i int) reflect.Value {
	return s.value.Method(i)
}

func (s *StructFieldEx) NumMethod() int {
	return s.value.NumMethod()
}

func (s *StructFieldEx) MethodByName(name string) reflect.Value {
	return s.value.MethodByName(name)
}

func (s *StructFieldEx) NumField() int {
	return s.value.NumField()
}

func (s *StructFieldEx) OverflowComplex(x complex128) bool {
	return s.value.OverflowComplex(x)
}

func (s *StructFieldEx) OverflowFloat(x float64) bool {
	return s.value.OverflowFloat(x)
}

func (s *StructFieldEx) OverflowInt(x int64) bool {
	return s.value.OverflowInt(x)
}

func (s *StructFieldEx) OverflowUint(x uint64) bool {
	return s.value.OverflowUint(x)
}

func (s *StructFieldEx) Pointer() uintptr {
	return s.value.Pointer()
}

func (s *StructFieldEx) Recv() (x reflect.Value, ok bool) {
	return s.value.Recv()
}

func (s *StructFieldEx) Send(x reflect.Value) {
	s.value.Send(x)
}

func (s *StructFieldEx) Set(x reflect.Value) {
	s.value.Set(x)
}

func (s *StructFieldEx) SetBool(x bool) {
	s.value.SetBool(x)
}

func (s *StructFieldEx) SetBytes(x []byte) {
	s.value.SetBytes(x)
}

func (s *StructFieldEx) SetComplex(x complex128) {
	s.value.SetComplex(x)
}

func (s *StructFieldEx) SetFloat(x float64) {
	s.value.SetFloat(x)
}

func (s *StructFieldEx) SetInt(x int64) {
	s.value.SetInt(x)
}

func (s *StructFieldEx) SetLen(n int) {
	s.value.SetLen(n)
}

func (s *StructFieldEx) SetCap(n int) {
	s.value.SetCap(n)
}

func (s *StructFieldEx) SetMapIndex(key reflect.Value, elem reflect.Value) {
	s.value.SetMapIndex(key, elem)
}

func (s *StructFieldEx) SetUint(x uint64) {
	s.value.SetUint(x)
}

func (s *StructFieldEx) SetPointer(x unsafe.Pointer) {
	s.value.SetPointer(x)
}

func (s *StructFieldEx) SetString(x string) {
	s.value.SetString(x)
}

func (s *StructFieldEx) Slice(i int, j int) reflect.Value {
	return s.value.Slice(i, j)
}

func (s *StructFieldEx) Slice3(i int, j int, k int) reflect.Value {
	return s.value.Slice3(i, j, k)
}

func (s *StructFieldEx) String() string {
	return s.value.String()
}

func (s *StructFieldEx) TryRecv() (x reflect.Value, ok bool) {
	return s.value.TryRecv()
}

func (s *StructFieldEx) TrySend(x reflect.Value) bool {
	return s.value.TrySend(x)
}

func (s *StructFieldEx) Type() reflect.Type {
	return s.value.Type()
}

func (s *StructFieldEx) CanUint() bool {
	return s.value.CanUint()
}

func (s *StructFieldEx) Uint() uint64 {
	return s.value.Uint()
}

func (s *StructFieldEx) UnsafeAddr() uintptr {
	return s.value.UnsafeAddr()
}

func (s *StructFieldEx) UnsafePointer() unsafe.Pointer {
	return s.value.UnsafePointer()
}

func (s *StructFieldEx) Grow(n int) {
	s.value.Grow(n)
}

func (s *StructFieldEx) Clear() {
	s.value.Clear()
}

func (s *StructFieldEx) Convert(t reflect.Type) reflect.Value {
	return s.value.Convert(t)
}

func (s *StructFieldEx) CanConvert(t reflect.Type) bool {
	return s.value.CanConvert(t)
}

func (s *StructFieldEx) Comparable() bool {
	return s.value.Comparable()
}

func (s *StructFieldEx) Equal(u reflect.Value) bool {
	return s.value.Equal(u)
}

type StructTagExImpl interface {
	GetIgnore() bool
	GetOmitEmpty() bool
	GetRequired() bool
	GetName() string
	Unmatch(value any) bool
	Match(value any) bool
}

type StructTagEx struct {
	ignore    bool
	omitEmpty bool
	required  bool
	name      string
}

func (s *StructTagEx) GetIgnore() bool {
	return s.ignore
}

func (s *StructTagEx) GetOmitEmpty() bool {
	return s.omitEmpty
}

func (s *StructTagEx) GetRequired() bool {
	return s.required
}

func (s *StructTagEx) GetName() string {
	return s.name
}

func (s *StructTagEx) Unmatch(value any) bool {
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

func (s *StructTagEx) Match(value any) bool {
	return !s.Unmatch(value)
}

func GetStructTagEx(key string, sTag reflect.StructTag) StructTagExImpl {
	// skip tag name 'db'
	switch strings.TrimSpace(key) {
	case "mapstructure", "json", "yaml":
		if val, ok := sTag.Lookup(key); ok {
			val = strings.TrimSpace(val)
			tokens := strings.Split(val, ",")
			if len(tokens) > 0 {
				var name string
				var ignore, omitEmpty, required bool
				for i, token := range tokens {
					KeepVoid(i)

					token = strings.TrimSpace(token)
					switch token {
					case "-", "ignore":
						ignore = true
						break

					case "omitempty":
						omitEmpty = true
						break

					case "required":
						required = true
						break

					default:
						name = token
					}
				}

				return &StructTagEx{
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

type StructFieldAction func(name string, sFieldX StructFieldExImpl) error

func (s StructFieldAction) Call(name string, sFieldX StructFieldExImpl) error {
	return s(name, sFieldX)
}

type ForEachStructFieldsOptions struct {
	// Validation indicates whether the field should be matched.
	// Validation is matching ignored, omitempty, required attributes.
	Validation bool `mapstructure:"validation" json:"validation" yaml:"validation"`
}

func NewForEachStructFieldsOptions() *ForEachStructFieldsOptions {
	return &ForEachStructFieldsOptions{
		Validation: false,
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

			var sTagX StructTagExImpl

			for j, sTagName := range sTagNames {
				KeepVoid(j)

				if sTagX = GetStructTagEx(sTagName, sTag); sTagX != nil {
					pName = sTagX.GetName()
					pValid = sTagX.Match(vField)
					break
				}
			}

			if !pValid && options.Validation {
				continue
			}

			sFieldX := NewStructFieldEx(sField, sTagX, vField)
			if err := action.Call(pName, sFieldX); err != nil && !errors.Is(err, ForEachStop) {
				return err
			}
		}
		return nil
	}

	return errors.New("value is not struct")
}
