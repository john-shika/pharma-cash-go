package nokocore

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ShikaObjectDataTypeKind int

const (
	ShikaObjectDataTypeUndefined ShikaObjectDataTypeKind = iota
	ShikaObjectDataTypeNull
	ShikaObjectDataTypeBool
	ShikaObjectDataTypeInt
	ShikaObjectDataTypeUint
	ShikaObjectDataTypeInt8
	ShikaObjectDataTypeUint8
	ShikaObjectDataTypeInt16
	ShikaObjectDataTypeUint16
	ShikaObjectDataTypeInt32
	ShikaObjectDataTypeUint32
	ShikaObjectDataTypeInt64
	ShikaObjectDataTypeUint64
	ShikaObjectDataTypeUintptr
	ShikaObjectDataTypeFloat
	ShikaObjectDataTypeFloat32
	ShikaObjectDataTypeFloat64
	ShikaObjectDataTypeComplex64
	ShikaObjectDataTypeComplex128
	ShikaObjectDataTypeString
	ShikaObjectDataTypeArray
	ShikaObjectDataTypeObject
	ShikaObjectDataTypeNamespace
	ShikaObjectDataTypeInterface
	ShikaObjectDataTypeStruct
	ShikaObjectDataTypeClass
	ShikaObjectDataTypeFunction
	ShikaObjectDataTypeAttribute
	ShikaObjectDataTypeTime
)

func (shikaObjectDataTypeKind ShikaObjectDataTypeKind) ToString() string {
	switch shikaObjectDataTypeKind {
	case ShikaObjectDataTypeUndefined:
		return "undefined"

	case ShikaObjectDataTypeNull:
		return "null"

	case ShikaObjectDataTypeBool:
		return "bool"

	case ShikaObjectDataTypeInt:
		return "int"

	case ShikaObjectDataTypeUint:
		return "uint"

	case ShikaObjectDataTypeInt8:
		return "int8"

	case ShikaObjectDataTypeUint8:
		return "uint8"

	case ShikaObjectDataTypeInt16:
		return "int16"

	case ShikaObjectDataTypeUint16:
		return "uint16"

	case ShikaObjectDataTypeInt32:
		return "int32"

	case ShikaObjectDataTypeUint32:
		return "uint32"

	case ShikaObjectDataTypeInt64:
		return "int64"

	case ShikaObjectDataTypeUint64:
		return "uint64"

	case ShikaObjectDataTypeFloat:
		return "float"

	case ShikaObjectDataTypeFloat32:
		return "float32"

	case ShikaObjectDataTypeFloat64:
		return "float64"

	case ShikaObjectDataTypeComplex64:
		return "complex64"

	case ShikaObjectDataTypeComplex128:
		return "complex128"

	case ShikaObjectDataTypeString:
		return "string"

	case ShikaObjectDataTypeArray:
		return "array"

	case ShikaObjectDataTypeObject:
		return "object"

	case ShikaObjectDataTypeNamespace:
		return "namespace"

	case ShikaObjectDataTypeInterface:
		return "interface"

	case ShikaObjectDataTypeStruct:
		return "struct"

	case ShikaObjectDataTypeClass:
		return "class"

	case ShikaObjectDataTypeFunction:
		return "function"

	case ShikaObjectDataTypeAttribute:
		return "attribute"

	case ShikaObjectDataTypeTime:
		return "time"

	default:
		return "unknown"
	}
}

type ShikaHandleGetterFunc func() any
type ShikaHandleSetterFunc func(value any)

type ShikaObjectPropertyImpl interface {
	GetValue() any
	GetKind() ShikaObjectDataTypeKind
	SetValue(value any)
	IsConfigurable() bool
	IsEnumerable() bool
	IsWritable() bool
	IsValid() bool
}

type ShikaObjectProperty struct {
	Value        any
	Kind         ShikaObjectDataTypeKind
	Get          ShikaHandleGetterFunc
	Set          ShikaHandleSetterFunc
	Configurable bool
	Enumerable   bool
	Writable     bool
}

func NewShikaObjectProperty(value any, kind ShikaObjectDataTypeKind) ShikaObjectPropertyImpl {
	return &ShikaObjectProperty{
		Value:        value,
		Kind:         kind,
		Get:          nil,
		Set:          nil,
		Configurable: true,
		Enumerable:   true,
		Writable:     true,
	}
}

