package cmd

import (
	"context"
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
		Long: "Parse the configuration file and create a " +
			"database with that name",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := cmd.Context().Value(ContextKey("configFile")).(*ConfigFile)
			envStr, e := cmd.PersistentFlags().GetString("environment")
			if e != nil {
				// [TODO] handle error
				log.Print(e)
			}

			env := cfg.Environments[envStr]

			dbc, e := NewBaseDbConnection(env)
			if e != nil {
				// [TODO] handle error
				log.Print(e)
			}

			if dbc.DbExists(env["name"].(string)) {
				log.Println("Database exists")
			} else {
				dbc.TableCreate(env["name"].(string))
			}
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
