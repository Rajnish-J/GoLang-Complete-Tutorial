package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Print("Enter a value for N: ")
    fmt.Scanln(&N)

    fmt.Println("\nNumbers from 1 to", N, "with even digit sums:")

    for i := 1; i <= N; i++ {
        if isDigitSumEven(i) {
            fmt.Printf("%d ", i)
        }
    }

    fmt.Println()
}

// Function to check if sum of digits is even
func isDigitSumEven(num int) bool {
    sum := 0
    for num > 0 {
        sum += num % 10
        num /= 10
    }
    return sum%2 == 0
}
