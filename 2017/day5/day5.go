package main

import (
	"fmt"
	"strconv"

	"github.com/Alkestian/AdventOfCode/utils"
)

func main() {
	lines, err := utils.InputReader()
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	// fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) int {
	steps := 0

	for i := 0; i != len(lines); {
		jump, _ := strconv.Atoi(lines[i])
		lines[i] = strconv.Itoa(jump + 1)
		i += jump
		steps++
	}


	return steps
}

func part2(lines []string) int {
		steps := 0

	for i := 0; i != len(lines); {
		jump, _ := strconv.Atoi(lines[i])
		if jump >= 3 {
			lines[i] = strconv.Itoa(jump - 1)
		} else {
			lines[i] = strconv.Itoa(jump + 1)
		}
		i += jump
		steps++
		if i == len(lines) {
			break
		}
	}


	return steps
}