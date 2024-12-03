package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Error parsing input")
	}
	input := string(file)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	re := regexp.MustCompile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
	matches := re.FindAllString(input, -1)

	total := 0

	for _, match := range matches {
		trimmed := strings.TrimSuffix(strings.TrimPrefix(match, "mul("), ")")
		parts := strings.Split(trimmed, ",")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])

		total += num1 * num2
	}
	
	return total
}

func part2(input string) int {
	re := regexp.MustCompile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)|do\(\)|don\'t\(\)`)
	matches := re.FindAllString(input, -1)

	total := 0
	do := true
	for _, match := range matches {
		if match == "do()" {
			do = true
		} else if match == "don't()"{
			do = false
		} else {
			if do {
				trimmed := strings.TrimSuffix(strings.TrimPrefix(match, "mul("), ")")
				parts := strings.Split(trimmed, ",")
				num1, _ := strconv.Atoi(parts[0])
				num2, _ := strconv.Atoi(parts[1])

				total += num1 * num2
			}
		}
	}
	
	return total
}