package migrationfile

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Migration struct {
	Action     string                 `json:"action"`
	Parameters map[string]interface{} `json:"parameters"`
}

func Migrate(file string) {
	filePath := fmt.Sprintf("db/migrations/%s", file)

	fileIO, e := os.Open(filePath)
	if e != nil {
		log.Fatal(e)
	}

	var migration Migration
	e = json.NewDecoder(fileIO).Decode(&migration)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(migration.Parameters["columns"])
}
