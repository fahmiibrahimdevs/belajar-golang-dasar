package main

import "fmt"

func main() {
	type NoKTP string

	var ktpFahmi NoKTP = "31750321"
	fmt.Println(ktpFahmi)
	fmt.Println(NoKTP("38037421"))
}