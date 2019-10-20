package __3_duplicate_lines

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	file := os.Args[1]

	if len(os.Args) > 2 {
		fmt.Printf("%s: only one file, please\n", os.Args[0])
	} else {
		f, err := os.Open(file)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			return
		}
		countLines(f, counts)
		_ = f.Close()
	}

	for line, n := range counts {
		if n > 2 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
