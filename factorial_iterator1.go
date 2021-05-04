// Package main implements factorial.
package main

import "fmt"

// facItr1 returns n!
func facItr(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n--
	}
	return result
}

func main() {
	fmt.Println(facItr(5))
}
