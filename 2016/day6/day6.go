package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) string {
	var result strings.Builder
	for i := 0; i < len(lines[0]); i++ {
		letterCounter := make(map[byte]int)
		for _, line := range lines {
			letterCounter[line[i]]++
		}
		result.WriteByte(sortMap(letterCounter, "a"))
	}
	return result.String()
}

func part2(lines []string) string {
	var result strings.Builder
	for i := 0; i < len(lines[0]); i++ {
		letterCounter := make(map[byte]int)
		for _, line := range lines {
			letterCounter[line[i]]++
		}
		result.WriteByte(sortMap(letterCounter, "b"))
	}
	return result.String()
}

func sortMap(dict map[byte]int, order string) byte {
	type kv struct {
		Key byte
		Value int
	}
	var sorter []kv
	for k, v := range dict {
		sorter = append(sorter, kv{k, v})
	}
	if order == "a" {
		sort.Slice(sorter, func(i, j int) bool {
			if sorter[i].Value == sorter[j].Value {
				return sorter[i].Key < sorter[j].Key
			}
			return sorter[i].Value > sorter[j].Value
		})
	} else {
		sort.Slice(sorter, func(i, j int) bool {
			if sorter[i].Value == sorter[j].Value {
				return sorter[i].Key > sorter[j].Key
			}
			return sorter[i].Value < sorter[j].Value
		})
	}
	return sorter[0].Key
}