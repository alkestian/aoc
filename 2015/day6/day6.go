package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error with input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	grid := make(map[int]map[int]bool)

	// Initialize the grid (1000x1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make(map[int]bool)
		for j := 0; j < 1000; j++ {
			grid[i][j] = false
		}
	}

	// Process each line of instruction
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		var xStart, yStart, xEnd, yEnd int
		var instruction string

		// Parse the instruction and the start, end coordinates
		if pieces[0] == "turn" {
			instruction = pieces[1] // either "on" or "off"
			start := strings.Split(pieces[2], ",")
			end := strings.Split(pieces[4], ",")
			xStart, _ = strconv.Atoi(start[0])
			yStart, _ = strconv.Atoi(start[1])
			xEnd, _ = strconv.Atoi(end[0])
			yEnd, _ = strconv.Atoi(end[1])
		} else if pieces[0] == "toggle" {
			instruction = "toggle"
			start := strings.Split(pieces[1], ",")
			end := strings.Split(pieces[3], ",")
			xStart, _ = strconv.Atoi(start[0])
			yStart, _ = strconv.Atoi(start[1])
			xEnd, _ = strconv.Atoi(end[0])
			yEnd, _ = strconv.Atoi(end[1])
		}

		// Apply the instruction to the specified range
		for i := xStart; i <= xEnd; i++ {
			for j := yStart; j <= yEnd; j++ {
				if instruction == "on" {
					grid[i][j] = true
				} else if instruction == "off" {
					grid[i][j] = false
				} else if instruction == "toggle" {
					grid[i][j] = !grid[i][j]
				}
			}
		}
	}

	// Count how many lights are on
	lightsOn := 0
	for _, row := range grid {
		for _, value := range row {
			if value {
				lightsOn++
			}
		}
	}
	return lightsOn
}

func part2(lines []string) int {
	grid := make(map[int]map[int]int)

	// Initialize the grid (1000x1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make(map[int]int)
		for j := 0; j < 1000; j++ {
			grid[i][j] = 0
		}
	}

	// Process each line of instruction
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		var xStart, yStart, xEnd, yEnd int
		var instruction string

		// Parse the instruction and the start, end coordinates
		if pieces[0] == "turn" {
			instruction = pieces[1] // either "on" or "off"
			start := strings.Split(pieces[2], ",")
			end := strings.Split(pieces[4], ",")
			xStart, _ = strconv.Atoi(start[0])
			yStart, _ = strconv.Atoi(start[1])
			xEnd, _ = strconv.Atoi(end[0])
			yEnd, _ = strconv.Atoi(end[1])
		} else if pieces[0] == "toggle" {
			instruction = "toggle"
			start := strings.Split(pieces[1], ",")
			end := strings.Split(pieces[3], ",")
			xStart, _ = strconv.Atoi(start[0])
			yStart, _ = strconv.Atoi(start[1])
			xEnd, _ = strconv.Atoi(end[0])
			yEnd, _ = strconv.Atoi(end[1])
		}

		// Apply the instruction to the specified range
		for i := xStart; i <= xEnd; i++ {
			for j := yStart; j <= yEnd; j++ {
				if instruction == "on" {
					grid[i][j]++
				} else if instruction == "off" {
					if grid[i][j] >= 1 {
						grid[i][j]--
					}
				} else if instruction == "toggle" {
					grid[i][j] += 2
				}
			}
		}
	}

	// Count how many lights are on
	lightsOn := 0
	for _, row := range grid {
		for _, value := range row {
			lightsOn += value
		}
	}
	return lightsOn
}