func (shikaObjectProperty *ShikaObjectProperty) GetValue() any {

	// calling getter function
	if shikaObjectProperty.Get != nil {
		return shikaObjectProperty.Get()
	}

	// get value directly
	return shikaObjectProperty.Value
}

func (shikaObjectProperty *ShikaObjectProperty) GetKind() ShikaObjectDataTypeKind {
	return shikaObjectProperty.Kind
}

func (shikaObjectProperty *ShikaObjectProperty) SetValue(value any) {

	// calling setter function
	if shikaObjectProperty.Set != nil {
		shikaObjectProperty.Set(value)
		return
	}

	// set value directly
	shikaObjectProperty.Value = value
}

func (shikaObjectProperty *ShikaObjectProperty) IsConfigurable() bool {
	return shikaObjectProperty.Configurable
}

func (shikaObjectProperty *ShikaObjectProperty) IsEnumerable() bool {
	return shikaObjectProperty.Enumerable
}

func (shikaObjectProperty *ShikaObjectProperty) IsWritable() bool {
	return shikaObjectProperty.Writable
}

func (shikaObjectProperty *ShikaObjectProperty) IsValid() bool {
	if shikaObjectProperty.Kind == ShikaObjectDataTypeString {
		if v, ok := Cast[string](shikaObjectProperty.Value); ok {
			return v != ""
		}
		return false
	}
	return shikaObjectProperty.Kind != ShikaObjectDataTypeUndefined &&
		shikaObjectProperty.Kind != ShikaObjectDataTypeNull
}

type ShikaObjectAttributeImpl interface {
	GetName() string
	GetParametersLength() int
	GetParameters() []any
}

type ShikaObjectAttribute struct {
	Name       string
	Parameters []any
}

func NewShikaObjectAttribute(name string, parameters ...any) ShikaObjectAttributeImpl {
	return &ShikaObjectAttribute{
		Name:       name,
		Parameters: parameters,
	}
}

func (shikaObjectAttribute *ShikaObjectAttribute) GetName() string {
	return shikaObjectAttribute.Name
}

func (shikaObjectAttribute *ShikaObjectAttribute) GetParametersLength() int {
	return len(shikaObjectAttribute.Parameters)
}

func (shikaObjectAttribute *ShikaObjectAttribute) GetParameters() []any {
	return shikaObjectAttribute.Parameters
}

type ShikaVarObjectImpl interface {
	GetName() string
	GetOwnProperty() ShikaObjectPropertyImpl
	GetProperties() []ShikaVarObjectImpl
	SetOwnProperty(property ShikaObjectPropertyImpl)
	SetProperties(properties []ShikaVarObjectImpl)
	PropertiesLength() int
	GetPropertyKeys() []string
	GetPropertyValues() []ShikaObjectPropertyImpl
	HasPropertyKey(key string) bool
	ContainPropertyKeys(keys ...string) bool
	GetPropertyByName(name string) ShikaObjectPropertyImpl
	SetPropertyByName(name string, property ShikaObjectPropertyImpl)
	RemovePropertyByName(name string)
	GetAttributesLength() int
	GetAttributes() []ShikaObjectAttributeImpl
	SetAttributes(attributes []ShikaObjectAttributeImpl)
	HasAttributeByName(name string) bool
	ContainAttributeNames(names ...string) bool
	GetAttributeByName(name string) ShikaObjectAttributeImpl
	SetAttributeByName(name string, attribute ShikaObjectAttributeImpl)
	RemoveAttributeByName(name string)
}

type ShikaVarObject struct {
	Name        string
	OwnProperty ShikaObjectPropertyImpl
	Properties  []ShikaVarObjectImpl
	Attributes  []ShikaObjectAttributeImpl
}

func NewShikaVarObject(name string) ShikaVarObjectImpl {
	return &ShikaVarObject{
		Name:        name,
		OwnProperty: nil,
		Properties:  []ShikaVarObjectImpl{},
		Attributes:  []ShikaObjectAttributeImpl{},
	}
}

func (shikaVarObject *ShikaVarObject) GetName() string {
	return shikaVarObject.Name
}

