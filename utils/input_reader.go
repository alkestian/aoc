package utils

import (
	"bufio"
	"fmt"
	"os"
)

func InputReader() ([]string, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening input file: %w", err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()
	return lines, nil
}