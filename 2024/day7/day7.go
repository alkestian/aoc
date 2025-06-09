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
	// sampleLines := []string{
	// 	"190: 10 19",
	// 	"3267: 81 40 27",
	// 	"83: 17 5",
	// 	"156: 15 6",
	// 	"7290: 6 8 6 15",
	// 	"161011: 16 10 13",
	// 	"192: 17 8 14",
	// 	"21037: 9 7 18 13",
	// 	"292: 11 6 16 20",
	// }

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
	fmt.Printf("The calibration value is %d for part 2.", part2(lines))
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

func part2(lines []string) int {
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

        if len(pieces) == 2 {
            if nums[0]*nums[1] == target || nums[0]+nums[1] == target ||
               strconv.Itoa(nums[0])+strconv.Itoa(nums[1]) == strconv.Itoa(target) {
                total_calibration += target
            }
        } else {
            combinations := generateCombinationsPart2(len(nums) - 1)
            for _, combination := range combinations {
                if calculate(nums, combination) == target {
                    total_calibration += target
                    break
                }
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

func generateCombinationsPart2(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	if n == 1 {
		return [][]string{{"+"}, {"*"}, {"||"}}
	}
	smaller := generateCombinationsPart2(n - 1)
	var result [][]string
	for _, combination := range smaller {
		result = append(result, append([]string{"+"}, combination...))
		result = append(result, append([]string{"*"}, combination...))
		result = append(result, append([]string{"||"}, combination...))
	}
	return result
}

func calculate(pieces []int, operators []string) int {
    resultTotals := []int{pieces[0]}

    for i, operator := range operators {
        if operator == "+" {
            resultTotals = append(resultTotals, resultTotals[len(resultTotals)-1]+pieces[i+1])
        } else if operator == "*" {
            resultTotals = append(resultTotals, resultTotals[len(resultTotals)-1]*pieces[i+1])
        } else if operator == "||" {
            stringy, err := strconv.Atoi(strconv.Itoa(resultTotals[len(resultTotals)-1]) + strconv.Itoa(pieces[i+1]))
            if err != nil {
                log.Fatalf("Error processing concatenation: %v", err)
            }
            resultTotals = append(resultTotals, stringy)
        }
    }
    return resultTotals[len(resultTotals)-1]
}

