package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "cqjxjnds"

	fmt.Println(part1(input))
}

func part1(input string) string {
	var password []string
	increment := true

	for i := len(input); i > 0; i-- {
		if increment {
			if input[i] == 122 {
				password = append(password, string(97))
			} else {
				password = append(password, string(input[i]))
				increment = false
			}
		} else {
			password = append(password, string(input[i]))
		}
	}

	return strings.Join(password, "")
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