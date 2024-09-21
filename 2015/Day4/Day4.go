package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	input := "iwrupvqb"
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	finished := false
	counter := 1
	for !finished {
		md5hash := md5.Sum([]byte(input + fmt.Sprintf("%d", counter)))
		md5Str := hex.EncodeToString(md5hash[:])
		if md5Str[:5] == "00000" {
			break
		}
		counter++
	}
	return counter	
}

func part2(input string) int {
	finished := false
	counter := 1
	for !finished {
		md5hash := md5.Sum([]byte(input + fmt.Sprintf("%d", counter)))
		md5Str := hex.EncodeToString(md5hash[:])
		if md5Str[:6] == "000000" {
			break
		}
		counter++
	}
	return counter	
}