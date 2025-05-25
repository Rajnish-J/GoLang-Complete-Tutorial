package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	// Accept three numbers from the user
	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)
	fmt.Print("Enter third number: ")
	fmt.Scanln(&c)

	// Using Ternary-style logic
	fmt.Print("The Biggest among three numbers is: ")
	biggest := a
	if b > biggest {
		biggest = b
	}
	if c > biggest {
		biggest = c
	}
	fmt.Print(biggest)
}
