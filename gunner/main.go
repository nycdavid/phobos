package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"../dbconnector"
)

type DbConfig struct {
	Environments map[string]Environment `json:"environments"`
}

type Environment struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var Tasks = map[string]func(string){
	"db:setup":   DbSetup,
	"db:migrate": DbMigrate,
}

func main() {
	taskName := os.Args[1]
	environment := os.Args[2]

	Tasks[taskName](environment)
}

func DbSetup(environment string) {
	var conf DbConfig
	readDbConfig(&conf)

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

type SchemaMigration struct {
	Version int
}

func DbMigrate(env string) {
	dbo := dbconnector.NewDBO(env)

	rows, e := dbo.Conn.Query("SELECT version FROM schema_migrations")
	if e != nil {
		log.Fatal(e)
	}

	var schemaMigration SchemaMigration
	rows.Scan(&schemaMigration)

	files, e := ioutil.ReadDir("db/migrations/")
	if e != nil {
		log.Fatal(e)
	}

	var migrations []string
	version := strconv.Itoa(schemaMigration.Version)
	for _, file := range files {
		if version < file.Name() {
			migrations = append(migrations, file.Name())
		}
	}
}

func readDbConfig(conf *DbConfig) {
	file, e := os.Open("db/config.json")
	if e != nil {
		log.Fatal(e)
	}

	e = json.NewDecoder(file).Decode(&conf)
	if e != nil {
		log.Fatal(e)
	}
}

func connectDb(url string) *sql.DB {
	dbo, e := sql.Open("postgres", url)
	if e != nil {
		log.Fatal(e)
	}

	return dbo
}
