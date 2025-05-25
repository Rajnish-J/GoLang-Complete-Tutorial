package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter the main string: ")
    mainStr, _ := reader.ReadString('\n')
    mainStr = strings.TrimSpace(mainStr)

    fmt.Print("Enter the substring: ")
    subStr, _ := reader.ReadString('\n')
    subStr = strings.TrimSpace(subStr)

    index := strings.Index(mainStr, subStr)

    if index != -1 {
        fmt.Printf("Substring found at index: %d\n", index)
    } else {
        fmt.Println("Substring not found.")
    }
}