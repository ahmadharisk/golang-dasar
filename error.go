package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Pembagian(100, 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi is zero")
	}
	return nilai / pembagi, nil
}
