package main

import (
    "fmt"
)

func main() {
    var numbers [5]int
    var evenCount, oddCount int

    fmt.Println("Enter 5 integers:")

    for i := 0; i < 5; i++ {
        fmt.Printf("Enter number %d: ", i+1)
        fmt.Scan(&numbers[i])
    }

    for _, num := range numbers {
        if num%2 == 0 {
            evenCount++
        } else {
            oddCount++
        }
    }

    fmt.Printf("Even numbers count: %d\n", evenCount)
    fmt.Printf("Odd numbers count: %d\n", oddCount)
}