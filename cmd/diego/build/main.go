package build

import (
	"errors"
	"fmt"
	"github.com/daewu14/golang-base/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
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
	// Replace the URL below with the Git repository you want to clone
	repoURL := "https://github.com/daewu14/golang-base.git"

	// base project name
	baseProjectName := "github.com/daewu14/golang-base"

	// Set the destination directory where the repository will be cloned
	destinationDir := projectName

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
	sCommand := "cd " + destinationDir + " && go mod edit -module=" + destinationDir + " && go mod tidy && rm -rf .git"

	err = exec.Command("sh", "-c", sCommandReplace).Run()
	if err != nil {
		logger.Fatal("Error buildProject sCommandReplace : " + err.Error())
	}
	err = exec.Command("sh", "-c", sCommand).Run()
	if err != nil {
		logger.Fatal("Error buildProject sCommand : " + err.Error())
	}
}
