package cmd

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

// When db does not exist
func Test_DbExists(t *testing.T) {
	configFile := NewConfigFile("test/db/config.json")

	testEnv := configFile.Environments["test"]
	_, e := NewDbConnection(testEnv)

	if e == nil {
		t.Error("Expected connection error")
	}
}

// When db does exist
func Test_AfterCreationDbExists(t *testing.T) {
	configFile := NewConfigFile("test/db/config.json")

	testEnvConfig := configFile.Environments["test"]

	dataSourceName := fmt.Sprintf(
		"host=%s user=postgres dbname=postgres password=password sslmode=disable",
		testEnvConfig["host"].(string),
	)

	dbconn, e := sql.Open("postgres", dataSourceName)
	if e != nil {
		t.Errorf("Error opening database: %s", e)
	}

	defer cleanUpDb(t)

	_, e = dbconn.Exec("CREATE DATABASE test")
	if e != nil {
		cleanUpDb(t)
		t.Error(e)
	}

	testEnv := configFile.Environments["test"]
	_, e = NewDbConnection(testEnv)

	if e != nil {
		cleanUpDb(t)
		t.Error(e)
	}
}

func TestDbConnection_TableExists(t *testing.T) {
	defer cleanUpDb(t)

	configFile := NewConfigFile("test/db/config.json")
	testEnvConfig := configFile.Environments["test"]

	dbo, e := sql.Open(
		"postgres",
		"host=localhost user=postgres dbname=postgres password=password sslmode=disable",
	)

	if e != nil {
		t.Errorf("Error connecting to database: %s", e)
	}

	_, e = dbo.Exec("CREATE DATABASE test;")
	if e != nil {
		cleanUpDb(t)
		t.Errorf("Error creating db: %s", e)
	}

	dbconn, e := NewDbConnection(testEnvConfig)
	if e != nil {
		cleanUpDb(t)
		t.Errorf("Error connecting to database: %s", e)
	}

	_, e = dbconn.Dbo.Exec(`CREATE TABLE foobar (
		version integer
	);`)

	if e != nil {
		cleanUpDb(t)
		t.Error(e)
	}

	if !(dbconn.TableExists("foobar")) {
		cleanUpDb(t)
		t.Error("Expected table foobar to exist")
	}
}

func cleanUpDb(t *testing.T) {
	dbconn, e := sql.Open(
		"postgres",
		"host=localhost user=postgres password=password dbname=postgres sslmode=disable",
	)
	if e != nil {
		t.Error(e)
	}

	_, e = dbconn.Exec("REVOKE CONNECT ON DATABASE test FROM public;")
	if e != nil {
		t.Error(e)
	}
	_, e = dbconn.Exec(`SELECT pg_terminate_backend(pg_stat_activity.pid)
		FROM pg_stat_activity
		WHERE pg_stat_activity.datname = 'test';
	`)
	if e != nil {
		t.Error(e)
	}
	_, e = dbconn.Exec("DROP DATABASE test;")
	if e != nil {
		t.Error(e)
	}
}
