package main

import (
	"day2/lib"
	"day2/models"
	"fmt"
)

func addNumber(a int, b int) int {
	return a + b
}

func sayName(name string) {
	fmt.Printf("nama mu %s \n", name)
}

func main() {
	addNumber(1, 2)
	sayName("golep")

	fmt.Println(lib.AllUpperCase("anu"))
	fmt.Println(lib.Multiply(2, 3))

	person := models.Person{
		Name: "golep",
		Age:  20,
	}
	fmt.Println(person)
}
