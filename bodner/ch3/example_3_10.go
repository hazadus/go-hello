// Using a map
package main

import "fmt"

func main() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	totalWins["Bears"]++

	fmt.Println("totalWins[\"Lions\"] =", totalWins["Lions"])
	fmt.Println("totalWins[\"Bears\"] =", totalWins["Bears"])
	fmt.Println(totalWins)

	// Use this syntax to check whether map has the key:
	v, ok := totalWins["Pinguins"]
	fmt.Println(v, ok)
}
