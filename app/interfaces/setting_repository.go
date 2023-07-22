package interfaces

import "github.com/daewu14/golang-base/app/models"

type ISettingRepo interface {
	FindByKey(value string) (error, models.Setting)
	GetWhere(query interface{}, args ...interface{}) (error, []models.Setting)
}
