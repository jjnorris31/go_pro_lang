package __2_cmd_line_arguments

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("index: %d | args: %s\n", index, arg)
	}
}
