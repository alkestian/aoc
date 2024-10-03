package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	happiness := parser(lines)

}
//Generate perumtations
//Calculate max happiness of all permutations
func parser(lines []string) map[string]map[string]int {
	var happiness map[string]map[string]int;
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		key, neighbour, pos, score := pieces[0], pieces[len(pieces)-1], pieces[2], pieces[3]
		scoreInt, _ := strconv.Atoi(score)
		if pos == "gain" {
			happiness[key][neighbour] = scoreInt
		} else if pos == "lose" {
			happiness[key][neighbour] = -scoreInt
		}
	}
	return happiness
}