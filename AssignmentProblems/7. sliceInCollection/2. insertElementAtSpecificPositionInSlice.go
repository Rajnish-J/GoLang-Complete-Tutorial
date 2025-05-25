package main

import (
	"fmt"
)

// insertAt inserts value `val` at index `index` in slice `s`
func insertAt(s []int, index int, val int) []int {
	if index < 0 || index > len(s) {
		fmt.Println("Index out of range")
		return s
	}
	s = append(s[:index], append([]int{val}, s[index:]...)...)
	return s
}

func main() {
	slice := []int{10, 20, 30, 40}
	fmt.Println("Original slice:", slice)

	// Insert 25 at index 2
	slice = insertAt(slice, 2, 25)
	fmt.Println("After insertion:", slice)
}
