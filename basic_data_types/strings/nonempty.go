package main

import "fmt"
// nonempty is an example on an in-place slice algorithm

func nonempty(strings []string) []string {
	i := 0;
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}