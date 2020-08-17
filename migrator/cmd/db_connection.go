package cmd

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DbConnection struct {
	Environment map[string]interface{}
	Dbo         *sql.DB
}

func NewDbConnection(environment map[string]interface{}) (*DbConnection, error) {
	dataSourceName := fmt.Sprintf(
		"host=%s user=%s dbname=postgres password=%s sslmode=disable",
		environment["host"].(string),
		environment["user"].(string),
		environment["password"].(string),
	)

	dbo, e := sql.Open("postgres", dataSourceName)
	if e != nil {
		return &DbConnection{}, e
	}

	e = dbo.Ping()
	if e != nil {
		return &DbConnection{}, e
	}

	return &DbConnection{Dbo: dbo, Environment: environment}, nil
}

func (dbc *DbConnection) DbExists(dbname string) bool {
	res, e := dbc.Dbo.Exec(fmt.Sprintf(`
		SELECT *
		FROM pg_database
		WHERE datname='%s'
	`, dbname))

	if e != nil {
		log.Fatal(e)
	}

	rowsAffected, e := res.RowsAffected()

	if e != nil {
		log.Fatal(e)
	}

	if rowsAffected == 0 {
		return false
	}

	return true
}

func (dbc *DbConnection) TableExists(tableName string) bool {
	res, e := dbc.Dbo.Exec(fmt.Sprintf(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
		AND table_name = '%s';
	`, tableName))
	if e != nil {
		log.Fatal(e)
	}

	ct, e := res.RowsAffected()
	if e != nil {
		log.Fatal(e)
	}

	if ct == 0 {
		return false
	}

	return true
}

func (dbc *DbConnection) TableCreate(tableName string) error {
	_, e := dbc.Dbo.Exec(fmt.Sprintf(`
		CREATE DATABASE %s;
	`, tableName))
	if e != nil {
		return e
	}

	log.Print("Database created...")

	return nil
}

func (dbc *DbConnection) CurrentDb() (string, error) {
	row, e := dbc.Dbo.QueryRow("SELECT current_database()")
	if e != nil {
		return "", e
	}

	var name string
	row.Scan(&name)

	return name, nil
}

func (dbc *DbConnection) ChangeDb() error {
	_, e := dbc.Dbo.Exec()
}
