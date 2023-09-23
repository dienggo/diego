package generate

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/dienggo/diego/pkg/file"
	"github.com/dienggo/diego/pkg/helper"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
)

type task struct{}

func (c task) Command() *cli.Command {
	return &cli.Command{
		Name:    "task",
		Aliases: []string{"bg-task"},
		Usage:   "Generate/make task by boilerplate",
		Action: func(context *cli.Context) error {
			if context.Args().Get(0) == "" {
				return errors.New("task name can not be empty")
			}
			c.build(helper.ReplaceSpecialCharacters(context.Args().Get(0), "_"))
			return nil
		},
	}
}

func (c task) build(taskName string) {

	path := "app/background"
	exist := helper.PathExist(path)
	if !exist {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Error("Error: ", err.Error())
			return
		}
	}

	moduleName := helper.GetModuleName()

	// The name of the file to create.
	fileName := "app/background/" + taskName + ".go"

	oStructName := cases.Title(language.English, cases.NoLower).String(strings.Replace(fileName, "app/background/", "", -1))

	oStructName = buildStructName(oStructName)

	structName := helper.StrFirstLetterToLower(oStructName)
	// Write content to the file.
	content :=
		`package background

import (
	"context"
	"` + moduleName + `/pkg/task"
	"time"
)

// New` + oStructName + ` is method caller on app provider
func New` + oStructName + `() task.ITask { return ` + structName + `{} }

type ` + structName + ` struct{}

// Task is main method to handle when task executed
func (t ` + structName + `) Task() task.Scheduler {
	// sleep delay run every 10 second after method completed / without overlapping
	// use .Overlaps(true) to set with overlapping
	duration := 10 * time.Second // just for example, please update soon
	return *task.Scheduler{Run: func(ctx context.Context) {
		// TODO : write your logic here
		
	}}.Every(task.Duration{SleepDuration: &duration}).Overlaps(false).Name("` + oStructName + ` - Kernel")
}
`
	err := file.Create(fileName, content)
	if err != nil {
		return
	}

	// Open the input file for reading
	inputFile, err := os.Open("provider/task.go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer inputFile.Close()

	// Create a temporary output file
	outputFile, err := os.Create("provider/task.tmp.go")
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
			"return []task.ITask{ // List-Of-Generated-Task-By-Diego",
			"return []task.ITask{ // List-Of-Generated-Task-By-Diego\n\t\tbackground.New"+oStructName+"(),",
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

	os.Remove("provider/task.go")
	// Rename the output file to the original file name
	err = os.Rename("provider/task.tmp.go", "provider/task.go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File values replaced successfully.")
}
