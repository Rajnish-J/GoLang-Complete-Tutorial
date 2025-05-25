package main

import (
	"fmt"
)

func removeDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

func main() {
	original := []int{1, 2, 3, 2, 4, 3, 5, 1, 6}
	fmt.Println("Original slice:", original)

	unique := removeDuplicates(original)
	fmt.Println("Slice after removing duplicates:", unique)
}
