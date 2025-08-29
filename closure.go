package main

import "fmt"

func main() {
	counter := 0

	// kemampuan sebuah func dapat mengakses variable di sekitarnya
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
