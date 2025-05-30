package main

import "fmt"

func main() {
	name := "Fahmi"

	switch name {
	case "Fahmi":
		fmt.Println("Hello Fahmi")
	case "Ibrahim":
		fmt.Println("Hello Ibrahim")
	default:
		fmt.Println("Hi, Boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang!")
	case false:
		fmt.Println("Nama sesuai standar")
	}
}