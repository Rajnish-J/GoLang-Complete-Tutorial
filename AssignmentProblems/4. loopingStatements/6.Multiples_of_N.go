package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Print("Enter a number N: ")
    fmt.Scanln(&N)

    fmt.Println("\nMultiples of", N, "from 1 to 100:")

    i := 1
    for i <= 100 {
        if i%N == 0 {
            fmt.Printf("%d ", i)
        }
        i++ // while-style increment
    }

    fmt.Println()
}