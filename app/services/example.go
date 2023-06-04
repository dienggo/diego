package services

import (
	"go_base_project/app/base"
	"go_base_project/app/dto/dto_request"
	"go_base_project/app/interfaces"
	"go_base_project/app/repositories"
)

// NewExample : make new instance for Example struct
func NewExample(settingByKey dto_request.SettingByKey) Example {
	return Example{
		Request:     settingByKey,
		SettingRepo: repositories.SettingRepo{},
	}
}

type Example struct {
	Request     dto_request.SettingByKey
	SettingRepo interfaces.ISettingRepo
}

func (s Example) Do() base.ServiceResponse {

	// validate data
	errValidate := s.Request.Validate()
	if errValidate != nil {
		return base.SRFail(errValidate, nil)
	}

	// do inquiry on database search
	err, setting := s.SettingRepo.FindByKey(s.Request.Key)
	if err != nil {
		return base.SRFail(err, nil)
	}

	return base.SRSuccess(setting)
}
