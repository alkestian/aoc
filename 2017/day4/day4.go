package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/Alkestian/AdventOfCode/utils"
)

func main() {
	lines, err := utils.InputReader()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

	 pattern := regexp.MustCompile(`[a-zA-Z]+`)

	 fmt.Printf("Part 1: %d\n", part1(lines, pattern))
	 fmt.Printf("Part 2: %d\n", part2(lines, pattern))
}

func part1(lines []string, pattern *regexp.Regexp) int {
	valid := 0
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

func part2(lines []string, pattern *regexp.Regexp) int {
	valid := 0
	for _, line := range lines {
		words := pattern.FindAllString(line, -1)
		var sortedWords []string
		var matchFound bool
		for _, word := range words {
			if matchFound {
				break
			}
			letters := strings.Split(word, "")
			slices.Sort(letters)
			word := strings.Join(letters, "")
			if slices.Contains(sortedWords, word) {
				matchFound = true
			}
			sortedWords = append(sortedWords, word)
		}
		if !matchFound {
			valid++
		}
	}
	return valid
}