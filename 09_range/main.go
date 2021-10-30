package main

import "fmt"

func main() {
	// ids := []int{33,76,34,34,67}

	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n",i, id)
	// }

	// sum := 0
	// for _, id := range ids {

	// }

	emails := map[string]string{"Bob":"bob@gmai.com", "sharon":"sharon@gmail.com"}

	for k, v := range emails {
		fmt.Println("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}