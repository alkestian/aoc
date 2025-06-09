package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byt, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
		
	input := string(byt)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	houses := 0
	xCo := 0
	yCo := 0
	var visited []string;

	isVisited := func(coord string, visited []string) bool {
		for _, v := range visited {
			if v == coord {
				return true
			}
		}
		return false
	}

	for _, val := range input {
		if string(val) == "v" {
			xCo--
			coord := fmt.Sprintf("%d,%d", xCo, yCo)
			if !isVisited(coord, visited) {
				visited = append(visited, coord)
				houses++
			}
		} else if string(val) == "^" {
			xCo++
			coord := fmt.Sprintf("%d,%d", xCo, yCo)
			if !isVisited(coord, visited) {
				visited = append(visited, coord)
				houses++
			}
		} else if string(val) == "<" {
			yCo--
			coord := fmt.Sprintf("%d,%d", xCo, yCo)
			if !isVisited(coord, visited) {
				visited = append(visited, coord)
				houses++
			}
		} else {
			yCo++
			coord := fmt.Sprintf("%d,%d", xCo, yCo)
			if !isVisited(coord, visited) {
				visited = append(visited, coord)
				houses++
			}
		}
	}
	return houses
}

func part2(input string) int {
	houses := 0
	xCo := 0
	yCo := 0
	roboXCo := 0
	roboYCo := 0

	var visited []string;

	coord := fmt.Sprintf("%d,%d", xCo, yCo)
	visited = append(visited, coord)
	houses++

	isVisited := func(coord string, visited []string) bool {
		for _, v := range visited {
			if v == coord {
				return true
			}
		}
		return false
	}

	for i, val := range input {
		if i == 0 || i % 2 == 0 {
			if string(val) == "v" {
				xCo--
				coord := fmt.Sprintf("%d,%d", xCo, yCo)
				if !isVisited(coord, visited) {
					visited = append(visited, coord)
					houses++
				}
			} else if string(val) == "^" {
				xCo++
				coord := fmt.Sprintf("%d,%d", xCo, yCo)
				if !isVisited(coord, visited) {
					visited = append(visited, coord)
					houses++
				}
			} else if string(val) == "<" {
				yCo--
				coord := fmt.Sprintf("%d,%d", xCo, yCo)
				if !isVisited(coord, visited) {
					visited = append(visited, coord)
					houses++
				}
			} else {
				yCo++
				coord := fmt.Sprintf("%d,%d", xCo, yCo)
				if !isVisited(coord, visited) {
					visited = append(visited, coord)
					houses++
				}
			}
		} else {
			if string(val) == "v" {
				roboXCo--
				coord := fmt.Sprintf("%d,%d", roboXCo, roboYCo)
				if !isVisited(coord, visited) {
					visited = append(visited, coord)
					houses++
				}
			} else if string(val) == "^" {
				roboXCo++
				coord := fmt.Sprintf("%d,%d", roboXCo, roboYCo)
				if !isVisited(coord, visited) {
					visited = append(visited, coord)
					houses++
				}
			} else if string(val) == "<" {
				roboYCo--
				coord := fmt.Sprintf("%d,%d", roboXCo, roboYCo)
				if !isVisited(coord, visited) {
					visited = append(visited, coord)
					houses++
				}
			} else {
				roboYCo++
				coord := fmt.Sprintf("%d,%d", roboXCo, roboYCo)
				if !isVisited(coord, visited) {
					visited = append(visited, coord)
					houses++
				}
			}
		}
	}

	return houses
}