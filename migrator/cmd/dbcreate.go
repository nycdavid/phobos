package cmd

import (
	"context"
	"fmt"
	"log"

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
			fmt.Println("Inside Run")
			cfg := cmd.Context().Value(ContextKey("configFile")).(*ConfigFile)
			env, e := cmd.PersistentFlags().GetString("environment")
			if e != nil {
				// [TODO] handle error
				log.Print(e)
			}
			fmt.Println(cfg.Environments[env])
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

func ConfigCtx(cfgPath string) context.Context {
	cfg := NewConfigFile(cfgPath)

	return context.WithValue(
		context.Background(),
		ContextKey("configFile"),
		cfg,
	)
}
