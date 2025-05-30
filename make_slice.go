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

	newSlice := make([]string, 2, 5) // length: 2, capacity: 5
	newSlice[0] = "Fahmi"
	newSlice[1] = "Fahmi"
	// newSlice[1] = "Fahmi" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	
	newSlice2 := append(newSlice, "Fahmi")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}