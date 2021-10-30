package main

import "fmt"

func main() {
	var name string = "Brad"
	var age int32 = 37
	//var isCool = true

	//Shorthand
	size := 1.3

	name, email := "brad", "gra@gd.com"

	fmt.Println(name, email, age)
	fmt.Printf("%T\n", size)
}