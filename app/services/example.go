package services

import (
	"github.com/dienggo/diego/app/dto/dto_request"
	"github.com/dienggo/diego/app/repositories"
)

// NewExample : make new instance for Example struct
func NewExample(settingByKey dto_request.SettingByKey) Example {
	return Example{
		Request:     settingByKey,
		SettingRepo: repositories.Setting{},
	}
}

type Example struct {
	Request     dto_request.SettingByKey
	SettingRepo repositories.ISetting
}

func (s Example) Do() ServiceResponse {

	// validate data
	errValidate := s.Request.Validate()
	if errValidate != nil {
		return SRFail(errValidate, nil)
	}

	// do inquiry on database search
	err, setting := s.SettingRepo.FindByKey(s.Request.Key)
	if err != nil {
		return SRFail(err, nil)
	}

	return SRSuccess(setting)
}
