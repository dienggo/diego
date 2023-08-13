package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewHttp(StartHttp func()) http { return http{StartHttp} }

type http struct {
	startHttp func()
}

func (c http) Command() *cli.Command {
	return &cli.Command{
		Name:  "http",
		Usage: "http [command]",
		Subcommands: cli.Commands{
			&cli.Command{
				Name:    "serve",
				Aliases: []string{"server", "start", "run"},
				Usage:   "Run http server",
				Action: func(context *cli.Context) error {
					c.startHttp()
					return nil
				},
			},
		},
	}
}
