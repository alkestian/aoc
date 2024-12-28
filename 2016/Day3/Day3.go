package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("There were %d matches in part 1\n", part1(lines))
	fmt.Printf("There were %d matches in part 2\n", part2(lines))
}

func part1(lines []string) int {
	counter := 0
	for _, line := range lines {
		pieces := strings.Split(strings.Trim(line, " "), "  ")
		side, max := findBiggestSide(convertToIntSlice(pieces))
		if side[0] + side[1] > max {
			counter++
		}
	}
	return counter
}

func part2(lines []string) int {
	counter := 0
	for i := 0; i < len(lines); i += 3 {
        line1 := strings.Fields(lines[i])
        line2 := strings.Fields(lines[i+1])
        line3 := strings.Fields(lines[i+2])
		triangles := make([][]string, 3)
		for j := 0; j < 3; j++ {
			triangles[j] = []string{strings.TrimSpace(line1[j]), strings.TrimSpace(line2[j]), strings.TrimSpace(line3[j])}
		}
		fmt.Println(triangles)
		for _, triangle := range triangles {
			side, max := findBiggestSide(convertToIntSlice(triangle))
			if side[0] + side[1] > max {
				counter++
			}
		}
	}
	return counter
}

func findBiggestSide(sides []int) ([2]int, int) {
	sort.Ints(sides)
	smallerSides := [2]int{sides[0], sides[1]}
	return smallerSides, sides[2]
}

func convertToIntSlice(line []string) []int {
	var nums []int
	for _, piece := range line {
		num, _ := strconv.Atoi(strings.Trim(piece, " "))
		nums = append(nums, num)
	}
	return nums
}