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

    count := 0
    for _, char := range input {
        switch strings.ToLower(string(char)) {
        case "a", "e", "i", "o", "u":
            count++
        }
    }

    fmt.Println("Number of vowels:", count)
}
