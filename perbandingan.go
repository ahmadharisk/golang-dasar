package main

import "fmt"

func main() {
	kucing1 := "kucing"
	kucing2 := "kucing"

	benar := kucing2 == kucing1
	salah := kucing2 != kucing1

	fmt.Println(benar)
	fmt.Println(salah)
}
