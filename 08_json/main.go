/*
	Objective: decode JSON data into a struct.
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

func main() {
	var jsonString = `
		{
			"full_name": "Marco Diaz",
			"email": "hola.diaz@gmail.com",
			"age": 20
		}	
	`

	var result Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result.FullName, result.Email, result.Age)
}
