package main

import "fmt"

func main() {
	// pointer
	// length
	// capacity

	var hewan = [...]string{"anjing", "babi", "cacing", "dara", "elang", "flamingo", "gajah"}
	slice := hewan[1:4]
	slice1 := hewan[:]

	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println(len(hewan))

	hewan2 := append(hewan[5:], "hiu", "ikan")
	fmt.Println(hewan2)
	fmt.Println(hewan)

	newSlice := make([]int, 2, 5)
	newSlice[0] = 1
	newSlice[1] = 2
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, 3)
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fromSlice := hewan[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
}
