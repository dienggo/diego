package generate

import (
	"errors"
	command2 "github.com/dienggo/diego/pkg/command"
	"github.com/urfave/cli/v2"
)

func New() generate {
	return generate{}
}

type generate struct{}

func (gen generate) Command() *cli.Command {
	return &cli.Command{
		Name:        "generate",
		Aliases:     []string{"make", "gen"},
		Usage:       "Generate/make component by boilerplate",
		Subcommands: gen.subCommands(),
	}
}

func (gen generate) subCommands() (cmd []*cli.Command) {
	cmd = append(cmd, &cli.Command{
		Name:  "controller",
		Usage: "Generate/make controller by boilerplate",
		Action: func(context *cli.Context) error {
			if context.Args().Get(0) == "" {
				return errors.New("controller name can not be empty")
			}
			Controller(context.Args().Get(0))
			return nil
		},
	})

	registered := []command2.ICommand{
		service{},
		middleware{},
		command{},
		mockRepo{},
		ucase{},
	}

	for _, iCommand := range registered {
		cmd = append(cmd, iCommand.Command())
	}
	return cmd
}
