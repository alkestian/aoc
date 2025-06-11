package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		defer file.Close()
		fmt.Printf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	var checksum int
	pattern, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		fmt.Printf("error compiling regex: %v", err)
		return -1
	}

	for _, line := range lines {
		matches := pattern.FindAllString(line, -1)
		min, max := 1000000, 0
		for _, match := range matches {
			num, err := strconv.Atoi(match); if err != nil {
				fmt.Printf("error converting string to int: %v", err)
				return -1
			}
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
		checksum += max - min
	}

	return checksum
}

func part2 (lines []string) int {
	var checksum int
	pattern, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		fmt.Printf("error compiling regex: %v", err)
		return -1
	}

    for _, line := range lines {
        matches := pattern.FindAllString(line, -1)
        nums := make([]int, len(matches))
        for i, m := range matches {
            nums[i], _ = strconv.Atoi(m)
        }
        found := false
        for i := 0; i < len(nums) && !found; i++ {
            for j := range nums {
                if i == j {
                    continue
                }
                if nums[i]%nums[j] == 0 {
                    checksum += nums[i] / nums[j]
                    found = true
                    break
                }
            }
        }
    }

	return checksum
}