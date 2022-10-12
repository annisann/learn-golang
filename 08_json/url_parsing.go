package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "http://developer.com:80/hello?name=claire&age=45"

	var u, err = url.Parse(urlString) // parsing string to url
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)
	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	// parse query in url to map[string][]string
	name := u.Query()["name"][0]
	age := u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)
}
