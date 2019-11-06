package main

import (
	"fmt"
)

// reverse in-place
func main() {
	ex := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(ex[:])
	fmt.Printf("%v", ex)
}

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}