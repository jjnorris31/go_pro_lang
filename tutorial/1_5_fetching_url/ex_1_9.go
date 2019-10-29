package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" +  url
		}

		resp, err := http.Get(url)

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Printf("status: %s\n", resp.Status)
		_ = resp.Body.Close()

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading from %s: %v\n", url, err)
			os.Exit(1)
		}

	}
}
