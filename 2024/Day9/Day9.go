package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//sampleInput := "2333133121414131402"
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Error parsing input")
	}
	input := string(file)
	fileBlocks := generateFileBlocks(input)
	fmt.Printf("The pre-arrangement string is: %s\n", fileBlocks)
	arranged := arrangeFileBlocks(fileBlocks)
	fmt.Printf("The post-arrangement string is: %s\n", arranged)
	checksum := generateChecksum(arranged)
	fmt.Printf("Your checksum sum is %d\n", checksum)
	//Answer too low?
}

func generateChecksum(arranged string) int {
	result := 0
	for i := 0; i < len(arranged); i++ {
		result += (i * toInt(arranged[i]))
	}
	return result
}

func arrangeFileBlocks(fileBlocks string) string {
	runes := []rune(fileBlocks)
	
	for findFirstPeriod(string(runes)) < findLastNum(string(runes)) {
		per := findFirstPeriod(string(runes))
		num := findLastNum(string(runes))
		
		runes[per], runes[num] = runes[num], runes[per]
	}

	return strings.Trim(string(runes), ".")
}

func findFirstPeriod(fb string) int {
	for i := 0; i < len(fb); i++ {
		if fb[i] == '.' {
			return i
		}
	}
	return -1
}

func findLastNum(fb string) int {
	for i := len(fb) - 1; i >= 0; i-- {
		if fb[i] != '.' {
			return i
		}
	}
	return -1
}

func generateFileBlocks(diskMap string) string {
	var builder strings.Builder
	counter := 0
	for i := 0; i < len(diskMap); i++ {
		if i % 2 == 0 {
			builder.WriteString(strings.Repeat(strconv.Itoa(counter), toInt(diskMap[i])))
			counter++
		} else {
			builder.WriteString(strings.Repeat(".", toInt(diskMap[i])))
		}
	}
	return builder.String()
}

func toInt(s byte) int {
	conv, err := strconv.Atoi(string(s))
	if err != nil {
		log.Fatal("Error parsing string to int")
	}
	return conv
}