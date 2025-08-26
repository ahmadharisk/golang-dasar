package main

import "fmt"

func main() {
	// const tidak dapat diubat
	const harimau string = "harimau"
	const kucing = "kucing"

	// pembuatan multiple const
	const (
		semut string = "semut"
		gagak        = "gagak"
	)

	fmt.Println(harimau)
	fmt.Println(kucing)
	fmt.Println(semut)
	fmt.Println(gagak)
}
