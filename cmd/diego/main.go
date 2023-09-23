package main

import (
	"fmt"
	"os"

	"github.com/dienggo/diego/cmd/diego/build"
	"github.com/dienggo/diego/cmd/diego/generate"
	"github.com/dienggo/diego/cmd/diego/migration"
	"github.com/dienggo/diego/cmd/diego/update"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	name        = "diego"
	usage       = "diego"
	greetings   = "Hi, welcome to diego framework!"
	version     = "v1.4.2"
	nextVersion = "v1.4.3"
)

type ICommand interface {
	Command() *cli.Command
}

func main() {
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
			{
				Name:  "Ahdiyat Lalu",
				Email: "ahdiyatlalu@gmail.com",
			},
			{
				Name:  "Gema Antika Hariadi",
				Email: "gemaantikahr@gmail.com",
			},
		},
		Commands: commands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Error("Error Command : ", err.Error())
	}
}

func commandRegistry() []ICommand {
	return []ICommand{
		build.New(),
		generate.New(),
		migration.New(),
		update.New(nextVersion, version),
	}
}

func commands() (cmnds cli.Commands) {

	for _, command := range commandRegistry() {
		cmnds = append(cmnds, command.Command())
	}

	return cmnds
}
