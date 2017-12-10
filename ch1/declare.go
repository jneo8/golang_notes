package main

import "fmt"

func main() {
	// A short variable declaration, is the most compact, but it may only use with in a function, not for package level variables.
	s1 := "test"
	fmt.Println(s1)

        // This form relies on default init to the zero value for strings, which is ""
        var s2 string
	fmt.Println(s2)

        // This form is rarely used except when declaring multiple variables.
        var s3 = ""
	fmt.Println(s3)

        // This form is explicit about the variable's type, which is redundant when it is the same as that of the initial value but not necessary in other class where they are not of the same type.
        var s4 string = ""
	fmt.Println(s4)
}
