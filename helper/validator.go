package helper

import validator2 "github.com/go-playground/validator/v10"

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

// ValidateStruct : to validate struct
func ValidateStruct(s interface{}) error {
	return validate().Struct(s)
}
