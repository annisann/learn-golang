package main

import (
	"fmt"
	"os"
)

/*
	INSTRUCTIONS:
		Create a CLI service to presend friends data in class using `struct` and `function`.

		Data:
			1. Name
			2. Address
			3. Job
			4. Reasons why choose Golang


		Expected output: data with student ID = 1
		P.S. To get argument from terminal, use `os.Args`
*/

type Student struct {
	ID      int
	Name    string
	Address string
	Job     string
	Reasons string
}

var students = []Student{}

func addStudents(args []string) {
	var newStudent = Student{
		ID:      len(students) + 1,
		Name:    args[1],
		Address: args[2],
		Job:     args[3],
		Reasons: args[4],
	}
	students = append(students, newStudent)
}

func main() {
	inp := os.Args
	addStudents(inp)
	fmt.Println(students)
}
