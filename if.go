package main

import "fmt"

func main() {
	hewan := "ayam"

	if hewan == "kukang" {
		fmt.Println("halo ", hewan)
	} else if hewan == "kijang" {
		fmt.Println("hale ", hewan)
	} else {
		fmt.Println("siapa kamu?")
	}

	// if short statement
	if length := len(hewan); length > 5 {
		fmt.Println("nama hewan terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}

}
