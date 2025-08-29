package main

import "fmt"

func main() {
	person := Person{"peppo"}
	SayHello(person)

	animal := Animal{"ular"}
	SayHello(animal)
}

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello ", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
