package main

import (
    "fmt"
)

func main() {
    var numbers [5]int
    var sum int

    fmt.Println("Enter 5 integers:")

    for i := 0; i < 5; i++ {
        fmt.Printf("Enter number %d: ", i+1)
        fmt.Scan(&numbers[i])
        sum += numbers[i]
    }

    fmt.Printf("Sum of the array elements: %d\n", sum)
}