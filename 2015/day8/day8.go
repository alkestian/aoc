package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error parsing input")
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
	
}

func part1(lines []string) int {
	charCount := 0
	literalCount := 0
	for _, line := range lines {
		literalCount += len(line)
		for i := 0; i < len(line); i++ {
			if line[i] == 92 {
				if i+1 < len(line) {
					if line[i+1] == 92 || line[i+1] == 34 { 
						charCount++
						i++ 
					} else if line[i+1] == 'x' { 
						i += 3
						charCount++ 
					}
				}
			} else if line[i] == 34 {

			} else {
				charCount++ 
			}
		}
	}
	return literalCount - charCount
}

func part2(lines []string) int {
	initLiterals := 0 
	newLength := 0
	for _, line := range lines {
		newLength += 2
		initLiterals += len(line)
		for i := 0; i < len(line); i++ {
			if line[i] == 34 || line[i] == 92 {
				newLength += 2
			} else {
				newLength++
			}
		}
	}
	return newLength - initLiterals
}
