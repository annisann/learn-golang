package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080" // save localhost port `localhost:8080`

func greet(w http.ResponseWriter, r *http.Request) { // send a response message
	/*
		ResponseWriter => interface	to send back response to client who sent request.
		Request        => struct to fetch request data: form value, headers, url, etc.
	*/
	msg := "Hello World! X"
	fmt.Fprint(w, msg)
}

func main() {
	http.HandleFunc("/", greet) // for routing

	http.ListenAndServe(PORT, nil) // run server application
}
