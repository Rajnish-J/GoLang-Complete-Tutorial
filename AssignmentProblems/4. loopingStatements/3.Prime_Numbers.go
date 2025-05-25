package main

import (
    "fmt"
)

func main() {
    fmt.Println("Prime numbers between 1 and 100:")

    for num := 2; num <= 100; num++ {
        isPrime := true

        for i := 2; i*i <= num; i++ {
            if num%i == 0 {
                isPrime = false
                break
            }
        }

        if isPrime {
            fmt.Printf("%d ", num)
        }
    }
    fmt.Println()
}
