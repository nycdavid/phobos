package cmd

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestDbCreate_UsesDevelopmentByDefault(t *testing.T) {
	dbcreateCmd := NewDbCreateCommand()

	fVal := dbcreateCmd.Flag("environment").Value.String()
	if fVal != "development" {
		t.Errorf("Expected %s got %s", "development", fVal)
	}
}

func TestDbCreate_ConfigFileInCtx(t *testing.T) {
	defer cleanUpDb(t)
	os.Args[1] = "-e=test"

	cfgPath := "test/db/config.json"
	ctx := ConfigCtx(cfgPath)

	dbcreateCmd := NewDbCreateCommand()
	dbcreateCmd.ExecuteContext(ctx)

	configFile := dbcreateCmd.Context().Value(ContextKey("configFile"))

	if configFile == nil {
		t.Errorf("Expected non-nil config file")
	}
}

func TestDbCreate_CreatesDatabase(t *testing.T) {
	defer cleanUpDb(t)

	cfgPath := "/Users/davidko/projects/phobos/db/config.json"
	ctx := ConfigCtx(cfgPath)

	os.Args[1] = "-e=test"

	dbcreateCmd := NewDbCreateCommand()
	dbcreateCmd.ExecuteContext(ctx)

	dbo, e := sql.Open(
		"postgres",
		"host=localhost port=5432 database=test password=password user=postgres database=test sslmode=disable",
	)
	if e != nil {
		t.Error(e)
	}

	e = dbo.Ping()

	if e != nil {
		t.Error(e)
	}
}
