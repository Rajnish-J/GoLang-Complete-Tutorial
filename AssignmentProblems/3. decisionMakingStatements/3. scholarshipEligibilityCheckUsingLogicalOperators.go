package main

import (
	"fmt"
)

func main() {
	var marks, age int

	// Accept marks and age from the user
	fmt.Print("Enter your marks: ")
	fmt.Scanln(&marks)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	// Check eligibility
	if marks > 70 && age > 20 {
		fmt.Println("You are eligible for scholarship")
	} else {
		fmt.Println("You are not eligible for scholarship")
	}
}
