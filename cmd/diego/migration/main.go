package migration

import (
	"errors"
	"github.com/daewu14/golang-base/config"
	"github.com/daewu14/golang-base/pkg/database"
	"github.com/daewu14/golang-base/pkg/environment"
	"github.com/daewu14/golang-base/pkg/helper"
	"github.com/daewu14/golang-base/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"strings"
)

func New() migration {
	return migration{}
}

type migration struct{}

func (m migration) Command() *cli.Command {
	return &cli.Command{
		Name:    "migration",
		Aliases: []string{"migrate", "m"},
		Usage:   "Run migration depend on goose",
		Action: func(context *cli.Context) error {
			if context.Args().Get(0) == "" {
				return errors.New("argument migration can not be empty")
			}
			args := ""
			args = strings.Join(context.Args().Slice(), " ")
			var cdMigration = false
			var canCdMigration = []string{"up", "up-by-one", "up-to", "down", "down-to", "validate", "redo", "reset", "status"}
			cdMigration = helper.SliceContain(canCdMigration, context.Args().Get(0))
			migrate(args, cdMigration)
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Creates new migration file with the current timestamp",
				Action: func(context *cli.Context) error {
					if context.Args().Get(0) == "" {
						return errors.New("argument migration create can not be empty")
					}
					args := ""
					args = strings.Join(context.Args().Slice(), " ")
					migrate("create "+args, true)
					return nil
				},
			},
		},
	}
}

func migrate(get string, cdMigration bool) {
	environment.Load()
	database.Open()
	defer database.Close()

	prefix := ""
	if cdMigration {
		prefix = "cd migrations && "
	}
	cmd := exec.Command("sh", "-c", prefix+`goose `+config.Database().Connection+` "`+database.GetMainDsn()+`" `+get)

	// Set the output to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		logger.Error("Error Migration", logger.SetField("error", err.Error()))
		return
	}
}
