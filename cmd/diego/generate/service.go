package generate

import (
	"errors"
	"github.com/dienggo/diego/pkg/file"
	"github.com/dienggo/diego/pkg/helper"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

type service struct{}

func (c service) Command() *cli.Command {
	return &cli.Command{
		Name:  "service",
		Usage: "Generate/make service by boilerplate",
		Action: func(context *cli.Context) error {
			if context.Args().Get(0) == "" {
				return errors.New("service name can not be empty")
			}
			c.build(helper.ReplaceSpecialCharacters(context.Args().Get(0), "_"))
			return nil
		},
	}
}

func (c service) build(serviceName string) {
	// The name of the file to create.
	fileName := "app/services/" + serviceName + ".go"

	structName := cases.Title(language.English, cases.NoLower).String(strings.Replace(fileName, "app/services/", "", -1))

	structName = buildStructName(structName)

	moduleName := helper.GetModuleName()

	// Write content to the file.
	content :=
		`package services

import "` + moduleName + `/pkg/app"

// New` + structName + ` : to simplify service call
func New` + structName + `() ` + structName + ` {
	return ` + structName + `{}
}

type ` + structName + ` struct{
	app.Service
}

func (s ` + structName + `) Do() app.ServiceResponse {
	// TODO : Implement me
	panic("Implement me")
}
`
	err := file.Create(fileName, content)
	if err != nil {
		return
	}
}
