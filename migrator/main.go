package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	env := os.Args[1]
	if env != "" {
		env = os.Args[1]
	}

	dbo := dbconnector.NewDBO(env)
	version := currentVersion(dbo)

	files := migrationsToRun(version)

	for i, file := range files {
		migrationfile.Migrate(file, dbo)

		newVersion := version + i + 1
		_, e := dbo.Conn.Query(fmt.Sprintf(
			"UPDATE schema_migrations SET version = %d",
			newVersion,
		))

		if e != nil {
			log.Fatal(e)
		}
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
	row := dbo.Conn.QueryRow("SELECT schema_migrations.version FROM schema_migrations")

	var version uint8 = 0
	e := row.Scan(&version)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(version)

	return int(version)
}
