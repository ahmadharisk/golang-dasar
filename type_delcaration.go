package main

import "fmt"

func main() {
	type NoKTP string

	var ktpPeppo NoKTP = "28736478265472"
	fmt.Println(ktpPeppo)

	var contoh string = "2874236542635"
	var contohKTP NoKTP = NoKTP(contoh)
	fmt.Println(contohKTP)
}
