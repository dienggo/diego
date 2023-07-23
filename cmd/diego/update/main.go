package update

import (
	"github.com/daewu14/golang-base/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

func New() update {
	return update{}
}

type update struct{}

func (c update) Command() *cli.Command {
	return &cli.Command{
		Name:    "update",
		Aliases: []string{"updt", "upgrade"},
		Usage:   "Get update diego version",
		Action:  c.update,
	}
}

func (c update) update(context *cli.Context) error {
	cmd := exec.Command("go", "clean")
	c.setOutputCmd(cmd)
	cmd = exec.Command("go", "install", "github.com/daewu14/golang-base/cmd/diego@latest")
	c.setOutputCmd(cmd)
	cmd = exec.Command("sh", "-c", "diego -v")
	c.setOutputCmd(cmd)
	return nil
}

func (c update) setOutputCmd(cmd *exec.Cmd) {
	// Set the output to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the git clone command
	err := cmd.Run()
	if err != nil {
		logger.Error("Error get update : " + err.Error())
	}
}
