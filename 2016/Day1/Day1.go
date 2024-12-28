package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
    fileContent, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Println("Error reading file")
        return
    }
    input := string(fileContent)
	coord := [2]int{0, 0}
	direction := "N"
	instructions := strings.Split(input, ", ")
	//fmt.Println("Part 1:", part1(coord, instructions, direction))
	fmt.Println("Part 2:", part2(coord, instructions, direction))
}

func part1(coord [2]int, instructions []string, direction string) int {
	for _, instruction := range instructions {
		direction = turn(string(instruction[0]), direction)
		distance, _ := strconv.Atoi(instruction[1:])
		switch (direction) {
			case "N":
				coord[1] += distance
			case "S":
				coord[1] -= distance
			case "E":
				coord[0] += distance
			case "W":
				coord[0] -= distance
		}
		
	}
	return int(math.Abs(float64(coord[0])) + math.Abs(float64(coord[1])))
}

func part2(coord [2]int, instructions []string, direction string) int {
	visited := make(map[[2]int]bool)
	visited[coord] = true
	for _, instruction := range instructions {
		direction = turn(string(instruction[0]), direction)
		distance, _ := strconv.Atoi(instruction[1:])
		for i := 0; i < distance; i++ {
			switch direction {
			case "N":
				coord[1]++
			case "S":
				coord[1]--
			case "E":
				coord[0]++
			case "W":
				coord[0]--
			}
			if visited[coord] {
				return int(math.Abs(float64(coord[0])) + math.Abs(float64(coord[1])))
			}
			visited[coord] = true
		}
	}
	return int(math.Abs(float64(coord[0])) + math.Abs(float64(coord[1])))
}

func turn(direction string, currentDirection string) string {
	rightTurns := map[string]string{
		"N": "E",
		"E": "S",
		"S": "W",
		"W": "N",
	}
	leftTurns := map[string]string{
		"N": "W",
		"W": "S",
		"S": "E",
		"E": "N",
	}

	if direction == "R" {
		return rightTurns[currentDirection]
	} else {
		return leftTurns[currentDirection]
	}
}