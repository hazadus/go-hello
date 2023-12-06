package main

import (
	"fmt"

	print "github.com/hazadus/go-hello/bodner/ch9/package_example/formatter"
	"github.com/hazadus/go-hello/bodner/ch9/package_example/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
