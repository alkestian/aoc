package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("error reading file: %w", err))
	}
	banks := strings.Fields(string(data))

	fmt.Println("Part 1:", part1(banks))

}

func part1(banks []string) int {
	cycles := 0
	bank, max := findMaxBank(banks)

	return cycles
}

func findMaxBank(blocks []string) (int, int) {
	max := 0 
	index := 0
	for i, b := range blocks {
		b, _ := strconv.Atoi(b)
		if b >= max {
			max = b
			index = i
		}
	}
	return index, max
}