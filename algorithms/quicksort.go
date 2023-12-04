// Quicksort algorithm demo
package main

import "fmt"

// SortNumbers ...
// Sort `numbers` using quicksort algorithm
func SortNumbers(numbers []int) []int {
	if numbers == nil {
		return []int{}
	}

	if len(numbers) < 2 {
		return numbers
	}

	pivot := numbers[0]
	var less []int
	var more []int

	for _, i := range numbers[1:] {
		if i <= pivot {
			less = append(less, i)
		} else {
			more = append(more, i)
		}
	}

	return append(append(SortNumbers(less), pivot), SortNumbers(more)...)
}

func main() {
	array := []int{99, 88, 77, 66, 55, 44, 33, 22, 11}
	fmt.Println(SortNumbers(array))
}
