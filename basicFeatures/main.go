package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var firstName string = "Alexander" // camelCase is preferred
	lastName := "Goldovsky"            // Shorthand assignment
	var fullName string                // declare but not assign
	fullName = firstName + lastName    // concatenate
	const age uint = 38                // not using a variable will give an error

	fmt.Println("Hello Go World!")
	fmt.Println(rand.Int())
	fmt.Println(firstName, lastName)
	fmt.Println(fullName, age)

	// Defining Multiple Constants with a Single Statement
	const price, tax float32 = 275, 27.50
	const quantity, inStock = 2, true
	fmt.Println("Total:", quantity*(price+tax))
	fmt.Println("In stock: ", inStock)
}
