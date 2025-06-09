package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()

	lines := bufio.NewScanner(file)
	//fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines *bufio.Scanner) int {
	niceStrings := 0
	for lines.Scan() {
		line := lines.Text()
		niceLine := true
        if niceLine {
            niceLine = fragmentChecker(line)
        }
        if niceLine {
            niceLine = vowelChecker(line)
        }
        if niceLine {
            niceLine = doubleChecker(line)
        }
		if niceLine {
			niceStrings++
		}
	}
	return niceStrings
}

func part2(lines *bufio.Scanner) int {
	niceStrings := 0
	for lines.Scan(){
		line := lines.Text()
		if letterBetweenChecker(line) && pairChecker(line) {
			niceStrings++
		}
	}
	return niceStrings	
}

func fragmentChecker(line string) bool {
	fragments := []string{"pq", "cd", "ab", "xy"}
	for _, fragment := range fragments {
		if strings.Contains(line, fragment) {
			return false
		} 
	}
	return true
}

func vowelChecker(line string) bool {
	vowels := "aeiou"
	count := 0
	for _, char := range line {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}	
	return count >= 3
}

func doubleChecker(line string) bool {
	for i := range len(line) - 1 {
		if line[i] == line[i + 1] {
			return true
		}
	}
	return false
}

func letterBetweenChecker(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i + 2] {
			return true
		}
	}
	return false
}

func pairChecker(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		pair := line[i : i+2]
		for j := i + 2; j < len(line)-1; j++ {
			if line[j:j+2] == pair {
				return true
			}
		}
	}
	return false
}
