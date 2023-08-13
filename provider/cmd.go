package provider

import (
	"github.com/dienggo/diego/cmd"
	"github.com/dienggo/diego/pkg/database"
	"github.com/dienggo/diego/pkg/environment"
	"github.com/dienggo/diego/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
)

// Don't remove or edit manually all comments on this method
// commands : Slice of command registry
func commands() cli.Commands {
	return cli.Commands{ // List-Of-Generated-Commands-By-Diego
		cmd.NewHttp(new(App).Start).Command(),
	}
}

// RunCmd : main method to run cmd app
func RunCmd() {
	// Load environment
	environment.Load()

	logger.SetJSONFormatter()
	app := &cli.App{
		Name:     "github.com/dienggo/diego",
		Version:  "v1.0.0",
		Commands: commands(),
		Usage:    "",
	}

	if err := app.Run(os.Args); err != nil {
		logger.Error("Main Error Command", logger.SetField("error", err.Error()))
	}

	database.Close()
}
