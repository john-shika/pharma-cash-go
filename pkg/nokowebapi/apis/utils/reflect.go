package utils

import (
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

		default:
			return false
		}
	}

	return true
}