func (shikaVarObject *ShikaVarObject) GetOwnProperty() ShikaObjectPropertyImpl {
	return shikaVarObject.OwnProperty
}

func (shikaVarObject *ShikaVarObject) GetProperties() []ShikaVarObjectImpl {
	return shikaVarObject.Properties
}

func (shikaVarObject *ShikaVarObject) SetOwnProperty(property ShikaObjectPropertyImpl) {
	shikaVarObject.OwnProperty = property
}

func (shikaVarObject *ShikaVarObject) SetProperties(properties []ShikaVarObjectImpl) {
	shikaVarObject.Properties = properties
}

func (shikaVarObject *ShikaVarObject) PropertiesLength() int {
	return len(shikaVarObject.Properties)
}

func (shikaVarObject *ShikaVarObject) GetPropertyKeys() []string {
	size := len(shikaVarObject.Properties)
	keys := make([]string, 0, size)
	for i, shikaVarObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaVarObj)

		keys = append(keys, shikaVarObj.GetName())
	}
	return keys
}

func (shikaVarObject *ShikaVarObject) GetPropertyValues() []ShikaObjectPropertyImpl {
	size := len(shikaVarObject.Properties)
	values := make([]ShikaObjectPropertyImpl, 0, size)
	for i, shikaVarObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaVarObj)

		values = append(values, shikaVarObj.GetOwnProperty())
	}
	return values
}

func (shikaVarObject *ShikaVarObject) HasPropertyKey(key string) bool {
	for i, shikaVarObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaVarObj)

		if shikaVarObj.GetName() == key {
			return true
		}
	}
	return false
}

func (shikaVarObject *ShikaVarObject) ContainPropertyKeys(keys ...string) bool {
	for i, key := range keys {
		KeepVoid(i, key)

		if !shikaVarObject.HasPropertyKey(key) {
			return false
		}
	}
	return true
}

func (shikaVarObject *ShikaVarObject) GetPropertyByName(name string) ShikaObjectPropertyImpl {
	for i, shikaVarObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaVarObj)

		if shikaVarObj.GetName() == name {
			return shikaVarObj.GetOwnProperty()
		}
	}
	return nil
}

func (shikaVarObject *ShikaVarObject) SetPropertyByName(name string, property ShikaObjectPropertyImpl) {

	// replace existing property
	for i, shikaObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaObj)

		if shikaObj.GetName() == name {
			shikaVarObject.Properties[i].SetOwnProperty(property)
			return
		}
	}

	// create new property
	shikaVarObj := NewShikaVarObject(name)
	shikaVarObj.SetOwnProperty(property)
	shikaVarObject.Properties = append(shikaVarObject.Properties, shikaVarObj)
}

func (shikaVarObject *ShikaVarObject) RemovePropertyByName(name string) {
	for i, shikaVarObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaVarObj)

		if shikaVarObj.GetName() == name {
			j := i + 1
			if j <= len(shikaVarObject.Properties) {
				temp := shikaVarObject.Properties[j:]
				shikaVarObject.Properties = append(shikaVarObject.Properties[:i], temp...)
				return
			}
			shikaVarObject.Properties = shikaVarObject.Properties[:i]
			return
		}
	}
}

func (shikaVarObject *ShikaVarObject) GetAttributesLength() int {
	return len(shikaVarObject.Attributes)
}

func (shikaVarObject *ShikaVarObject) GetAttributes() []ShikaObjectAttributeImpl {
	return shikaVarObject.Attributes
}

func (shikaVarObject *ShikaVarObject) SetAttributes(attributes []ShikaObjectAttributeImpl) {
	shikaVarObject.Attributes = attributes
}

func (shikaVarObject *ShikaVarObject) HasAttributeByName(name string) bool {
	for i, shikaObjAttr := range shikaVarObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			return true
		}
	}
	return false
}

func (shikaVarObject *ShikaVarObject) ContainAttributeNames(names ...string) bool {
	for i, name := range names {
		KeepVoid(i, name)

		if !shikaVarObject.HasAttributeByName(name) {
			return false
		}
	}
	return true
}

func (shikaVarObject *ShikaVarObject) GetAttributeByName(name string) ShikaObjectAttributeImpl {
	for i, shikaObjAttr := range shikaVarObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			return shikaObjAttr
		}
	}
	return nil
}

