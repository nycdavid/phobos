package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/nycdavid/phobos/dbconnector"
	"github.com/nycdavid/phobos/migrator/migrationfile"

	_ "github.com/lib/pq"
)

type SchemaMigration struct {
	Version int
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	action := os.Args[1]
	env := os.Args[2]
	if action != "" {
		action = os.Args[1]
	}
	if env != "" {
		env = os.Args[2]
	}

	root := os.Getenv("PROJECT_ROOT")
	configPath := path.Join(root, "db/config.json")
	configFile := NewConfigFile(configPath)

	if action == "db_create" {
		dbName := configFile.Environments[env]["dbname"]
		dbo, _ := sql.Open(
			"postgres",
			"host=localhost dbname=postgres sslmode=disable user=davidko",
		)

		// Create database
		_, e := dbo.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
		if e != nil {
			log.Fatalf("Error creating database: %s", e)
		}

		dbo.Close()
		dbconn, e := NewDbConnection(configFile.Environments[env])
		if e != nil {
			log.Fatal(e)
		}

		// Create schema_migrations table
		_, e = dbconn.Dbo.Exec(`CREATE TABLE schema_migrations (
			version int
		)`)

		if e != nil {
			log.Fatalf("Error creating table: %s", e)
		}
	} else {
		dbconn, e := NewDbConnection(configFile.Environments[env])
		if e != nil {
			log.Fatal(e)
		}

		version := currentVersion(dbconn)

		files := migrationsToRun(version)

		for i, file := range files {
			migrationfile.Migrate(file, dbconn)

			newVersion := version + i + 1
			_, e := dbo.Conn.Query(fmt.Sprintf(
				"UPDATE schema_migrations SET version = %d",
				newVersion,
			))

			if e != nil {
				log.Fatalf("Error setting version in schema_migrations: %s", e)
			}
		}
	}
}

func migrationsToRun(version int) []string {
	files, e := ioutil.ReadDir("db/migrations")
	if e != nil {
		log.Fatalf("Error reading migrations: %s", e)
	}

	var migrations []string
	for _, file := range files {
		fileVersion, e := strconv.Atoi(strings.Split(file.Name(), "_")[0])
		if e != nil {
			log.Fatal(e)
		}

		if version < fileVersion {
			migrations = append(migrations, file.Name())
		}
	}

	return migrations
}

func currentVersion(dbo *dbconnector.DBO) int {
	row := dbo.Conn.QueryRow("SELECT schema_migrations.version FROM schema_migrations")

	var version uint8 = 0
	e := row.Scan(&version)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(version)

	return int(version)
}
