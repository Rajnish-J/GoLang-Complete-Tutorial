package main

import (
	"fmt"
	"sort"
)

func mergeAndSort(a, b []int) []int {
	merged := append(a, b...)
	sort.Ints(merged)
	return merged
}

func main() {
	slice1 := []int{5, 2, 9, 1}
	slice2 := []int{8, 3, 7, 4}

	result := mergeAndSort(slice1, slice2)
	fmt.Println("Merged and sorted slice:", result)
}
