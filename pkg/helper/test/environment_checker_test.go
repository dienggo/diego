package test

import (
	"go_base_project/pkg/environment"
	"go_base_project/pkg/helper"
	"strings"
	"testing"
)

func TestCheckEnvKeys(t *testing.T) {
	environment.Load()
	result := helper.IsExistAllEnvKeys("OK", "DB_CONNECTION", "DD")
	println("status", result.Status)
	println("notExistKeys", strings.Join(result.NotExistKeys, ", "))
	println("existKeys", strings.Join(result.ExistKeys, ", "))
}
