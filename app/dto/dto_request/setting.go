package dto_request

import (
	"github.com/dienggo/diego/pkg/validates"
)

type SettingByKey struct {
	Key string `json:"key" validate:"required"`
}

func (s SettingByKey) Validate() error {
	return validates.ValidateStruct(s)
}
