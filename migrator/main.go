package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"../dbconnector"
	"./migrationfile"

	_ "github.com/lib/pq"
)

type SchemaMigration struct {
	Version int
}

func main() {
	env := os.Args[1]
	if env != "" {
		env = os.Args[1]
	}

	dbo := dbconnector.NewDBO(env)
	version := currentVersion(dbo)

	files := migrationsToRun(version)

	for _, file := range files {
		migrationfile.Migrate(file)
	}
}

func migrationsToRun(version int) []string {
	files, e := ioutil.ReadDir("db/migrations")
	if e != nil {
		log.Fatal(e)
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
	row, e := dbo.Conn.Query("SELECT schema_migrations.version FROM schema_migrations")
	if e != nil {
		log.Fatal(e)
	}

	var schemaMig SchemaMigration
	row.Scan(&schemaMig)

	return schemaMig.Version
}
