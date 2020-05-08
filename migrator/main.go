package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Migration struct {
	Action     string                 `json:"action"`
	Parameters map[string]interface{} `json:"parameters"`
}

var actions = map[string]string{
	"create_table": "CREATE TABLE",
}

func main() {
	connStr := "postgres://localhost/postgres?sslmode=disable"
	db, e := sql.Open("postgres", connStr)

	if e != nil {
		log.Fatal(e)
	}

	var mig Migration
	dat, e := os.Open("migrator/create_table.json")
	if e != nil {
		log.Fatal(e)
	}
	e = json.NewDecoder(dat).Decode(&mig)
	if e != nil {
		log.Fatal(e)
	}

	var migrationStr bytes.Buffer
	migrationStr.WriteString(actions[mig.Action])
	migrationStr.WriteString(" ")
	migrationStr.WriteString(mig.Parameters["name"].(string))
	migrationStr.WriteString(" ")
	migrationStr.WriteString("")
	fmt.Println(mig.Parameters["columns"].([]map[string]string))

	_, e = db.Query("CREATE TABLE items (id int);")

	if e != nil {
		log.Fatal(e)
	}
}
