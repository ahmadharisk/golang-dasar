package main

import "fmt"

func main() {
	var alamat1 = &Address{}
	var alamat2 = alamat1

	alamat2.Country = "indonesia"
	fmt.Println(alamat1)
	fmt.Println(alamat2)

	// untuk data kosong menggunakan new bisa juga
	var alamat3 = new(Address)
	alamat4 := alamat3

	alamat4.Country = "jepang"
	fmt.Println(alamat3)
	fmt.Println(alamat4)
}

type Address struct {
	City, Province, Country string
}
