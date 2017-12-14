package main

import (
    "fmt"
)

func main() {
    nums := []int{1, 2, 3}
    for idx, value := range nums {
        fmt.Println("index: ", idx, " Value: ", value)
    }
}
