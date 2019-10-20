package __2_cmd_line_arguments

import (
	"fmt"
	"os"
)


func main() {
	fmt.Printf("program name: %s\n", os.Args[0])
	for _, args := range os.Args[1:] {
		fmt.Printf("%s \n", args)
	}
}
