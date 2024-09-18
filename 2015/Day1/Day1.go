package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./Day1Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	byt, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
		
	input := string(byt)
	
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(brackets string) int {
	floor := 0
	for _, bracket := range brackets{
		if bracket == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}
	return floor
}

func part2(input string) int {
	floor := 0
	for idx, bracket := range input{
		if bracket == '(' {
			floor += 1
		} else {
			floor -= 1
		}
		if floor == -1 {
			return idx + 1
		}
	}
	return floor
}

