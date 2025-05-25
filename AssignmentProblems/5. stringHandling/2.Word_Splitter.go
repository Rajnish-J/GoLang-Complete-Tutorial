package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a sentence: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    words := strings.Fields(input)

    for i, word := range words {
        if i == 0 {
            fmt.Println(strings.ToUpper(word))
        } else {
            fmt.Println(word)
        }
    }
}
