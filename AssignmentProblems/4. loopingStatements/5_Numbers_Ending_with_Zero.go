package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Print("Enter a value for N: ")
    fmt.Scanln(&N)

    fmt.Println("\nNumbers from 1 to", N, "that end with 0:")

    i := 1
    for i <= N {
        if i%10 == 0 {
            fmt.Printf("%d ", i)
        }
        i++ // loop acts like a while
    }

    fmt.Println()
}
