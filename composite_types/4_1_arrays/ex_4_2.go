package main

import (
	"os"
	"fmt"
	"flag"
	"crypto/sha256"
	"crypto/sha512"
)

var t = flag.Int("t", 0, "hash type")

func main() {
	
	flag.Parse() // parse before use
	args := flag.Args() // get the non flags arguments
	counter := 0

	if len(args) != 2 {
		fmt.Printf("ex_4_2: give only two arguments\n")
		os.Exit(1)
	} 
	
	switch *t {
	case 256:
		c1 := sha256.Sum256([]byte(args[0]))
		c2 := sha256.Sum256([]byte(args[1]))
		for i := 0; i < len(c1); i++ {
			if c1[i] != c2[i] {
				counter++;
			}
		}
		fmt.Printf("\"%s\" and \"%s\" sha256 hash differs in %d bytes", args[0], args[1], counter)
	case 512:
		c1 := sha512.Sum512([]byte(args[0]))
		c2 := sha512.Sum512([]byte(args[1]))
		for i := 0; i < len(c1); i++ {
			if c1[i] != c2[i] {
				counter++;
			}
		}
		fmt.Printf("\"%s\" and \"%s\" sha512 hash differs in %d bytes", args[0], args[1], counter)
	case 384:
		c1 := sha512.Sum384([]byte(args[0]))
		c2 := sha512.Sum384([]byte(args[1]))
		for i := 0; i < len(c1); i++ {
			if c1[i] != c2[i] {
				counter++;
			}
		}
		fmt.Printf("\"%s\" and \"%s\" sha384 hash differs in %d bytes", args[0], args[1], counter)
	default:
		fmt.Printf("ex_4_2: only support 256, 384 and 512 for sha crypto\n")
		os.Exit(1)
	}

}
