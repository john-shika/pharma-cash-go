package utils

import (
	"github.com/google/uuid"
	"nokowebapi/nokocore"
	"reflect"
)

func SqlValueIsNull(value any) bool {
	if value != nil {
		val := nokocore.PassValueIndirectReflect(value)
		if !val.IsValid() {
			return true
		}

		switch val.Kind() {
		case reflect.Struct:
			valid := true
			options := nokocore.NewForEachStructFieldsOptions()
			nokocore.NoErr(nokocore.ForEachStructFieldsReflect(value, options, func(name string, sFieldX nokocore.StructFieldExImpl) error {
				if name == "Valid" && sFieldX.Kind() == reflect.Bool && !sFieldX.Bool() {
					valid = false
					return nil
				}

				return nil
			}))

			return !valid

		case reflect.Array, reflect.Slice:
			size := val.Len()
			elem := val.Type().Elem()

			// fix uuid array
			switch elem.Kind() {
			case reflect.Uint8:
				if size == 16 {
					if UUID, ok := val.Interface().(uuid.UUID); ok {
						return UUID == uuid.Nil
					}
				}

			default:
				break
			}

			return false

		default:
			return false
		}
	}

	return true
}
