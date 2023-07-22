package validates

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

// FormattedErrors : to generate format error messages
func FormattedErrors(err error) []error {
	var ve validator2.ValidationErrors
	var eErrors []error
	if errors.As(err, &ve) {
		for _, fe := range ve {
			msg := fe.Field() + " field is " + fe.Tag()
			eErrors = append(eErrors, errors.New(msg))
		}
	}
	if len(eErrors) == 0 {
		return nil
	}
	return eErrors
}

// FormattedError : to generate format error message
func FormattedError(err error) error {
	formats := FormattedErrors(err)
	if len(formats) == 0 {
		return nil
	}
	return formats[0]
}

// ValidateStruct : to validate struct
func ValidateStruct(s interface{}) error {
	return Validator().Struct(s)
}

// ValidateStructFormatted : to validate struct with formatted error message
func ValidateStructFormatted(s any) error {
	return FormattedError(ValidateStruct(s))
}

// ValidateStructFormattedErrors : to validate struct with formatted error messages
func ValidateStructFormattedErrors(s any) []error {
	return FormattedErrors(ValidateStruct(s))
}
