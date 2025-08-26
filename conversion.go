package main

import (
	"fmt"
)

func main() {
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// kenapa nilai16 menjadi negatif?
	// karena melebihi kapasitas (number overflow) akan kembali kebelakang dan terus looping sampai sama dengan nilai yang dikonversi
	// 32768 maka -32768
	// 32769 maka -32767

	var lele = "lele"
	var l uint8 = lele[0]
	var lString = string(l)

	fmt.Println(l)
	fmt.Println(lString)
}
