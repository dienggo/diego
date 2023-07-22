package generate

import (
	"go_base_project/pkg/file"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func Controller(controllerName string) {
	// The name of the file to create.
	fileName := "app/controllers/" + controllerName + ".go"
	structName := cases.Title(language.English, cases.NoLower).String(strings.Replace(fileName, "app/controllers/", "", -1))
	structName = strings.Replace(structName, ".go", "", -1)
	// Write content to the file.
	content :=
		`package controllers

import "github.com/gin-gonic/gin"

type ` + structName + ` struct{}

func (controller ` + structName + `) Main(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong 2",
	})
}

`
	err := file.Create(fileName, content)
	if err != nil {
		return
	}
}
