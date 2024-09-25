package main

import (
	"fmt"
	"strings"
)

func main() {
	sequence := "1321131112"

	for i := 0; i < 50; i++{
		sequence = lookAndSay(sequence)
	}

	fmt.Println(len(sequence))
}

func lookAndSay(sequence string) string {
	var result []string
	var currentChar rune
	var count int

	for i, char := range sequence {
		if i == 0 {
			currentChar = char
			count = 1
		} else {
			if char == currentChar {
				count++
			} else {
				result = append(result, fmt.Sprintf("%d%c", count, currentChar))
				currentChar = char
				count = 1
			}
		}
	}
	
	if count > 0 {
		result = append(result, fmt.Sprintf("%d%c", count, currentChar))
	}

	return strings.Join(result, "")
}
