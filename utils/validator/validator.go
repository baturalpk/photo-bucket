package validator

import (
	"github.com/go-playground/validator/v10"
)

func init() {
	_ = GetValidatorInstance()
}

var v *validator.Validate

func GetValidatorInstance() *validator.Validate {
	if v == nil {
		v = validator.New()
		return v
	}

	return v
}
