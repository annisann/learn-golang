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

type Employee struct {
	ID       int
	FullName string
	Email    string
	Age      int
	Division string
}

func CreateEmployee() {
	employee := Employee{}

	// sql statement to create data on database
	sqlStatement :=
		`
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *
	` // `Returning *` => return all values from newly created data

	// execute sql query
	err = db.QueryRow(sqlStatement, "Annisa N Nabila", "annisa@gmail.com", 23, "IT").
		// get return value and store in `employee`
		Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data: %+v\n", employee)
}

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

	// create a connection to database && check if info on Open() is true
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database.")

	CreateEmployee()
}
