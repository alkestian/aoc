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
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	register := 1
	register_b := 0
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
		pieces := strings.Split(lines[i], " ")
		switch pieces[0] {
			case "hlf":
				register /= 2
				continue
			case "tpl":
				register *= 3
				continue
			case "inc":
				if pieces[1] == "a" {
					register++
				} else {
					register_b++
				}
				continue
			case "jmp":
				num, _ := strconv.Atoi(pieces[1][1:])
				if pieces[1][0] == '+' {
					i += num - 1
				} else {
					i -= num + 1
				}
				continue
			case "jie":
				if register % 2 == 0 {
					num, _ := strconv.Atoi(pieces[2][1:])
					if pieces[2][0] == '+' {
						i += num - 1
					} else {
						i -= num + 1
					}
				}
				continue
			case "jio":
				if register == 1 {
					num, _ := strconv.Atoi(pieces[2][1:])
					if pieces[2][0] == '+' {
						i += num - 1
					} else {
						i -= num + 1
					}
				}
				continue
			default:
				log.Fatal("Somehow found an instruction that doesn't exist")
		}
	}
	fmt.Println("Final register a value:", register)
	fmt.Println("Final register b value:", register_b)
}