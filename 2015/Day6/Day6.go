package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
}

func part1(lines []string) map[int]map[int]int {
	grid := make(map[int]map[int]int)

	for i := 0; i < 1000; i++ {
		grid[i] = make(map[int]int)
		for j := 0; j < 1000; j++ {
			grid[i][j] = 0
		}
	}

	for _, line := range lines {
		pieces := strings.Split(line, " ")
		if pieces[1] == "on" || pieces[1] == "off" {
			instruction := pieces[0:2]
			start := pieces[2]
		} else {
			instruction := pieces[0]
			start := pieces[1]
		}
		end := pieces[len(pieces) - 1]

		
	}

	return grid
}