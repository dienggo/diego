package test

import (
	"go_base_project/helper"
	"testing"
)

type Example struct {
	Key string `json:"key" validate:"required"`
}

func TestValidator(t *testing.T) {
	var example Example

	err := helper.ValidateStructFormatted(example)

	if err != nil {
		println(err.Error())
		return
	}
}
