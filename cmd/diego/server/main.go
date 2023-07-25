package server

import (
	"github.com/dienggo/diego/provider"
	"github.com/urfave/cli/v2"
)

func New() server { return server{} }

type server struct{}

func (server server) Command() *cli.Command {
	return &cli.Command{
		Name:    "serve",
		Aliases: []string{"server"},
		Usage:   "Run server",
		Action: func(context *cli.Context) error {
			new(provider.App).Start()
			return nil
		},
	}
}
