package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
			fmt.Println("Creating database...")
			fmt.Println(cmd.Flag("environment"))
		},
	}

	cmd.PersistentFlags().StringVarP(
		&Environment,
		"environment",
		"e",
		"development",
		"Application environment to use",
	)

	return cmd
}
