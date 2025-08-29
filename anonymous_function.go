package main

import "fmt"

func main() {
	blacklist := func(name string) bool {
		return name == "blacklist"
	}

	registerUser("blacklist", blacklist)
	registerUser("peppo", blacklist)
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked, ", name)
	} else {
		fmt.Println("Welcome, ", name)
	}
}
