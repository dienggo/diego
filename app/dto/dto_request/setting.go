package dto_request

import (
	"go_base_project/pkg/validates"
)

type SettingByKey struct {
	Key string `json:"key" validate:"required"`
}

func (s SettingByKey) Validate() error {
	return validates.ValidateStruct(s)
}
