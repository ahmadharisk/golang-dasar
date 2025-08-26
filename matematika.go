package main

import "fmt"

func main() {
	a := 10
	b := 20
	d := 3
	e := 4
	f := 2
	c := a + b - d*e/f
	fmt.Println(c)

	// augmented assignment
	a += 20
	fmt.Println(a)
	a -= 2
	fmt.Println(a)
	a *= 3
	fmt.Println(a)
	a /= 2
	fmt.Println(a)
	a %= 3
	fmt.Println(a)

	// unary operator
	a++
	a--
}
