package validators

import "github.com/go-playground/validator/v10"

type EchoValidatorImpl interface {
	Validate(data any) error
}

type EchoValidator struct {
	validator *validator.Validate
}

func NewEchoValidator() EchoValidatorImpl {
	return &EchoValidator{
		validator: validator.New(),
	}
}

func (v *EchoValidator) Validate(data any) error {
	return v.validator.Struct(data)
}
