package main

import "fmt"

func print(c chan string) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan string)

	names := []string{"Annisa", "Nuri", "Nabila"}

	for _, v := range names {
		go func(name string) {
			fmt.Println("Name", name)
			result := fmt.Sprintf("My name is %s", name)
			c <- result
		}(v)
	}

	/*
		melooping data yang disimpan pada slice of string
		& memanggil 3 Goroutine
	*/
	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)
}
