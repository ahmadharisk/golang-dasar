package main

import (
	"fmt"
	"goland-dasar/helper"
)

func main() {
	result := helper.SayHello("peppo")
	fmt.Println(result)

	fmt.Println(helper.Application)
	helper.AccessModifier()
}
