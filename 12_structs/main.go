package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

func (p Person) greet() string {
	return "hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age ++
}

func (p *Person) getMarried(spouseLastName string) {
	// if (p.gender == "m")
}

func main() {
	// person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 23}
	person1 := Person{"Samantha", "Smith", "Boston", "f", 23}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// person1.age ++
	// fmt.Println(person1.age)

	person1.hasBirthday()
	fmt.Println(person1.greet())
}