// Dup2 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files. Adapted from
// github.com/adonovan/gopl.io/tree/master/ch1/dup2.
//
// Level: beginner
// Topics: io, maps, bufio.Scanner
package main

import (
	"bufio"
	"fmt"
	"os"
)

// go run . <(./random.sh 1000)
func main() {
	counts := make(map[string]int)
	if len(os.Args) > 1 {
		for _, name := range os.Args[1:] {
			f, err := os.Open(name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: failed to open file %q: %v\n", name, err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	} else {
		countLines(os.Stdin, counts)
	}

	for k, v := range counts {
		fmt.Printf("%d\t%s\n", v, k)
	}
}

func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}
