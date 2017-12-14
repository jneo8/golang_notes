// Dup1 prints the text of each every line that appers more than
// once in the standard input, preceded by its count.

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // declared
    scanner := bufio.NewScanner(os.Stdin)
    var text string
    counts := make(map[rune]int)

    // Input
    for text != "q" {
        fmt.Print("Enter :")
        scanner.Scan()
        text = scanner.Text()
        for idx, s := range text {
            counts[s]++
            fmt.Printf("%d : %c \n", idx, s)
        }

        fmt.Printf("\n\n----------\n\n")

        for line, n := range counts {
            fmt.Printf("Value: %c, Count: %d\n", line, n)
        }
    }

}
