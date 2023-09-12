package build

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/dienggo/diego/pkg/helper"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func New() build {
	return build{}
}

type build struct{}

func (b build) Command() *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build new project by boilerplate,\nexample : diego b project_name",
		Action: func(context *cli.Context) error {
			if context.Args().Get(0) == "" {
				return errors.New("project name can not be empty")
			}
			buildProject(context.Args().Get(0))
			return nil
		},
	}
}

func buildProject(projectName string) {
	// Repository URL
	repoURL := "https://github.com/dienggo/diego.git"

	// base project name
	baseProjectName := "github.com/dienggo/diego"

	// Set the destination directory where the repository will be cloned
	destinationDir := helper.ReplaceSpecialCharacters(projectName, "_")

	cmd := exec.Command("git", "clone", repoURL, destinationDir)

	// Set the output to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the git clone command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error cloning repository:", err)
		return
	}

	fmt.Println("Repository cloned successfully.")

	sCommandReplace := "cd " + destinationDir + " && find . -type f -name '*.go' -exec sed -i '' 's#" + baseProjectName + "#" + destinationDir + "#g' {} +\n"
	sCommand := "cd " + destinationDir + " && go mod edit -module=" + destinationDir + " && go get && rm -rf .git && rm -rf cmd/diego"

	terminal := "sh"
	if runtime.GOOS == "windows" {
		terminal = "cmd"
	}

	if runtime.GOOS == "linux" {
		sCommandReplace = `grep -rlZ "` + baseProjectName + `" . | xargs -0 sed -i '` + destinationDir + `'`
	}

	err = exec.Command(terminal, "-c", sCommandReplace).Run()
	if err != nil {
		err = nil
		buildOnUbuntuOS(baseProjectName, destinationDir)
	}
	err = exec.Command(terminal, "-c", sCommand).Run()
	if err != nil {
		log.Fatal("Error buildProject sCommand : ", err.Error())
	}
}

func buildOnUbuntuOS(baseProjectName string, destinationDir string) {
	// Define the directory where you want to replace the string.
	directory := destinationDir

	// Define the old and new strings.
	oldString := baseProjectName
	newString := destinationDir

	// Run the find and sed commands using exec.Command.
	cmdFind := exec.Command("find", directory, "-type", "f", "-name", "*.go")
	cmdSed := exec.Command("xargs", "sed", "-i", fmt.Sprintf("s|%s|%s|g", oldString, newString))

	// Pipe the output of cmdFind to cmdSed.
	cmdSed.Stdin, _ = cmdFind.StdoutPipe()

	// Start both commands.
	if err := cmdFind.Start(); err != nil {
		fmt.Println("Error starting find:", err)
		os.Exit(1)
	}

	if err := cmdSed.Start(); err != nil {
		fmt.Println("Error starting sed:", err)
		os.Exit(1)
	}

	// Wait for both commands to finish.
	if err := cmdFind.Wait(); err != nil {
		fmt.Println("Error waiting for find:", err)
		os.Exit(1)
	}

	if err := cmdSed.Wait(); err != nil {
		fmt.Println("Error waiting for sed:", err)
		os.Exit(1)
	}

	fmt.Println("String replacement complete.")
}
