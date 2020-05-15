package migrationfile

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nycdavid/phobos/dbconnector"
	"github.com/nycdavid/phobos/migrator/migrationaction"
)

type MigrationFile struct {
	Action     string                 `json:"action"`
	Parameters map[string]interface{} `json:"parameters"`
}

func Migrate(file string, dbo *dbconnector.DBO) {
	migrationFile := readFile(file)

	migrationaction.Actions[migrationFile.Action](dbo, migrationFile.Parameters)
}

func readFile(file string) *MigrationFile {
	filePath := fmt.Sprintf("db/migrations/%s", file)

	fileIO, e := os.Open(filePath)
	if e != nil {
		log.Fatal(e)
	}

	var migration MigrationFile
	e = json.NewDecoder(fileIO).Decode(&migration)
	if e != nil {
		log.Fatal(e)
	}

	return &migration
}
