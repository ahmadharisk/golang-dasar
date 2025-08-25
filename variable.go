package main

import "fmt"

func main() {
	var name string

	name = "haris"
	fmt.Println(name)

	name = "ahmad haris kurniawan"
	fmt.Println(name)

	// bisa mendeteksi tipe data
	var age = 12
	fmt.Println(age)

	// tidak perlu menggunakan var cukup :=
	negara := "indonesia"
	fmt.Println(negara)

	// membuat var lebih dari satu
	var (
		kota = "jakarta pusat"
		hobi = "membuat program"
	)

	fmt.Println(kota)
	fmt.Println(hobi)
}
