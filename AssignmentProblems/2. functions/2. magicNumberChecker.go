package main

import (
	"fmt"
)

// Helper function to sum the digits of a number
func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// isMagicNumber returns true if repeated sum of digits results in 1
func isMagicNumber(n int) bool {
	for n >= 10 {
		n = sumOfDigits(n)
	}
	return n == 1
}

func main() {
	number := 10
	if isMagicNumber(number) {
		fmt.Println("Magic Number")
	} else {
		fmt.Println("Not Magic")
	}
}
