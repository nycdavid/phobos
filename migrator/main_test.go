package main

import (
	"database/sql"
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

func TestMain_MigratingEmptyDatabase(t *testing.T) {
	os.Setenv("PROJECT_ROOT", project_root)
	defer cleanUpDb(t)

	dbo, e := sql.Open(
		"postgres",
		"host=localhost database=postgres user=davidko sslmode=disable",
	)
	if e != nil {
		cleanUpDb(t)
		t.Error(e)
	}

	_, e = dbo.Exec("CREATE DATABASE test;")
	if e != nil {
		cleanUpDb(t)
		t.Error(e)
	}

	os.Args = []string{"main.go", "migrate", "test"}

	code := main()
	if code != 0 {
		cleanUpDb(t)
		t.Error("Main failed")
	}

	dbo, e = sql.Open(
		"postgres",
		"host=localhost database=test user=davidko sslmode=disable",
	)
	if e != nil {
		cleanUpDb(t)
		t.Error(e)
	}

	res, e := dbo.Exec(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
		AND table_name = 'users';
	`)
	if e != nil {
		cleanUpDb(t)
		t.Error(e)
	}

	ct, e := res.RowsAffected()
	if e != nil {
		cleanUpDb(t)
		t.Error(e)
	}

	if ct == 0 {
		t.Error("Expected users table to exist")
	}
}
