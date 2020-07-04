package main

import (
	"os"
	"path"
	"testing"

	_ "github.com/lib/pq"
)

const (
	project_root = "/Users/davidko/projects/phobos"
)

func TestSetup(t *testing.T) {
	os.Setenv("PROJECT_ROOT", project_root)
	os.Args = []string{"main.go", "db_create", "test"}

	main()

	configPath := path.Join(project_root, "db/config.json")
	configFile := NewConfigFile(configPath)
	testEnv := configFile.Environments["test"]
	dbo, _ := NewDbConnection(testEnv)
	defer cleanUpDb(t)

	if !(dbo.DbExists(testEnv["dbname"].(string))) {
		cleanUpDb(t)
		t.Error("Expected database to exist")
	}

	if !(dbo.TableExists("schema_migrations")) {
		cleanUpDb(t)
		t.Error("Expected table schema_migrations to exist")
	}
}
