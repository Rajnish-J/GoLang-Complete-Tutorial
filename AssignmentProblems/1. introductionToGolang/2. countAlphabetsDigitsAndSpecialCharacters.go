package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a line of text: ")
	text, _ := reader.ReadString('\n')

	letters := 0
	digits := 0
	specials := 0

	for _, char := range text {
		if unicode.IsLetter(char) {
			letters++
		} else if unicode.IsDigit(char) {
			digits++
		} else if !unicode.IsSpace(char) {
			specials++
		}
	}

	fmt.Println("Total letters:", letters)
	fmt.Println("Total digits:", digits)
	fmt.Println("Total special characters:",specials)
}