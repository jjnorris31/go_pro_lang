package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {
	args := os.Args[1:]
	fmt.Printf("%s\n", basename(args[0]))
}

// basename removes directory components and a .suffix
func basename(s string) string {
	
	slash := strings.LastIndex(s, "/")
	s = s[slash + 1:]

	
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	
	return s
}
