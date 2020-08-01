package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "A database migrator",
		Long:  "Migrator is a Rails-inspired database migrator",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
