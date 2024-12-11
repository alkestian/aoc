package main

import "fmt"

func main() {
	grid := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	rowLen := len(grid[0])
	colLen := len(grid)
	y, x := findTheCursor(grid, rowLen, colLen)
	fmt.Printf("Part 1: %d unique points visited.\n", part1(grid, y, x, colLen, rowLen))

}

func part1(grid []string, y int, x int, rowLen int, colLen int) int {
	return -1
}

func roatetCursor(cursor rune) rune {
	if cursor == '^' {
		return '>'
	} else if cursor == '>' {
		return 'v'
	} else if cursor == 'v' {
		return '<'
	} else {
		return '^'
	}
}

func findTheCursor(grid []string, rowLen int, colLen int) (int, int) {
	for y := 0; y < colLen; y++ {
		for x := 0; x < rowLen; x++ {
			if grid[y][x] == '^' {
				return y, x
			}
		}
	}
	return -1, -1
}

//Implement a second grid to track Xs or else will run into issues if cursor moves over a populated square.
func countVisitedPoints(grid []string, rowLen int, colLen int) int {
	visitedPoints := 0
	for y := 0; y < colLen; y++ {
		for x := 0; x < rowLen; x++ {
			if grid[y][x] == 'X' {
				visitedPoints++
			}
		}
	}
	return visitedPoints
}