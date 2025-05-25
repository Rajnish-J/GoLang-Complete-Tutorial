package main

import (
	"fmt"
)

func rotateLeft(s []int, n int) []int {
	if len(s) == 0 {
		return s
	}
	n = n % len(s) // handle rotation greater than slice length
	return append(s[n:], s[:n]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Original slice:", slice)

	rotated := rotateLeft(slice, 2)
	fmt.Println("After rotating left by 2:", rotated)
}
