package main

import (
	"fmt"
)

func findCommonElements(a, b []int) []int {
	elementMap := make(map[int]bool)
	common := []int{}

	// Store elements of slice a in map
	for _, val := range a {
		elementMap[val] = true
	}

	// Check which elements of b are in the map
	seen := make(map[int]bool)
	for _, val := range b {
		if elementMap[val] && !seen[val] {
			common = append(common, val)
			seen[val] = true // avoid duplicates
		}
	}

	return common
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 5}
	slice2 := []int{4, 5, 6, 7, 5}

	common := findCommonElements(slice1, slice2)
	fmt.Println("Common elements:", common)
}
