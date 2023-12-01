// The for-range loop
package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12, 14, 16}

	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}

	for k, v := range uniqueNames {
		fmt.Println(k, v)
	}

	samples := []string{"Hello!", "Emojis🤪🤨🤩"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}
