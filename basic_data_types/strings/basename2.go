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
	
	// LastIndex busca en el string la primera aparición
	// del parámetro que le pasamos de derecha a izquierda
	slash := strings.LastIndex(s, "/")
	s = s[slash + 1:]

	
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	
	return s
}
