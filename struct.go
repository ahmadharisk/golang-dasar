package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// assign func ke struct customer
func (customer Customer) sayHello() {
	fmt.Println("hello ", customer.Name)
}

func main() {
	var peppo Customer
	peppo.Name = "John Doe"
	peppo.Address = "China"
	peppo.Age = 18

	fmt.Println(peppo)

	silpiana := Customer{
		Name:    "Sil Pi",
		Address: "Vietnam",
		Age:     18,
	}
	fmt.Println(silpiana)

	sharp := Customer{"sharp", "Vietnam", 18}
	fmt.Println(sharp)

	peppo.sayHello()
}
