package test

import (
	"testing"

	"github.com/dienggo/diego/pkg/validates"
)

type Example struct {
	Key  string `json:"key" validate:"required"`
	Test string `json:"test" validate:"required"`
}

func TestValidator(t *testing.T) {
	var example Example

	err := validates.ValidateStructFormatted(example)

	if err != nil {
		println(err.Error())
		return
	}
}

func TestValidatorFormattedErrors(t *testing.T) {
	var example Example = Example{
		Test: "Okee",
	}

	errs := validates.ValidateStructFormattedErrors(example)

	if len(errs) > 0 {
		for _, val := range errs {
			println(val.Error())
		}
		return
	}
}

type Examples []struct {
	Key  string `json:"key" validate:"required"`
	Test string `json:"test" validate:"required"`
}

func TestStructOfSlice(t *testing.T) {

}
