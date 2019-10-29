package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse() // before the flags are used, update the flag variables from their values
	fmt.Print(strings.Join(flag.Args(), *sep)) // non-flag arguments are available from flag.Args()
	if !*n {
		fmt.Println()
	}
}