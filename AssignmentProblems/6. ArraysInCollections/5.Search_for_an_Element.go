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

    var target int
    fmt.Print("Enter number to search: ")
    fmt.Scan(&target)

    index := search(numbers, target)

    if index != -1 {
        fmt.Printf("Number found at index: %d\n", index)
    } else {
        fmt.Println("Number not found.")
    }
}

// search returns the index of target in arr, or -1 if not found
func search(arr [5]int, target int) int {
    for i, num := range arr {
        if num == target {
            return i
        }
    }
    return -1
}