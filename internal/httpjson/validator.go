package httpjson

import "github.com/go-playground/validator/v10"

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func (v *Validator) Validate(i any) error {
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
