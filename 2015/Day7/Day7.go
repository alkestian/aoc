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

	a_val, memo, instructions := part1(lines)
	for i := range memo {
		delete(memo, i)
	}
	memo["b"] = a_val
	fmt.Println(part2(memo, instructions))
}


type Instruction struct {
	operation string
	operand1 string
	operand2 string
	shiftValue int
}

func part1(lines []string) (int, map[string]int, map[string]Instruction) {
	memo := make(map[string]int)
	instructions := make(map[string]Instruction)
	for _, line := range(lines){
		destination := strings.Split(line, " -> ")[1]
		pieces := strings.Split(line, " ")
		var instruction Instruction
		if pieces[0] == "NOT" {
			instruction.operation = "NOT"
			instruction.operand1 = pieces[1]
			instruction.operand2 = ""
			instructions[destination] = instruction
		} else if len(pieces) == 3 && pieces[1] == "->" {
			instruction.operand1 = pieces[0]
			instructions[destination] = instruction
		} else if pieces[1] == "RSHIFT" || pieces[1] == "LSHIFT" {
			shiftVal, _ := strconv.Atoi(pieces[2])
			instruction.operand1 = pieces[0]
			instruction.operation = pieces[1]
			instruction.shiftValue = shiftVal
			instructions[destination] = instruction
		} else {
			instruction.operand1 = pieces[0]
			instruction.operation = pieces[1]
			instruction.operand2 = pieces[2]
			instructions[destination] = instruction
		}		
	}
	return evaluateWire("a", memo, instructions), memo, instructions
}

func part2(memo map[string]int, instructions map[string]Instruction) int {
	return evaluateWire("a", memo, instructions)
}

func evaluateWire(wire string, memo map[string]int, instructions map[string]Instruction) int {
	var result int
	if value, found := memo[wire]; found {
		return value
	}

	if val, err := strconv.Atoi(wire); err == nil {
		return val
	}

	instruction := instructions[wire]

	switch instruction.operation {
	case "":
		result = evaluateWire(instruction.operand1, memo, instructions)
	case "NOT":
		result = ^evaluateWire(instruction.operand1, memo, instructions) & 0xFFFF
	case "AND":
		result = evaluateWire(instruction.operand1, memo, instructions) &
			evaluateWire(instruction.operand2, memo, instructions)
	case "OR":
		result = evaluateWire(instruction.operand1, memo, instructions) |
			evaluateWire(instruction.operand2, memo, instructions)
	case "LSHIFT":
		result = evaluateWire(instruction.operand1, memo, instructions) << instruction.shiftValue
	case "RSHIFT":
		result = evaluateWire(instruction.operand1, memo, instructions) >> instruction.shiftValue
	}

	memo[wire] = result

	return result
}