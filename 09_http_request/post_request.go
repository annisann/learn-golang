package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// data for request body or `POST` request
	data := map[string]interface{}{
		"title":  "Hello World!",
		"body":   "This is the body.",
		"userID": 1,
	}

	// convert to JSON
	requestJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}

	/*
		preparing the request: request method, request URL, data (`io.Reader`; opt).
		`bytes.NewBuffer()` => change `requestJSON` data type to `io.Reader` interface
	*/
	request, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJSON))
	request.Header.Set("Content-type", "application/json") // set content type as JSON
	if err != nil {
		log.Fatalln(err)
	}

	// send request & return response
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln()
	}

	log.Println(string(body))
}
