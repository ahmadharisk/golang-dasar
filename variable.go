package main

import "fmt"

func main() {
	var name string

	name = "peppo"
	fmt.Println(name)

	name = "peppo silpiana"
	fmt.Println(name)

	// bisa mendeteksi tipe data
	var age = 12
	fmt.Println(age)

	// tidak perlu menggunakan var cukup :=
	negara := "jepang"
	fmt.Println(negara)

	// membuat var lebih dari satu
	var (
		kota = "tokyo"
		hobi = "membuat program"
	)

	fmt.Println(kota)
	fmt.Println(hobi)
}
