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

	fmt.Println(part1(lines))
}

func part1(lines []string) int {
	var winningDistance int
	totalTime := 2503
	var distances []int
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		moveTime, _ := strconv.Atoi(pieces[6])
		restTime, _ := strconv.Atoi(pieces[13])
		speed, _ := strconv.Atoi(pieces[3])
		cycleTime := moveTime + restTime
		cycleDistance := speed * moveTime
		totalFullCycles := totalTime / cycleTime
		remainingSeconds := totalTime % cycleTime
		if remainingSeconds >= moveTime {
			distances = append(distances, ((totalFullCycles + 1) * cycleDistance))
		} else {
			distances = append(distances, (totalFullCycles * cycleDistance) + (remainingSeconds * speed))
		}
	}
	for _, distance := range distances {
		if distance > winningDistance {
			winningDistance = distance
		}
	}
	return winningDistance
}

func part2(lines []string) int {
	reindeers := make(map[string]map[string]int)
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		name := pieces[0]
		moveTime, _ := strconv.Atoi(pieces[6])
		restTime, _ := strconv.Atoi(pieces[13])
		speed, _ := strconv.Atoi(pieces[3])

		reindeers[name] = map[string]int{
			"speed":    speed,
			"moveTime": moveTime,
			"restTime": restTime,
		}
	}

	return 0
}