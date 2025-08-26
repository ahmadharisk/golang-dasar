package main

import "fmt"

func main() {
	var hewan [3]string

	hewan[0] = "anjing"
	hewan[1] = "burung"
	hewan[2] = "kucing"

	fmt.Println(hewan)

	// membuat array secara langsung
	var nilai = [3]int{1, 2, 3}
	fmt.Println(nilai)

	//... hanya dapat digunakan saat pembuatan langsung array
	var jarak = [...]int{1, 2, 3, 4}
	fmt.Println(jarak)
	fmt.Println(len(jarak))

	// tidak ada menghapus data array
}
