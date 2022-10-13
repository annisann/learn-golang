package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("Server started at localhost", server.Addr)
	server.ListenAndServe()
}

// convert data to JSON string
func OutputJSON(w http.ResponseWriter, o interface{}) {
	response, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(response)
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	// to check whether request is a valid basic auth request
	if !Auth(w, r) {
		return
	}

	// to validate `GET` method
	if !AllowOnlyGET(w, r) {
		return
	}

	// return student data by id
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	// return all students data
	OutputJSON(w, GetStudents())
}
