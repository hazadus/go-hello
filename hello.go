package main

import (
	"fmt"
)

func main() {
	// go run hello.go
	// https://www.youtube.com/watch?v=j0dThdwftmQ

	// Semicolons are not needed. The convention is to omit them.

	var firstName string = "Alexander" // camelCase is preferred
	lastName := "Goldovsky"            // Shorthand assignment
	var fullName string                // declare but not assign
	fullName = firstName + lastName    // concatenate
	age := 38                          // not using a variable will give an error
	var price float32 = 19.99
	var isSuccess bool = true

	fmt.Println("Hello Go World!")
	fmt.Println(firstName, lastName)
	fmt.Println(fullName, age)
	fmt.Println(price, isSuccess)

	// for loop syntax resembles C or JavaScript:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
