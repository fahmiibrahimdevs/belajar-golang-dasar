package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Fahmi"
	names[1] = "Ibrahim"
	names[2] = "Nexaryn"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		10, 8, 7, 100, 80,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}