func (shikaVarObject *ShikaVarObject) SetAttributeByName(name string, attribute ShikaObjectAttributeImpl) {

	// replace existing attribute
	for i, shikaObjAttr := range shikaVarObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			shikaVarObject.Attributes[i] = attribute
			return
		}
	}

	// create new attribute
	shikaVarObject.Attributes = append(shikaVarObject.Attributes, attribute)
}

func (shikaVarObject *ShikaVarObject) RemoveAttributeByName(name string) {
	for i, shikaObjAttr := range shikaVarObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			j := i + 1
			if j <= len(shikaVarObject.Attributes) {
				temp := shikaVarObject.Attributes[j:]
				shikaVarObject.Attributes = append(shikaVarObject.Attributes[:i], temp...)
				return
			}
			shikaVarObject.Attributes = shikaVarObject.Attributes[:i]
			return
		}
	}
}

func PassPtrShikaObjectPropertyReflect(value any) ShikaObjectPropertyImpl {
	val := GetValueReflect(value)
	if !val.IsValid() {
		return nil
	}
	if TypeEqualsReflect(val, new(ShikaObjectProperty)) {
		switch val.Kind() {
		case reflect.Pointer:
			return val.Interface().(ShikaObjectPropertyImpl)

		case reflect.Struct:
			// warning: ensure to use a proper pointer instead of a pseudo struct pointer.
			// convert the interface value to ShikaObjectProperty and return its pointer.
			fmt.Println("[WARN] Use a generic pointer instead of creating a new struct pointer with reflection.")
			temp := MakePointerReflect(val)
			return temp.Interface().(ShikaObjectPropertyImpl)

		default:
			panic("invalid data type")
		}
	}

	return nil
}

func GetShikaObjectProperty(obj any) ShikaObjectPropertyImpl {
	if obj == nil {
		return NewShikaObjectProperty(nil, ShikaObjectDataTypeNull)
	}

	if val, ok := obj.(StringableImpl); ok {
		return NewShikaObjectProperty(val.ToString(), ShikaObjectDataTypeString)
	}

	val := PassValueIndirectReflect(obj)
	if !val.IsValid() {
		return NewShikaObjectProperty(nil, ShikaObjectDataTypeNull)
	}

	switch val.Kind() {
	case reflect.Bool:
		return NewShikaObjectProperty(val.Bool(), ShikaObjectDataTypeBool)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return NewShikaObjectProperty(val.Int(), ShikaObjectDataTypeInt64)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return NewShikaObjectProperty(val.Uint(), ShikaObjectDataTypeUint64)

	case reflect.Uintptr:
		return NewShikaObjectProperty(val.Uint(), ShikaObjectDataTypeUintptr)

	case reflect.Float32, reflect.Float64:
		return NewShikaObjectProperty(val.Float(), ShikaObjectDataTypeFloat64)

	case reflect.Complex64, reflect.Complex128:
		return NewShikaObjectProperty(val.Complex(), ShikaObjectDataTypeComplex128)

	case reflect.String:
		return NewShikaObjectProperty(val.String(), ShikaObjectDataTypeString)

	case reflect.Struct:

		// ensure to use a proper pointer instead of a pseudo struct pointer.
		if shikaObjectProperty := PassPtrShikaObjectPropertyReflect(obj); shikaObjectProperty != nil {
			return shikaObjectProperty
		}

		// TODO: storing all converters in another struct

		/// START CONVERTERS

		if IsTimeUtcISO8601(val) {
			return NewShikaObjectProperty(ToTimeUtcStringISO8601(val), ShikaObjectDataTypeTime)
		}

		if IsURL(val) {
			return NewShikaObjectProperty(ToURLString(val), ShikaObjectDataTypeString)
		}

		if IsUUID(val) {
			return NewShikaObjectProperty(ToUUIDString(val), ShikaObjectDataTypeString)
		}

		if IsIP(val) {
			return NewShikaObjectProperty(ToIPString(val), ShikaObjectDataTypeString)
		}

		/// END CONVERTERS

		var temp []ShikaVarObjectImpl

		// create foreach struct fields options
		options := NewForEachStructFieldsOptions()
		options.Validation = true

		// get struct fields
		NoErr(ForEachStructFieldsReflect(val, options, func(name string, sFieldX StructFieldExImpl) error {
			shikaVarObject := NewShikaVarObject(name)
			shikaObjectProperty := GetShikaObjectProperty(sFieldX.Interface())
			shikaVarObject.SetOwnProperty(shikaObjectProperty)
			temp = append(temp, shikaVarObject)
			return nil
		}))

		return NewShikaObjectProperty(temp, ShikaObjectDataTypeObject)

	case reflect.Array, reflect.Slice:
		size := val.Len()
		values := make([]ShikaObjectPropertyImpl, size)
		for i := 0; i < size; i++ {
			elem := val.Index(i).Interface()
			values[i] = GetShikaObjectProperty(elem)
		}
		return NewShikaObjectProperty(values, ShikaObjectDataTypeArray)

	case reflect.Map:
		size := val.Len()
		iter := val.MapRange()
		values := make([]ShikaVarObjectImpl, size)
		for i := 0; iter.Next(); i++ {
			key := ToStringReflect(iter.Key())
			value := iter.Value().Interface()
			temp := NewShikaVarObject(key)
			temp.SetOwnProperty(GetShikaObjectProperty(value))
			values[i] = temp
		}
		return NewShikaObjectProperty(values, ShikaObjectDataTypeObject)

	default:
		return NewShikaObjectProperty(nil, ShikaObjectDataTypeUndefined)
	}
}

