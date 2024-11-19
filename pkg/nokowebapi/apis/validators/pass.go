package validators

import (
	"nokowebapi/nokocore"
	"strings"
)

type ValidatePassErrorImpl interface {
	Fields() []string
	Error() string
}

type ValidatePassError struct {
	fields []string
}

func NewValidatePassError(fields []string) ValidatePassErrorImpl {
	return &ValidatePassError{
		fields: fields,
	}
}

func (v *ValidatePassError) Fields() []string {
	return v.fields
}

func (v *ValidatePassError) Error() string {
	return strings.Join(v.fields, "\n")
}

func ValidatePass(pass string) error {
	var temp []string
	nokocore.KeepVoid(temp)

	size := len(pass)
	if size < 8 {
		message := "Password must be at least 8 characters long."
		temp = append(temp, message)
	}

	if !strings.ContainsAny(pass, nokocore.AsciiLower) {
		message := "Password must contain at least one lowercase letter."
		temp = append(temp, message)
	}

	if !strings.ContainsAny(pass, nokocore.AsciiUpper) {
		message := "Password must contain at least one uppercase letter."
		temp = append(temp, message)
	}

	if !strings.ContainsAny(pass, nokocore.Digits) {
		message := "Password must contain at least one digit."
		temp = append(temp, message)
	}

	if !strings.ContainsAny(pass, nokocore.Punctuation) {
		message := "Password must contain at least one special character."
		temp = append(temp, message)
	}

	if len(temp) > 0 {
		return NewValidatePassError(temp)
	}

	return nil
}
