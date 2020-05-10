package migrationfile

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type MigrationFile struct {
	Action     string                 `json:"action"`
	Parameters map[string]interface{} `json:"parameters"`
}

func Migrate(file string) {
	migrationFile := readFile(file)

	fmt.Println(migrationFile.Parameters["columns"])
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
