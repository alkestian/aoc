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
	sampleLines := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}

	lines := []string{}
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error parsing input")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}	


	fmt.Printf("The calibration value is %d for part 1.", part1(lines))

}

func part1(lines []string) int {
	total_calibration := 0
	for _, line := range lines {
		targetStr, equation := strings.Split(line, ":")[0], strings.Trim(strings.Split(line, ":")[1], " ")
		target, _ := strconv.Atoi(targetStr)
		pieces := strings.Split(equation, " ")
		nums := []int{}
		for _, piece := range pieces {
			num, _ := strconv.Atoi(piece)
			nums = append(nums, num)
		}
		if len(pieces) > 2 {
			combinations := generateCombinations(len(nums) - 1)
			for _, combination := range combinations {
				if calculate(nums, combination) == target {
					total_calibration += target
					break
				}
			}
		} else {
			if nums[0] * nums[1] == target || nums[0] + nums[1] == target {
				total_calibration += target
			}
		}
	}
	return total_calibration
}

func generateCombinations(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	if n == 1 {
		return [][]string{{"+"}, {"*"}}
	}
	smaller := generateCombinations(n - 1)
	var result [][]string
	for _, combination := range smaller {
		result = append(result, append([]string{"+"}, combination...))
		result = append(result, append([]string{"*"}, combination...))
	}
	return result
}

func calculate(pieces []int, operators []string) int {
	//append each stage to resultTotals, the last entry will be the final answer
	result := pieces[0]
	resultTotals := []int{}
	for i, operator := range operators {
		if operator == "+" {
			result += pieces[i + 1]
			resultTotals = append(resultTotals, )
		} else if operator == "*" {
			result *= pieces[i + 1]
		}
	}
	return result
}

