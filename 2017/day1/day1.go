package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("error opening file: %v", err))
	}
	captcha := string(file)
	part1(captcha)
	part1(captcha)
}

func part1(captcha string) {	
	var total int
	for i, x := range captcha {
		if x == rune(captcha[(i+1) % len(captcha)]) {
			total += int(x - '0')
		}	
	}

	fmt.Println(total)
}

func part2(captcha string) {
	var total int
	for i, x := range captcha {
		if x == rune(captcha[(i + len(captcha)/2) % len(captcha)]) {
			total += int(x - '0')
		}	
	}

	fmt.Println(total)
}