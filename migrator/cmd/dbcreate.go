package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	dbcreateCmd.PersistentFlags().StringVarP(
		&Environment,
		"environment",
		"e",
		"development",
		"Application environment to use",
	)

	rootCmd.AddCommand(dbcreateCmd)
}

var Environment string

var dbcreateCmd = &cobra.Command{
	Use:   "dbcreate",
	Short: "Create database from config file",
	Long: `Parse the configuration file and create a
		database with that name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating database...")
		fmt.Println(cmd.Flag("environment").Value)
	},
}
