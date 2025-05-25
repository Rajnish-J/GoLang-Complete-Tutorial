package main

import (
    "fmt"
)

func main() {
    var multiplier, max int

    fmt.Print("Enter the multiplier: ")
    fmt.Scanln(&multiplier)

    fmt.Print("Enter the max value: ")
    fmt.Scanln(&max)

    fmt.Printf("\nMultiplication Table of %d up to %d:\n", multiplier, max)
    for i := 1; i <= max; i++ {
        fmt.Printf("%d x %d = %d\n", multiplier, i, multiplier*i)
    }
}
