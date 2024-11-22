package nokocore

import (
	"errors"
	"fmt"
)

type Exception struct {
	error
}

func (e Exception) GetNameType() string {
	return "Exception"
}

func NewThrow(message string, err error, more ...error) error {
	var args []any
	var format string
	KeepVoid(format, args)

	args = []any{errors.New(message), err}
	format = "%w, %w"

	for i, e := range more {
		KeepVoid(i)

		args = append(args, e)
		format += ", %w"
	}

	return fmt.Errorf(format, args...)
}
