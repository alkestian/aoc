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
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	
	fmt.Println(part1(scanner))
	fmt.Println(part2(scanner))
}

func part1(scanner *bufio.Scanner) int {
	totalPaper := 0

	for scanner.Scan() {
		dims := strings.Split(scanner.Text(), "x")
		dimInts := make([]int, 3)                

		for i, val := range dims {
			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Error converting to int:", err)
			}
			dimInts[i] = num
		}

		area1 := dimInts[0] * dimInts[1] * 2
		area2 := dimInts[1] * dimInts[2] * 2
		area3 := dimInts[2] * dimInts[0] * 2

		totalPaper += area1 + area2 + area3 + min(area1, area2, area3)/2 
	}	

	return totalPaper
}

func part2(scanner *bufio.Scanner) int {
	ribbon := 0

	for scanner.Scan() {
		dims := strings.Split(scanner.Text(), "x")
		dimInts := make([]int, 3)    

		for i, val := range dims {
			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Error converting to int:", err)
			}
			dimInts[i] = num
		}

		noMax := withoutMax(dimInts)

		for _, val := range noMax {
			ribbon += 2*val
		}

		ribbon += dimInts[0] * dimInts[1] * dimInts[2]

	}

	return ribbon
}

func withoutMax(dims []int) [2]int {
	var noMax [2]int;
	var maxVal = dims[0]
	maxCount := 0
	for _, val := range dims{
		if val > maxVal {
			maxVal = val
		}
	}
	for _, val := range dims {
		if val == maxVal {
			maxCount++
		}
	}
	counter := 0
	for _, val := range dims{
		if val != maxVal {
			noMax[counter] = val
			counter++
		} else if maxCount > 1 {
			noMax[counter] = val
			counter++
			maxCount--
		}
	}
	return noMax
}