package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // `_` => will not use any syntax
)

// information from postgresql database
const (
	host     = "localhost"
	port     = 1234
	user     = "postgres"
	password = "270719"
	dbname   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

func main() {
	// connect info from psql
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// validate info
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// create a connection to database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database.")
}
