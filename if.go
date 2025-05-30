package main

import "fmt"

func main() {
	name := "Fah"

	if name == "Fahmis" {
		fmt.Println("Hallo", name)
	} else if name == "Ibrahim" {
		fmt.Println("Yooo", name)
	} else {
		fmt.Println("Tidak ada")
	}

	// Short Statement

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sesuai standar")
	}

}