// Converting strings to slices
package main

import "fmt"

func main() {
	var s string = "Hello, 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)

	fmt.Println("bs =", bs) // [72 101 108 108 111 44 32 240 159 140 158]
	fmt.Println("rs =", rs) // [72 101 108 108 111 44 32 127774]
}
