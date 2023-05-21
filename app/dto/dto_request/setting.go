package dto_request

import validation "github.com/go-ozzo/ozzo-validation"

type SettingByKey struct {
	Key string `json:"key" form:"key"`
}

func (s SettingByKey) Validate() error {
	return validation.ValidateStruct(&s,
		// make request 'key` is required
		validation.Field(&s.Key, validation.Required),
	)
}