func shikaJsonEncodeIndentPermutate(shikaObjectProperty ShikaObjectPropertyImpl, indent int, start int) string {
	var err error
	var t time.Time
	KeepVoid(err, t)

	end := start + indent
	whiteSpaceStart := strings.Repeat(" ", start)
	whiteSpaceEnd := strings.Repeat(" ", end)
	KeepVoid(end, whiteSpaceStart, whiteSpaceEnd)

	if shikaObjectProperty == nil {
		return "undefined"
	}

	switch shikaObjectProperty.GetKind() {
	case ShikaObjectDataTypeUndefined:
		return "undefined"

	case ShikaObjectDataTypeNull:
		return "null"

	case ShikaObjectDataTypeBool:
		v := Unwrap(Cast[bool](shikaObjectProperty.GetValue()))
		return strconv.FormatBool(v)

	case ShikaObjectDataTypeInt:
		// TODO: not implemented yet
		v := Unwrap(Cast[int](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)

	case ShikaObjectDataTypeUint:
		// TODO: not implemented yet
		v := Unwrap(Cast[uint](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)

	case ShikaObjectDataTypeInt8:
		v := Unwrap(Cast[int8](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)

	case ShikaObjectDataTypeUint8:
		v := Unwrap(Cast[uint8](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)

	case ShikaObjectDataTypeInt16:
		v := Unwrap(Cast[int16](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)

	case ShikaObjectDataTypeUint16:
		v := Unwrap(Cast[uint16](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)

	case ShikaObjectDataTypeInt32:
		v := Unwrap(Cast[int32](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)

	case ShikaObjectDataTypeUint32:
		v := Unwrap(Cast[uint32](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)

	case ShikaObjectDataTypeInt64:
		v := Unwrap(Cast[int64](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(v, 10)

	case ShikaObjectDataTypeUint64:
		v := Unwrap(Cast[uint64](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(v, 10)

	case ShikaObjectDataTypeUintptr:
		v := Unwrap(Cast[uintptr](shikaObjectProperty.GetValue()))
		return strconv.Quote(ToStringReflect(v))

	case ShikaObjectDataTypeFloat:
		// TODO: not implemented yet
		v := Unwrap(Cast[float64](shikaObjectProperty.GetValue()))
		return strconv.FormatFloat(v, 'f', -1, 64)

	case ShikaObjectDataTypeFloat32:
		v := Unwrap(Cast[float32](shikaObjectProperty.GetValue()))
		return strconv.FormatFloat(float64(v), 'f', -1, 32)

	case ShikaObjectDataTypeFloat64:
		v := Unwrap(Cast[float64](shikaObjectProperty.GetValue()))
		return strconv.FormatFloat(v, 'f', -1, 64)

	case ShikaObjectDataTypeComplex64:
		v := Unwrap(Cast[complex64](shikaObjectProperty.GetValue()))
		return strconv.FormatComplex(complex128(v), 'f', -1, 64)

	case ShikaObjectDataTypeComplex128:
		v := Unwrap(Cast[complex128](shikaObjectProperty.GetValue()))
		return strconv.FormatComplex(v, 'f', -1, 128)

	case ShikaObjectDataTypeString:
		v := Unwrap(Cast[string](shikaObjectProperty.GetValue()))
		return strconv.Quote(v)

	case ShikaObjectDataTypeArray:
		shikaObjectProperties := Unwrap(Cast[[]ShikaObjectPropertyImpl](shikaObjectProperty.GetValue()))
		size := len(shikaObjectProperties)
		values := make([]string, size)
		for i := 0; i < size; i++ {
			elem := shikaObjectProperties[i]
			v := shikaJsonEncodeIndentPermutate(elem, indent, end)
			if len(v) > 0 {
				values[i] = whiteSpaceEnd + v
				continue
			}
			values[i] = whiteSpaceEnd + "undefined"
		}
		if len(values) > 0 {
			return "[\n" + strings.Join(values, ",\n") + "\n" + whiteSpaceStart + "]"
		}
		return "[]"

	case ShikaObjectDataTypeObject:
		shikaVarObjects := Unwrap(Cast[[]ShikaVarObjectImpl](shikaObjectProperty.GetValue()))
		size := len(shikaVarObjects)
		values := make([]string, size)
		for i := 0; i < size; i++ {
			elem := shikaVarObjects[i]
			k := strconv.Quote(ToCamelCase(elem.GetName()))
			v := elem.GetOwnProperty()
			values[i] = whiteSpaceEnd + k + ": " + shikaJsonEncodeIndentPermutate(v, indent, end)
		}
		if len(values) > 0 {
			return "{\n" + strings.Join(values, ",\n") + "\n" + whiteSpaceStart + "}"
		}
		return "{}"

	case ShikaObjectDataTypeTime:
		if t, err = GetTimeUtcISO8601(shikaObjectProperty.GetValue()); err != nil {
			return "undefined"
		}
		v := ToTimeUtcStringISO8601(t)
		return strconv.Quote(v)

	default:
		return "undefined"
	}
}

func ShikaJsonEncodeIndent(obj any, indent int) string {
	shikaObjectProperty := GetShikaObjectProperty(obj)
	return shikaJsonEncodeIndentPermutate(shikaObjectProperty, indent, 0)
}

func ShikaJsonEncode(obj any) string {
	return ShikaJsonEncodeIndent(obj, 4)
}

func shikaYamlEncodeIndentPermutate(shikaObjectProperty ShikaObjectPropertyImpl, indent int, start int) string {
	var err error
	var t time.Time
	KeepVoid(err, t)

	end := start + indent
	whiteSpaceStart := strings.Repeat(" ", start)
	whiteSpaceEnd := strings.Repeat(" ", end)
	KeepVoid(end, whiteSpaceStart, whiteSpaceEnd)

	header := "---\n"
	if start > 0 {
		header = "\n"
	}

	if shikaObjectProperty == nil {
		return "undefined"
	}

	switch shikaObjectProperty.GetKind() {
	case ShikaObjectDataTypeUndefined:
		return "undefined"

	case ShikaObjectDataTypeNull:
		return "null"

	case ShikaObjectDataTypeBool:
		v := Unwrap(Cast[bool](shikaObjectProperty.GetValue()))
		return strconv.FormatBool(v)

	case ShikaObjectDataTypeInt:
		// TODO: not implemented yet
		v := Unwrap(Cast[int](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)

	case ShikaObjectDataTypeUint:
		// TODO: not implemented yet
		v := Unwrap(Cast[uint](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)

	case ShikaObjectDataTypeInt8:
		v := Unwrap(Cast[int8](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)

	case ShikaObjectDataTypeUint8:
		v := Unwrap(Cast[uint8](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)

	case ShikaObjectDataTypeInt16:
		v := Unwrap(Cast[int16](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)

	case ShikaObjectDataTypeUint16:
		v := Unwrap(Cast[uint16](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)

	case ShikaObjectDataTypeInt32:
		v := Unwrap(Cast[int32](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)

	case ShikaObjectDataTypeUint32:
		v := Unwrap(Cast[uint32](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)

	case ShikaObjectDataTypeInt64:
		v := Unwrap(Cast[int64](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(v, 10)

	case ShikaObjectDataTypeUint64:
		v := Unwrap(Cast[uint64](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(v, 10)

	case ShikaObjectDataTypeUintptr:
		v := Unwrap(Cast[uintptr](shikaObjectProperty.GetValue()))
		return strconv.Quote(ToStringReflect(v))

	case ShikaObjectDataTypeFloat:
		// TODO: not implemented yet
		v := Unwrap(Cast[float64](shikaObjectProperty.GetValue()))
		return strconv.FormatFloat(v, 'f', -1, 64)

	case ShikaObjectDataTypeFloat32:
		v := Unwrap(Cast[float32](shikaObjectProperty.GetValue()))
		return strconv.FormatFloat(float64(v), 'f', -1, 32)

	case ShikaObjectDataTypeFloat64:
		v := Unwrap(Cast[float64](shikaObjectProperty.GetValue()))
		return strconv.FormatFloat(v, 'f', -1, 64)

	case ShikaObjectDataTypeComplex64:
		v := Unwrap(Cast[complex64](shikaObjectProperty.GetValue()))
		return strconv.FormatComplex(complex128(v), 'f', -1, 64)

	case ShikaObjectDataTypeComplex128:
		v := Unwrap(Cast[complex128](shikaObjectProperty.GetValue()))
		return strconv.FormatComplex(v, 'f', -1, 128)

	case ShikaObjectDataTypeString:
		v := Unwrap(Cast[string](shikaObjectProperty.GetValue()))
		return strconv.Quote(v)

	case ShikaObjectDataTypeArray:

		// modified white space for indentation string
		start = start - 2             // truncate the hyphen and whitespace
		start = Unwrap(Max(start, 0)) // truncate safety number
		end = start + indent
		end = end - 2             // truncate the hyphen and whitespace
		end = Unwrap(Max(end, 0)) // truncate safety number
		whiteSpaceStart = strings.Repeat(" ", start)
		whiteSpaceEnd = strings.Repeat(" ", end)
		KeepVoid(end, whiteSpaceStart, whiteSpaceEnd)

		shikaObjectProperties := Unwrap(Cast[[]ShikaObjectPropertyImpl](shikaObjectProperty.GetValue()))
		size := len(shikaObjectProperties)
		values := make([]string, size)
		for i := 0; i < size; i++ {
			elem := shikaObjectProperties[i]
			v := shikaYamlEncodeIndentPermutate(elem, indent, end)
			if len(v) > 0 {
				if v[0] == '\n' {
					s := len(whiteSpaceStart) + 3 // truncate the newline, hyphen and whitespace
					values[i] = whiteSpaceStart + "- " + v[s:]
					continue
				}
				values[i] = whiteSpaceStart + "- " + v
				continue
			}
			values[i] = whiteSpaceStart + "- " + "undefined"
		}
		if len(values) > 0 {
			return header + strings.Join(values, "\n")
		}
		return "[]"

	case ShikaObjectDataTypeObject:
		shikaVarObjects := Unwrap(Cast[[]ShikaVarObjectImpl](shikaObjectProperty.GetValue()))
		size := len(shikaVarObjects)
		values := make([]string, size)
		for i := 0; i < size; i++ {
			elem := shikaVarObjects[i]
			k := ToSnakeCase(elem.GetName())
			v := elem.GetOwnProperty()
			values[i] = whiteSpaceStart + k + ": " + shikaYamlEncodeIndentPermutate(v, indent, end)
		}
		if len(values) > 0 {
			return header + strings.Join(values, "\n")
		}
		return "{}"

	case ShikaObjectDataTypeTime:
		if t, err = GetTimeUtcISO8601(shikaObjectProperty.GetValue()); err != nil {
			return "undefined"
		}
		v := ToTimeUtcStringISO8601(t)
		return strconv.Quote(v)

	default:
		return "undefined"
	}
}

func ShikaYamlEncodeIndent(obj any, indent int) string {
	shikaObjectProperty := GetShikaObjectProperty(obj)
	return shikaYamlEncodeIndentPermutate(shikaObjectProperty, indent, 0)
}

func ShikaYamlEncode(obj any) string {
	return ShikaYamlEncodeIndent(obj, 4)
}
