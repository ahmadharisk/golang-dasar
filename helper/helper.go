package helper

import "fmt"

var version = "1.0.0" // tidak bisa diakses
var Application = "golang"

func sayGoodBye(name string) string { // tidak bisa di akses
	return "goodbye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func AccessModifier() {
	fmt.Println(version)
	fmt.Println(sayGoodBye("peppo"))
}
