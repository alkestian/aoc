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
		log.Fatal("Error parsing input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("There are %d safe reports in part 1.\n", part1(lines))
	fmt.Printf("There are %d safe reports in part 2.\n", part2(lines))
}

func part1(lines []string) int {
	safeCount := 0
	for _, line := range lines {
		var nums []int
		pieces := strings.Split(line, " ")
		for _, x := range pieces {
			num, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal("Error casting str to int")
			}
			nums = append(nums, num)
		}
		if checkAscOrDesc(nums) && checkSafeDistance(nums) {
			safeCount++
		}
	}
	return safeCount
}

func part2(lines []string) int {
	safeCount := 0
	for _, line := range lines {
		var nums []int
		pieces := strings.Split(line, " ")
		for _, x := range pieces {
			num, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal("Error casting str to int")
			}
			nums = append(nums, num)
		}
		if checkAscOrDesc(nums) && checkSafeDistance(nums) {
			safeCount++
			continue
		}
		for idx := range nums {
			emptySlice := append([]int{}, nums[:idx]...)
			emptySlice = append(emptySlice, nums[idx+1:]...)
			if checkAscOrDesc(emptySlice) && checkSafeDistance(emptySlice) {
				safeCount++
				break
			}
		}
	}
	return safeCount
}

func checkAscOrDesc(nums []int) bool {
	ascNums := append([]int{}, nums...)
	descNums := append([]int{}, nums...)
	sort.Ints(ascNums)
	sort.Slice(descNums, func(i, j int) bool {
		return descNums[i] > descNums[j]
	})
	isAsc, isDesc := true, true
	for i := range nums {
		if nums[i] != ascNums[i] {
			isAsc = false
		}
		if nums[i] != descNums[i] {
			isDesc = false
		}
		if !isAsc && !isDesc {
			break
		}
	}
	if !isAsc && !isDesc {
		return false
	}
	return true
}

func checkSafeDistance(nums []int) bool {
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] == nums[i+1] || nums[i] - nums[i +1] < -3 || nums[i] - nums[i+1] > 3 {
			return false
		}
	}
	return true
}