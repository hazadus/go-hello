package print

import "fmt"

// Format returns string with `num` inside.
func Format(num int) string {
	return fmt.Sprintf("The number is %d", num)
}
