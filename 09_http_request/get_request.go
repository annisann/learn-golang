package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://go.dev/tour/basics/7")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response.Body)
	body, err := ioutil.ReadAll(response.Body) // body value to slice of byte
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close() // don't forget to close to prevents memory leak

	sb := string(body) // slice of byte => string
	log.Println(sb)
}
