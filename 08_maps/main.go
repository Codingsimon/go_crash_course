package main

import "fmt"

func main() {
	// emails := make(map[string]string)

	// //Assign Key Values
	// emails["Bob"] = "bob@lgmail.com"
	// emails["Sharon"] = "sharon@lgmail.com"

	// delete(emails, "Bob")

	emails := map[string]string{"Bob":"bob@gmai.com", "sharon":"sharon@gmail.com"}

	emails["Mike"] = "mike@gmdf.com"

	fmt.Println(emails)
}