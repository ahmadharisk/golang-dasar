package main

import "fmt"

func main() {
	hewan := "ayam"

	switch hewan {
	case "ayam":
		fmt.Println("pasti ayam")
	case "sapi":
		fmt.Println("pasti sapi")
	default:
		fmt.Println("siapa ini")
	}

	switch length := len(hewan); length > 5 {
	case true:
		fmt.Println("nama hewan terlalu panjang")
	case false:
		fmt.Println("nama hewan pas")
	}
}
