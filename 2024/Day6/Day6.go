package main

import "fmt"

func main() {
	grid := [][]string{
		{"." , ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{"." , ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"." , ".", "#", ".", ".", ".", ".", ".", ".", "."},
		{"." , ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{"." , ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"." , "#", ".", ".", "^", ".", ".", ".", ".", "."},
		{"." , ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{"#" , ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"." , ".", ".", ".", ".", ".", "#", ".", ".", "."},
	}

	rowLen := len(grid[0])
	colLen := len(grid)
	y, x := findTheCursor(grid, rowLen, colLen)
	fmt.Printf("Part 1: %d unique points visited.\n", part1(grid, y, x, colLen, rowLen, "^"))

}

func part1(grid [][]string, y int, x int, rowLen int, colLen int, cursor string) int {
	grid2 := grid
	for {
		initialX, initialY := x, y
		updateGrid(grid2, initialY, initialX)
		if cursor == "^" && y - 1 >= 0 {
			if grid[y-1][x] == "#" {
				cursor = rotateCursor(cursor)
			} else {
				y = y - 1
			}
		} else if cursor == ">" && x + 1 <= rowLen {
			if grid[y][x+1] == "#" {
				cursor = rotateCursor(cursor)
			} else {
				x = x + 1
			}
		} else if cursor == "v" && y + 1 <= colLen {
			if grid[y+1][x] == "#" {
				cursor = rotateCursor(cursor)
			} else {
				y = y + 1
			}
		} else if cursor == "<" && x - 1 >= 0 {
			if grid[y][x-1] == "#" {
				cursor = rotateCursor(cursor)
			} else {
				x = x - 1
			}
		} else {
			break
		}
	}
	return countVisitedPoints(grid2, rowLen, colLen)
}

func rotateCursor(cursor string) string {
	if cursor == "^" {
		return ">"
	} else if cursor == ">" {
		return "v"
	} else if cursor == "v" {
		return "<"
	} else {
		return "^"
	}
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