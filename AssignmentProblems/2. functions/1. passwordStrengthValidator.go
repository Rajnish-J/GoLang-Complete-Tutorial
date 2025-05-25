package main

import (
	"fmt"
	"unicode"
)

func isStrongPassword(pwd string) bool {
	if len(pwd) < 8 {
		return false
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, ch := range pwd {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}

func main() {
	password := "Rajnish@07"
	if isStrongPassword(password) {
		fmt.Println("Strong")
	} else {
		fmt.Println("Weak")
	}
}
