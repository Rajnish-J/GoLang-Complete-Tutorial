package main

import (
	"fmt"
	"strings"
)

// classifyItem classifies the input string using switch-case
func classifyItem(item string) {
	// Normalize the input
	switch strings.ToLower(item) {
	case "apple", "orange", "grape":
		fmt.Println("This is a fruit")
	case "potato", "onion":
		fmt.Println("This is a vegetable")
	default:
		fmt.Println("I cannot classify this as a fruit or vegetable")
	}
}

func main() {
	var input string
	fmt.Print("Enter an item: ")
	fmt.Scanln(&input)

	classifyItem(input)
}
