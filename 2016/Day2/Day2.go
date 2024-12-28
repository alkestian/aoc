package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
	//instructions := []string{"ULL", "RRDDD", "LURDL", "UUUUD"}

	fmt.Println("Pin code:", part1(instructions))
	fmt.Println("Pin code:", string(part2(instructions)))
}

func part1(instructions []string) []int {
	keypad := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	coord := []int{1, 1}
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

	return pinCode
}

func part2(instructions []string) []rune {
    keypad := [][]rune{
        {0, 0, '1', 0, 0},
        {0, '2', '3', '4', 0},
        {'5', '6', '7', '8', '9'},
        {0, 'A', 'B', 'C', 0},
        {0, 0, 'D', 0, 0},
    }

	coord := []int{2, 0}
	pinCode := []rune{}

	for _, line := range instructions {
		for _, c := range line {
			switch c {
			case 'U':
				if coord[0] > 0 && keypad[coord[0]-1][coord[1]] != 0 {
					coord[0]--
				}
			case 'D':
				if coord[0] < 4 && keypad[coord[0]+1][coord[1]] != 0 {
					coord[0]++
				}
			case 'L':
				if coord[1] > 0 && keypad[coord[0]][coord[1]-1] != 0 {
					coord[1]--
				}
			case 'R':
				if coord[1] < 4 && keypad[coord[0]][coord[1]+1] != 0 {
					coord[1]++
				}
			}
		}
		pinCode = append(pinCode, keypad[coord[0]][coord[1]])
	}

	return pinCode
}