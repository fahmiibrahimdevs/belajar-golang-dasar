package main

import "fmt"

func main() {
	nilaiAkhir := 90
	absensi := 70

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)

}