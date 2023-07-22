package test

import (
	"github.com/daewu14/golang-base/pkg/environment"
	"github.com/daewu14/golang-base/pkg/helper"
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
