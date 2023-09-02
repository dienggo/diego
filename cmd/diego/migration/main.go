package migration

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/dienggo/diego/config"
	"github.com/dienggo/diego/pkg/database"
	"github.com/dienggo/diego/pkg/environment"
	"github.com/urfave/cli/v2"
)

func New() migration {
	return migration{}
}

type migration struct{}

func (m migration) Command() *cli.Command {
	return &cli.Command{
		Name:      "migration",
		Aliases:   []string{"migrate", "m"},
		Usage:     "Run migration depend on goose\ncheck goose repository : https://github.com/pressly/goose",
		UsageText: m.help(),
		Action: func(context *cli.Context) error {
			if context.Args().Get(0) == "" {
				println("EXAMPLE COMMANDS:\n" + m.help())
				return nil
			}
			args := ""
			args = strings.Join(context.Args().Slice(), " ")
			migrate(args, true)
			return nil
		},
	}
}

func (m migration) help() string {
	return "" +
		"diego migration up                   Migrate the DB to the most recent version available\n" +
		"diego migration up-by-one            Migrate the DB up by 1\n" +
		"diego migration up-to VERSION        Migrate the DB to a specific VERSION\n" +
		"diego migration down                 Roll back the version by 1\n" +
		"diego migration down-to VERSION      Roll back to a specific VERSION\n" +
		"diego migration redo                 Re-run the latest migration\n" +
		"diego migration reset                Roll back all migrations\n" +
		"diego migration status               Dump the migration status for the current DB\n" +
		"diego migration version              Print the current version of the database\n" +
		"diego migration create NAME [sql|go] Creates new migration file with the current timestamp\n" +
		"diego migration fix                  Apply sequential ordering to migrations\n" +
		"diego migration validate             Check migration files without running them"
}

func migrate(get string, cdMigration bool) {
	environment.Load()
	database.Open()
	defer database.Close()

	prefix := ""
	if cdMigration {
		prefix = "cd migrations && "
	}

	terminal := "sh"
	if runtime.GOOS == "windows" {
		terminal = "cmd"
	}

	cmd := exec.Command(terminal, "-c", prefix+`goose `+config.Database().Connection+` "`+database.GetMainDsn()+`" `+get)

	// Set the output to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Error("Error Migration", err.Error())
		return
	}
}
