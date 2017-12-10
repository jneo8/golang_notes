//Echo1 prints its command-line arguments.
// go run echo.go 1 2 3 4
package main

import (
	"fmt"
	"os"
)

func main() {
	// declares two variables, s and sep
	var s, sep string
	// os.Args is a slice of strings, Slices are a fundamental notion in go.
        // os.Args[0] is the name of command itself.
	fmt.Println(os.Args)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(s)
	}
}
