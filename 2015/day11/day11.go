package main

import (
	"fmt"
	"strings"
)

func main() {
	// input := "cqjxjnds"
	part2input := "cqjxxyzz"

	fmt.Println(part1(part2input))
}

func part1(input string) string {
	var password []string
	increment := true

	for i := len(input) - 1 ; i >= 0; i-- {
		if increment {
			if input[i] == 122 {
				password = append(password, string(97))
			} else {
				password = append(password, string(input[i] + 1))
				increment = false
			}
		} else {
			password = append(password, string(input[i]))
		}
	}

	for i, j := 0, len(password)-1; i < j; i, j = i+1, j-1 {
		password[i], password[j] = password[j], password[i]
	}

	passwordStr := strings.Join(password, "")

	if forbiddenLetterChecker(passwordStr) && threeStraightChecker(passwordStr) && twoDoubleChecker(passwordStr) {
		return passwordStr
	}

	return part1(passwordStr)
}

func forbiddenLetterChecker(password string) bool {
	for i := 0; i < len(password); i++ {
		if password[i] == 'i' || password[i] == 'o' || password[i] == 'l' {
			return false
		}
	}

	return true
}

func threeStraightChecker(password string) bool {
	for i := 0; i < len(password) - 2; i++ {
		if password[i + 2] == password[i + 1] + 1 && password[i + 1] == password[i] + 1 {
			return true
		}
	}

	return false
}

func twoDoubleChecker(password string) bool {
	doubleCounter := 0

	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			doubleCounter++
			i++
		}
		if doubleCounter >= 2 {
			return true
		}
	}

	return false
}