package dbconnector

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

type DbConfig struct {
	Envs map[string]Environment `json:"environments"`
}

type Environment struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type DBO struct {
	Conn *sql.DB
}

func NewDBO(env string) *DBO {
	var conf DbConfig
	readDbConfig(&conf)

	dbo := connectDb(conf.Envs[env].URL)

	return &DBO{Conn: dbo}
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
