package main

import (
    "fmt"
)

func main() {
    var n int

    fmt.Print("Enter a number: ")
    fmt.Scanln(&n)

    fmt.Println("\nNumber Pattern:")
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print(i)
        }
        fmt.Println()
    }
}
