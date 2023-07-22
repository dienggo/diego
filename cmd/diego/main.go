package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go_base_project/cmd/diego/build"
	"go_base_project/cmd/diego/generate"
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
		Version: "1.0.0",
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
	}
}

func commands() (cmnds cli.Commands) {

	for _, command := range commandRegistry() {
		cmnds = append(cmnds, command.Command())
	}

	return cmnds
}
