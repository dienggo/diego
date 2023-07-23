package generate

import (
	"errors"
	"github.com/daewu14/golang-base/pkg/file"
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
			c.build(context.Args().Get(0))
			return nil
		},
	}
}

func (c service) build(serviceName string) {
	// The name of the file to create.
	fileName := "app/services/" + serviceName + ".go"

	structName := cases.Title(language.English, cases.NoLower).String(strings.Replace(fileName, "app/services/", "", -1))

	structName = buildStructName(structName)
	// Write content to the file.
	content :=
		`package services

type ` + structName + ` struct{}

func (s ` + structName + `) Do() ServiceResponse {
	//TODO implement me
	panic("implement me")
}
`
	err := file.Create(fileName, content)
	if err != nil {
		return
	}
}
