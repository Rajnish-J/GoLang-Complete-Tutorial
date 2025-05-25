package main

import "fmt"

func main() {
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	var age int
	fmt.Scan(&age)

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}