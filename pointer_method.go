package main

import "fmt"

func main() {
	peppo := Man{"Peppo"}
	peppo.Married()

	fmt.Println(peppo)
}

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
