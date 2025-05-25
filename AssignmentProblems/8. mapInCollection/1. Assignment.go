package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findDefaulters(input string) []string {
	countMap := make(map[string]int)
	defaulters := []string{}

	// Split the input into individual account numbers
	accounts := strings.Split(input, ",")

	// Count occurrences
	for _, acc := range accounts {
		acc = strings.TrimSpace(acc) // Remove leading/trailing spaces
		countMap[acc]++
	}

	// Collect accounts that appeared more than once
	for acc, count := range countMap {
		if count >= 2 {
			defaulters = append(defaulters, acc)
		}
	}

	return defaulters
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter comma-separated account numbers:")

	if scanner.Scan() {
		input := scanner.Text()
		defaulters := findDefaulters(input)
		fmt.Println("Defaulters for 2+ months:", defaulters)
	}
}
