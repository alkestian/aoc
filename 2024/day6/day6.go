package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var grid [][]string
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error parsing input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		grid = append(grid, line)
	}
	
	rowLen := len(grid[0])
	colLen := len(grid)
	y, x := findTheCursor(grid, rowLen, colLen)
	fmt.Printf("Part 1: %d unique points visited.\n", part1(grid, y, x, colLen, rowLen, "^"))

}

func part1(grid [][]string, y int, x int, rowLen int, colLen int, cursor string) int {
	grid2 := make([][]string, rowLen)
    for i := range grid {
        grid2[i] = append([]string{}, grid[i]...)
    }

	for {
		if y < 0 || y >= rowLen || x < 0 || x >= colLen {
			break
		}
		
		updateGrid(grid2, y, x)

		// Move based on the current cursor direction
		switch cursor {
		case "^":
			if y-1 >= 0 {
				if grid[y-1][x] == "#" {
					cursor = rotateCursor(cursor) // Rotate if blocked
				} else {
					y-- // Move up
				}
			} else {
				y--
			} 
		case ">":
			if x+1 < colLen {
				if grid[y][x+1] == "#" {
					cursor = rotateCursor(cursor) // Rotate if blocked
				} else {
					x++ // Move right
				}
			} else {
				x++
			}
		case "v":
			if y+1 < rowLen {
				if grid[y+1][x] == "#" {
					cursor = rotateCursor(cursor) // Rotate if blocked
				} else {
					y++ // Move down
				}
			} else {
				y++
			}
		case "<":
			if x-1 >= 0 {
				if grid[y][x-1] == "#" {
					cursor = rotateCursor(cursor) // Rotate if blocked
				} else {
					x-- // Move left
				}
			} else {
				x--
			}
		}
	}
	return countVisitedPoints(grid2, rowLen, colLen)
}



func rotateCursor(cursor string) string {
	newCursor := ""
	if cursor == "^" {
		newCursor = ">"
	} else if cursor == ">" {
		newCursor = "v"
	} else if cursor == "v" {
		newCursor = "<"
	} else if cursor == "<" {
		newCursor = "^"
	}
	return newCursor
}

func findTheCursor(grid [][]string, rowLen int, colLen int) (int, int) {
	for y := 0; y < colLen; y++ {
		for x := 0; x < rowLen; x++ {
			if grid[y][x] == "^"|| grid[y][x] == ">" || grid[y][x] == "<" || grid[y][x] == "v" {
				return y, x
			}
		}
	}
	return -1, -1
}

func updateGrid(grid [][]string, y int, x int) [][]string {
	grid[y][x] = "X"
	return grid
}

func countVisitedPoints(grid [][]string, rowLen int, colLen int) int {
	visitedPoints := 0
	for y := 0; y < colLen; y++ {
		for x := 0; x < rowLen; x++ {
			if grid[y][x] == "X" {
				visitedPoints++
			}
		}
	}
	return visitedPoints
}