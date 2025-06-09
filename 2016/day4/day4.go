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
	fmt.Println("Part 2:")
	for name, sid := range part2(lines) {
		fmt.Println(name, sid)
	}

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

func part2(lines []string) map[string]int {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	realNames := make(map[string]int)

	for _, line := range lines {
		parts := strings.Split(line, "[")
		parts = strings.Split(parts[0], "-")
		parts, sid := parts[:len(parts)-1], parts[len(parts)-1]
		sectorId, _ := strconv.Atoi(sid)
		unparted := strings.Join(parts, "-")
		shiftKey := alphabet[sectorId%26:] + alphabet[:sectorId%26]
		cipher := make(map[string]string)
		cipher["-"] = " "
		for i, char := range alphabet {
			cipher[string(char)] = string(shiftKey[i])
		}
		var decrypted strings.Builder
		for _, char := range unparted {
			decrypted.WriteString(cipher[string(char)])
		}
		realNames[decrypted.String()] = sectorId
	}

	return realNames
}