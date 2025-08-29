package main

import "fmt"

func main() {
	runApp(true)
	fmt.Println("harusnya ter cetak meskipun terjadi panic")
}

func endApp() {
	fmt.Println("endApp")

	// recover ditaro sini
	message := recover()
	fmt.Println("terjadi panic, ", message)
}
func runApp(error bool) {
	defer endApp()

	if error {
		panic("error")
	}
}
