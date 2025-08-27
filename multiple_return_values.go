package main

import "fmt"

func main() {
	plusResut, minusResult := plusMinus(2, 3)
	fmt.Println(plusResut, minusResult)

	// ketika tidak butuh salah satu value diganti dengan _
	_, minusss := plusMinus(2, 3)
	fmt.Println(minusss)
}

func plusMinus(a int, b int) (int, int) {
	return a + b, a - b
}
