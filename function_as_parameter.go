package main

import "fmt"

func main() {
	sayHelloWithFilter("peppo", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("ini anjing", filter)
}

type Filter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) {
func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "xxx"
	}

	return name
}
