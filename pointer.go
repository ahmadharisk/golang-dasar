package main

import "fmt"

func main() {
	address1 := Address{"medan", "sumatera", "indonesia"}
	address2 := address1

	address2.City = "New York"

	fmt.Println(address1)
	fmt.Println(address2)

	address3 := &address1
	address3.City = "New York"
	fmt.Println("address1, ", address1)
	fmt.Println("address3, ", address3)
}

type Address struct {
	City, Province, Country string
}
