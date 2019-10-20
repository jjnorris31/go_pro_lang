package __2_cmd_line_arguments

import (
	"fmt"
	"os"
)

func main() {

	/*args := os.Args[1:]
	for i := 0; i < len(args); i++{
		fmt.Printf("%s ", args[i])
	}*/

	for index, arg := range os.Args[1:] {
		fmt.Printf("index: %d | args: %s\n", index, arg)
	}
}
