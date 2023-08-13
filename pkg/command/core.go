package command

import "github.com/urfave/cli/v2"

type ICommand interface {
	Command() *cli.Command
}
