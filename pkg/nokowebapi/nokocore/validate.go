package nokocore

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// TODO:: added ipaddr ipv4 ipv6, datetime, uuid

type ValidatorImpl interface {
	Validate(value any) error
}

type Validator struct {
	Void any
}

func NewValidator() ValidatorImpl {
	return &Validator{
		Void: nil,
	}
}

func (e *Validator) Validate(value any) error {
	return ValidateStruct(value)
}

type ValidateErrorImpl interface {
	Fields() []string
	Error() string
}

type ValidateError struct {
	fields []string
}

func NewValidateError(fields []string) ValidateErrorImpl {
	return &ValidateError{
		fields: fields,
	}
}

func (v *ValidateError) Fields() []string {
	return v.fields
}

func (v *ValidateError) Error() string {
	return strings.Join(v.fields, "\n")
}

func CheckPassword(password string) error {
	var temp []string
	KeepVoid(temp)

	size := len(password)
	if size < 8 {
		message := "password must be at least 8 characters long"
		temp = append(temp, message)
	}

	if !strings.ContainsAny(password, AlphaLower) {
		message := "password must contain at least one lowercase letter"
		temp = append(temp, message)
	}

	if !strings.ContainsAny(password, AlphaUpper) {
		message := "password must contain at least one uppercase letter"
		temp = append(temp, message)
	}

	if !strings.ContainsAny(password, Digits) {
		message := "password must contain at least one digit"
		temp = append(temp, message)
	}

	if !strings.ContainsAny(password, Punctuation) {
		message := "password must contain at least one special character"
		temp = append(temp, message)
	}

	if len(temp) > 0 {
		return NewValidateError(temp)
	}

	return nil
}

