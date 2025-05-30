package main

import "fmt"

func main() {
	a := 15
	b := 30
	c := a + b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	i := 10
	i += 10 // i = i + 10
	fmt.Println(i)
	
	i += 8 // i = i + 8
	fmt.Println(i)
	
	j := 1
	j++ // j = j + 1
	fmt.Println(j)
	j++ // j = j + 1
	fmt.Println(j)

	j-- // j = j - 1
	fmt.Println(j)
	j-- // j = j - 1
	fmt.Println(j)
}