package main

import "fmt"

func main() {
	names := [...]string{"Fahmi", "Ibrahim", "Nexaryn", "Ridho", "Arif", "Geva"}

	slice1 := names[4:6]
	fmt.Println(slice1) // [Arif Geva]
	
	slice2 := names[:4]
	fmt.Println(slice2) // [Fahmi Ibrahim Nexaryn Ridho]

	slice3 := names[2:]
	fmt.Println(slice3) // [Nexaryn Ridho Arif Geva]

	slice4 := names[:]
	fmt.Println(slice4) // [Fahmi Ibrahim Nexaryn Ridho Arif Geva]

	fmt.Println(slice1[1])
	fmt.Println(slice2[0])
}