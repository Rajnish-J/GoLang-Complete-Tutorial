package main

import (
	"fmt"
)

func main() {
	var a, b int

	// Read two integers from the user
	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	// Compare and print the largest
	if a > b {
		fmt.Printf("The largest number is: %d\n", a)
	} else if b > a {
		fmt.Printf("The largest number is: %d\n", b)
	} else {
		fmt.Println("Both numbers are equal.")
	}
}
