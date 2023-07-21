package dto_request

type SettingByKey struct {
	Key string `json:"key" validate:"required"`
}

func (s SettingByKey) Validate() error {
	return validates.ValidateStruct(s)
}
