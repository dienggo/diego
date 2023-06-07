package helper

import (
	"errors"
	validator2 "github.com/go-playground/validator/v10"
)

var validator *validator2.Validate
var isValidatorInstance bool = false

// validate : make singleton validator instance
func validate() *validator2.Validate {
	if !isValidatorInstance {
		validator = validator2.New()
		isValidatorInstance = true
	}
	return validator
}

// Validator : main method of validator
func Validator() *validator2.Validate {
	return validate()
}

// FormattedError : to generate format error message
func FormattedError(err error) error {
	var ve validator2.ValidationErrors
	var sErrors []string
	if errors.As(err, &ve) {
		for _, fe := range ve {
			sErrors = append(sErrors, fe.Field()+" field is "+fe.Tag())
		}
	}
	if len(sErrors) == 0 {
		return nil
	}
	return errors.New(sErrors[0])
}

// ValidateStruct : to validate struct
func ValidateStruct(s interface{}) error {
	return validate().Struct(s)
}

// ValidateStructFormatted : to validate struct with formatted error message
func ValidateStructFormatted(s any) error {
	return FormattedError(ValidateStruct(s))
}
