package main

import "fmt"

func main() {
	kosong := NewMap("")
	fmt.Println(kosong)
	if kosong == nil {
		fmt.Println("kosong is nil")
	}

	newMap := NewMap("peppo")
	fmt.Println(newMap)
}

// nil hanya bisadigunakan di
// interface, func, map, slice, pointer, channel
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
