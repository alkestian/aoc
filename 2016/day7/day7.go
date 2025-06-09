package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error parsing input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	fmt.Println("Part 1:", part1(lines))
    fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) int {
    abbaCounter := 0
    for _, line := range lines {
        var pieces []string
        var hypernets []string

        parts := strings.Split(line, "[")
        pieces = append(pieces, parts[0])
        for i := 1; i < len(parts); i++ {
            subParts := strings.Split(parts[i], "]")
            hypernets = append(hypernets, subParts[0])
            if len(subParts) > 1 {
                pieces = append(pieces, subParts[1])
            }
        }

        skip := false
        for _, hypernet := range hypernets {
            if containsABBA(hypernet) {
                skip = true
                break
            }
        }
        if skip {
            continue
        }

        for _, piece := range pieces {
            if containsABBA(piece) {
                abbaCounter++
                break
            }
        }
    }
    return abbaCounter
}

func part2(lines []string) int {
    abaCounter := 0
    for _, line := range lines {
        var pieces []string
        var hypernets []string

        parts := strings.Split(line, "[")
        pieces = append(pieces, parts[0])
        for i := 1; i < len(parts); i++ {
            subParts := strings.Split(parts[i], "]")
            hypernets = append(hypernets, subParts[0])
            if len(subParts) > 1 {
                pieces = append(pieces, subParts[1])
            } 
        }

        var abas []string 

        for _, piece := range pieces {
            //check for aba
            for i := 0; i <= len(piece) - 3; i++ {
                segment := piece[i:i+3]
                if segment[0] == segment[2] && segment[0] != segment[1] {
                    abas = append(abas, segment)
                } 
            }
        }

        match := false
        for _, aba := range abas {
            bab := string(aba[1]) + string(aba[0]) + string(aba[1]) // Generate corresponding BAB
            for _, hypernet := range hypernets {
                if strings.Contains(hypernet, bab) {
                    match = true
                    break
                }
            }
            if match {
                abaCounter++
                break
            }
        }
    }
    return abaCounter
}

func containsABBA(s string) bool {
    for i := 0; i <= len(s)-4; i++ {
        if s[i] != s[i+1] && s[i:i+2] == reverse(s[i+2:i+4]) {
            return true
        }
    }
    return false
}

func reverse(x string) string {
	runes := []rune(x)
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}