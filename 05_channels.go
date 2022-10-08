package main

import "fmt"

func main() {
	c := make(chan string)

	// membuat go routine
	go introduce("Annisa", c)
	go introduce("Nuri", c)
	go introduce("Nabila", c)

	msg1 := <-c // func main menerima data dari func introduce dan disimpan pada `msg1`
	fmt.Println(msg1)

	msg2 := <-c // func main menerima data dari func introduce dan disimpan pada `msg2`
	fmt.Println(msg2)

	msg3 := <-c // func main menerima data dari func introduce dan disimpan pada `msg3`
	fmt.Println(msg3)

	close(c) // channel sudah tidak digunakan untuk berkomunikasi lagi (required)
}

func introduce(name string, c chan string) {
	result := fmt.Sprintf("My name is %s", name)

	c <- result // mengirim `result` ke `c`
}
