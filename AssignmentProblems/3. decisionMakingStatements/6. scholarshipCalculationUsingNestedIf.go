package main

import (
	"fmt"
	"strings"
)

func calculateScholarship(marks int, class int, city string) int {
	var baseAmount int
	city = strings.ToLower(city)

	// Set base amount based on class
	if class == 10 {
		baseAmount = 15000
	} else if class == 11 {
		baseAmount = 25000
	} else {
		fmt.Println("Invalid class")
		return 0
	}

	// Calculate scholarship
	if marks > 30 && marks < 60 {
		scholarship := baseAmount
		if city == "chennai" {
			scholarship -= 1000
		} else {
			scholarship -= 3000 // for class 10
			if class == 11 {
				scholarship = baseAmount - 2000
			}
		}
		return scholarship
	} else if marks >= 60 && marks < 90 {
		return baseAmount / 2
	} else if marks >= 90 {
		return baseAmount
	}
	return 0 // No scholarship
}

func main() {
	// Scenario 1
	marks1 := 80
	city1 := "Chennai"
	class1 := 10
	fmt.Printf("Scenario 1: Scholarship = %d\n", calculateScholarship(marks1, class1, city1))

	// Scenario 2
	marks2 := 70
	city2 := "Delhi"
	class2 := 10
	fmt.Printf("Scenario 2: Scholarship = %d\n", calculateScholarship(marks2, class2, city2))

	// Scenario 3
	marks3 := 20
	city3 := "Chennai"
	class3 := 10
	fmt.Printf("Scenario 3: Scholarship = %d\n", calculateScholarship(marks3, class3, city3))
}
