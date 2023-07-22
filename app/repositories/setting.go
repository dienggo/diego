package repositories

import (
	"github.com/daewu14/golang-base/app/models"
	"github.com/daewu14/golang-base/pkg/database"
)

type SettingRepo struct{}

func (repo SettingRepo) FindByKey(value string) (error, models.Setting) {
	var settingData models.Setting
	result := database.Main().Unscoped().Where("`key` = ?", value).First(&settingData)
	return result.Error, settingData
}

func (repo SettingRepo) GetWhere(query interface{}, args ...interface{}) (error, []models.Setting) {
	var settingData []models.Setting
	result := database.Main().Unscoped().Where(query, args).Find(&settingData)
	return result.Error, settingData
}
