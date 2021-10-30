package main

import "fmt"

func main() {
	a := 5
	b := &a

	// fmt.Println(a,b)
	// fmt.Printf("%T\n", a)

	// fmt.Println(*b)
	// fmt.Println(*&a)

	*b = 10
	fmt.Println(a)
}