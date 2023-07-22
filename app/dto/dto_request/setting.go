package dto_request

import (
	"github.com/daewu14/golang-base/pkg/validates"
)

type SettingByKey struct {
	Key string `json:"key" validate:"required"`
}

func (s SettingByKey) Validate() error {
	return validates.ValidateStruct(s)
}
