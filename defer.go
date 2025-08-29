package main

import "fmt"

func main() {
	runMyApp()
}

func loggin() {
	fmt.Println("run func xxx")
}

func runMyApp() {
	defer loggin()
	fmt.Println("run app")
}
