package sqlx

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"net/url"
	"nokowebapi/nokocore"
	"reflect"
	"strconv"
	"strings"
	"time"
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
	nokocore.KeepVoid(temp)

	size := len(password)
	if size < 8 {
		message := "password must be at least 8 characters long"
		temp = append(temp, message)
	}

	if !strings.ContainsAny(password, nokocore.AlphaLower) {
		message := "password must contain at least one lowercase letter"
		temp = append(temp, message)
	}

	if !strings.ContainsAny(password, nokocore.AlphaUpper) {
		message := "password must contain at least one uppercase letter"
		temp = append(temp, message)
	}

	if !strings.ContainsAny(password, nokocore.Digits) {
		message := "password must contain at least one digit"
		temp = append(temp, message)
	}

	if !strings.ContainsAny(password, nokocore.Punctuation) {
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
	nokocore.KeepVoid(ok, err)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Struct:
		options := nokocore.NewForEachStructFieldsOptions()
		var errorStack []error
		err = nokocore.ForEachStructFieldsReflect(val, options, func(name string, sFieldX nokocore.StructFieldExImpl) error {
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
				isAlphaLower := false
				isAlphaUpper := false
				isAlphabet := false
				isAlphaNum := false
				isAscii := false
				isEmail := false
				isPhone := false
				isPassword := false
				isDateTime := false
				isDateTimeISO8601 := false
				isDateOnly := false
				isTimeOnly := false
				isURL := false
				isUUID := false
				isDecimal := false

				// additional options for number or numeric
				var minimum float64
				var maximum float64

				setMinimum := false
				setMaximum := false

				for i, token := range tokens {
					nokocore.KeepVoid(i)

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

					case "alphaLower":
						isAlphaLower = true
						break

					case "alphaUpper":
						isAlphaUpper = true
						break

					case "alphabet":
						isAlphabet = true
						break

					case "alphaNum":
						isAlphaNum = true
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

					case "datetime":
						isDateTime = true
						break

					case "datetimeISO8601":
						isDateTimeISO8601 = true
						break

					case "dateOnly":
						isDateOnly = true
						break

					case "timeOnly":
						isTimeOnly = true
						break

					case "url":
						isURL = true
						break

					case "uuid":
						isUUID = true
						break

					case "decimal":
						isDecimal = true
						break

					case "omitempty":
						omitEmpty = true
						break
					}
				}

				if ignore {
					return nil
				}

				if nokocore.IsNoneOrEmptyWhiteSpace(vField) {
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

					num := nokocore.ToInt(vField.Interface())

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

					num := nokocore.ToFloat(vField.Interface())

					if setMinimum && num < minimum {
						temp := strconv.FormatFloat(minimum, 'f', -1, 64)
						return errors.New(fmt.Sprintf("field '%s' is less than %s", name, temp))
					}

					if setMaximum && num > maximum {
						temp := strconv.FormatFloat(maximum, 'f', -1, 64)
						return errors.New(fmt.Sprintf("field '%s' is greater than %s", name, temp))
					}
				}

				if isAlphaLower {
					if err = ValidateAlphaLower(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isAlphaUpper {
					if err = ValidateAlphaUpper(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isAlphabet {
					if err = ValidateAlphabet(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isAlphaNum {
					if err = ValidateAlphaNum(vField); err != nil {
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

				if isDateTime {
					if err = ValidateDateTime(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isDateTimeISO8601 {
					if err = ValidateDateTimeISO8601(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isDateOnly {
					if err = ValidateDateOnly(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isTimeOnly {
					if err = ValidateTimeOnly(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isURL {
					if err = ValidateURL(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isUUID {
					if err = ValidateUUID(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}
				}

				if isDecimal {
					if err = ValidateDecimal(vField); err != nil {
						return errors.New(fmt.Sprintf("field '%s' is %s", name, err.Error()))
					}

					num := nokocore.ToDecimal(vField.Interface())

					if setMinimum {
						minDecimal := nokocore.ToDecimal(minimum)
						if num.LessThan(minDecimal) {
							temp := strconv.FormatFloat(minimum, 'f', -1, 64)
							return errors.New(fmt.Sprintf("field '%s' is less than %s", name, temp))
						}
					}

					if setMaximum {
						maxDecimal := nokocore.ToDecimal(maximum)
						if num.GreaterThan(maxDecimal) {
							temp := strconv.FormatFloat(maximum, 'f', -1, 64)
							return errors.New(fmt.Sprintf("field '%s' is greater than %s", name, temp))
						}
					}
				}

				switch vField.Kind() {
				case reflect.Struct:
					return ValidateStruct(vField)

				case reflect.String:
					size := int64(vField.Len())

					minInt64 := int64(minimum)
					if setMinimum && size < minInt64 {
						return errors.New(fmt.Sprintf("field '%s' is less than %d", name, minInt64))
					}

					maxInt64 := int64(maximum)
					if setMaximum && size > maxInt64 {
						return errors.New(fmt.Sprintf("field '%s' is greater than %d", name, maxInt64))
					}

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
				nokocore.KeepVoid(i)

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
	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Bool:
		return nil

	case reflect.String:
		if !nokocore.BooleanRegex().MatchString(val.String()) {
			return errors.New("invalid boolean format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateNumber(value any) error {
	val := nokocore.PassValueIndirectReflect(value)
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
		if !nokocore.NumberRegex().MatchString(val.String()) {
			return errors.New("invalid number format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateNumeric(value any) error {
	val := nokocore.PassValueIndirectReflect(value)
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
		if !nokocore.NumericRegex().MatchString(val.String()) {
			return errors.New("invalid numeric format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAlphaLower(value any) error {
	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		if !nokocore.AlphaLowerRegex().MatchString(val.String()) {
			return errors.New("invalid alpha lower format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAlphaUpper(value any) error {
	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		if !nokocore.AlphaUpperRegex().MatchString(val.String()) {
			return errors.New("invalid alpha upper format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAlphabet(value any) error {
	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		if !nokocore.AlphabetRegex().MatchString(val.String()) {
			return errors.New("invalid alphabet format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAlphaNum(value any) error {
	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		if !nokocore.AlphaNumRegex().MatchString(val.String()) {
			return errors.New("invalid alpha number format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAscii(value any) error {
	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		if !nokocore.AsciiRegex().MatchString(val.String()) {
			return errors.New("invalid ascii format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateEmail(value any) error {
	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		if !nokocore.EmailRegex().MatchString(val.String()) {
			return errors.New("invalid email format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidatePhone(value any) error {
	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		if !nokocore.PhoneRegex().MatchString(val.String()) {
			return errors.New("invalid phone format")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidatePassword(value any) error {
	var err error
	nokocore.KeepVoid(err)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.String:
		return CheckPassword(val.String())

	default:
		return errors.New("invalid data type")
	}
}

func ValidateDateTime(value any) error {
	var ok bool
	var err error
	var temp time.Time
	nokocore.KeepVoid(ok, err, temp)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return nil

	case reflect.Struct:
		switch val.Interface().(type) {
		case time.Time, time.Duration:
			return nil

		default:
			return errors.New("invalid data type")
		}

	case reflect.String:
		if temp, err = time.Parse(nokocore.DateTimeFormat, val.String()); err != nil {
			return errors.New("invalid datetime format, please using YYYY-MM-DD HH:mm:ss")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateDateTimeISO8601(value any) error {
	var ok bool
	var err error
	var temp time.Time
	nokocore.KeepVoid(ok, err, temp)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return nil

	case reflect.Struct:
		switch val.Interface().(type) {
		case time.Time, time.Duration:
			return nil

		default:
			return errors.New("invalid data type")
		}

	case reflect.String:
		if temp, err = time.Parse(nokocore.DateTimeFormatISO8601, val.String()); err != nil {
			return errors.New("invalid datetime format, please using YYYY-MM-DDTHH:mm:ss.sssZ or YYYY-MM-DDTHH:mm:ss+/-HH:mm")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateDateOnly(value any) error {
	var ok bool
	var err error
	var temp time.Time
	nokocore.KeepVoid(ok, err, temp)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return nil

	case reflect.Struct:
		switch val.Interface().(type) {
		case time.Time, time.Duration, NullDateOnly, DateOnlyImpl, DateOnly:
			return nil

		default:
			return errors.New("invalid data type")
		}

	case reflect.String:
		if temp, err = time.Parse(nokocore.DateOnlyFormat, val.String()); err != nil {
			return errors.New("invalid date format, please using YYYY-MM-DD")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateTimeOnly(value any) error {
	var ok bool
	var err error
	var temp time.Time
	nokocore.KeepVoid(ok, err, temp)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return nil

	case reflect.Struct:
		switch val.Interface().(type) {
		case time.Time, time.Duration, NullTimeOnly, TimeOnlyImpl, TimeOnly:
			return nil

		default:
			return errors.New("invalid data type")
		}

	case reflect.String:
		if temp, err = time.Parse(nokocore.TimeOnlyFormat, val.String()); err != nil {
			return errors.New("invalid time format, please using HH:mm:ss")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateURL(value any) error {
	var ok bool
	var err error
	var temp *url.URL
	nokocore.KeepVoid(ok, err, temp)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Struct:

		switch val.Interface().(type) {
		case url.URL:
			return nil

		default:
			return errors.New("invalid data type")
		}

	case reflect.String:
		if temp, err = url.Parse(val.String()); err != nil {
			return errors.New("invalid URL format, please using http://<domain>.<tld>/$request_uri or https://<domain>.<tld>/$request_uri")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateUUID(value any) error {
	var ok bool
	var err error
	var temp uuid.UUID
	nokocore.KeepVoid(ok, err, temp)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Struct:

		switch val.Interface().(type) {
		case uuid.UUID:
			return nil

		default:
			return errors.New("invalid data type")
		}

	case reflect.String:
		if temp, err = uuid.Parse(val.String()); err != nil {
			return errors.New("invalid UUID format, please using 00000000-0000-0000-0000-000000000000")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateDecimal(value any) error {
	var ok bool
	var err error
	var temp decimal.Decimal
	nokocore.KeepVoid(ok, err, temp)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Struct:

		switch val.Interface().(type) {
		case decimal.NullDecimal, decimal.Decimal:
			return nil

		default:
			return errors.New("invalid data type")
		}

	case reflect.String:
		if temp, err = decimal.NewFromString(val.String()); err != nil {
			return errors.New("invalid decimal format, please using 0.00")
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}
