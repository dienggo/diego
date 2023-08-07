package main

import (
	"fmt"
	"github.com/dienggo/diego/cmd/diego/build"
	"github.com/dienggo/diego/cmd/diego/generate"
	"github.com/dienggo/diego/cmd/diego/migration"
	"github.com/dienggo/diego/cmd/diego/update"
	"github.com/dienggo/diego/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	name        = "diego"
	usage       = "diego"
	greetings   = "Hi, welcome to diego framework!"
	version     = "v1.2.2"
	nextVersion = "v1.2.3"
)

type ICommand interface {
	Command() *cli.Command
}

func main() {
	logger.SetJSONFormatter()
	app := &cli.App{
		Name:  name,
		Usage: usage,
		Action: func(c *cli.Context) error {
			fmt.Println(greetings)
			return nil
		},
		Version: version,
		Authors: []*cli.Author{
			{
				Name:  "Daewu Bintara",
				Email: "daewu.bintara1996@gmail.com",
			},
		},
		Commands: commands(),
	}

	if err := app.Run(os.Args); err != nil {
		logger.Error("Error Command", logger.SetField("error", err.Error()))
	}
}

func commandRegistry() []ICommand {
	return []ICommand{
		build.New(),
		generate.New(),
		migration.New(),
		update.New(nextVersion, version),
		// server.New(),
	}
}

func commands() (cmnds cli.Commands) {

	for _, command := range commandRegistry() {
		cmnds = append(cmnds, command.Command())
	}

	return cmnds
}
