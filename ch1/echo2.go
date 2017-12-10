package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// forloop use range
	// range produces a pair if values
	// the index and the value of the element at that index
	for idx, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println("idx: ", idx, ",", s)
	}
	fmt.Println(s)
}
