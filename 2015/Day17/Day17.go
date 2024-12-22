package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	goal := 150
	var containers []int
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error parsing input")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		num, _ := strconv.Atoi(scanner.Text())
		containers = append(containers, num)
	}
	combos := generateCombinations(containers)
	validCombos := 0
	var correctCombos [][]int
	minLength := 100
	correctLengthCombos := 0
	for _, combo := range combos {
		length := len(combo)
		sum := 0
		fmt.Println(combo)
		for i := 0; i < len(combo); i++ {
			sum += combo[i]
		}
		if sum == goal {
			validCombos++
			correctCombos = append(correctCombos, combo)
			if length < minLength {
				minLength = length
			}
		}
	}
	for _, combo := range correctCombos {
		sum := 0
		for i := 0; i < len(combo); i++ {
			sum += combo[i]
		}
		if sum == goal && len(combo) == minLength {
			correctLengthCombos++
		}
	}
	fmt.Println("There are", validCombos, "combinations in this list.")
	fmt.Println("There are", correctLengthCombos, "combinations with the minimum number of containers.")
}

func generateCombinations(containers []int) [][]int {
	var result [][]int
	var helper func([]int, int, []int)

	helper = func(containers []int, index int, current []int) {
		if index == len(containers) {
			combination := make([]int, len(current))
			copy(combination, current)
			result = append(result, combination)
			return
		}

		helper(containers, index+1, current)

		current = append(current, containers[index])
		helper(containers, index+1, current)
	}

	helper(containers, 0, []int{})
	return result
}
