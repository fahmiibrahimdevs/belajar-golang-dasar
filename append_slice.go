package main

import "fmt"

func main() {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	daySlice1 := days[5:] // [Sabtu Minggu]
	daySlice1[0] = "Sabtu Baru" // diubah Sabtu -> jadi Sabtu Baru
	daySlice1[1] = "Minggu Baru" // diubah Minggu -> jadi Minggu Baru
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jum'at Sabtu Baru Minggu Baru]

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2) // [Ups Minggu Baru Libur Baru]
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jum'at Sabtu Baru Minggu Baru]
}