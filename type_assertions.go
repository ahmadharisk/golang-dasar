package main

import (
	"fmt"
)

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println(value)
	case int:
		fmt.Println(value)
	default:
		fmt.Println(result)
	}
}

func random() any {
	return "ok"
}
