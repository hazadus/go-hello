// Using `copy()`
package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	// copy() returns number of elements copied
	num := copy(y, x)
	x[3] = 40
	fmt.Println("y =", y, num)
	fmt.Println("x =", x, num)
}
