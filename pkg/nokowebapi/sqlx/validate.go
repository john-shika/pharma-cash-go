package sqlx

import (
	"database/sql"
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

type ValidatorForStructImpl interface {
	Validate(value any, attributes map[string]string) error
	Keys() []string
	Is(key string) bool
}

type ValidatorForStruct struct {
	validate func(value any, attributes map[string]string) error
	keys     []string
}

func NewValidatorForStruct(keys []string, validate func(value any, attributes map[string]string) error) ValidatorForStructImpl {
	return &ValidatorForStruct{
		keys:     keys,
		validate: validate,
	}
}

func (v *ValidatorForStruct) Validate(value any, attributes map[string]string) error {
	if v.validate != nil {
		return v.validate(value, attributes)
	}

	return nil
}

func (v *ValidatorForStruct) Keys() []string {
	return v.keys
}

func (v *ValidatorForStruct) Is(key string) bool {
	for i, k := range v.keys {
		nokocore.KeepVoid(i)

		if strings.EqualFold(k, key) {
			return true
		}
	}

	return false
}

func ValidateStruct(value any) error {
	var ok bool
	var err error
	nokocore.KeepVoid(ok, err)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	validators := []ValidatorForStructImpl{
		NewValidatorForStruct([]string{"boolean"}, func(value any, attributes map[string]string) error {
			return ValidateBoolean(value)
		}),
		NewValidatorForStruct([]string{"number"}, func(value any, attributes map[string]string) error {
			var ok bool
			var err error
			var minVal int64
			var maxVal int64
			var minimum string
			var maximum string
			nokocore.KeepVoid(ok, err, minVal, maxVal, minimum, maximum)

			if err = ValidateNumber(value); err != nil {
				return err
			}

			numVal := nokocore.ToIntReflect(value)
			if minimum, ok = attributes["min"]; ok {
				if minVal, err = strconv.ParseInt(minimum, 10, 64); err != nil {
					return errors.New("invalid min value")
				}

				if numVal < minVal {
					return errors.New("value is less than min value")
				}
			}

			if maximum, ok = attributes["max"]; ok {
				if maxVal, err = strconv.ParseInt(maximum, 10, 64); err != nil {
					return errors.New("invalid max value")
				}

				if numVal > maxVal {
					return errors.New("value is greater than max value")
				}
			}

			return nil
		}),
		NewValidatorForStruct([]string{"numeric"}, func(value any, attributes map[string]string) error {
			var ok bool
			var err error
			var minVal float64
			var maxVal float64
			var minimum string
			var maximum string
			nokocore.KeepVoid(ok, err, minVal, maxVal, minimum, maximum)

			if err = ValidateNumeric(value); err != nil {
				return err
			}

			numVal := nokocore.ToFloatReflect(value)
			if minimum, ok = attributes["min"]; ok {
				if minVal, err = strconv.ParseFloat(minimum, 64); err != nil {
					return errors.New("invalid min value")
				}

				if numVal < minVal {
					return errors.New("value is less than min value")
				}
			}

			if maximum, ok = attributes["max"]; ok {
				if maxVal, err = strconv.ParseFloat(maximum, 64); err != nil {
					return errors.New("invalid max value")
				}

				if numVal > maxVal {
					return errors.New("value is greater than max value")
				}
			}

			return nil
		}),
		NewValidatorForStruct([]string{"alphaLower"}, func(value any, attributes map[string]string) error {
			return ValidateAlphaLower(value)
		}),
		NewValidatorForStruct([]string{"alphaUpper"}, func(value any, attributes map[string]string) error {
			return ValidateAlphaUpper(value)
		}),
		NewValidatorForStruct([]string{"alphabet"}, func(value any, attributes map[string]string) error {
			return ValidateAlphabet(value)
		}),
		NewValidatorForStruct([]string{"alphaNum"}, func(value any, attributes map[string]string) error {
			return ValidateAlphaNum(value)
		}),
		NewValidatorForStruct([]string{"ascii"}, func(value any, attributes map[string]string) error {
			return ValidateAscii(value)
		}),
		NewValidatorForStruct([]string{"email"}, func(value any, attributes map[string]string) error {
			return ValidateEmail(value)
		}),
		NewValidatorForStruct([]string{"phone"}, func(value any, attributes map[string]string) error {
			return ValidatePhone(value)
		}),
		NewValidatorForStruct([]string{"password"}, func(value any, attributes map[string]string) error {
			return ValidatePassword(value)
		}),
		NewValidatorForStruct([]string{"datetime"}, func(value any, attributes map[string]string) error {
			return ValidateDateTime(value)
		}),
		NewValidatorForStruct([]string{"datetimeISO", "datetimeISO8601"}, func(value any, attributes map[string]string) error {
			return ValidateDateTimeISO8601(value)
		}),
		NewValidatorForStruct([]string{"dateOnly"}, func(value any, attributes map[string]string) error {
			return ValidateDateOnly(value)
		}),
		NewValidatorForStruct([]string{"timeOnly"}, func(value any, attributes map[string]string) error {
			return ValidateTimeOnly(value)
		}),
		NewValidatorForStruct([]string{"url"}, func(value any, attributes map[string]string) error {
			return ValidateURL(value)
		}),
		NewValidatorForStruct([]string{"uuid"}, func(value any, attributes map[string]string) error {
			return ValidateUUID(value)
		}),
		NewValidatorForStruct([]string{"decimal"}, func(value any, attributes map[string]string) error {
			var ok bool
			var err error
			var minVal decimal.Decimal
			var maxVal decimal.Decimal
			var minimum string
			var maximum string
			nokocore.KeepVoid(ok, err, minVal, maxVal, minimum, maximum)

			if err = ValidateDecimal(value); err != nil {
				return err
			}

			numVal := nokocore.ToDecimalReflect(value)
			if minimum, ok = attributes["min"]; ok {
				if minVal, err = decimal.NewFromString(minimum); err != nil {
					return errors.New("invalid min value")
				}

				if numVal.LessThan(minVal) {
					return errors.New("value is less than min value")
				}
			}

			if maximum, ok = attributes["max"]; ok {
				if maxVal, err = decimal.NewFromString(maximum); err != nil {
					return errors.New("invalid max value")
				}

				if numVal.GreaterThan(maxVal) {
					return errors.New("value is greater than max value")
				}
			}

			return nil
		}),
	}

	switch val.Kind() {
	case reflect.Struct:
		var errorStack []error
		options := nokocore.NewForEachStructFieldsOptions()
		err = nokocore.ForEachStructFieldsReflect(val, options, func(name string, sFieldX nokocore.StructFieldExImpl) error {
			err = func() error {
				vField := sFieldX.GetValue()
				sTag := sFieldX.GetTag()

				validate := sTag.Get("validate")
				tokens := strings.Split(validate, ",")

				ignore := false
				omitEmpty := false

				attributes := make(map[string]string)
				for i, token := range tokens {
					nokocore.KeepVoid(i)

					token = strings.TrimSpace(token)
					if token, ok = strings.CutPrefix(token, "min="); ok {
						attributes["min"] = token
						continue
					}

					if token, ok = strings.CutPrefix(token, "max="); ok {
						attributes["max"] = token
						continue
					}

					switch token {
					case "-", "ignore":
						ignore = true
						break

					case "omitempty":
						omitEmpty = true
						break

					default:
						break
					}
				}

				if ignore {
					return nil
				}

				nameCamelCase := nokocore.ToCamelCase(name)
				if nokocore.IsNoneOrEmptyWhiteSpaceReflect(vField) {
					if !omitEmpty {
						if name != nameCamelCase {
							return errors.New(fmt.Sprintf("field '%s' is required, use '%s' for form data", nameCamelCase, name))
						}

						return errors.New(fmt.Sprintf("field '%s' is required", name))
					}

					// skip
					return nil
				}

				for i, token := range tokens {
					nokocore.KeepVoid(i)

					for j, validator := range validators {
						nokocore.KeepVoid(j)

						if validator.Is(token) {
							if err = validator.Validate(vField, attributes); err != nil {
								return fmt.Errorf("field '%s' is %w", nameCamelCase, err)
							}
						}
					}
				}

				switch vField.Kind() {
				case reflect.Struct:
					return ValidateStruct(vField)

				case reflect.String:
					var ok bool
					var minVal int
					var maxVal int
					var minimum string
					var maximum string
					nokocore.KeepVoid(ok, minVal, maxVal, minimum, maximum)

					size := vField.Len()

					if minimum, ok = attributes["min"]; ok {
						if minVal, err = strconv.Atoi(minimum); err != nil {
							return errors.New("invalid min value")
						}

						if size < minVal {
							return errors.New(fmt.Sprintf("field '%s' is less than %d", name, minVal))
						}
					}

					if maximum, ok = attributes["max"]; ok {
						if maxVal, err = strconv.Atoi(maximum); err != nil {
							return errors.New("invalid max value")
						}

						if size > maxVal {
							return errors.New(fmt.Sprintf("field '%s' is greater than %d", name, maxVal))
						}
					}

					return nil

				default:
					return nil
				}
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
	var err error
	nokocore.KeepVoid(err)

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

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateBoolean(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateNumber(value any) error {
	var err error
	nokocore.KeepVoid(err)

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

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateNumber(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateNumeric(value any) error {
	var err error
	nokocore.KeepVoid(err)

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

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateNumeric(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAlphaLower(value any) error {
	var err error
	nokocore.KeepVoid(err)

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

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateAlphaLower(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAlphaUpper(value any) error {
	var err error
	nokocore.KeepVoid(err)

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

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateAlphaUpper(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAlphabet(value any) error {
	var err error
	nokocore.KeepVoid(err)

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

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateAlphabet(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAlphaNum(value any) error {
	var err error
	nokocore.KeepVoid(err)

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

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateAlphaNum(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateAscii(value any) error {
	var err error
	nokocore.KeepVoid(err)

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

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateAscii(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateEmail(value any) error {
	var err error
	nokocore.KeepVoid(err)

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

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateEmail(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidatePhone(value any) error {
	var err error
	nokocore.KeepVoid(err)

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

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidatePhone(temp); err != nil {
				return err
			}
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

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidatePassword(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateDateTime(value any) error {
	var ok bool
	var err error
	var check time.Time
	nokocore.KeepVoid(ok, err, check)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return nil

	case reflect.Struct:
		switch val.Interface().(type) {
		case time.Time, time.Duration, sql.NullTime, sql.Null[time.Time], sql.Null[time.Duration]:
			return nil

		default:
			return errors.New("invalid data type")
		}

	case reflect.String:
		if check, err = time.Parse(nokocore.DateTimeFormat, val.String()); err != nil {
			return errors.New("invalid datetime format, please using YYYY-MM-DD HH:mm:ss")
		}

		return nil

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateDateTime(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateDateTimeISO8601(value any) error {
	var ok bool
	var err error
	var check time.Time
	nokocore.KeepVoid(ok, err, check)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return nil

	case reflect.Struct:
		switch val.Interface().(type) {
		case time.Time, time.Duration, sql.NullTime, sql.Null[time.Time], sql.Null[time.Duration]:
			return nil

		default:
			return errors.New("invalid data type")
		}

	case reflect.String:
		if check, err = time.Parse(nokocore.DateTimeFormatISO8601, val.String()); err != nil {
			return errors.New("invalid datetime format, please using YYYY-MM-DDTHH:mm:ss.sssZ or YYYY-MM-DDTHH:mm:ss+/-HH:mm")
		}

		return nil

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateDateTimeISO8601(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateDateOnly(value any) error {
	var ok bool
	var err error
	var check time.Time
	nokocore.KeepVoid(ok, err, check)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return nil

	case reflect.Struct:
		switch val.Interface().(type) {
		case time.Time, time.Duration, sql.NullTime, sql.Null[time.Time], sql.Null[time.Duration]:
			return nil

		case NullDateOnly, DateOnlyImpl, DateOnly:
			return nil

		default:
			return errors.New("invalid data type")
		}

	case reflect.String:
		if check, err = time.Parse(nokocore.DateOnlyFormat, val.String()); err != nil {
			return errors.New("invalid date format, please using YYYY-MM-DD")
		}

		return nil

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateDateOnly(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateTimeOnly(value any) error {
	var ok bool
	var err error
	var check time.Time
	nokocore.KeepVoid(ok, err, check)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		return errors.New("invalid value")
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return nil

	case reflect.Struct:
		switch val.Interface().(type) {
		case time.Time, time.Duration, sql.NullTime, sql.Null[time.Time], sql.Null[time.Duration]:
			return nil

		case NullTimeOnly, TimeOnlyImpl, TimeOnly:
			return nil

		default:
			return errors.New("invalid data type")
		}

	case reflect.String:
		if check, err = time.Parse(nokocore.TimeOnlyFormat, val.String()); err != nil {
			return errors.New("invalid time format, please using HH:mm:ss")
		}

		return nil

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateTimeOnly(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateURL(value any) error {
	var ok bool
	var err error
	var check *url.URL
	nokocore.KeepVoid(ok, err, check)

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
		if check, err = url.Parse(val.String()); err != nil {
			return errors.New("invalid URL format, please using http://<domain>.<tld>/$request_uri or https://<domain>.<tld>/$request_uri")
		}

		return nil

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateURL(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateUUID(value any) error {
	var ok bool
	var err error
	var check uuid.UUID
	nokocore.KeepVoid(ok, err, check)

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
		if check, err = uuid.Parse(val.String()); err != nil {
			return errors.New("invalid UUID format, please using 00000000-0000-0000-0000-000000000000")
		}

		return nil

	case reflect.Array, reflect.Slice:
		size := val.Len()

		// fix uuid array
		if size == 16 {
			var v uuid.UUID
			nokocore.KeepVoid(v)

			if v, ok = val.Interface().(uuid.UUID); ok {
				return nil
			}

			return errors.New("invalid UUID format, please using 00000000-0000-0000-0000-000000000000")
		}

		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateUUID(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}

func ValidateDecimal(value any) error {
	var ok bool
	var err error
	var check decimal.Decimal
	nokocore.KeepVoid(ok, err, check)

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
		if check, err = decimal.NewFromString(val.String()); err != nil {
			return errors.New("invalid decimal format, please using 0.00")
		}

		return nil

	case reflect.Array, reflect.Slice:
		size := val.Len()
		for i := 0; i < size; i++ {
			temp := val.Index(i)
			if err = ValidateDecimal(temp); err != nil {
				return err
			}
		}

		return nil

	default:
		return errors.New("invalid data type")
	}
}
