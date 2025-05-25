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

    // Assume the first number is the largest initially
    largest := numbers[0]

    // Loop through the array to find the largest
    for i := 1; i < len(numbers); i++ {
        if numbers[i] > largest {
            largest = numbers[i]
        }
    }

    fmt.Printf("The largest number is: %d\n", largest)
}