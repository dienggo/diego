package dto_request

import "go_base_project/helper"

type SettingByKey struct {
	Key string `json:"key" validate:"required"`
}

func (s SettingByKey) Validate() error {
	return helper.ValidateStruct(s)
}
