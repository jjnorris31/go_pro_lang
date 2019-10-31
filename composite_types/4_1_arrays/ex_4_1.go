package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf("sha256: give only two arguments\n")
		os.Exit(1)
	} 

	counter := 0
	c1 := sha256.Sum256([]byte(args[0]))
	c2 := sha256.Sum256([]byte(args[1]))

	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			counter++;
		}
	}

	fmt.Printf("\"%s\" and \"%s\" sha256 hashes differs in %d bytes", args[0], args[1], counter)
}