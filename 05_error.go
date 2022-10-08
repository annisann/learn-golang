package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	var err error

	// parse string to int
	number, err = strconv.Atoi("123")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}
}
