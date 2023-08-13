package generate

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/dienggo/diego/pkg/file"
	"github.com/dienggo/diego/pkg/helper"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
)

type command struct{}

func (c command) Command() *cli.Command {
	return &cli.Command{
		Name:    "command",
		Aliases: []string{"cmd"},
		Usage:   "Generate/make command by boilerplate",
		Action: func(context *cli.Context) error {
			if context.Args().Get(0) == "" {
				return errors.New("command name can not be empty")
			}
			c.build(helper.ReplaceSpecialCharacters(context.Args().Get(0), "_"))
			return nil
		},
	}
}

func (c command) build(cmdName string) {

	// The name of the file to create.
	fileName := "cmd/" + cmdName + ".go"

	oStructName := cases.Title(language.English, cases.NoLower).String(strings.Replace(fileName, "cmd/", "", -1))

	oStructName = buildStructName(oStructName)

	structName := helper.StrFirstLetterToLower(oStructName)
	// Write content to the file.
	content :=
		`package cmd

import "github.com/urfave/cli/v2"

func New` + oStructName + `() ` + structName + ` { return ` + structName + `{} }

type ` + structName + ` struct{}

func (c ` + structName + `) Command() *cli.Command {
	return &cli.Command{
		Name:    "",
		Aliases: []string{""},
		Usage:   "",
		Action: func(context *cli.Context) error {
			// TODO : Write command logic here
			return nil
		},
	}
}
`
	err := file.Create(fileName, content)
	if err != nil {
		return
	}

	// Open the input file for reading
	inputFile, err := os.Open("provider/cmd.go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer inputFile.Close()

	// Create a temporary output file
	outputFile, err := os.Create("provider/cmd.tmp.go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	for scanner.Scan() {
		line := scanner.Text()

		// Replace the desired value in the line
		// For example, replace "oldValue" with "newValue"
		newLine := strings.Replace(line,
			"return cli.Commands{ // List-Of-Generated-Commands-By-Diego",
			"return cli.Commands{ // List-Of-Generated-Commands-By-Diego\n\t\tcmd.New"+oStructName+"().Command(),",
			-1)

		// Write the modified line to the output file
		_, err := writer.WriteString(newLine + "\n")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	// Flush any remaining buffered data and close the writer
	writer.Flush()

	os.Remove("provider/cmd.go")
	// Rename the output file to the original file name
	err = os.Rename("provider/cmd.tmp.go", "provider/cmd.go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File values replaced successfully.")
}
