package main

import "fmt"

func main() {
	nilaiAkhir := 90
	absensi := 80

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := absensi >= 80

	result := lulusNilaiAkhir && lulusAbsensi
	fmt.Println(result)

}
