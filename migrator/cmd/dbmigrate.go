package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	dbmigrateCmd := NewDbMigrateCommand()
	rootCmd.AddCommand(dbmigrateCmd)
}

func NewDbMigrateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dbmigrate",
		Short: "Migrate database using migration files",
		Long: "Reads in migration JSON files and migrates to the " +
			"correct version of the database",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Starting migration(s)...")
		},
	}

	return cmd
}
