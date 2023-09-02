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

type middleware struct{}

func (c middleware) Command() *cli.Command {
	return &cli.Command{
		Name:  "middleware",
		Usage: "Generate/make middleware by boilerplate",
		Action: func(context *cli.Context) error {
			if context.Args().Get(0) == "" {
				return errors.New("middleware name can not be empty")
			}
			c.build(helper.ReplaceSpecialCharacters(context.Args().Get(0), "_"))
			return nil
		},
	}
}

func (c middleware) build(middlewareName string) {
	// The name of the file to create.
	fileName := "app/middleware/" + middlewareName + ".go"

	structName := cases.Title(language.English, cases.NoLower).String(strings.Replace(fileName, "app/middleware/", "", -1))

	structName = buildStructName(structName)
	// Write content to the file.
	content :=
		`package middleware

import (
	"net/http"
)

func ` + structName + `(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO : write your business logic here 
		h.ServeHTTP(w, r)
	})
}`
	err := file.Create(fileName, content)
	if err != nil {
		return
	}
}
