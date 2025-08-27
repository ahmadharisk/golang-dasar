package main

import "fmt"

func main() {

	hewan := map[string]string{
		"name": "kucing",
	}

	var buku map[string]string = map[string]string{}
	buku["judul"] = "membuka dunia"
	buku["pengarang"] = "haris"

	fmt.Println(hewan)
	fmt.Println(buku["judul"])
	fmt.Println(len(buku))

	pakaian := make(map[string]string)
	pakaian["jenis"] = "baju"
	pakaian["bahan"] = "lembut"
	pakaian["del"] = "akan dihapus"

	fmt.Println(pakaian)
	delete(pakaian, "del")
	fmt.Println(pakaian)
}
