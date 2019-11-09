package main

//Exercise 4.3: rewrite reverse to use an array pointer instead of a slice

import (
	"fmt"
)
func main() {
	someArray := [5]int{1, 2, 3, 4, 5}
	reverse(&someArray)
	fmt.Printf("%v", someArray)
}

func reverse(arr *[5]int) {
	for i, j := 0, len(arr) - 1; i != j; i, j = i + 1, j - 1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}