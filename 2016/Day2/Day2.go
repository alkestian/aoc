package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	keypad := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	coord := []int{1, 1}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	var instructions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	// sampleInstructions := []string{"ULL", "RRDDD", "LURDL", "UUUUD"}

	pinCode := []int{}

	for _, line := range instructions {
		for _, c := range line {
			switch c {
			case 'U':
				if coord[0] > 0 {
					coord[0]--
				}
			case 'D':
				if coord[0] < 2 {
					coord[0]++
				}
			case 'L':
				if coord[1] > 0 {
					coord[1]--
				}
			case 'R':
				if coord[1] < 2 {
					coord[1]++
				}
			}
		}
		pinCode = append(pinCode, keypad[coord[0]][coord[1]])
	}

	fmt.Println("Pin code:", pinCode)
}