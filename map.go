package main

import "fmt"

func main() {
	// person := map[string]string{
	// 	"name": "Fahmi Ibrahim",
	// 	"address": "Jakarta",
	// }
	// fmt.Println(person["name"])
	// fmt.Println(person["address"])
	// fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Fahmi Ibrahim"
	book["wrong"] = "ups"
	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)
}