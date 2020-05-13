package migrationaction

import (
	"fmt"

	"../../dbconnector"
)

var Actions = map[string]func(*dbconnector.DBO, map[string]interface{}){
	"create_table": CreateTable,
}

func CreateTable(dbo *dbconnector.DBO, params map[string]interface{}) {
	fmt.Println(params)
}
