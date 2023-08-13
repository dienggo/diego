package generate

import (
	"errors"
	"github.com/urfave/cli/v2"
)

func New() generate {
	return generate{}
}

type generate struct{}

func (gen generate) Command() *cli.Command {
	return &cli.Command{
		Name:    "generate",
		Aliases: []string{"make", "gen"},
		Usage:   "Generate/make component by boilerplate",
		Subcommands: []*cli.Command{
			{
				Name:  "controller",
				Usage: "Generate/make controller by boilerplate",
				Action: func(context *cli.Context) error {
					if context.Args().Get(0) == "" {
						return errors.New("controller name can not be empty")
					}
					Controller(context.Args().Get(0))
					return nil
				},
			},
			service{}.Command(),
			middleware{}.Command(),
			command{}.Command(),
		},
	}
}
