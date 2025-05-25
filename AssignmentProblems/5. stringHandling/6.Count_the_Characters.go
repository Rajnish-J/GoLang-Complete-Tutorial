package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a string: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    charCount := make(map[rune]int)

    for _, char := range input {
        charCount[char]++
    }

    fmt.Println("Character frequencies:")
    for char, count := range charCount {
        fmt.Printf("%c: %d\n", char, count)
    }
}