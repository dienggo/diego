package test

import (
	"github.com/daewu14/golang-base/app/dto/dto_request"
	"github.com/daewu14/golang-base/app/models"
	"github.com/daewu14/golang-base/app/services"
	"github.com/daewu14/golang-base/pkg/environment"
	"testing"
)

type mockResult struct {
	err     error
	result  models.Setting
	results []models.Setting
}

// MockSettingRepo : manual create mocking repository
type MockSettingRepo struct {
	res mockResult
}

func (r MockSettingRepo) MockResult(result mockResult) MockSettingRepo {
	r.res = result
	return r
}

func (r MockSettingRepo) FindByKey(value string) (error, models.Setting) {
	return r.res.err, r.res.result
}

func (r MockSettingRepo) GetWhere(query interface{}, args ...interface{}) (error, []models.Setting) {
	return r.res.err, r.res.results
}

func TestExampleService(t *testing.T) {
	environment.Load()

	repo := MockSettingRepo{}.MockResult(mockResult{
		err: nil,
		result: models.Setting{
			Key:   "max_value_cod",
			Value: "20000",
		},
	})

	service := services.Example{
		Request:     dto_request.SettingByKey{Key: "max_value_cod"},
		SettingRepo: repo,
	}.Do()

	var setting models.Setting
	err := service.ResultParse(&setting)
	if err != nil {
		println("error :", err.Error())
		return
	}

	println("result setting value : " + setting.Value)
}
