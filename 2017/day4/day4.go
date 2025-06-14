package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		defer file.Close()
		fmt.Println(fmt.Errorf("error opening input file: %w", err))
	}

	 scanner := bufio.NewScanner(file)
	 var lines []string
	 for scanner.Scan() {
		lines = append(lines, scanner.Text())
	 }

	 fmt.Printf("Part 1: %d\n", part1(lines))
}

func part1(lines []string) int {
	valid := 0
	pattern := regexp.MustCompile(`[a-zA-Z]+`)
	for _, line := range lines {
		words := pattern.FindAllString(line, -1)
		var checkedWords []string
		var matchFound bool
		for _, word := range words {
			if matchFound {
				break
			}
			if slices.Contains(checkedWords, word) {
					matchFound = true
				}
			checkedWords = append(checkedWords, word)
		}
		if !matchFound {
			valid++
		}
	}
	return valid
}