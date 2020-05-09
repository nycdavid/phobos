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
	Action     string                `json:"action"`
	Parameters CreateTableParameters `json:"parameters"`
}

type CreateTableParameters struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

type Column struct {
	Name       string `json:"name"`
	DataType   string `json:"data_type"`
	PrimaryKey bool   `json:"primary_key"`
}

var actions = map[string]string{
	"create_table": "CREATE TABLE",
}

var types = map[string]string{
	"jsonb":  "jsonb",
	"number": "int",
	"string": "text",
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

	var b bytes.Buffer
	write(&b, actions[mig.Action])
	write(&b, mig.Parameters.Name)
	writeColumns(&b, mig.Parameters.Columns)

	fmt.Println(b.String())

	_, e = db.Query(b.String())

	if e != nil {
		log.Fatal(e)
	}
}

func write(b *bytes.Buffer, str string) {
	b.WriteString(str)
	b.WriteString(" ")
}

func writeColumns(b *bytes.Buffer, columns []Column) {
	b.WriteString("(\n")
	for i, column := range columns {
		dataType := types[column.DataType]
		var columnLine bytes.Buffer
		columnLine.WriteString(fmt.Sprintf("%s %s", column.Name, dataType))

		if column.PrimaryKey {
			columnLine.WriteString(" PRIMARY KEY")
		}

		if i != len(columns)-1 {
			columnLine.WriteString(",")
		}

		writeLine(b, 1, columnLine.String())
	}
	b.WriteString(")")
}

func writeLine(b *bytes.Buffer, indents int, parts ...string) {
	for i := 0; i < indents; i++ {
		b.WriteString("  ")
	}

	for i, part := range parts {
		b.WriteString(part)
		b.WriteString(" ")

		if i == len(parts)-1 {
			b.WriteString("\n")
		}
	}
}
