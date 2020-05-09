package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DbConfig struct {
	Environments map[string]Environment `json:"environments"`
}

type Environment struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var Tasks = map[string]func(){
	"db:setup":   DbSetup,
	"db:migrate": DbMigrate,
}

func main() {
	taskName := os.Args[1]

	Tasks[taskName]()
}

func DbSetup() {
	environment := os.Args[2]

	file, e := os.Open("db/config.json")
	if e != nil {
		log.Fatal(e)
	}

	var conf DbConfig
	e = json.NewDecoder(file).Decode(&conf)
	if e != nil {
		log.Fatal(e)
	}

	url := conf.Environments[environment].URL
	dbo, e := sql.Open("postgres", url)
	if e != nil {
		log.Fatal(e)
	}

	query := "CREATE TABLE schema_migrations (version int);"
	_, e = dbo.Query(query)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("Created schema_migrations table....")

	_, e = dbo.Query("INSERT INTO schema_migrations VALUES (0);")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("Created version row....")
}

func DbMigrate() {
}