func ValidateStruct(value any) error {
	var ok bool
	var err error
	KeepVoid(ok, err)

	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Struct:
		options := NewForEachStructFieldsOptions()
		var errorStack []error
		err = ForEachStructFieldsReflect(val, options, func(name string, sFieldX StructFieldExImpl) error {
			err = func() error {
				vField := sFieldX.GetValue()
				sTag := sFieldX.GetTag()

				validate := sTag.Get("validate")
				tokens := strings.Split(validate, ",")

				ignore := false
				omitEmpty := false
				isBoolean := false
				isNumber := false
				isNumeric := false
				isAlpha := false
				isAscii := false
				isEmail := false
				isPhone := false
				isPassword := false

				// additional options for number or numeric
				var minimum float64
				var maximum float64

				setMinimum := false
				setMaximum := false

				for i, token := range tokens {
					KeepVoid(i)

					token = strings.TrimSpace(token)
					if token, ok = strings.CutPrefix(token, "min="); ok {
						if minimum, err = strconv.ParseFloat(token, 64); err != nil {
							return errors.New(fmt.Sprintf("invalid options for field '%s': min=%s", name, token))
						}

						setMinimum = true
						continue
					}

					if token, ok = strings.CutPrefix(token, "max="); ok {
						if maximum, err = strconv.ParseFloat(token, 64); err != nil {
							return errors.New(fmt.Sprintf("invalid options for field '%s': max=%s", name, token))
						}

						setMaximum = true
						continue
					}

					switch token {
					case "-", "ignore":
						ignore = true
						break

					case "boolean":
						isBoolean = true
						break

					case "number":
						isNumber = true
						break

					case "numeric":
						isNumeric = true
						break

					case "alpha":
						isAlpha = true
						break

					case "ascii":
						isAscii = true
						break

					case "email":
						isEmail = true
						break

					case "phone":
						isPhone = true
						break

					case "password":
						isPassword = true
						break

					case "omitempty":
						omitEmpty = true
						break
					}
				}

				if ignore {
					return nil
				}

				if IsNoneOrEmptyWhiteSpace(vField) {
					if !omitEmpty {
						return errors.New(fmt.Sprintf("field '%s' is required", name))
					}

					// skip
					return nil
				}

				if isBoolean {
					if err = ValidateBoolean(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isNumber {
					if err = ValidateNumber(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}

					num := ToInt(vField.Interface())

					minInt64 := int64(minimum)
					if setMinimum && num < minInt64 {
						return errors.New(fmt.Sprintf("field '%s' is less than %d", name, minInt64))
					}

					maxInt64 := int64(maximum)
					if setMaximum && num > maxInt64 {
						return errors.New(fmt.Sprintf("field '%s' is greater than %d", name, maxInt64))
					}
				}

				if isNumeric {
					if err = ValidateNumeric(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}

					num := ToFloat(vField.Interface())

					if setMinimum && num < minimum {
						temp := strconv.FormatFloat(minimum, 'f', -1, 64)
						return errors.New(fmt.Sprintf("field '%s' is less than %s", name, temp))
					}

					if setMaximum && num > maximum {
						temp := strconv.FormatFloat(maximum, 'f', -1, 64)
						return errors.New(fmt.Sprintf("field '%s' is greater than %s", name, temp))
					}
				}

				if isAlpha {
					if err = ValidateAlpha(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isAscii {
					if err = ValidateAscii(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isEmail {
					if err = ValidateEmail(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isPhone {
					if err = ValidatePhone(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isPassword {
					return ValidatePassword(vField)
				}

				switch vField.Kind() {
				case reflect.Struct:
					return ValidateStruct(vField)

				default:
					break
				}

				return nil
			}()

			if err != nil {
				errorStack = append(errorStack, err)
			}

			return nil
		})

		if err != nil {
			return errors.New(fmt.Sprintf("failed to validate struct: %s", err.Error()))
		}

		if len(errorStack) > 0 {
			var fields []string
			for i, err := range errorStack {
				KeepVoid(i)

				var validateErr *ValidateError
				if errors.As(err, &validateErr) {
					fields = append(fields, validateErr.Fields()...)
					continue
				}

				fields = append(fields, err.Error())
			}

			return NewValidateError(fields)
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateBoolean(value any) error {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Bool:
		return nil

	case reflect.String:
		if !BooleanRegex().MatchString(val.String()) {
			return errors.New("invalid boolean format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateNumber(value any) error {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return nil

	case reflect.Float32, reflect.Float64:
		return nil

	case reflect.Complex64, reflect.Complex128:
		return nil

	case reflect.String:
		if !NumberRegex().MatchString(val.String()) {
			return errors.New("invalid number format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateNumeric(value any) error {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return nil

	case reflect.Float32, reflect.Float64:
		return nil

	case reflect.Complex64, reflect.Complex128:
		return nil

	case reflect.String:
		if !NumericRegex().MatchString(val.String()) {
			return errors.New("invalid numeric format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAlpha(value any) error {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		if !AlphaRegex().MatchString(val.String()) {
			return errors.New("invalid alpha format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAscii(value any) error {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		if !AsciiRegex().MatchString(val.String()) {
			return errors.New("invalid ascii format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateEmail(value any) error {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		if !EmailRegex().MatchString(val.String()) {
			return errors.New("invalid email format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidatePhone(value any) error {
	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		if !PhoneRegex().MatchString(val.String()) {
			return errors.New("invalid phone format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidatePassword(value any) error {
	var err error
	KeepVoid(err)

	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		//if err = CheckPassword(val.String()); err != nil {
		//	var validateErr *ValidateError
		//	if errors.As(err, &validateErr) {
		//		var fields []string
		//		for i, field := range validateErr.Fields() {
		//			nokocore.KeepVoid(i)
		//
		//			// custom message error
		//			fields = append(fields, fmt.Sprintf("field 'password' is invalid value, %s", field))
		//		}
		//
		//		return NewValidateError(fields)
		//	}
		//}
		//
		//return nil
		return CheckPassword(val.String())

	default:
		return errors.New("invalid data type")
	}
}
