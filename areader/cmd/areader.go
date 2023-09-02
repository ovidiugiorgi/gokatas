// Areader reads three bytes from A into a slice.
package main

import (
	"fmt"
	"strings"

	"github.com/jreisinger/gokatas/areader"
)

func main() {
	var a areader.A
	p := make([]byte, 5)
	n, _ := a.Read(p)
	if n < len(p) {
		fmt.Printf("expected areader.A to read %v bytes; got %v\n", len(p), n)
	}
	sp := string(p)
	if sp != strings.Repeat("A", len(p)) {
		fmt.Printf("expected areader.A to add %v repeated 'A' characters; got %q\n", len(p), sp)
	} else {
		fmt.Println(sp)
	}
}
