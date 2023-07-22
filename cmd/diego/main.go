package main

import (
	"fmt"
	"github.com/daewu14/golang-base/cmd/diego/build"
	"github.com/daewu14/golang-base/cmd/diego/generate"
	"github.com/daewu14/golang-base/cmd/diego/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type ICommand interface {
	Command() *cli.Command
}

func main() {
	app := &cli.App{
		Name:  "diego",
		Usage: "Base REST API lite project",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
		Version: "v1.0.3",
		Authors: []*cli.Author{
			{
				Name:  "Daewu Bintara",
				Email: "daewu.bintara1996@gmail.com",
			},
		},
		Commands: commands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func commandRegistry() []ICommand {
	return []ICommand{
		build.New(),
		generate.New(),
		server.New(),
	}
}

func commands() (cmnds cli.Commands) {

	for _, command := range commandRegistry() {
		cmnds = append(cmnds, command.Command())
	}

	return cmnds
}
