// This program declares a constant, a function, and a couple of variables

package main
import "fmt"

// boilingF is a package-level declaration(as is main)
const boilingF = 212.0

func main() {
    // f & c are variables local to function main 
    var f = boilingF
    var c = (f - 32) * 5 / 9
    fmt.Printf("boiling point = %f or %g 。C\n", f, c)
}
