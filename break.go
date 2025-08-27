package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("udah 5 nih")
			break
		}
		fmt.Println("perulangan ke ", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke ", i)
	}
}
