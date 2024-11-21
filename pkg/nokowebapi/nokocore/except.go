package nokocore

import (
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

	if message != "" {
		format = "%w, %s"
		args = []any{err, message}

	} else {
		format = "%w"
		args = []any{err}
		
	}

	for i, e := range more {
		KeepVoid(i, e)

		format += ", %w"
		args = append(args, e)
	}

	return fmt.Errorf(format, args...)
}
