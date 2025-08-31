package main

import "fmt"

func main() {
	// cara 1
	address := &Address{}
	ChangeCountryToIndonesia(address)

	// cara 2
	address2 := Address{}
	ChangeCountryToIndonesia(&address2)

	fmt.Println(address)
	fmt.Println(address2)
}

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
