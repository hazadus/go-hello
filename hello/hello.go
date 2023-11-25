package main

import (
	"fmt"
)

func main() {
	// Semicolons are not needed. The convention is to omit them.

	// for loop syntax resembles C or JavaScript:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// PrintNumber writes a `number` to the console.
func PrintNumber(number int) {
	fmt.Println(number)
}
