package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

type ContextKey string

func init() {
	dbcreateCmd := NewDbCreateCommand()
	rootCmd.AddCommand(dbcreateCmd)
}

var Environment string

func NewDbCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dbcreate",
		Short: "Create database from config file",
		Long: `Parse the configuration file and create a
		database with that name`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cfg := NewConfigFile("/Users/davidko/projects/phobos/db/config.json")

	ctx := context.WithValue(
		context.Background(),
		ContextKey("configFile"),
		cfg,
	)
	cmd.ExecuteContext(ctx)

	cmd.PersistentFlags().StringVarP(
		&Environment,
		"environment",
		"e",
		"development",
		"Application environment to use",
	)

	return cmd
}
