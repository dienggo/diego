package test

import (
	"go_base_project/helper"
	"go_base_project/provider"
	"strings"
	"testing"
)

func TestCheckEnvKeys(t *testing.T) {
	provider.LoadEnv()
	result := helper.IsExistAllEnvKeys("OK", "DB_CONNECTION", "DD")
	println("status", result.Status)
	println("notExistKeys", strings.Join(result.NotExistKeys, ", "))
	println("existKeys", strings.Join(result.ExistKeys, ", "))
}
