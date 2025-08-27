package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("counter is", counter)
		counter++
	}

	fmt.Println("done :)")

	// for dengan statement
	for hitung := 1; hitung <= 10; hitung++ {
		fmt.Println("hitungan ", hitung)
	}

	fmt.Println("done hitung")

	hewan := []string{"ayam", "bebek", "cacing"}
	for i := 0; i < len(hewan); i++ {
		fmt.Println("hewan ", hewan[i])
	}

	for i, namaHewan := range hewan {
		fmt.Println(i+1, ". hewan ", namaHewan)
	}

	// ketika tidak butuh index nya ganti dengan _
	for _, namaHewan := range hewan {
		fmt.Println(". hewan ", namaHewan)
	}
}
