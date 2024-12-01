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
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("An error occured parsing input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var column1, column2 []int

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("Error parsing ints")
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("Error parsing ints")
		}
		column1, column2 = append(column1, num1), append(column2, num2)
	}

	fmt.Printf("The total distance between lists is: %d.\n", part1(column1, column2))
	fmt.Printf("The similarity score between the lists is: %d", part2(column1, column2))
}

func part1(column1 []int, column2 []int) int {
	totalDistance := 0
	sortedColumn1, sortedColumn2 := column1, column2
	sort.Ints(sortedColumn1)
	sort.Ints(sortedColumn2)
	for index, value := range column1 {
		distance := value - column2[index]
		if distance < 0 {
			totalDistance += -distance
		} else {
			totalDistance += distance
		}
	}
	return totalDistance
}

func checkIfKeyExists(hashMap map[int]int, key int) bool {
	_, exists := hashMap[key]
	return exists
}

func part2(column1 []int, column2 []int) int {
	hashMap := make(map[int]int)
	similarityScore := 0
	for _, outerValue := range column1 {
		if !checkIfKeyExists(hashMap, outerValue) {
			hashMap[outerValue] = 0
		}
		for _, innerValue := range column2 {
			if innerValue == outerValue {
				hashMap[outerValue]++
			}
		}
	}
	for key, value := range hashMap {
		similarityScore += key * value
	}
	return similarityScore
}