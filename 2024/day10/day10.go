package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

type coordinate struct {
	y int
	x int
	value int
}
func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error parsing input")
	}
	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var nums []int
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			nums = append(nums, num)
		}
		grid = append(grid, nums)
	}
	

	queue := list.New()
	findAllThe0s(grid, queue)
	fmt.Printf("There are %d valid trails in part 1.\n", traverseTheTrails(grid, queue))
	queue2 := list.New()
	findAllThe0s(grid, queue2)
	fmt.Printf("There are %d valid trails in part 2.\n", traverseTheTrails2(grid, queue2))
}

func findAllThe0s(grid [][]int, queue *list.List) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 0 {
				queue.PushBack(coordinate{y, x, 0})
			}
		}
	}
}

func traverseTheTrails(grid [][]int, queue *list.List) int {
	colLen := len(grid)
	rowLen := len(grid[0])
	totalScore := 0

	directions := []struct{ dy, dx int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)

		start := front.Value.(coordinate)

		visited := make(map[coordinate]bool)
		localQueue := list.New()
		localQueue.PushBack(start)

		reachableNines := make(map[coordinate]bool)

		for localQueue.Len() > 0 {
			curr := localQueue.Front()
			localQueue.Remove(curr)

			coord := curr.Value.(coordinate)
			if visited[coord] {
				continue
			}
			visited[coord] = true

			if coord.value == 9 {
				reachableNines[coord] = true
				continue
			}

			for _, d := range directions {
				ny, nx := coord.y+d.dy, coord.x+d.dx
				if ny >= 0 && ny < colLen && nx >= 0 && nx < rowLen {
					if grid[ny][nx] == coord.value+1 {
						localQueue.PushBack(coordinate{ny, nx, grid[ny][nx]})
					}
				}
			}
		}

		totalScore += len(reachableNines)
	}

	return totalScore
}

func traverseTheTrails2(grid [][]int, queue *list.List) int {
	colLen := len(grid)
	rowLen := len(grid[0])
	totalScore := 0
	viableTrails := 0

	directions := []struct{ dy, dx int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)

		start := front.Value.(coordinate)

		localQueue := list.New()
		localQueue.PushBack(start)

		reachableNines := make(map[coordinate]bool)

		for localQueue.Len() > 0 {
			curr := localQueue.Front()
			localQueue.Remove(curr)

			coord := curr.Value.(coordinate)

			if coord.value == 9 {
				viableTrails++
			}

			for _, d := range directions {
				ny, nx := coord.y+d.dy, coord.x+d.dx
				if ny >= 0 && ny < colLen && nx >= 0 && nx < rowLen {
					if grid[ny][nx] == coord.value+1 {
						localQueue.PushBack(coordinate{ny, nx, grid[ny][nx]})
					}
				}
			}
		}

		totalScore += len(reachableNines)
	}

	return viableTrails
}