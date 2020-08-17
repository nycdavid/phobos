package cmd

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
)

var Actions = map[string]func(*sql.DB, map[string]interface{}){
	"create_table":     CreateTable,
	"add_column":       AddColumn,
	"make_foreign_key": MakeForeignKey,
}

func CreateTable(dbo *sql.DB, params map[string]interface{}) {
	tableName := params["name"].(string)
	columns := params["columns"].([]interface{})

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", tableName))
	for i, _ := range columns {
		column := columns[i].(map[string]interface{})

		b.WriteString(fmt.Sprintf(
			"  %s",
			column["name"].(string),
		))

		if column["serial"] != nil && column["serial"].(bool) {
			b.WriteString(" SERIAL")
		} else {
			b.WriteString(fmt.Sprintf(" %s", column["data_type"].(string)))
		}

		if column["primary_key"] != nil && column["primary_key"].(bool) {
			b.WriteString(" PRIMARY KEY")
		}

		if i != len(columns)-1 {
			b.WriteString(",")
		}

		b.WriteString("\n")
	}
	b.WriteString(");")

	_, e := dbo.Query(b.String())

	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(b.String())
}

func AddColumn(dbo *sql.DB, params map[string]interface{}) {
	tableName := params["name"].(string)
	columns := params["columns"].([]interface{})

	var b bytes.Buffer

	for i, _ := range columns {
		column := columns[i].(map[string]interface{})
		b.WriteString(fmt.Sprintf(
			"ALTER TABLE %s ADD COLUMN %s %s;\n",
			tableName,
			column["name"].(string),
			column["data_type"].(string),
		))
	}

	_, e := dbo.Query(b.String())
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(b.String())
}

func MakeForeignKey(dbo *sql.DB, params map[string]interface{}) {
	tableName := params["table_name"].(string)
	foreignKeyColumn := params["column"].(string)
	referenceTable := params["references"].(string)

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("ALTER TABLE %s\n", tableName))
	b.WriteString(fmt.Sprintf(
		"  ADD CONSTRAINT fk_%s_%s FOREIGN KEY (%s) REFERENCES %s (id);",
		tableName,
		referenceTable,
		foreignKeyColumn,
		referenceTable,
	))

	_, e := dbo.Query(b.String())
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(b.String())
}
