package interfaces

import "go_base_project/app/models"

type ISettingRepo interface {
	FindByKey(value string) (error, models.Setting)
	GetWhere(query interface{}, args ...interface{}) (error, []models.Setting)
}
