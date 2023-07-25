package repositories

import (
	"github.com/dienggo/diego/app/models"
	"github.com/dienggo/diego/pkg/database"
)

type ISetting interface {
	FindByKey(value string) (error, models.Setting)
	GetWhere(query interface{}, args ...interface{}) (error, []models.Setting)
}

type Setting struct{}

func (repo Setting) FindByKey(value string) (error, models.Setting) {
	var settingData models.Setting
	result := database.Main().Unscoped().Where("`key` = ?", value).First(&settingData)
	return result.Error, settingData
}

func (repo Setting) GetWhere(query interface{}, args ...interface{}) (error, []models.Setting) {
	var settingData []models.Setting
	result := database.Main().Unscoped().Where(query, args).Find(&settingData)
	return result.Error, settingData
}
