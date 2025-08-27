package main

import "fmt"

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)

	d, _, _ := getCompleteName()
	fmt.Println(d)
}

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "silpiana"
	middleName = "oke"
	lastName = "rian"

	return firstName, middleName, lastName
}

// apabila tipe data sama dapat di buat lebih efisien
func getCompleteName2() (firstName, middleName, lastName string) {
	firstName = "silpiana"
	// bisa tidak mengisi semua return value

	return firstName, middleName, lastName
}
