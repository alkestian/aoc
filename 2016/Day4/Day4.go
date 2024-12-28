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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error parsing input")
	} 
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part 1:", part1(lines))

}

func part1(lines []string) int {
	idSum := 0
	for _, line := range lines {
		alphabetCounter := make(map[rune]int)
		parts := strings.Split(line, "[")
		checksum := strings.Split(parts[1], "]")[0]
		parts = strings.Split(parts[0], "-")
		parts, sectorId := parts[:len(parts)-1], parts[len(parts)-1]
		for _, part := range parts {
			for _, char := range part {
				alphabetCounter[char]++
			}
		}
		type kv struct {
            Key   rune
            Value int
        }
        var sortedAlphabet []kv
        for k, v := range alphabetCounter {
            sortedAlphabet = append(sortedAlphabet, kv{k, v})
        }
        sort.Slice(sortedAlphabet, func(i, j int) bool {
            if sortedAlphabet[i].Value == sortedAlphabet[j].Value {
                return sortedAlphabet[i].Key < sortedAlphabet[j].Key
            }
            return sortedAlphabet[i].Value > sortedAlphabet[j].Value
        })
		var builder strings.Builder
		for i := range len(checksum) {
			builder.WriteRune(sortedAlphabet[i].Key)
		}
		if builder.String() == checksum {
			num, err := strconv.Atoi(sectorId)
			if err != nil {
				log.Fatal("error handling sectorId")
			}
			idSum += num
		}
	}

	return idSum
}