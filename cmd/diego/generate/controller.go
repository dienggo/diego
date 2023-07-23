package generate

import (
	"github.com/daewu14/golang-base/pkg/file"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
)

func buildStructName(structName string) (s string) {
	s = strings.Replace(structName, ".go", "", -1)

	s = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, " ")

	split := strings.Split(s, " ")
	for i, s2 := range split {
		split[i] = cases.Title(language.English, cases.NoLower).String(s2)
	}
	s = strings.Join(split, "")

	return s
}

func Controller(controllerName string) {
	// The name of the file to create.
	fileName := "app/controllers/" + controllerName + ".go"

	structName := cases.Title(language.English, cases.NoLower).String(strings.Replace(fileName, "app/controllers/", "", -1))

	structName = buildStructName(structName)
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
