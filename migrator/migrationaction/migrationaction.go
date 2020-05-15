package migrationaction

import (
	"bytes"
	"fmt"
	"log"

	"github.com/nycdavid/phobos/dbconnector"
)

var Actions = map[string]func(*dbconnector.DBO, map[string]interface{}){
	"create_table": CreateTable,
	"add_column":   AddColumn,
}

func CreateTable(dbo *dbconnector.DBO, params map[string]interface{}) {
	tableName := params["name"].(string)
	columns := params["columns"].([]interface{})

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", tableName))
	for i, _ := range columns {
		column := columns[i].(map[string]interface{})

		b.WriteString(fmt.Sprintf(
			"  %s %s",
			column["name"].(string),
			column["data_type"].(string),
		))

		if column["primary_key"] != nil && column["primary_key"].(bool) {
			b.WriteString(" PRIMARY KEY")
		}

		if i != len(columns)-1 {
			b.WriteString(",")
		}

		b.WriteString("\n")
	}
	b.WriteString(");")

	_, e := dbo.Conn.Query(b.String())

	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(b.String())
}

func AddColumn(dbo *dbconnector.DBO, params map[string]interface{}) {
	tableName := params["name"].(string)
	columns := params["columns"].([]interface{})

	var b bytes.Buffer
	for i, _ := range columns {
		column := columns[i].(map[string]interface{})
		b.WriteString(fmt.Sprintf(
			"ALTER TABLE %s ADD COLUMN %s;\n",
			tableName,
			column["name"].(string),
		))
	}

	fmt.Printf(b.String())
}
