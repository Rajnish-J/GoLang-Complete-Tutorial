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
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	words := strings.Fields(sentence)

	longest := ""
	maxLen := 0

	for _, word := range words {
		if len(word) >= maxLen {
			longest = word
			maxLen = len(word)
		}
	}

	fmt.Println("Longest word:", longest)
}