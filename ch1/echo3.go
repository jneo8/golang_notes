package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// see https://golang.org/pkg/strings/#Join
	fmt.Println(strings.Join(os.Args[1:], ","))
}
