package config

import (
	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "simple backend with golang",
}

func init() {
	initCommand.AddCommand(ServerCmd)
	initCommand.AddCommand(orm.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)
	return initCommand.Execute()
}
