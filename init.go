package main

import (
	"fmt"
	"goland-dasar/database"
	_ "goland-dasar/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
