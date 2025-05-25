package main

import (
    "fmt"
)

func main() {
    var numbers [5]int

    fmt.Println("Enter 5 integers:")

    for i := 0; i < 5; i++ {
        fmt.Printf("Enter number %d: ", i+1)
        fmt.Scan(&numbers[i])
    }

    // Reverse the array
    for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
        numbers[i], numbers[j] = numbers[j], numbers[i]
    }

    fmt.Println("Reversed array:")
    for _, num := range numbers {
        fmt.Printf("%d ", num)
    }
    fmt.Println()
}