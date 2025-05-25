package main

import (
	"fmt"
)

func main() {
	var number int

	// Prompt user for input
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	// Check if the number is even or odd
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
