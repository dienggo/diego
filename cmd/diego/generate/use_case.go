package generate

import (
	"errors"
	"github.com/dienggo/diego/pkg/file"
	"github.com/dienggo/diego/pkg/helper"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
)

type ucase struct{}

func (c ucase) Command() *cli.Command {
	return &cli.Command{
		Name:    "ucase",
		Aliases: []string{"use-case", "uc", "u-case"},
		Usage:   "Generate/make ucase by boilerplate",
		Action: func(context *cli.Context) error {
			if context.Args().Get(0) == "" {
				return errors.New("Use case name can not be empty")
			}
			name := helper.ReplaceSpecialCharacters(context.Args().Get(0), "_")
			folder := helper.ReplaceSpecialCharacters(context.Args().Get(1), "_")
			c.build(name, folder)
			return nil
		},
	}
}

func (c ucase) build(uCaseName string, folder string) {
	if folder != "" {
		path := "app/ucase/" + folder
		exist := helper.PathExist(path)
		if !exist {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				log.Error("Error: ", err.Error())
				return
			}
		}
		c.buildPartial(uCaseName, path)
	} else {

		if !strings.Contains(uCaseName, "case") && !strings.Contains(uCaseName, "Case") {
			uCaseName = uCaseName + "Case"
		}

		folder = uCaseName
		path := "app/ucase/" + folder
		exist := helper.PathExist(path)
		if !exist {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				log.Error("Error: ", err.Error())
				return
			}
			err = os.MkdirAll(path+"/test", os.ModePerm)
			if err != nil {
				log.Error("Error: ", err.Error())
				return
			}

			registryTemplateUcase := []string{
				"upsert",
				"detail",
				"delete",
				"list",
			}

			for _, s := range registryTemplateUcase {
				c.buildPartial(s, path)
			}

		} else {
			log.Error("Error: ", uCaseName, " is already exist!")
		}
	}
}

func (c ucase) buildPartial(name string, folder string) {
	pkg := helper.LastOfSlice[string](strings.Split(folder, "/"))

	// The name of the file to create.
	fileName := folder + "/" + name + ".go"

	structName := cases.Title(language.English, cases.NoLower).String(strings.Replace(fileName, folder+"/", "", -1))

	structName = buildStructName(structName)

	moduleName := helper.GetModuleName()

	CaseStructName := "Case" + structName

	// Write content to the file.
	content :=
		`package ` + pkg + `

import "` + moduleName + `/pkg/app"

// ` + structName + ` is a method call without parsing the required parameters
func ` + structName + `() app.UseCase {
	return ` + CaseStructName + `{}
}

type ` + CaseStructName + ` struct{}

// Handle is main method of case handler
func (c ` + CaseStructName + `) Handle(uch app.UseCaseHandler) {
	//TODO implement me
	panic("implement me")
}
`
	err := file.Create(fileName, content)
	if err != nil {

	}
}
