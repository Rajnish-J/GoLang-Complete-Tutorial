package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Print("Enter a string: ")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()

    reversed := ""
    for i := len(input) - 1; i >= 0; i-- {
        reversed += string(input[i])
    }

    fmt.Printf("Reversed string: %s\n", reversed)
}