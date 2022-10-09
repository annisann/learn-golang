package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

/*
API
GET & POST => get `Employee` data and create new `Employee` data
*/
var PORT = ":8080"

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Annisa", Age: 23, Division: "Data"},
	{ID: 2, Name: "Mitchell", Age: 39, Division: "Legal"},
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	// Header => menentukan bentuk dari data response yang ingin dikirimkan [json].
	// w.Header().Set("Content-Type", "application/json")

	// check if method is `GET`
	if r.Method == "GET" {
		// parse HTML file
		tpl, err := template.ParseFiles("web-server/template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employees)
		// json.NewEncoder(w).Encode(employees) // convert `employees` to JSON
		return
	}

	// if method sent by client is not `GET`
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// check if method is `POST`
	if r.Method == "POST" {
		// FormValue => get input value from client
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age) // convert string to int

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}
	http.Error(w, "Invalid method", http.StatusBadGateway)
}

func main() {
	// GET
	http.HandleFunc("/employees", getEmployees)

	// POST
	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}
