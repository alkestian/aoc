package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "cxdnnyjw"

	fmt.Println("Part 2:", part2(input))
}

func part1(doorId string) string {
	var password strings.Builder
	index := 0
	for password.Len() < 8 { 
		prehash := doorId + strconv.Itoa(index)
		hash := md5.New()
		hash.Write([]byte(prehash))
		hashString := hex.EncodeToString(hash.Sum(nil))
		if hashString[:5] == "00000" {
			password.WriteByte(hashString[5])
		} else {
			fmt.Println(index, "...", password.String())
		}
		index++
	}
	return password.String()
}

func part2(doorId string) string {
	var password [8]string
	index := 0
	for {
		matchCounter := 0
		for i := 0; i < 8; i++ {
			if password[i] != "" {
				matchCounter++
			}
		}
		if matchCounter >= 8 {
			return strings.Join(password[:], "")
		}
		prehash := doorId + strconv.Itoa(index)
		hash := md5.New()
		hash.Write([]byte(prehash))
		hashString := hex.EncodeToString(hash.Sum(nil))
		if hashString[:5] == "00000" {
			position, _ := strconv.Atoi(string(hashString[5]))
			if position < 8 {
				password[position] = string(hashString[6])
			}
		} else {
			fmt.Println(index, "...", password)
		}
		index++
	}
}