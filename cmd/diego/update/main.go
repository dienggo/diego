package update

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

func New(nextVersion string, currentVersion string) update {
	return update{nextVersion, currentVersion}
}

type update struct {
	nextVersion    string
	currentVersion string
}

func (c update) Command() *cli.Command {
	return &cli.Command{
		Name:    "update",
		Aliases: []string{"updt", "upgrade"},
		Usage:   "Get update diego version (beta)",
		Action:  c.update,
	}
}

func (c update) update(context *cli.Context) error {
	cmd := exec.Command("go", "clean")
	_ = c.setOutputCmd(cmd)
	cmd = exec.Command("go", "install", "github.com/dienggo/diego/cmd/diego@"+c.nextVersion)
	err := c.setOutputCmd(cmd)
	if err != nil {
		cmd = exec.Command("go", "install", "github.com/dienggo/diego/cmd/diego@"+c.currentVersion)
		_ = c.setOutputCmd(cmd)
	}
	cmd = exec.Command("diego", "-v")
	_ = c.setOutputCmd(cmd)
	return nil
}

func (c update) setOutputCmd(cmd *exec.Cmd) error {
	// Set the output to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the git clone command
	err := cmd.Run()
	if err != nil {
		log.Error("Error get update : " + err.Error())
		return err
	}
	return nil
}
